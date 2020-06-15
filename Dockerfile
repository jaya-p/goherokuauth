# Docker build: docker build -t gha .
# Set app: heroku git:remote --app <app-name>
# Build for heroku: heroku container:push <app-name>
# Push for heroku: heroku container:release <app-name>
# See logs in heroku: heroku logs --tail

# This Dockerfile use multi stage approach. STAGE 1 is to build the binary
#    and STAGE 2 is to have smallest image possible by including only necessary binary

# STAGE 1 is to build the binary
# Use golang-based image for container; golang version 1.14.4
FROM golang:1.14.4-alpine AS builder

# GO111MODULE=on to enable go modules for downloading dependencies
# CGO_ENABLED=0 to include all linked library included in the output binary
ENV GO111MODULE=on \  
  CGO_ENABLED=0 

# Add git executable in container
RUN apk add --no-cache git

# Copy webserver go file in local computer to container
COPY ./driver/mainrestapi.go /go/src/

# Set working directory in container
WORKDIR /go

# Build the application
RUN go build -o bin/main src/mainrestapi.go

# STAGE 2 is to have smallest image possible by including only necessary binary
# Use smallest base image
FROM scratch

# Copy application binary from STAGE 1 image to STAGE 2 image
COPY --from=builder /go/bin/main /

# Expose listening port for application
EXPOSE 8080

# Run the application
ENTRYPOINT ["/main"]
