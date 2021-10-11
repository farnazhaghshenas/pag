FROM golang:1.16-alpine
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

RUN go build -o /pack-and-go

## Deploy

WORKDIR /

COPY --from=build /pack-and-go /pack-and-go

FROM alpine:3.9
RUN apk add --update tzdata
RUN apk add --update ca-certificates
RUN apk add --update bash

EXPOSE 8080

CMD [ "/pack-and-go" ]
