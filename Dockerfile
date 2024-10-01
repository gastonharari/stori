FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o processtransactions ./cmd

FROM alpine:latest

WORKDIR /root/

# Copy the binary from the build stage
COPY --from=build /app/processtransactions .

# Make sure the binary is executable
RUN chmod +x /root/processtransactions

ENTRYPOINT ["/root/processtransactions"]
