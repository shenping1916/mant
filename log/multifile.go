package log

import (
	"fmt"
)

type MultiFileObject struct {
	Files []*FileObject
}

// NewMultiFileObject is an initialization constructor
// that returns a MultiFileObject pointer object.
func NewMultiFileObject(path string, levels []string, rotate,compress,daily bool, lines,size int64, keepdays int) *MultiFileObject {
	multiobj := &MultiFileObject{}
	multiobj.Files = make([]*FileObject, 0, len(levels))

	for index, level := range levels {
		fName := fmt.Sprintf("%s/%s.log", path, level)

		f := NewFileObject(fName, rotate, compress, daily, WithMaxLinesOption(lines), WithMaxSizeOption(size), WithMaxDaysOption(keepdays))
		f.level = index

		multiobj.Files = append(multiobj.Files, f)
	}

	return multiobj
}

// Write method is used to write a byte array to all files.
// Automatically execute rotate logic and delete logic before writing.
// TODO: To be optimized
func (m *MultiFileObject) Writing(p []byte) error {
	for i, j := 0, len(m.Files); i < j; i ++ {
		f := m.Files[i]
		if f != nil && byte('0' + f.level) == p[0:1][0] {
			if err := f.Writing(p); err != nil {
				return err
			}
		}
	}

	return nil
}

// Flush will flush the cache of all file handles
func (m *MultiFileObject) Flush() {
	for i, j := 0, len(m.Files); i < j; i ++ {
		f := m.Files[i]
		if f != nil {
			f.Flush()
		}
	}
}

// Close all file handle resources.
func (m *MultiFileObject) Close() {
	for i, j := 0, len(m.Files); i < j; i ++ {
		f := m.Files[i]
		if f != nil {
			f.Close()
		}
	}
}




