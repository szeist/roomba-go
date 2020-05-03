package osargs

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

const argReplaceBufLen int = 1024

func MaskCmdlineArg(argName string) error {
	if !argumentMaskIsSupportedByOS(runtime.GOOS) {
		return fmt.Errorf("Command line argument masking is supported under %s OS. Please load the pasword from environment variable", runtime.GOOS)
	}

	pwArgIdx := findArgIdx(argName)
	if pwArgIdx > -1 {
		maskedPassword := strings.Repeat("X", len(os.Args[pwArgIdx]))
		err := replaceOsArg(pwArgIdx, maskedPassword)
		if err != nil {
			return err
		}
	}
	return nil
}

func argumentMaskIsSupportedByOS(osName string) bool {
	return osName == "linux" || osName == "darwin"
}

func findArgIdx(argName string) int {
	for i, val := range os.Args {
		if val == "-"+argName || val == "--"+argName {
			return i + 1
		} else if strings.HasPrefix(val, "-"+argName+"=") || strings.HasPrefix(val, "--"+argName+"=") {
			return i
		}
	}
	return -1
}

func replaceOsArg(idx int, newValue string) error {
	argStr := (*reflect.StringHeader)(unsafe.Pointer(&os.Args[idx]))
	if argStr.Len > argReplaceBufLen {
		return fmt.Errorf("Cmdline arg length (%d) exceeds replace buffer length (%d)", argStr.Len, argReplaceBufLen)
	}
	arg := (*[argReplaceBufLen]byte)(unsafe.Pointer(argStr.Data))[:argStr.Len]

	n := copy(arg, newValue)
	if n < len(arg) {
		arg[n] = 0
	}

	return nil
}
