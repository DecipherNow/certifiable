#!/bin/bash

set -e

source /usr/local/lib/templates.sh

export SERVER_NAME="${1}"
export SUBJECT_ALT_NAME="${SERVER_NAME}"
export COMMON_NAME="${2}"
export PASSWORD="${3}"
export SUBJECT_ALT_NAME_1="${4}"
export SERVER_DIRECTORY="/usr/local/var/certifiable/servers/${COMMON_NAME}"
export INTERMEDIATE_DIRECTORY="/usr/local/var/certifiable/intermediate"

mkdir -p "${SERVER_DIRECTORY}"

pushd "${SERVER_DIRECTORY}" > /dev/null

    templates.render /usr/local/var/certifiable/templates/server.conf server.conf SUBJECT_ALT_NAME SUBJECT_ALT_NAME_1

    mkdir -p files crl csr newcerts
    chmod 700 files

    echo -n "${PASSWORD}" > "files/${COMMON_NAME}.password"
    openssl genrsa -aes256 -out "files/${COMMON_NAME}.key" -passout "file:files/${COMMON_NAME}.password" 2048 &> /dev/null
    openssl rsa -in "files/${COMMON_NAME}.key" -out "files/${COMMON_NAME}.nopass.key" -passin "file:files/${COMMON_NAME}.password"


    openssl req \
        -new \
        -config server.conf \
        -key "files/${COMMON_NAME}.key" \
        -out "${INTERMEDIATE_DIRECTORY}/csr/${COMMON_NAME}.csr" \
        -passin "file:files/${COMMON_NAME}.password" \
        -subj "/C=US/ST=Virginia/L=Alexandria/O=Decipher Technology Studios/OU=Engineering/CN=${COMMON_NAME}"

    pushd "${INTERMEDIATE_DIRECTORY}" > /dev/null

        openssl ca \
            -batch \
            -config intermediate.conf \
            -extensions server_cert \
            -days 3650 \
            -notext \
            -md sha256 \
            -in "csr/${COMMON_NAME}.csr" \
            -out "${SERVER_DIRECTORY}/files/${COMMON_NAME}.crt" \
            -passin file:private/intermediate.password

    popd > /dev/null

    cat "files/${COMMON_NAME}.crt" "${INTERMEDIATE_DIRECTORY}/certs/intermediate.chain.crt" > "files/${COMMON_NAME}.chain.crt"

    cp "files/${COMMON_NAME}.password" "files/${COMMON_NAME}.password.tmp"

    openssl pkcs12 \
        -export \
        -in "files/${COMMON_NAME}.chain.crt" \
        -inkey "files/${COMMON_NAME}.key" \
        -out "files/${COMMON_NAME}.p12" \
        -name "${COMMON_NAME}" \
        -passin "file:files/${COMMON_NAME}.password" \
        -passout "file:files/${COMMON_NAME}.password.tmp"

    rm "files/${COMMON_NAME}.password.tmp"

    pushd files > /dev/null

      tar -cvf ${COMMON_NAME}.tar *

    popd > /dev/null


#    keytool -importcert \
#        -noprompt \
#        -trustcacerts \
#        -alias intermediate \
#        -file /usr/local/var/certifiable/intermediate/certs/intermediate.chain.crt \
#        -keystore /usr/local/var/certifiable/servers/${SERVER_NAME}/files/truststore.jks \
#        -storepass file:/usr/local/var/certifiable/servers/${SERVER_NAME}/files/${SERVER_NAME}.password
#
#    keytool -importkeystore \
#        -noprompt \
#        -alias ${COMMON_NAME} \
#        -srckeystore /usr/local/var/certifiable/servers/${SERVER_NAME}/files/${SERVER_NAME}.p12 \
#        -srcstoretype PKCS12 \
#        -srckeypass file:/usr/local/var/certifiable/servers/${SERVER_NAME}/files/${SERVER_NAME}.password \
#        -srcstorepass file:/usr/local/var/certifiable/servers/${SERVER_NAME}/files/${SERVER_NAME}.password \
#        -destkeystore /usr/local/var/certifiable/servers/${SERVER_NAME}/files/${SERVER_NAME}.jks \
#        -destkeypass file:/usr/local/var/certifiable/servers/${SERVER_NAME}/files/${SERVER_NAME}.password \
#        -deststorepass file:/usr/local/var/certifiable/servers/${SERVER_NAME}/files/${SERVER_NAME}.password

popd > /dev/null
