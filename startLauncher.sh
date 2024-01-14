#!/bin/bash

configPath=$1

if [ -z $configPath ]
then
    echo "Usage: $(basename $0) <config file path>"
    return 1
fi

getAbsDirPath () {
    echo "$(realpath $1 | xargs dirname)"
}

launcherAbsDirPath="$(getAbsDirPath $0)"
launcherSrcAbsPath="$launcherAbsDirPath/sessionLauncher.go"
launcherBinAbsPath="$launcherAbsDirPath/tmux-session-launcher"

if [[ ! -f $launcherBinAbsPath ]] || [[ $launcherSrcAbsPath -nt $launcherBinAbsPath ]]
then
    echo "Binary file is not found or outdated. Starting compilation..."
    cd $launcherAbsDirPath
    go build
fi

configAbsPath="$(getAbsDirPath $configPath)/$(basename $configPath)"

dir=$(grep -Po '"projectDir":.*?[^\\]",' $configPath | awk -F ':' '{print $2}' | cut -d "\"" -f 2)
projectAbsPath=$(eval echo $dir)
cd $projectAbsPath

$launcherBinAbsPath $configAbsPath
