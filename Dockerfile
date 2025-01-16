FROM golang:1.24-rc-bookworm

WORKDIR /app

ENV ADMIN_PASSWORD="123"

COPY . .

RUN mkdir /app/out

EXPOSE 8080

CMD ["go", "run", "."]
