#!/bin/bash
# Generator Docker image Dummy HTTP Server
# By Tedezed

list_ports="80 443 8000 8443 8080 8888 3000 5000 6000"

for p in $(echo ${list_ports})
do
    echo "[INFO] Docker port: $p"
    docker build --build-arg="DUMMY_PORT=${p}" -t tedezed/dummy-http-server:$p .
    docker push tedezed/dummy-http-server:$p
done
