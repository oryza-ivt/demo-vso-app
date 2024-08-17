# Stage 1: Build the Go binary
FROM golang:1.23 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files to download dependencies
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app for Linux with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./main.go

# Stage 2: Create the final image using a distroless base image
FROM gcr.io/distroless/base-debian10

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/main /app/main

# Expose the port the service will run on
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
