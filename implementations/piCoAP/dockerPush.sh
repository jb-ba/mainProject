#!/usr/bin/env bash
export VERSION="v0.0.9arm"
sudo docker build . -t jbba/picoap:$VERSION
sudo docker push jbba/picoap:$VERSION
