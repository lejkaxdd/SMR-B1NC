#!/bin/bash

echo "Starting Gin server..."
./gin 
echo ""
echo "Starting gRPC server..."
go run ./rpcServer/main.go

