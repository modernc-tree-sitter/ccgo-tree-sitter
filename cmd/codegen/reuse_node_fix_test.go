package main

import (
	"strings"
	"testing"
)

func TestFixReuseNodeLoopPatternA(t *testing.T) {
	for _, maxLit := range []string{"0xffffffff", "4294967295"} {
		in := "\tfor *(*uintptr)(unsafe.Pointer(bp + 104)) != 0 {\n" +
			"\t\tv1 = self18 + 1368\n" +
			"\t\tif (*ReusableNode)(unsafe.Pointer(v1)).Fstack.Fsize > uint32(0) {\n" +
			"\t\t\tv4 = (*(*StackEntry)(unsafe.Pointer((*ReusableNode)(unsafe.Pointer(v1)).Fstack.Fcontents + uintptr((*ReusableNode)(unsafe.Pointer(v1)).Fstack.Fsize-uint32(1))*16))).Fbyte_offset\n" +
			"\t\t} else {\n" +
			"\t\t\tv4 = uint32(" + maxLit + ")\n" +
			"\t\t}\n" +
			"\t\tv2 = v4\n"
		out := fixReuseNodeLoop(in)
		if strings.Contains(out, "for *(*uintptr)(unsafe.Pointer(bp + 104)) != 0") {
			t.Fatalf("pattern A not fixed (max=%s):\n%s", maxLit, out)
		}
		for _, want := range []string{
			"for {",
			"*(*Subtree)(unsafe.Pointer(bp + 104)) = ",
			".Ftree",
			"if *(*uintptr)(unsafe.Pointer(bp + 104)) == 0 {",
			"break",
			"v4 = ",
		} {
			if !strings.Contains(out, want) {
				t.Fatalf("missing %q (max=%s) in\n%s", want, maxLit, out)
			}
		}
	}
}

func TestFixReuseNodeLoopPatternB(t *testing.T) {
	in := "\tfor *(*uintptr)(unsafe.Pointer(bp)) != 0 {\n" +
		"\t\tbyte_offset = reusable_node_byte_offset(tls, self+1368)\n" +
		"\t\tend_byte_offset = byte_offset\n"
	out := fixReuseNodeLoop(in)
	if strings.Contains(out, "for *(*uintptr)(unsafe.Pointer(bp)) != 0") {
		t.Fatalf("pattern B not fixed:\n%s", out)
	}
	if !strings.Contains(out, "reusable_node_tree(tls, self+1368)") {
		t.Fatalf("missing reusable_node_tree:\n%s", out)
	}
}
