#!/bin/bash

templates.render() {
    local source="${1}"
    local destination="${2}"
    local variables=("${@:3}")

    cp "${source}" "${destination}"

    count=1
    for variable in ${variables[@]}
    do
        value="${!variable}"
        if [ "$value" != "" ]; then sed -i.bak '$aDNS.'"$count"' = '"${value}" "${destination}"; fi
        count=$((count+1))
    done
}
