FROM golang:1.22

WORKDIR /usr/src/coreapp

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY wait-for-it.sh wait-for-it.sh
RUN chmod +x wait-for-it.sh

COPY . .
RUN go build -o coreapp ./cmd/main.go
CMD ["./wait-for-it.sh", "rabbitmq:5672", "-t", "70", "--", "./coreapp"]