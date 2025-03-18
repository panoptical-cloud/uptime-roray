#!/bin/bash

echo " "
echo "Generating proto files..."
echo " "

echo "protoc --go_out=. panmon_agent.proto"
protoc --go_out=. panmon_agent.proto
echo " "

cp panmon_agent.pb.go ../remote_agent/api
cp panmon_agent.pb.go ../gobff/api
