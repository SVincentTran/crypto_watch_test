# Golang Img
FROM golang:1.19.2-bullseye

# Creates an work directory
WORKDIR /app

# Copies source code from root directory into /app
COPY . .

# Installs go dependencies
RUN go mod download

# Builds app
RUN go build -o /gocrypto

# Expose port
EXPOSE 3000

# Excute command
CMD ["/gocrypto"]