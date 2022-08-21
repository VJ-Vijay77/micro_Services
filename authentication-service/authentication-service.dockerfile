FROM alpine:latest

RUN mkdir /app

COPY authApp /app

RUN apk add libcap && setcap 'cap_net_bind_service=+ep' /app/authApp

CMD [ "/app/authApp" ]
