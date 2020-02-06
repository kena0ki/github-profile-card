FROM golang:1.13

WORKDIR /api
COPY . .

CMD cd api && go run main.go
