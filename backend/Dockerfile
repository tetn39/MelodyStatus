FROM golang:1.23

WORKDIR /go/src

COPY go* ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]