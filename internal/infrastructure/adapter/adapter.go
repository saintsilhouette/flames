package adapter

import (
	"fmt"
	"io"
	"log/slog"
)

// IOAdapter is the additional layer for the i/o operations.
type IOAdapter struct {
	w      io.Writer
	logger *slog.Logger
}

// New instantiates a new IOAdapter entity.
func New(w io.Writer, logger *slog.Logger) *IOAdapter {
	return &IOAdapter{
		w:      w,
		logger: logger,
	}
}

// Output writes data to the provided io.Writer.
func (a *IOAdapter) Output(content string) {
	_, err := fmt.Fprint(a.w, content)
	if err != nil {
		a.logger.Info(err.Error())
	}
}
