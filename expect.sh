#!/usr/bin/expect -f

set timeout 30
   
spawn zssh -p [lindex $argv 0] [lindex $argv 1]
expect {
        "(yes/no)?"
        {send "yes\n";exp_continue}
        "password:"
        {send "[lindex $argv 2]\n"}
        "Select group:"
        {send "/172.31.160.93\n"}
        "Select account:"
        {interact}
#        "Last login"
#        {interact}
}

expect {
        "Select account:"
        {send "7\n"}
}
expect {
        "Last login"
        {interact}
}
interact
