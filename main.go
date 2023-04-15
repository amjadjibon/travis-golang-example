package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v8"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Host     string `env:"HOST" envDefault:"0.0.0.0"`
	Port     int    `env:"PORT" envDefault:"8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

func NewConfig() *Config {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func SetLogger(cfg *Config) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if cfg.LogLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if cfg.LogLevel == "info" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else if cfg.LogLevel == "warn" {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else if cfg.LogLevel == "error" {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	} else {
		panic("invalid log level")
	}
}

type Handlers struct {
	cfg *Config
}

func (h *Handlers) Index() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return index
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Debug().
		Str("path", r.URL.Path).
		Msg("index")

	_, _ = fmt.Fprint(w, "Welcome!\n")
}

func (h *Handlers) Hello() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return hello
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Debug().
		Str("path", r.URL.Path).
		Str("name", ps.ByName("name")).
		Msg("hello")
	_, _ = fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func main() {
	cfg := NewConfig()
	SetLogger(cfg)
	handlers := &Handlers{cfg: cfg}
	router := httprouter.New()
	router.GET("/", handlers.Index())
	router.GET("/hello/:name", handlers.Hello())
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	log.Info().
		Str("addr", addr).
		Msg("starting server")

	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
