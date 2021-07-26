FROM golang:1.16-buster

RUN go version

ENV GOPATH=/ 

COPY ./ ./

RUN go mod download 
RUN go build -o tg-weather-bot.exe ./cmd/main.go

CMD ["./tg-weather-bot.exe"]