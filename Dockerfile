FROM golang:1.17.2-alpine3.13 AS build

WORKDIR /app

COPY . .

RUN go build -o /go-poc

FROM alpine:3.13.0

WORKDIR /

COPY --from=build /go-poc /go-poc

EXPOSE 9091

CMD [ "/go-poc" ]