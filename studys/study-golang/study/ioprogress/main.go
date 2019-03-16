//create: 2015/09/18 10:37:57 change: 2015/09/18 13:54:03 author:lijiao
package main

import(
	"github.com/mitchellh/ioprogress"
	"github.com/lijiaocn/GoPkgs/virtio"
	"io"
	"os"
	"time"
)

func main() {
	
	var size int = 1000
	r := virtio.VirtReader{Index:0, Capaticy: size, Step:2, Delay: 30 *time.Millisecond}

	progresBar := &ioprogress.Reader{
		Reader: &r,
		Size:   int64(size),
	}

	io.Copy(os.Stdout, progresBar)
}



