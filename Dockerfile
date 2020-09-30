FROM golang as builder

LABEL maintainer="Himadri"

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the source files from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

############### Start a new stage from scratch ###################
FROM alpine

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 50000
EXPOSE 50000

# Execute the application
CMD [ "./main" ]