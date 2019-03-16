#!/bin/bash
rm -rf ~/.vim
rm -ff ~/.vimrc

ln -s `pwd`/.vim    ~/.vim
ln -s `pwd`/.vimrc  ~/.vimrc

#git submodule update --init --recursive

pushd  ~/.vim/
	bash ./install_bundle.sh
popd

pushd  ~/.vim/bundle/YouCompleteMe
	git submodule update --init --recursive
	./install.py --clang-completer --go-completer
popd

pushd ~/.vim/bundle/vimproc.vim
	make
popd
