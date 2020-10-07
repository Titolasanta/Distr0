FROM golang:1.14

ENV SERVER_PORT=11111
ENV SERVER_LISTEN_BACKLOG=5
ENV CLI_ID=1
ENV CLI_SERVER_ADDRESS=server:12345
ENV CLI_LOOP_LAPSE=1m2s
ENV CLI_LOOP_PERIOD=10s

RUN mkdir -p /build
WORKDIR /build/
COPY . .
# CGO_ENABLED must be disabled to run go binary in Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/setup


FROM busybox:latest
COPY --from=0 /build/bin/setup /setup