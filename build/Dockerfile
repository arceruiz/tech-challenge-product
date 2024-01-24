#mauricio1998/tech-challenge-product

FROM golang as builder

RUN mkdir app
COPY ./ app
WORKDIR app
RUN CGO_ENABLED=0 go build -o dist/tech-challenge-product cmd/client/main.go

FROM alpine as runner

RUN mkdir app
COPY ./dist/tech-challenge-product app/
RUN chmod +x app

COPY ./internal/config/config.yaml /app
EXPOSE 3001
WORKDIR /app

ENTRYPOINT [ "./tech-challenge-product", "--config-dir", "." ]