# https://github.com/cosmtrek/air/blob/master/Dockerfile
FROM golang:1.22.2-alpine

ENV TZ /usr/share/zoneinfo/Asia/Tokyo
ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

EXPOSE 8080

# Install libraries for testing
RUN apk add --no-cache gcc musl-dev

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]