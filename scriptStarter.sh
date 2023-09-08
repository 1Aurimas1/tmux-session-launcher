#!/bin/bash

configPath=$1

if [ -z $configPath ]
then
    echo "Usage: $(basename $0) <config file path>"
    return 1
fi

dir=$(grep -Po '"projectDir":.*?[^\\]",' $configPath | awk -F ':' '{print $2}' | cut -d "\"" -f 2)
absolutePath=$(eval echo $dir)
cd $absolutePath

go run "$(dirname $configPath)/sessionLauncher.go" $configPath
