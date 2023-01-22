FROM golang:1.19
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w' -o cluster1

RUN touch /tmp/goerr
#COPY my_server.key /app
#COPY my_server.pem /app
RUN CGO_ENABLED=0 GOOS=linux go build -o cluster1
CMD ["/app/cluster1 server1"]

