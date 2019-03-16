//create: 2015/07/21 09:51:54 change: 2015/07/21 09:58:41 author:lijiao
package dockeropt

import(
	"testing"
)

func TestGoClientExample(t *testing.T){
	err := GoClientExample()
	if err != nil{
		t.Fatal(err)
	}
}
