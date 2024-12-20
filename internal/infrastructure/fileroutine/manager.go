package fileroutine

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
)

// ImagesManager manages the image directory and the naming of added files.
type ImagesManager struct {
	Directory   string
	LastFile    string
	FilePattern *regexp.Regexp
}

// New instantiates a new ImagesManager entity.
func New(directory string) (*ImagesManager, error) {
	manager := &ImagesManager{
		Directory:   directory,
		FilePattern: regexp.MustCompile(`^fractal_(\d+)\.png$`),
	}

	if err := manager.EnsureDirectory(); err != nil {
		return nil, ManagerCreationError
	}

	return manager, nil
}

// EnsureDirectory either checks that provided
// directory exists or creates it.
func (im *ImagesManager) EnsureDirectory() error {
	if _, err := os.Stat(im.Directory); os.IsNotExist(err) {
		if err := os.MkdirAll(im.Directory, 0o755); err != nil {
			return DirectoryCreationError
		}
	} else if err != nil {
		return ExistanceUncertaintyError
	}

	return nil
}

// GetNextFileName determines the next available
// filename that satisfies the pattern fractal_xxx.png.
func (im *ImagesManager) GetNextFileName() (string, error) {
	files, err := os.ReadDir(im.Directory)
	if err != nil {
		return "", DirectoryReadingError
	}

	indices := make([]int, 0, 16)

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		matches := im.FilePattern.FindStringSubmatch(f.Name())
		index, err := strconv.Atoi(matches[1])

		if err != nil {
			continue
		}

		indices = append(indices, index)
	}

	sort.Ints(indices)

	nextIndex := 1
	if len(indices) != 0 {
		nextIndex = indices[len(indices)-1] + 1
	}

	nextFile := fmt.Sprintf("filename_%d.png", nextIndex)

	return nextFile, nil
}

// CreateImageFile creates file for the image
// and writes it.
func (im *ImagesManager) CreateImageFile(img *image.RGBA) error {
	fileName, err := im.GetNextFileName()
	if err != nil {
		return err
	}

	fullPath := path.Join(im.Directory, fileName)

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return ImageEncodeError
	}

	return nil
}
