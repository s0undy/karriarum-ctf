FROM golang:1.21.0

WORKDIR /api

COPY ./api .

RUN GOOS=linux GOARCH=amd64 go build -o api

CMD [ "./api" ]