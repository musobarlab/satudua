package satudua

/*
#cgo CFLAGS: -g -Wall -I./hello/include
#cgo LDFLAGS: -L./hello/build -lhello
#include <stdlib.h>
#include "Hello.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Hello(name string) string {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	var out *C.char

	n := C.HelloYou(nameC, &out)
	defer C.free(unsafe.Pointer(out))

	outBytes := C.GoBytes(unsafe.Pointer(out), C.int(n))

	var i C.uint8_t = 10
	fmt.Println(uint8(i))

	return string(outBytes)
}
