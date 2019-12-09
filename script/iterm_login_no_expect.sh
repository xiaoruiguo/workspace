#!/bin/bash
# Usage:   iterm_login.sh  PORT  USER  PASSWORD 
# Example: iterm_login.sh  22   lijiao@192.168.88.2  123456

PATH=$PATH:/usr/local/bin

ssh -p [lindex \$argv 0] [lindex \$argv 1]
