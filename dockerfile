FROM golang:1.22-alpine AS builder

COPY . /github.com/Artenso/command-runner/
WORKDIR /github.com/Artenso/command-runner/

RUN apk add --no-cache bash
RUN go mod download
RUN go build -o ./bin/command-runner cmd/command_runner/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/Artenso/command-runner/bin/command-runner .
COPY --from=builder /github.com/Artenso/command-runner/config.json .

EXPOSE 50051
EXPOSE 8000

CMD ["./command-runner"]