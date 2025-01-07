package valider

/*
#include <stdlib.h>
#include <string.h>
#include "valider.h"

int validate_string(const char* str) {
	return valid(str);
}
*/
import "C"
import "unsafe"

func IsValid(text string) int {
	cStr := C.CString(text)
	defer C.free(unsafe.Pointer(cStr))
	
	return int(C.validate_string(cStr))
}