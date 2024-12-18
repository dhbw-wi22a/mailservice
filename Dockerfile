# 1. Usage of go image while build process
FROM golang:latest As builder

# 2. Set work directory
WORKDIR /mailservice

# 3. Copy go mod and sum files
COPY go.mod go.sum ./

# 4. Get go modules
RUN go mod download

# 5. Copy all files to the container
COPY . .

# 6. Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o mailservice .

# 7. Small image for production
FROM alpine:latest

# 8. Set directory
WORKDIR /root/

# 9. Copy go binary from builder
COPY --from=builder /mailservice .

# 10. Expose port
EXPOSE 8999

# 11. Startcommand
CMD ["./mailservice"]