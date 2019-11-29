#!/bin/bash

ls -l
glide update
rice embed-go
go build -o out/pkcalc
ls -l
#docker build --rm -f "docker/pkcalc.dockerfile" -t demo:latest .
