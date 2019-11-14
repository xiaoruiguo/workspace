#! /bin/sh
#
# loop.sh
# Copyright (C) 2019 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#

function help(){
	echo "usage: $0 [-f file] [-n name] [arguments...]"
}

while :; do
	case $1 in
		-h|-\?|--help)
			help
			exit
			;;
		-f|--file)
			if [[ $2 == "" || ${2:0:1} == "-" ]];then
				echo 'ERROR: "--file" requires a non-empty option argument.' 2>&1
				exit 1
			fi
			file=$2
			shift
			;;
		-n|--name)
			if [[ $2 == "" || ${2:0:1} == "-" ]];then
				echo 'ERROR: "--name" requires a non-empty option argument.' 2>&1
				exit 1
			fi
			name=$2
			shift
			;;
		-|--)
			shift
			break
			;;
		*)
			break
			;;
	esac
	shift
done

echo "file is: $file"
echo "name is: $name"
echo "arguments: $*"
