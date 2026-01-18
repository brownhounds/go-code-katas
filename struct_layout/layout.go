package struct_layout

import (
	"fmt"
	"reflect"
	"strings"
)

// @Note:
// When padding is unavoidable, it is generally preferable for it to appear
// at the end of the struct (tail padding) rather than between fields.
//
// Internal padding creates "holes" that waste cache lines and make the
// effective memory layout harder to reason about. Tail padding keeps all
// frequently accessed fields tightly packed and predictable.
//
// Reordering fields does not always reduce the total size of a struct
// (alignment may still require padding), but it can move padding to the end,
// which is typically the least harmful place for it.
func PrintLayout[T any](name string) string {
	var zero T
	t := reflect.TypeOf(zero)

	var b strings.Builder
	fmt.Fprintf(&b, "== %s ==\n", name)
	fmt.Fprintf(&b, "type: %s\n", t.String())
	fmt.Fprintf(&b, "size: %d bytes\n", t.Size())
	fmt.Fprintf(&b, "align: %d bytes\n", t.Align())

	if t.Kind() != reflect.Struct {
		fmt.Fprintf(&b, "(not a struct)\n")
		return b.String()
	}

	n := t.NumField()
	var prevEnd uintptr = 0
	var totalPadding uintptr = 0

	for i := range n {
		f := t.Field(i)
		off := f.Offset
		fs := f.Type.Size()
		fa := f.Type.Align()

		if off > prevEnd {
			pad := off - prevEnd
			totalPadding += pad
			fmt.Fprintf(&b, "  pad:  +%d bytes (between fields)\n", pad)
		}

		fmt.Fprintf(&b, "  field %-12s offset=%-3d size=%-3d align=%-2d type=%s\n",
			f.Name, off, fs, fa, f.Type.String(),
		)

		prevEnd = off + fs
	}

	if uintptr(t.Size()) > prevEnd {
		tail := uintptr(t.Size()) - prevEnd
		totalPadding += tail
		fmt.Fprintf(&b, "  pad:  +%d bytes (tail padding)\n", tail)
	}

	fmt.Fprintf(&b, "padding total: %d bytes\n", totalPadding)
	return b.String()
}

type Bad struct {
	A bool   // 1 byte
	B int64  // 8 bytes, needs 8-byte alignment -> padding inserted after A
	C bool   // 1 byte -> more padding at end
	D int32  // 4 bytes
	E uint16 // 2 bytes
}

type Good struct {
	B int64
	D int32
	E uint16
	A bool
	C bool
}

type WithPointersBad struct {
	ID    int64
	Flag  bool
	Ptr   *byte
	Slice []byte
}

type WithPointersGood struct {
	ID    int64
	Slice []byte
	Ptr   *byte
	Flag  bool
}
