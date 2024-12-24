package fileroutine

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
)

// ImagesManager manages the image directory and the naming of added files.
type ImagesManager struct {
	Directory string
}

// New instantiates a new ImagesManager entity.
func New(directory string) *ImagesManager {
	return &ImagesManager{
		Directory: directory,
	}
}

// ensureDirectory either checks that provided
// directory exists or creates it.
func (im *ImagesManager) ensureDirectory() error {
	if _, err := os.Stat(im.Directory); os.IsNotExist(err) {
		if err := os.MkdirAll(im.Directory, 0o755); err != nil {
			return DirectoryCreationError
		}
	} else if err != nil {
		return ExistanceUncertaintyError
	}

	return nil
}

// getNextFileName determines the suitable name for the
// rendered image.
func (im *ImagesManager) getNextFileName() (string, error) {
	files, err := os.ReadDir(im.Directory)
	if err != nil {
		return "", DirectoryReadingError
	}

	indexes := make([]int, 0, len(files))

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		underscore := strings.LastIndex(file.Name(), "_")
		if underscore == -1 {
			continue
		}

		dot := strings.LastIndex(file.Name(), ".")
		if dot == -1 {
			continue
		}

		index, err := strconv.Atoi(file.Name()[underscore+1 : dot])
		if err != nil {
			return "", InvalidImageIndex
		}

		indexes = append(indexes, index)
	}

	slices.Sort(indexes)

	var index int

	if len(indexes) == 0 {
		index = 1
	} else {
		index = indexes[len(indexes)-1] + 1
	}

	return fmt.Sprintf("fractal_%d.png", index), nil
}

// CreateImageFile creates file for the image
// and writes it.
func (im *ImagesManager) CreateImageFile(img *image.RGBA) error {
	if err := im.ensureDirectory(); err != nil {
		return DirectoryCreationError
	}

	fileName, err := im.getNextFileName()
	if err != nil {
		return err
	}

	fullPath := path.Join(im.Directory, fileName)

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return ImageEncodeError
	}

	return nil
}
