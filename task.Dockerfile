# Stage 1
FROM golang:1.22-alpine as build

COPY task /task
COPY common /common

WORKDIR /task

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/task main.go

# Stage 2
FROM scratch

COPY --from=build /task/bin .

EXPOSE 3001

CMD ["/task"]

