package filedelivery

import (
	"io"
	"mime/multipart"
	"path/filepath"

	fileusecase "github.com/aryyawijaya/go-storage-with-clean-arch/modules/file/use-case"
)

func (fh *FileHandler) getFileContent(file *multipart.FileHeader) (content []byte, err error) {
	// Read file content
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	// Create file
	content, err = io.ReadAll(src)
	if err != nil {
		return
	}

	return
}

func (fh *FileHandler) getFileExt(file *multipart.FileHeader) string {
	return filepath.Ext(file.Filename)
}

func (fh *FileHandler) newFileDto(content []byte, ext string) *fileusecase.FileDto {
	return &fileusecase.FileDto{
		Content: content,
		Ext:     ext,
	}
}
