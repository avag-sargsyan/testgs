FROM golang:1.20 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o app ./cmd/app/main.go

# Use multistage build to reduce image size
FROM alpine:latest

# Set necessary environment variables here
#ENV GS_HTTP_ADDRESS=8080
#ENV GS_PACK_SIZES=250,500,1000,2000,5000

COPY --from=builder /src/app /app

CMD ["/app"]
