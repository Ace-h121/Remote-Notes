FROM golang:1.23rc1-bookworm

WORKDIR /app

RUN mkdir /RecieveFile /SendFile

COPY RecieveFile/ RecieveFile

COPY SendFile/ SendFile

COPY go.mod go.sum ./

RUN go mod download

RUN mkdir /app/notes

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]