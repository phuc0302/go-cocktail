package cocktail

import (
	"net/http"
	"net/url"
	"os"
)

/** Create directories. */
func CreateDir(path string, perm os.FileMode) {
	os.MkdirAll(path, perm)
}

/** Check if file exist at path or not. */
func DirExisted(dirPath string) bool {
	fileInfo, err := os.Stat(dirPath)
	return fileInfo.IsDir() && err == nil
}

/** Check if file exist at path or not. */
func FileExisted(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

/** Extract input form. */
func ParseForm(request *http.Request) url.Values {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	return request.Form
}

/** Extract multipart form. */
func ParseMultipartForm(request *http.Request) url.Values {
	err := request.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		panic(err)
	}

	// request.URL.Query()
	params := url.Values(request.MultipartForm.Value)
	for k, v := range request.URL.Query() {
		params[k] = v
	}
	return params
}

/** Format path. */
func FormatPath(path string) string {
	/* Condition validation: Turn empty string into "/" */
	if len(path) == 0 {
		return "/"
	}

	// Try var buffer bytes.Buffer
	var buf []byte
	n := len(path)
	r := 1
	w := 1

	if path[0] != '/' {
		r = 0
		buf = make([]byte, n+1)
		buf[0] = '/'
	}

	trailing := (n > 2 && path[n-1] == '/')
	for r < n {
		switch {
		case path[r] == '/': // Empty path element, trailing slash is added after the end
			r++

		case path[r] == '.' && r+1 == n:
			trailing = true
			r++

		case path[r] == '.' && path[r+1] == '/': // . element
			r++

		case path[r] == '.' && path[r+1] == '.' && (r+2 == n || path[r+2] == '/'): // .. element: remove to last /
			r += 2

			if w > 1 { // can backtrack
				w--

				if buf == nil {
					for w > 1 && path[w] != '/' {
						w--
					}
				} else {
					for w > 1 && buf[w] != '/' {
						w--
					}
				}
			}

		default:
			// real path element.
			// add slash if needed
			if w > 1 {
				bufApp(&buf, path, w, '/')
				w++
			}

			// copy element
			for r < n && path[r] != '/' {
				bufApp(&buf, path, w, path[r])
				w++
				r++
			}
		}
	}

	// re-append trailing slash
	if trailing && w > 1 {
		bufApp(&buf, path, w, '/')
		w++
	}

	if buf == nil {
		return path[:w]
	}
	return string(buf[:w])
}

// internal helper to lazily create a buffer if necessary
func bufApp(buf *[]byte, s string, w int, c byte) {
	if *buf == nil {
		if s[w] == c {
			return
		}

		*buf = make([]byte, len(s))
		copy(*buf, s[:w])
	}
	(*buf)[w] = c
}
