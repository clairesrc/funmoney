#!/usr/bin/env sh

daemon() {
    chsum1=""

    while [[ true ]]
    do
        chsum2=`find \( -path "./.git" -o -path "./db" \) ! -prune -o ${pwd} -type f -exec md5sum {} \;`
        if [[ $chsum1 != $chsum2 ]] ; then           
            if [ -n "$chsum1" ]; then
                ENV=dev ./build-and-run.sh
            fi
            chsum1=$chsum2
        fi
        sleep 2
    done
}
./build-and-run.sh && daemon