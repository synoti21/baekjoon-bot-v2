FROM golang:1.21.10

WORKDIR /app

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o main .

CMD ["./main"]