#!/bin/bash

pushd docker

    docker build . -t deciphernow/certifiable:di2e
    docker push deciphernow/certifiable:di2e

popd
