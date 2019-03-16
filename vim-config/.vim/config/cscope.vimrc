"cscope 配置
":help cscope-suggestions
"
"####################################
"
"      自动加载cscope.out
"
"####################################

"可以映射到快捷键，自动寻找ctags和cscope.out并加载"
function! LoadCtagsAndCscope()
	let max = 10     "最大向上查找层数
	let dir = './'   "当前路径
	let i = 0
	let break = 0    "是否找到文件
	while isdirectory(dir) && i < max
		if filereadable(dir.'cscope.out')
			execute 'cs add ' . dir . 'cscope.out'
			:echomsg dir.'cscope.out'
			let break = 1
		endif
"		if filereadable(dir.'tags')
"			:echomsg dir.'tags'
"			execute 'set tags=' . dir . 'tags'
"			let break = 1
"		endif
		if break == 1
			execute 'lcd' . dir
			break
		endif
		let dir = dir . '../'
		let i = i + 1
	endwhile
	if i == max
		:echomsg "Not Found: cscope.out"
	endif
endf
