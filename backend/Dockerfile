FROM golang:1.22

LABEL version=1.0.0
WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o bin/pizzapizza ./main.go

EXPOSE 3000

CMD ["./bin/pizzapizza"]
