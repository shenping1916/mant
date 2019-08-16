package log

import (
	"fmt"
	"strconv"
)

type MultiFileObject struct {
	Files []*FileObject
}

// NewMultiFileObject is an initialization constructor
// that returns a MultiFileObject pointer object.
func NewMultiFileObject(path string, levels []string, rotate, compress, daily bool, lines, size int64, keepDays int) *MultiFileObject {
	obj := &MultiFileObject{}
	obj.Files = make([]*FileObject, 0, len(levels))

	for _, level := range levels {
		fName := fmt.Sprintf("%s/%s.log", path, level)
		f := NewFileObject(fName, rotate, compress, daily, WithMaxLinesOption(lines), WithMaxSizeOption(size), WithMaxDaysOption(keepDays))

		obj.Files = append(obj.Files, f)
	}

	return obj
}

// Writing is used to write a byte array to all files.
// Automatically execute rotate logic and delete logic before writing.
func (m *MultiFileObject) Writing(p []byte) error {
	level, _ := strconv.Atoi(string(p[0:1][0]))
	f := m.Files[level]
	if f != nil {
		f.Writing(p[:])
	}

	return nil
}

// Flush will flush the cache of all file handles
func (m *MultiFileObject) Flush() {
	for i, j := 0, len(m.Files); i < j; i++ {
		f := m.Files[i]
		if f != nil {
			f.Flush()
		}
	}
}

// Close all file handle resources.
func (m *MultiFileObject) Close() {
	for i, j := 0, len(m.Files); i < j; i++ {
		f := m.Files[i]
		if f != nil {
			f.Close()
		}
	}
}
