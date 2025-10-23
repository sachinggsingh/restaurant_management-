FROM  golang:1.25

WORKDIR /app

COPY . .

RUN go build -o dockergin 

EXPOSE 8080

ENTRYPOINT [ "./dockergin" ]