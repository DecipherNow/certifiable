#!/bin/bash

pushd docker

    docker build . -t deciphernow/certifiable
    docker push deciphernow/certifiable

popd
