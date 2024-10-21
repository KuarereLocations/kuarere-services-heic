FROM golang:1.22-bullseye

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app

COPY go.mod .
COPY go.sum .

# Installs Go dependencies
RUN go mod download

COPY . .

# Builds your app with optional configuration
RUN go build -o /app/bin/kuarere-service /app/cmd/http/

# Tells Docker which network port your container listens on
EXPOSE 4500

# Specifies the executable command that runs when the container starts
CMD ["/app/bin/kuarere-service" ]