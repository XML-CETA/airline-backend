FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o airline

FROM alpine
COPY .env .
COPY --from=build_container /app/airline /usr/bin
EXPOSE 3000
ENTRYPOINT ["airline"]

