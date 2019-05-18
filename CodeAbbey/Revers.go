package main
import (
	"strings"
	"fmt"
)
func main() {
	var prl = "shelf turn interrogative stay coal make stove where"
	prlArray := strings.Fields(prl)
	var result []string
	for i:=len(prlArray)-1; i>=0; i--{
		hue := prlArray[i]
		result = append(result,hue)
	}
	for j:=0;j<len(result);j++  {
		z := result[j]
		x := strings.SplitAfter(z, "")
		t := len(x)
		for k:=0;k<len(x)/2;k++  {
			y:=x[k]
			x[k]=x[t-1]
			x[t-1]=y
			t--
		}
		fmt.Println(x)
	}
}
