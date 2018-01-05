#!/bin/bash

set -e

export USER_NAME="${1}"
export COMMON_NAME="${2}"
export PASSWORD="${3}"
export USER_DIRECTORY="/usr/local/var/certifiable/users/${USER_NAME}"
export INTERMEDIATE_DIRECTORY="/usr/local/var/certifiable/intermediate"

mkdir -p "${USER_DIRECTORY}"

pushd "${USER_DIRECTORY}" > /dev/null

    mkdir -p files crl csr private newcerts
    chmod 700 files

    echo -n "${PASSWORD}" > "files/${USER_NAME}.password"
    openssl genrsa -aes256 -out "files/${USER_NAME}.key" -passout "file:files/${USER_NAME}.password" 2048 &> /dev/null

    openssl req \
        -new \
        -key "files/${USER_NAME}.key" \
        -out "${INTERMEDIATE_DIRECTORY}/csr/${USER_NAME}.csr" \
        -passin "file:files/${USER_NAME}.password" \
        -subj "/C=US/ST=Virginia/L=Alexandria/O=Decipher Technology Studios/OU=Engineering/CN=${COMMON_NAME}"

    pushd "${INTERMEDIATE_DIRECTORY}" > /dev/null

        openssl ca \
            -batch \
            -config intermediate.conf \
            -extensions usr_cert \
            -days 365 \
            -notext \
            -md sha256 \
            -in "csr/${USER_NAME}.csr" \
            -out "${USER_DIRECTORY}/files/${USER_NAME}.crt" \
            -passin file:private/intermediate.password

    popd > /dev/null

    cat "files/${USER_NAME}.crt" "${INTERMEDIATE_DIRECTORY}/certs/intermediate.chain.crt" > "files/${USER_NAME}.chain.crt"

    cp "files/${USER_NAME}.password" "files/${USER_NAME}.password.tmp"

    openssl pkcs12 \
        -export \
        -clcerts \
        -in "files/${USER_NAME}.chain.crt" \
        -inkey "files/${USER_NAME}.key" \
        -out "files/${USER_NAME}.p12" \
        -name "${COMMON_NAME}" \
        -passin "file:files/${USER_NAME}.password" \
        -passout "file:files/${USER_NAME}.password.tmp"

    rm "files/${USER_NAME}.password.tmp"

popd > /dev/null
