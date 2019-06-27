#!/usr/bin/env bash


docker run -it --rm --name envoy -p 50051:50051 \
  quay.io/kubernetes-ingress-controller/grpc-fortune-teller:0.1

docker run -it --rm --name envoy -p 51051:51051 \                                                                                                                                                      [19/06/11|11:42AM]
  -v "$(pwd)/gen/descriptors/example.proto-descriptor:/data/example.proto-descriptor:ro" \
  -v "$(pwd)/envoy-config.yml:/etc/envoy/envoy.yaml:ro" \
  envoyproxy/envoy