FROM alpine:latest
WORKDIR /app
COPY . .
RUN chmod +x eden-inri
EXPOSE 9092
CMD ["./eden-inri"]