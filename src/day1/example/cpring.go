package main
import "C"
import "unsafe"

/*
初识Cgo：
warning: GOPATH set to GOROOT (E:\code\Golang) has no effect
# command-line-arguments
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%
 */
func main() {
	cstr := C.CString("Hello World")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
