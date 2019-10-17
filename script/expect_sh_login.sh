#!/bin/bash
# Usage:   expect_sh_login.sh  PORT  USER                PASSWORD  JUMPADDR
# Example: expect_sh_login.sh  22   lijiao@192.168.88.2  123456    10.10.1.2

PATH=$PATH:/usr/local/bin

# 解决使用 expect 时，无法用 lrzsz 的问题（在有的机器上好使。。）
export LC_CTYPE=en_US   

cat >/tmp/expect.sh << EOF
#!/usr/bin/expect
set timeout 30

spawn ssh -p [lindex \$argv 0] [lindex \$argv 1]
expect {
        "(yes/no)?"
        {send "yes\n";exp_continue}
        "password:"
        {send "[lindex \$argv 2]\n";exp_continue}
        "Select group: "
        {send "/[lindex \$argv 3]\n";exp_continue}
        "]"
        {interact}
}
interact
EOF

expect /tmp/expect.sh $*
