#!/bin/bash

templates.render() {
    local source="${1}"
    local destination="${2}"
    local variables=("${@:3}")

    cp "${source}" "${destination}"

    for variable in ${variables[@]}
    do
        value="${!variable}"
        sed -i.bak 's=%{'"${variable}"'}='"${value}"'=g' "${destination}"
    done
}
