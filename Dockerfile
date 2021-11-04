FROM golang:alpine AS builder
# Copy local code to the container image.
WORKDIR /app
COPY . .
# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o crudapp

FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /src/simple_service /app/
ENTRYPOINT ./simple_service