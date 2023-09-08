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

absScriptDirPath="$(getAbsDirPath $0)"
absLauncherPath="$absScriptDirPath/sessionLauncher.go"
absConfigPath="$(getAbsDirPath $configPath)/$(basename $configPath)"

dir=$(grep -Po '"projectDir":.*?[^\\]",' $configPath | awk -F ':' '{print $2}' | cut -d "\"" -f 2)
absoluteProjectPath=$(eval echo $dir)
cd $absolutePath

go run $absLauncherPath $absConfigPath
