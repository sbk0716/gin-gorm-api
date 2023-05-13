FROM golang:1.19

# Set destination for COPY
WORKDIR /go/src/app

COPY . .
# Download Go modules
RUN go mod download

EXPOSE 8000

CMD ["go", "run", "main.go"]