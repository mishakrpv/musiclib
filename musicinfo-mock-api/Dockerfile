FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main

EXPOSE 3256

ENTRYPOINT ["/app/main"]

CMD ["main"]