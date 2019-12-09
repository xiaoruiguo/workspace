#!/bin/bash
# Usage:   iterm_login.sh  PORT  USER  PASSWORD 
# Example: iterm_login.sh  22   lijiao@192.168.88.2  123456

PATH=$PATH:/usr/local/bin

cat >/tmp/expect.sh << XXX
#!/usr/bin/expect -f
set timeout 30

spawn /usr/local/bin/zssh -p [lindex \$argv 0] [lindex \$argv 1]
expect {
        "(yes/no)?"
        {send "yes\n";exp_continue}
        "password:"
        {send "[lindex \$argv 2]\n"}
        "Select group:"
        {send "/192.168.197.3\n"}
}

#expect {
#        "Select account:"
#        {send "7\n"}
#}

#expect {
#        "Last login"
#        {interact}
#}

interact
XXX

chmod +x /tmp/expect.sh
/tmp/expect.sh $*