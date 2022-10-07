#!/usr/bin/env sh

daemon() {
    chsum1=""

    while [[ true ]]
    do
        chsum2=`find ${pwd} ! -path '.git' ! -name '.gitignore' -type f -exec md5sum {} \;`
        if [[ $chsum1 != $chsum2 ]] ; then           
            if [ -n "$chsum1" ]; then
                ./build-and-run.sh
            fi
            chsum1=$chsum2
        fi
        sleep 2
    done
}
./build-and-run.sh && daemon
