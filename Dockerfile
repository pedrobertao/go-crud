# stage de build
FROM golang:1.20 AS build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o go-crud main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/go-crud ./

EXPOSE 3030

CMD [ "./go-crud" ]