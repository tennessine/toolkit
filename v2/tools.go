package toolkit

import (
	"crypto/rand"
	"errors"
	"net/http"
)

const randStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789_+"

type Tools struct {
	MaxFileSize int
}

func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}

type UploadedFile struct {
	NewFileName      string
	OriginalFileName string
	FileSize         int
}

func (t *Tools) UploadFiles(r *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}
	var uploadedFiles []*UploadedFile

	if t.MaxFileSize == 0 {
		t.MaxFileSize = 1024 * 1024 * 1024
	}

	err := r.ParseMultipartForm(int64(t.MaxFileSize))
	if err != nil {
		return nil, errors.New("the uploaded file is too big")
	}

	return uploadedFiles, nil
}
