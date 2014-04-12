package stringprep

/*
#cgo LDFLAGS: -lidn
#include <stdlib.h>
#include "stringprep.h"

int nameprep(char *in, size_t maxlen) {
	return stringprep_nameprep(in, maxlen);
}

int nameprep_no_unassigned(char *in, size_t maxlen) {
	return stringprep_nameprep_no_unassigned(in, maxlen);
}
*/
import "C"

import (
	"unsafe"
)

func Nameprep(in string) string {
	var cin *C.char = C.CString(in)
	defer C.free(unsafe.Pointer(cin))
	C.nameprep(cin, C.size_t(len(in)))
	return C.GoString(cin)
}

func NameprepNoUnassigned(in string) string {
	var cin *C.char = C.CString(in)
	defer C.free(unsafe.Pointer(cin))
	C.nameprep_no_unassigned(cin, C.size_t(len(in)))
	return C.GoString(cin)
}
