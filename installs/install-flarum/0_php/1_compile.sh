#!/bin/sh
sudo yum install -y gcc make libxml2-devel
pushd php-7.2.0
mkdir /opt/php-7.2.0
./configure --prefix=/opt/php-7.2.0 --with-openssl --with-zlib --enable-mbstring --with-pdo-mysql --with-gd --with-curl
make install
echo 'export PATH=$PATH:/opt/php-7.2.0/bin' >>~/.bashrc
echo 'export PATH=$PATH:/opt/php-7.2.0/bin' >>/etc/profile
source ~/.bashrc
