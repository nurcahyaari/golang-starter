## BUILD STAGE
FROM golang:1.16-alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app/golang-stater

COPY . .

RUN go generate .
RUN go build -o app .

## DISTRIBUTION
FROM alpine:latest

COPY --from=builder /app/golang-stater/config.yaml .
COPY --from=builder /app/golang-stater/app .

CMD /app
