# #base go image
# FROM golang:1.19-alpine AS builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# RUN chmod +x /app/brokerApp

# # building a tiny docker image
#! above lines are not needed when using Makefile
#todo Check the Makefile to know more about it 

FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

RUN apk add libcap && setcap 'cap_net_bind_service=+ep' /app/brokerApp

CMD [ "/app/brokerApp" ]
