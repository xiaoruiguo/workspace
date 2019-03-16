" Some Linux distributions set filetype in /etc/vimrc.
" Clear filetype flags before changing runtimepath to force Vim to reload them.
filetype off
filetype plugin indent off
set runtimepath+=$GOROOT/misc/vim
set path+=$GOPATH/src
filetype plugin indent on
syntax on
let g:go_version_warning = 0
let g:go_template_autocreate = 0
