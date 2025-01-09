package utils

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"

	"github.com/osamikoyo/void-init/components"
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

	if err := components.WriteMaingo("main.go");err != nil{
		return err
	}

	if err := components.WriteRoot("cmd/root.go");err != nil{
		return err
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