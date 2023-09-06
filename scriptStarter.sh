#!/bin/bash

configPath=$1

if [ -z "$configPath" ]
then
    echo "Provide config file as an argument"
    return 1
fi

dir=$(grep -Po '"projectDir":.*?[^\\]",' $configPath | awk -F ':' '{print $2}' | cut -d "\"" -f 2)
absolutePath=$(eval echo $dir)
cd $absolutePath

go run sessionLauncher.go $configPath
