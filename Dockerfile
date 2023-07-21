FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/sprox ./cmd/sprox

# make a new image from scratch with only the binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /bin/sprox /bin/sprox

ENTRYPOINT ["/bin/sprox"]