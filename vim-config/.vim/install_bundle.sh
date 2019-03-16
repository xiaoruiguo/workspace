#!/bin/bash
mkdir bundle
pushd bundle
git clone --depth 1 --no-tags https://github.com/vim-scripts/a.vim.git
git clone --depth 1 --no-tags https://github.com/vim-scripts/bufexplorer.zip.git
git clone --depth 1 --no-tags https://github.com/vim-scripts/DrawIt.git
git clone --depth 1 --no-tags https://github.com/vim-scripts/Mark--Karkat.git
git clone --depth 1 --no-tags https://github.com/vim-scripts/minibufexpl.vim.git
git clone --depth 1 --no-tags https://github.com/scrooloose/nerdtree.git
git clone --depth 1 --no-tags https://github.com/vim-scripts/Tagbar.git
git clone --depth 1 --no-tags https://github.com/fatih/vim-go.git
git clone --depth 1 --no-tags https://github.com/plasticboy/vim-markdown.git
git clone --depth 1 --no-tags https://github.com/honza/vim-snippets.git
git clone --depth 1 --no-tags https://github.com/aperezdc/vim-template.git
git clone --depth 1 --no-tags https://github.com/Valloric/YouCompleteMe.git
git clone --depth 1 --no-tags https://github.com/SirVer/ultisnips.git
git clone --depth 1 --no-tags https://github.com/Shougo/vimproc.vim.git
git clone --depth 1 --no-tags https://github.com/Shougo/vimshell.vim.git
git clone --depth 1 --no-tags https://github.com/leafgarland/typescript-vim.git
git clone --depth 1 --no-tags https://github.com/parkr/vim-jekyll.git
#git clone --depth 1 --no-tags https://github.com/vim-scripts/conque-repl.vim.git
#git clone --depth 1 --no-tags https://github.com/vim-scripts/neocomplcache.git
#git clone --depth 1 --no-tags https://github.com/vim-scripts/neosnippet.git
#git clone --depth 1 --no-tags https://github.com/vim-scripts/Conque-Shell.git
popd
