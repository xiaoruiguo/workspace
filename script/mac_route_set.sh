#! /bin/sh
#
# mac_route_set.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

action=$1

en0ip=`netstat -nr |grep en0 |grep default |awk '{ print $2}'`
en0ip=172.16.104.140
echo $en0ip

route -n $action -net 10.10.0.0  -netmask 255.255.0.0  $en0ip
route -n $action -net 10.19.0.0  -netmask 255.255.0.0  $en0ip
route -n $action -net 172.0.0.0  -netmask 255.0.0.0    $en0ip
route -n $action -net 172.16.0.0 -netmask 255.255.0.0  $en0ip
