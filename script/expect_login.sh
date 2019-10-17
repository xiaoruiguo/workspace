#!/usr/bin/expect

# Usage:   expect_login.sh  PORT  USER                PASSWORD  JUMPADDR
# Example: expect_login.sh  22   lijiao@192.168.88.2  123456    10.10.1.2

set timeout 30
spawn ssh -p [lindex $argv 0] [lindex $argv 1]
expect {
        "(yes/no)?"
        {send "yes\n";exp_continue}
        "password:"
        {send "[lindex $argv 2]\n";exp_continue}
        "Select group: "
        {send "/[lindex $argv 3]\n";exp_continue}
        "]"
        {interact}
}
