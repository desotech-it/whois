package util

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func AddContentInfoToResponseHeades(w http.ResponseWriter, f *os.File) error {
	fileinfo, err := f.Stat()
	if err == nil {
		extension := filepath.Ext(fileinfo.Name())
		mimeType := mime.TypeByExtension(extension)
		header := w.Header()
		header.Set("Content-Type", mimeType)
		header.Set("Content-Length", strconv.FormatInt(fileinfo.Size(), 10))
	}

	return err
}