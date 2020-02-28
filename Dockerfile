FROM golang:alpine

LABEL application='bts'

LABEL maintainer='Anyama Richard'

WORKDIR /app

ENV GOPATH=/app

# update system
RUN apk update && apk add --no-cache bash git openssh

# install dependencies
RUN go get github.com/dgrijalva/jwt-go
RUN	go get github.com/gorilla/context
RUN	go get github.com/gorilla/mux
RUN	go get github.com/jinzhu/gorm
RUN	go get github.com/joho/godotenv
RUN	go get github.com/lib/pq
RUN	go get golang.org/x/crypto/bcrypt

#copy source code to app dir
COPY . .

# build app
RUN go install main

EXPOSE 9000

CMD ["./bin/main"]
