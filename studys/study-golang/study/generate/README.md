---
createdate: 2017/05/15 14:23:36
changedate: 2017/05/15 14:35:19

---

Go提供了一个小功能，可以在源码的注释中写入一行命令，通过`go generate`执行这行命令。

	//go:generate echo "This should be  go generate's result ouput"
	package a

当执行:

	go generate github.com/lijiaocn/study-Golang/study/generate/pkgA

或者:

	go generate ./pkgA/a.go

`go:generate`后面的命令会被执行。

可以用这个功能对代码进行编译前的处理。
