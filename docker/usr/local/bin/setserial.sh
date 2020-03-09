#!/bin/bash

set -e

export INTERMEDIATE_DIRECTORY="/usr/local/var/certifiable/intermediate"

date +%s | od -A n -t x1 | sed 's/ *//g' | cut -c 5-20 > ${INTERMEDIATE_DIRECTORY}/serial