//create: 2017/10/26 22:00:09 change: 2017/10/26 23:29:07 lijiaocn@foxmail.com

package main

import (
	"github.com/lijiaocn/GoPkgs/version"
	"github.com/lijiaocn/build-in-docker/hello"
)

func main() {
	version.Show()
	hello.Hello()
}
