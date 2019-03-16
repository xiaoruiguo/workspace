#! /bin/sh
#
# download.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

INSTALL_DIR="`pwd`/openresty"
VERSION="openresty-1.13.6.2"

if [ ! -e "$VERSION.tar.gz" ];then
	wget "https://openresty.org/download/${VERSION}.tar.gz"
fi

tar -xvf "${VERSION}.tar.gz"
cd $VERSION


if [[ `uname`=="Darwin" ]]; then
	echo "not support"
	exit 0
	brew install pcre openssl
	./configure \ 
	--with-pcre-jit --with-http_ssl_module --with-http_realip_module --with-http_stub_status_module --with-http_v2_module --prefix=${INSTALL_DIR} 
	--with-cc-opt="-I/usr/local/opt/openssl/include/ -I/usr/local/opt/pcre/include/" \
	--with-ld-opt="-L/usr/local/opt/openssl/lib/ -L/usr/local/opt/pcre/lib/" \
	-j8
else
	./configure \
	--with-pcre-jit --with-http_ssl_module --with-http_realip_module --with-http_stub_status_module --with-http_v2_module --prefix=${INSTALL_DIR} 
fi

make -j2
make install     

export PATH=${INSTALL_DIR}/bin:$PATH
