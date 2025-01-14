FROM golang:alpine3.21

COPY . /app

WORKDIR /app/cmd/quiz-app

RUN go build -v

ENTRYPOINT ["/app/cmd/quiz-app/quiz-app"]