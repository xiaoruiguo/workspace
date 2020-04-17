#! /bin/sh
#
# relink.sh
# Copyright (C) 2020 lijiaocn <lijiaocn@foxmail.com wechat:lijiaocn>
#
# Distributed under terms of the GPL license.
#

# $1: link file
function delete_link() {
	if [[ ! -L $1 ]];then 
		echo "$1 is not symbol link"
		exit
	fi
	rm $1
}

if [[ $# != 1 ]];then
	echo "usage: file"
	exit
fi

file=`echo $1 |sed -e "s/.id_rsa//" |sed -e "s/.pub//"`
echo $file

delete_link id_rsa
delete_link id_rsa.pub

ln -s $file.id_rsa  id_rsa 
ln -s $file.id_rsa.pub  id_rsa.pub
