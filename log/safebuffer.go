package log

import (
	"bytes"
	"io"
	"sync"
)

// Buffer is a goroutine safe bytes.Buffer
type Buffer struct {
	buffer bytes.Buffer
	mutex  sync.Mutex
}

func NewBuffer() *Buffer {
	buf := new(Buffer)
	// Initialize byte buffer
	buf.buffer = bytes.Buffer{}
	// Preset buffer size to prevent memory redistribution caused by capacity expansion.
	buf.buffer.Grow(1024)
	// Initialize the mutex
	buf.mutex = sync.Mutex{}
	return buf
}

// Read reads the next len(p) bytes from the buffer or until the buffer
// is drained. The return value n is the number of bytes read.
func (b *Buffer) Read(p []byte) (n int, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.Read(p)
}

// Write appends the contents of p to the buffer, growing the buffer as needed. It returns
// the number of bytes written.
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.Write(p)
}

// String returns the contents of the unread portion of the buffer
// as a string.  If the Buffer is a nil pointer, it returns "<nil>".
func (b *Buffer) String() string {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.String()
}

func (b *Buffer) Bytes() []byte {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.Bytes()
}

// Cap returns the capacity of the buffer's underlying byte slice, that is, the
// total space allocated for the buffer's data.
func (b *Buffer) Cap() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.Cap()
}

// Grow grows the buffer's capacity, if necessary, to guarantee space for
// another n bytes.
func (b *Buffer) Grow(n int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.buffer.Grow(n)
}

// Len returns the number of bytes of the unread portion of the buffer;
// b.Len() == len(b.Bytes()).
func (b *Buffer) Len() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.Len()
}

// Next returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by Read.
func (b *Buffer) Next(n int) []byte {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.Next(n)
}

// ReadBytes reads until the first occurrence of delim in the input,
// returning a slice containing the data up to and including the delimiter.
func (b *Buffer) ReadByte() (c byte, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.ReadByte()
}

// ReadBytes reads until the first occurrence of delim in the input,
// returning a slice containing the data up to and including the delimiter.
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.ReadBytes(delim)
}

// ReadFrom reads data from r until EOF and appends it to the buffer, growing
// the buffer as needed.
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.ReadFrom(r)
}

// ReadRune reads and returns the next UTF-8-encoded
// Unicode code point from the buffer.
func (b *Buffer) ReadRune() (r rune, size int, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.ReadRune()
}

// ReadString reads until the first occurrence of delim in the input,
// returning a string containing the data up to and including the delimiter.
func (b *Buffer) ReadString(delim byte) (line string, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.ReadString(delim)
}

// Reset resets the buffer to be empty,
// but it retains the underlying storage for use by future writes.
// Reset is the same as Truncate(0).
func (b *Buffer) Reset() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.buffer.Reset()
}

// Truncate discards all but the first n unread bytes from the buffer
// but continues to use the same allocated storage.
func (b *Buffer) Truncate(n int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.buffer.Truncate(n)
}

// UnreadByte unreads the last byte returned by the most recent successful
// read operation that read at least one byte.
func (b *Buffer) UnreadByte() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.UnreadByte()
}

// UnreadRune unreads the last rune returned by ReadRune.
func (b *Buffer) UnreadRune() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.UnreadRune()
}

// WriteByte appends the byte c to the buffer, growing the buffer as needed.
func (b *Buffer) WriteByte(c byte) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.WriteByte(c)
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to the
// buffer, returning its length and an error, which is always nil but is
// included to match bufio.Writer's WriteRune.
func (b *Buffer) WriteRune(r rune) (n int, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.WriteRune(r)
}

// WriteString appends the contents of s to the buffer, growing the buffer as
// needed.
func (b *Buffer) WriteString(s string) (n int, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.WriteString(s)
}

// WriteTo writes data to w until the buffer is drained or an error occurs.
// The return value n is the number of bytes written.
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.buffer.WriteTo(w)
}
