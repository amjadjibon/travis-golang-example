FROM golang:1.20 as builder

WORKDIR /app
COPY . /app
RUN make build

# APP
FROM alpine:latest
RUN apk --no-cache add tzdata ca-certificates mailcap && addgroup -S app && adduser -S app -G app
RUN echo "Asia/Dhaka" > /etc/timezone
RUN cp /usr/share/zoneinfo/Asia/Dhaka /etc/localtime

USER app
WORKDIR /app
EXPOSE 8080
COPY --from=builder /app/bin/travis-golang-example /usr/local/bin/travis-golang-example
CMD ["travis-golang-example"]
