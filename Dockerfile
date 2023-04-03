# Use Go 1.17-alpine as base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the module dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

RUN apk update && apk add --no-cache gcc musl-dev

# Build the application
RUN go build -o app .

# Set the command to run the application
CMD ["/app/app"]