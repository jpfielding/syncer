package cmd

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

func HandleError(function interface{}, err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "%v\n", fmtError(function, err))
	os.Exit(1)
}

func fmtError(function interface{}, err error) error {
	if err == nil {
		return nil
	}
	fullName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	splitName := strings.Split(fullName, ".")
	localName := splitName[len(splitName)-1]
	return fmt.Errorf("%s: %v", localName, err)
}
