#!/bin/bash
rm -rf ~/.vim
rm -ff ~/.vimrc

ln -s `pwd`/.vim    ~/.vim
ln -s `pwd`/.vimrc  ~/.vimrc

pushd  ~/.vim/
	bash ./install_bundle.sh
popd

pushd  ~/.vim/bundle/YouCompleteMe
#	git submodule update --init --recursive
	export LDFLAGS="-L/usr/local/opt/openssl/lib"
	export CPPFLAGS="-I/usr/local/opt/openssl/include"
	export PKG_CONFIG_PATH="/usr/local/opt/openssl/lib/pkgconfig"
	./install.py --clang-completer --go-completer
popd

pushd ~/.vim/bundle/vimproc.vim
	make
popd
