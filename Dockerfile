FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache poppler-utils

COPY . .
RUN chmod +x eden-inri

EXPOSE 50003
CMD ["./eden-inri"]