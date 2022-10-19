#!/bin/sh

docker build -t mustache . && docker run -p 9191:8080 mustache