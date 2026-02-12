# Build Stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Run Stage (Distroless)
FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/main /
EXPOSE 8080
CMD ["/main"]
