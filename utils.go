package cocktail

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

func CreateDir(path string, perm os.FileMode) {
	os.MkdirAll(path, perm)
}

/**
 * Check if file exist at path or not.
 */
func DirExist(dirPath string) bool {
	fileInfo, err := os.Stat(dirPath)
	return fileInfo.IsDir() && err == nil
}

/**
 * Extract input form.
 */
func ExtractInputForm(request *http.Request) url.Values {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	return request.Form
}

/**
 * Extract multipart form.
 */
func ExtractMultipartForm(request *http.Request) (url.Values, FileParams) {
	err := request.ParseMultipartForm(5 << 20) // 5 MB
	if err != nil {
		panic(err)
	}

	// request.URL.Query()
	params := url.Values(request.MultipartForm.Value)
	for k, v := range request.URL.Query() {
		params[k] = v
	}
	return params, FileParams(request.MultipartForm.File)
}

/**
 * Check if file exist at path or not.
 */
func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

/**
 * Return error to client as json form.
 */
func WriteError(response http.ResponseWriter, httpStatus *Status) {
	response.Header().Set("Content-Type", "application/problem+json")
	response.WriteHeader(httpStatus.Status)

	cause, _ := json.Marshal(httpStatus)
	response.Write(cause)
}
