FROM alpine:latest

RUN mkdir /app

COPY serviceApp /app

CMD [ "/app/serviceApp"]