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

EXPOSE 9001

# Start Go binary when container starts
CMD ["./main"]