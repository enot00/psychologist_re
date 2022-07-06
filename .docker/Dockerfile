FROM golang:1.18-alpine

WORKDIR /app

COPY ../go.mod /app
COPY ../go.sum /app

RUN go mod download

COPY .. /app

WORKDIR /app/cmd/server

RUN go build -o mentalhealthapp

EXPOSE 8081

CMD [ "/app/cmd/server/mentalhealthapp" ]
