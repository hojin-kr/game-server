# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app
COPY . ./
COPY go.mod ./
COPY go.sum ./
COPY cmd ./
RUN go build -o /app/haru

# ENV PORT=8080
# EXPOSE 80

CMD [ "/app/haru" ]
