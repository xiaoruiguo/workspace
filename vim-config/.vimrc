" Last Change: 2012-09-14 16:59:38 星期五
if has('python3')
  silent! python3 1
endif
"pure"
source ~/.vim/config/quickfix.vimrc
source ~/.vim/config/timestamp.vimrc
source ~/.vim/config/cscope.vimrc

"common use"
source ~/.vim/config/markdown.vimrc
source ~/.vim/config/minibufexpl.vimrc
source ~/.vim/config/nerdtree.vimrc

"de.velop use"
source ~/.vim/config/a.vimrc
source ~/.vim/config/tagbar.vimrc
"source ~/.vim/config/neocomplcache.vimrc
"source ~/.vim/config/neosnippet.vimrc

""source ~/.vim/config/doxygen.vimrc
""source ~/.vim/config/slimv.vimrc
""source ~/.vim/config/vimwiki.vimrc
""source ~/.vim/config/code_complete.vimrc
""source ~/.vim/config/c.vimrc
source ~/.vim/config/golang.vimrc
source ~/.vim/config/godef.vimrc
source ~/.vim/config/ultisnips.vimrc
source ~/.vim/config/vim-template.vimrc
source ~/.vim/config/youcompleteme.vimrc

source ~/.vim/config/vim.vimrc

"website
source ~/.vim/config/jekyll.vimrc

"let $PYTHONCASEOK=""
"let g:UltiSnipsUsePythonVersion=3

call pathogen#infect()

"~~~~~~~~~~~~~~~~~~~ START 快捷键 ~~~~~~~~~~~~~~~~~~~~~~~"

nmap <F1> :call Usage()<cr>
nmap <F4> :call QuickfixToggle()<cr>
nmap <F6> :call IDEMode()<cr>
"LoadCtagsAndCscope在cscope.vimrc中实现
nmap <F10> :call LoadCtagsAndCscope()<cr>

"vim-go快捷键"
nmap <C-g>r :GoReferrers<CR><CR>

"cscope快捷键"
	"查找C语言符号"
nmap <C-\>s :cs find s <C-R>=expand("<cword>")<CR><CR>
	"查找定义
nmap <C-\>g :cs find g <C-R>=expand("<cword>")<CR><CR>
	"查找本函数调用的函数
nmap <C-\>d :cs find d <C-R>=expand("<cword>")<CR><CR>
	"查找调用本函数的函数
nmap <C-\>c :cs find c <C-R>=expand("<cword>")<CR><CR>
	"查找指定的字符串
nmap <C-\>t :cs find t <C-R>=expand("<cword>")<CR><CR>
	"查找egrep模式
nmap <C-\>e :cs find e <C-R>=expand("<cword>")<CR><CR>
	"查找并打开文件
nmap <C-\>f :cs find f <C-R>=expand("<cfile>")<CR><CR>
	"查找包含本文件的文件
nmap <C-\>i :cs find i ^<C-R>=expand("<cfile>")<CR>$<CR>

iab xtime <c-r>=strftime("%Y-%m-%d %H:%M:%S")

"显示自定义的快捷键等
let g:ShowUsage=0
function! Usage()
	if g:ShowUsage
		let g:ShowUsage=0
	else
		let g:ShowUsage=1
		echo "<F1> Help\n"
		echo "<F4> Quickfix Open/Close\n"
		echo "<F6> Tagbar  Open/Close\n"
		echo "<F10> Load Tags and Cscope"
		echo "[ <tab> 跳转到定义处"
	endif
endf

let g:IDEMode=0
function! IDEMode()
	if g:IDEMode
		exec ":NERDTreeClose"
		exec ":TagbarClose"
		let g:IDEMode=0
	else
		exec ":NERDTree"
		exec ":TagbarOpen"
		let g:IDEMode=1
	endif
endf

" ######### 括号、引号、中括号等自动匹配 ######### "
":inoremap ( ()<ESC>i
""":inoremap ) <c-r>=ClosePair(')')<CR>
":inoremap { {}<ESC>i
""":inoremap } <c-r>=ClosePair('}')<CR>
":inoremap [ []<ESC>i
""":inoremap ] <c-r>=ClosePair(']')<CR>
"":inoremap < <><ESC>i
"":inoremap > <c-r>=ClosePair('>')<CR>
":inoremap " ""<ESC>i
":inoremap ' ''<ESC>i
":inoremap ` ``<ESC>i

"~~~~~~~~~~~~~~~~~~~ END  快捷键 ~~~~~~~~~~~~~~~~~~~~~~~"
