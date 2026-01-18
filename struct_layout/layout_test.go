package struct_layout

import (
	"fmt"
	"testing"
)

func TestPrintLayout(t *testing.T) {
	fmt.Println(PrintLayout[Bad]("Bad (unordered fields)"))
	fmt.Println(PrintLayout[Good]("Good (reordered fields)"))
	fmt.Println(PrintLayout[WithPointersBad]("WithPointers Bad"))
	fmt.Println(PrintLayout[WithPointersGood]("WithPointers Good"))
}
