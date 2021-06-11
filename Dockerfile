FROM golang:1.16.5

WORKDIR /app


COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

EXPOSE 7777
