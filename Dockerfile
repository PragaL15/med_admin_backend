# Step 1: Use an official Golang image as the base image
FROM golang:1.20-alpine AS builder

# Step 2: Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Step 3: Install necessary packages
RUN apk update && apk add --no-cache git

# Step 4: Set the working directory inside the container
WORKDIR /app

# Step 5: Copy Go modules manifests
COPY go.mod go.sum ./

# Step 6: Download all Go module dependencies
RUN go mod download

# Step 7: Copy the entire project into the container
COPY . .

# Step 8: Build the Go application
RUN go build -o main .

# Step 9: Use a minimal base image for the final stage
FROM alpine:latest

# Step 10: Set the working directory for the final container
WORKDIR /root/

# Step 11: Copy the binary from the builder stage
COPY --from=builder /app/main .

# Step 12: Expose the application port
EXPOSE 8080

# Step 13: Set the command to run the application
CMD ["./main"]
