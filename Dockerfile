FROM alpine:latest
WORKDIR /app
COPY . .
RUN chmod +x eden-inri
EXPOSE 50003
CMD ["./eden-inri"]