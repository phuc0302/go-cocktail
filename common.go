package cocktail

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"os"
	"runtime"
	"strings"
)

// Constants
const (
	DELETE = "DELETE"
	GET    = "GET"
	HEAD   = "HEAD"
	PATCH  = "PATCH"
	POST   = "POST"
	PUT    = "PUT"
)

type (
	// Path params from url pattern
	PathParams map[string]string

	// Multipart files from multipart form
	FileParams map[string][]*multipart.FileHeader
)

/** Convert function pointer to human readable text. */
func getFunctionName(pc uintptr) string {
	fn := runtime.FuncForPC(pc)

	// Condition validation: return don't know if function is not available
	if fn == nil {
		return string(dunno)
	}

	// Convert function name to byte array for modification
	name := []byte(fn.Name())

	// Eliminate the path prefix
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}

	// Eliminate period prefix
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}

	// Convert center dot to dot
	name = bytes.Replace(name, centerDot, dot, -1)
	return string(name)
}

/** Return a nicely formated stack frame. */
func getStack(skip int) []string {
	srcPath := fmt.Sprintf("%s/src", os.Getenv("GOPATH"))
	traces := make([]string, 5)

	for i, j := skip, 0; ; i++ {
		// Condition validation: Stop if there is nothing else
		pc, file, line, ok := runtime.Caller(i)
		if !ok || j >= 5 {
			break
		}

		// Condition validation: Skip go root
		if !strings.HasPrefix(file, srcPath) {
			continue
		}

		// Trim prefix
		file = file[len(srcPath):]

		// Print this much at least. If we can't find the source, it won't show.
		traces[j] = fmt.Sprintf("%s: %s (%d)", file, getFunctionName(pc), line)
		j++
	}
	return traces
}
