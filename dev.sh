#!/usr/bin/env sh

tmux start-server

PWD=`pwd`
WATCH_BUILD_FLENAME="watch-build.sh"
TAIL_LOGS_FILENAME="tail-logs.sh"


# create a session with five panes
# tmux new-session -d -s FunMoneyDevSession -n Shell1 -d "/usr/bin/env sh -c \"echo 'first shell'\"; /usr/bin/env sh -i"
tmux new-session -d -s FunMoneyDevSession -n FunMoneyShell -d "/usr/bin/env sh -c \"$PWD/$WATCH_BUILD_FLENAME\""
tmux split-window -t FunMoneyDevSession:0 "/usr/bin/env sh -c \"$PWD/$TAIL_LOGS_FILENAME\""
tmux split-window -t FunMoneyDevSession:0 "/usr/bin/bash"


# change layout to tiled
tmux select-layout -t FunMoneyDevSession:0 tiled

tmux attach -tFunMoneyDevSession