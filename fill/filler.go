// Package fill provides fillers that fills each field with a value.
package fill

import "errors"

var (
	ErrCodecMismatch = errors.New("unsupported codec (could be invalid JSON format)")
)

// Filler tries to correspond input text to a struct.
type Filler interface {
	// Fill receives a struct v and corresponds input that is have internally to the struct.
	// Fill may return these errors:
	//
	//   - io.EOF: At the end of input.
	//   - ErrCodecMismatch: If v isn't a supported type.
	//
	Fill(v interface{}) error
}

// InteractiveFillerOpts represents options for InteractiveFiller.
type InteractiveFillerOpts struct {
	// DigManually is true, Fill asks whether to dig down if it encountered to a message field.
	DigManually,
	// BytesFromFile is true, Fill will read the contents of the file from the provided relative path.
	BytesFromFile,
	// AddRepeatedManually is true, Fill asks whether to add a repeated field value
	// if it encountered to a repeated field.
	AddRepeatedManually bool
}

// Filler tries to correspond input text to a struct interactively.
type InteractiveFiller interface {
	// Fill receives a struct v and corresponds input that is have internally to the struct.
	// Fill may return these errors:
	//
	//   - io.EOF: At the end of input.
	//   - ErrCodecMismatch: If v isn't a supported type.
	//
	Fill(v interface{}, opts InteractiveFillerOpts) error
}
