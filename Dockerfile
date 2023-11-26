# Install dependencies and compile to a binary
FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .

ENV GIN_MODE=release
# Define the target architecture so it can still be deployed to an x86 machine if compiled on an ARM machine
ENV GOOS=linux GOARCH=amd64
RUN go build -o main cmd/main.go

# Step 2: build a small image that runs the binary
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
EXPOSE 8080

CMD ["/app/main"]