#!/bin/bash

#go build -o out/pkcalc
docker build --rm -f "docker/pkcalc.dockerfile" -t demo:latest .
