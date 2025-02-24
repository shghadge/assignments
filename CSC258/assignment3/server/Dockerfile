# Start with the official Go base image
FROM golang:1.23-alpine as build-stage

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY ./server.go .

# Build the Go app
RUN go build -o main .

FROM scratch

COPY --from=build-stage /app/main .

# Expose port 8080
EXPOSE 8086

# Command to run the application
CMD ["./main"]