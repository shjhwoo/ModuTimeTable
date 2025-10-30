package util

import (
	"fmt"
	"runtime"
)

var funcInfoFormat = "{%s:%d} [%s]"

func getFuncInfo(pc uintptr, file string, line int) string {
	f := runtime.FuncForPC(pc)
	if f == nil {
		return fmt.Sprintf(funcInfoFormat, file, line, "unknown")
	}
	return fmt.Sprintf(funcInfoFormat, file, line, f.Name())
}

var wrapFormat = "%s\n%w"

func WrapWithStack(err error) error {
	pc, file, line, ok := runtime.Caller(1)

	if !ok {
		return fmt.Errorf(wrapFormat, "", err)
	}

	// {file:line} [funcName] msg
	stack := fmt.Sprintf("%s %s", getFuncInfo(pc, file, line), "")

	return fmt.Errorf(wrapFormat, stack, err)
}
