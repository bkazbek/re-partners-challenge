FROM golang:1.21.0 AS builder
RUN mkdir app
RUN chmod 777 -R ./app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o main

##DEPLOY##
FROM alpine:3.15 AS final
RUN mkdir app
WORKDIR /app
COPY --from=builder /app /app
RUN chmod +x ./main
ENTRYPOINT ["./main"]
EXPOSE 3000