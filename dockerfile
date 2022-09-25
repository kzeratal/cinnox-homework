FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY config.yml ./
RUN go mod download

COPY . ./
RUN go build -o /cinnox-homework

EXPOSE 8080

CMD [ "/cinnox-homework" ]