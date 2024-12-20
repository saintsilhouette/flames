package adapter

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

// IOAdapter is the additional layer for the i/o operations.
type IOAdapter struct {
	w       io.Writer
	r       io.Reader
	scanner *bufio.Scanner
	logger  *slog.Logger
}

// New instantiates a new IOAdapter entity.
func New(w io.Writer, r io.Reader, logger *slog.Logger) *IOAdapter {
	return &IOAdapter{
		w:       w,
		r:       r,
		scanner: bufio.NewScanner(r),
		logger:  logger,
	}
}

// Input reads data from the provided io.Reader.
func (a *IOAdapter) Input() string {
	a.scanner.Scan()
	input := a.scanner.Text()

	return strings.Trim(input, "\n")
}

// Output writes data to the provided io.Writer.
func (a *IOAdapter) Output(content string) {
	_, err := fmt.Fprint(a.w, content)
	if err != nil {
		a.logger.Info(err.Error())
	}
}
