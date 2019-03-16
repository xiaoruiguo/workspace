"#######################################
"
"           对vim本身的配置
"
"#######################################

"########  鼠标  ########
""set mouse=v
set backspace=indent,eol,start

"########  兼容性  ########
"为了使用vim的一些特性，设置为不兼容vi
set nocp

"########   编码   ########
"use the same encoding
if has("win32") || has("win32unix") || has("win95")
	"vim内部使用编码
	set encoding=utf-8
	"菜单语言编码
	set langmenu=en_US.utf-8
	"消息语言编码,与vim内部编码一致,不然提示信息乱码
	language message en_US.utf-8
	"终端使用的编码
	set termencoding=en_US.utf-8
else
	"vim内部使用编码
	set encoding=utf-8
	"菜单语言编码
	set langmenu=en_US.utf-8
	"消息语言编码,与vim内部编码一致,不然提示信息乱码
	language message en_US.utf-8
	"终端使用的编码
	"set termencoding=en_US.utf-8
	set termencoding=utf-8
endif

"########  文件   #########
"默认文件编码,新建一个文件时采用的编码
set fileencoding=utf-8

"打开一个文件时,按照给出的顺序探测文件编码
set fileencodings=utf-8,gbk,cp936,ucs-bom,big5,euc-jp,euc-kr,latin1,gb18030,default

"开启文件类型检测
filetype plugin indent on

"####### 个性化  ##########
"设置tab占用4个空格
set tabstop=4

"将tab键转换成空格
""set expandtab

"when delete a tab replaced by 4 banks,just put backspace one time
set smarttab

"开启智能缩进,类似于C语言的缩进
set smartindent

"open c/c++ auto indent
"in visual mode,choose texts ,then push =,the text will be auto indent
set cindent

"开启自动缩进,下一行使用与上一行同样的缩进
set autoindent

"auto indentation width
set shiftwidth=4

" 显示括号配对情况
set showmatch 

"设置一行的最大字符数,超过时将被分成两行
"set textwidth=80

"显示行号
set number

"when use Tab to autocomplete, show in a single line 
set wildmenu

"show the curser's positon ,at which line and which column
set ruler

"开启语法高亮
syntax on

"colorscheme
":colors murphy
":colors default
":colors ron
if has("win32") || has("win32unix") || has("win95")
:colors darkblue
else
:colors lijiaocn
endif


"highlight the search keyword
set hlsearch

"real-time search,搜索时实时显示匹配字符
set incsearch 

"开启状态栏信息
set laststatus=2                          

"命令行的高度，默认为1，这里设为2
set cmdheight=2

"close the preview windows of completeopt
set completeopt=longest,menu

" 状态行显示的内容 [包括系统平台、文件类型、坐标、所占比例、时间等]
set statusline=%F%m%r%h%w\ [FORMAT=%{&ff}]\ [TYPE=%Y]\ [POS=%l,%v][%p%%]\ %y%r%m%*%=\ %{strftime(\"%d/%m/%y\ -\ %H:%M\")}

" 突出显示当前行
set cursorline              

"显示Tab符，使用一高亮竖线代替
set list
set listchars=tab:\|\ ,

"设定文件浏览器目录为当前目录
"set autochdir 

"写入文件之前备份一份,写入成功后删除备份
set writebackup

"不保存备份文件"
set nobackup

"每行超过80个的字符用下划线标示
""au BufRead,BufNewFile *.wiki,*.md,*.asm,*.c,*.cpp,*.java,*.cs,*.sh,*.lua,*.pl,*.pm,*.py,*.rb,*.hs,*.vim 2match Underlined /.\%81v/
au BufRead,BufNewFile *.* 2match Underlined /.\%81v/

": maxmempattern
set maxmempattern=2000

"文件类型特殊配置"
au BufNewFile,BufRead *.vue set filetype=html
au BufNewFile,BufReadPre,BufReadPost,FileReadPre *.js set expandtab
au BufNewFile,BufReadPre,BufReadPost,FileReadPre *.js set tabstop=2
au BufNewFile,BufReadPre,BufReadPost,FileReadPre *.js set shiftwidth=2
au BufNewFile,BufReadPre,BufReadPost,FileReadPre *.vue set expandtab
au BufNewFile,BufReadPre,BufReadPost,FileReadPre *.vue set tabstop=2
au BufNewFile,BufReadPre,BufReadPost,FileReadPre *.vue set shiftwidth=2
