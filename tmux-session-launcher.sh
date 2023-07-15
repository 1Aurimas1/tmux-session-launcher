#!/bin/bash

if [ -z "$1" ]
  then
    echo "No argument supplied"
    return 1
fi

if [ -f "$1" ]; then
    source "$1"
    
    cd $PROJ_DIR
    git checkout $GIT_BRANCH
    tmux new-session -d -s $SESSION_NAME
    tmux send-keys -t $SESSION_NAME "nvim ." Enter
    tmux new-window
    tmux select-window -t $SESSION_NAME:2
    tmux send-keys -t $SESSION_NAME "cd backend" Enter
    # start server
    tmux split-window -h 
    tmux send-keys -t $SESSION_NAME "cd frontend" Enter
    # start server
    tmux new-window
    tmux send-keys -t $SESSION_NAME "cd backend" Enter
    tmux new-window
    tmux send-keys -t $SESSION_NAME "cd frontend" Enter
    tmux attach-session -t $SESSION_NAME
else
    echo ".env file not found."
fi

#cd ~/Projects/Other/dir_diff/dir_diff/
#git checkout web_gui
#tmux new-session -d -s dir_diff
#tmux send-keys -t dir_diff "nvim ." Enter
#tmux new-window
#tmux select-window -t dir_diff:2
#tmux send-keys -t dir_diff "cd backend" Enter
## start server
#tmux split-window -h 
#tmux send-keys -t dir_diff "cd frontend" Enter
## start server
#tmux new-window
#tmux send-keys -t dir_diff "cd backend" Enter
#tmux new-window
#tmux send-keys -t dir_diff "cd frontend" Enter
#tmux attach-session -t dir_diff
