FROM golang:alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /usr/bin/airline
EXPOSE 3000
ENTRYPOINT ["airline"]

