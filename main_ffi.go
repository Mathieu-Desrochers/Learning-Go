package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// calling C code
func Print(s string) {
	cs := C.CString(s)
	defer func() { C.free(unsafe.Pointer(cs)) }()

	C.fputs(cs, (*C.FILE)(C.stdout))
	C.fflush((*C.FILE)(C.stdout))
}
