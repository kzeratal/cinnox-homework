FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN go build -o /cinnox-homework

EXPOSE 8080

CMD [ "/cinnox-homework" ]