#!/usr/bin/env sh

daemon() {
    chsum1=""

    while [[ true ]]
    do
        chsum2=`find \( -path "./.git" -o -path "./db" \) ! -prune -o ${pwd} -type f -exec md5sum {} \;`
        if [[ $chsum1 != $chsum2 ]] ; then           
            if [ -n "$chsum1" ]; then
                ENV=dev ./build-and-run.sh
                # @TODO: live-reload: set env="DEV" env var here and have `build-and-run.sh` pass it down to `docker-compose up`. when that var is set to DEV in frontend container, the build script in frontend then appends some javascript that polls for changes in /app endpoint in the backend container. the baclend endpoint just needs an extra property added such as timestamp of server startup
            fi
            chsum1=$chsum2
        fi
        sleep 2
    done
}
./build-and-run.sh && daemon