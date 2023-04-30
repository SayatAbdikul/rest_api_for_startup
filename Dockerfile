# Use official Golang image as base image
FROM golang:latest

# Set working directory inside container
WORKDIR /app

# Copy necessary files to container
COPY . .

# Install MySQL driver
RUN go get -u github.com/go-sql-driver/mysql

# Build Go binary
RUN go build -o main .

# Expose port 8080
EXPOSE 9000

# Start Go binary when container starts
CMD ["./main"]