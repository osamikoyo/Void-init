package utils

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
)

func Init(name string) error {
	result := C.system(C.CString(fmt.Sprintf("go mod init %s", name)))
	if int(result) == 1 {
		return errors.New("code = 1")
	}

	result = C.system(C.CString("mkdir cmd pkg internal"))
	if int(result) == 1 {
		return errors.New("code = 1")
	}

	result = C.system(C.CString("touch main.go cmd/root.go internal/utils.go"))
	if int(result) == 1 {
		return errors.New("code = 1")
	}

	return nil
}

func Install(path string) error {
	result := C.system(C.CString(fmt.Sprintf("go install %s", path)))
	if int(result) == 1 {
		return errors.New("code = 1")
	}

	return nil
}