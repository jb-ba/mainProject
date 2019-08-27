#!/usr/bin/env bash
export VERSION="v0.0.6"
sudo docker build . -t jbba/syncserver:$VERSION
sudo docker push jbba/syncserver:$VERSION
