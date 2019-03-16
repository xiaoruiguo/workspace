" local syntax file - set colors on a per-machine basis:
" vim: tw=0 ts=4 sw=4
" Vim color file
" Maintainer:	lijiaocn
" Last Change:	2015-09-16 16:10:04

set background=dark
hi clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = "lijiaocn"
hi clear Visual
hi Visual		term=reverse cterm=reverse gui=reverse
""hi Pmenu        guibg=green gui=bold guifg=white
hi PmenuSel     ctermfg=green ctermbg=yellow 
hi PmenuThumb   ctermfg=white ctermbg=blue
hi PmenuSbar    ctermfg=blue ctermbg=white
