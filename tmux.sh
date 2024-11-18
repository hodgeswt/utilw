#!/bin/zsh

BASE_FOLDER="./"

get_base_index() {
    if ! tmux has-session 2>/dev/null; then
        tmux new-session -d -s temp_session
        BASE_INDEX=$(tmux show-options -g base-index | awk '{print $2}')
        tmux kill-session -t temp_session
    else
        BASE_INDEX=$(tmux show-options -g base-index | awk '{print $2}')
    fi
    echo $BASE_INDEX
}

BASE_INDEX=$(get_base_index)


tmux has-session -t "utilw" 2>/dev/null
if [ $? != 0 ]; then
    echo "Creating utilw session..."


    if [ "$BASE_INDEX" -eq 0]; then
        S_NVIM = 0
        S_SHELL = 1
    else
        S_NVIM = 1
        S_SHELL = 2
    fi

    tmux new-session -d -s "utilw" -c "$BASE_FOLDER" -n "nvim"
    tmux send-keys -t utilw:$S_NVIM "nix-shell" C-m

    tmux new-window -t utilw:$S_SHELL -c "$BASE_FOLDER" -n "nix-shell"
    tmux send-keys -t utilw:$S_SHELL "nix-shell" C-m
fi

echo "Attaching utilw session..."
tmux attach-session -t "utilw"
