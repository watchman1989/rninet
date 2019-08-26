package global

import (
	"fmt"
	"runtime"
)

var (
	SYSTEM = runtime.GOOS
	SLASH = "/"
)

func init () {
	fmt.Printf("SYSTEM: %s\n", SYSTEM)
	if SYSTEM == "windows" {
		SLASH = "\\"
	}
}