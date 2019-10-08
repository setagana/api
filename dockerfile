FROM golang:1.12 
WORKDIR /app
COPY . ./
RUN go get -d -v . && go build -o main .
EXPOSE 8080
CMD ["/app/main"]