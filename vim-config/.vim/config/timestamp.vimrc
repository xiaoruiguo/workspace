"###################################
"
"     自动填写文件创建和修改时间 
"
"###################################

"特别注意这里的正则表达式,'\'需要使用'\'进行转义,不同于直接在命令模式执行!
"这个问题折腾了近乎一晚上才发现!!
function! TimestampUpdate()
	exec ":% s#修改时间: \\d\\{4}\/\\d\\{2}\/\\d\\{2} \\d\\{1,2}:\\d\\{2}:\\d\\{2}#\修改时间: ".strftime('%Y\/%m\/%d %H:%M:%S')
endfun

function! GoTimestampUpdate()
	:silent! exec ":1,1 s#[cC]hange: \\d\\{4}\/\\d\\{2}\/\\d\\{2} \\d\\{1,2}:\\d\\{2}:\\d\\{2}#\Change: ".strftime('%Y\/%m\/%d %H:%M:%S')
endfun

function! MdTimestampUpdate()
	:silent! exec ":1,6 s#changedate: \"\\d\\{4}[\/-]\\d\\{2}[\/-]\\d\\{2}T\\d\\{1,2}:\\d\\{2}:\\d\\{2}\+0800#changedate: \"".strftime('%Y-%m-%dT%H:%M:%S+0800')
	:silent! exec ":1,$ s#http://127.0.0.1:4000#https://www.lijiaocn.com#"
endfun

autocmd BufWritePre,FileWritePre *.wiki  ks|call TimestampUpdate()|'s
autocmd BufWritePre,FileWritePre *.md  ks|call MdTimestampUpdate()|'s
autocmd BufWritePre,FileWritePre *.go  ks|call GoTimestampUpdate()|'s
