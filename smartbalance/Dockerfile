FROM golang:1.21.9-alpine3.19 AS builder


WORKDIR /application

COPY go.mod ./
RUN go mod download && go mod verify
RUN apk add bash 
RUN rm -rf /var/lib/apt/lists/*

COPY . ./

# RUN go build -v -o gin  -compiler gc ./main.go
# RUN cd ./rpcServer
# RUN go build -v -o rpc -compiler gc ./main.go

# ENTRYPOINT ["ls"]

# ENTRYPOINT  ["go", "run", "./main.go"]
# ENTRYPOINT ["./gin", "&&", "./rpc"]
# ENTRYPOINT  ["/bin/bash"]

# FROM alpine:3 AS release

# WORKDIR /application

# COPY --from=builder /application ./
# RUN mkdir /var/log/services; touch /var/log/services/boiler_room.log; touch /var/log/services/boiler_room_security.log
