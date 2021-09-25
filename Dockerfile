# Start from the latest golang base image
FROM golang:1.14

# Setting ENv
ENV HOST "db"

# Add Maintainer Info
LABEL maintainer="Rishabh Mishra <rishabhmishra131@gmail.com>"
 
# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags static_all -a -installsuffix cgo -ldflags '-extldflags "-static"' -o './bin/chatroom' *.go
 
# Expose port 8080 to the outside world
EXPOSE 8080

ADD database/database.sql /docker-entrypoint-initdb.d/ 

# Command to run the executable 
CMD ["bash", "-c", "./bin/chatroom"]
