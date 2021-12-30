FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./docs ./docs
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" ./cmd/apiserver

FROM scratch

COPY ./docs ./docs
COPY --from=builder /app/apiserver /usr/bin/
EXPOSE 8080

ENTRYPOINT ["apiserver"]
