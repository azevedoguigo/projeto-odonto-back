FROM golang:1.22-alpine

WORKDIR /projeto-odonto-back

COPY . .

RUN go mod download && go mod verify

CMD [ "go", "run", "main.go" ]

EXPOSE 8080