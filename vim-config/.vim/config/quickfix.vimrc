"######################################
"
"     quickfix功能的配置
"
"######################################

"可以将QuickfixToggle映射到快捷键,方便Quickfix窗口的打开/关闭
let g:QuickfixOpen=0
function! QuickfixToggle()
	if g:QuickfixOpen
		:cclose
		let g:QuickfixOpen=0
	else
		:botright copen
		let g:QuickfixOpen=1
	endif
endf
