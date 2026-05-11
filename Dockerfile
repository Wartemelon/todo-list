FROM golang:1.25.3 AS build
WORKDIR /build 
COPY go.mod go.sum ./
RUN go mod download 
COPY ./ ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o todo_list

FROM ubuntu:latest
WORKDIR /app 
COPY --from=build /build/todo_list ./
COPY --from=build /build/web ./web
RUN mkdir -p /database
EXPOSE 7540 
ENV TODO_PORT=7540 \
    TODO_DBFILE=/database/scheduler.db \
    TODO_PASSWORD=""
CMD ["./todo_list"]

