#!/bin/bash

. ../../../../LinuxShell/library.sh

dev="enp0s8"
cmdAdd="tc qdisc add dev $dev root handle 1:0 sfq divisor 1024 limit 1024 depth 127 perturb 300 quantum 3000 flows 127 headdrop"
cmdDel="tc qdisc delete dev $dev root handle 1:0 sfq" 

echo "Input Any Key to Add sfq on $dev:"
func_green_str "    $cmdAdd"
read

echo "Before:"
func_red_cmd "tc qdisc show dev $dev"

`$cmdAdd`

echo ""
echo "After:"
func_red_cmd "tc qdisc show dev $dev"

echo ""
echo "Input Any Key to Delete sfq on $dev"
func_green_str "    $cmdDel"
read

`$cmdDel`

echo "Deleted:"
func_red_cmd "tc qdisc show dev $dev"
