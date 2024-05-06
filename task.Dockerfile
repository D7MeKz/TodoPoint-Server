# Stage 1
FROM golang:1.22-alpine as build

COPY task222 /task
COPY ha /common

WORKDIR /task

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/taskA main.go

# Stage 2
FROM scratch

COPY --from=build /task222/bin .

EXPOSE 3001

CMD ["/task"]

