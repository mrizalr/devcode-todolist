# # mysql

# FROM mysql:8.0.32

# ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes

# COPY ./initdb.sql /tmp

# CMD [ "mysqld", "--init-file=/tmp/initdb.sql" ]


# golang

FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go build -o todo-api

EXPOSE 3030

ENTRYPOINT ["/app/todo-api"]