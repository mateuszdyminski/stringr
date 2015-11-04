package stringr

import "io"

// StringReader - struct which implements io.Reader interface.
type StringReader struct {
	Val    string
	offset int
}

// Read - reads specified number of bytes from string.
func (r *StringReader) Read(p []byte) (n int, e error) {
	if r.offset >= len(r.Val) {
		return 0, io.EOF
	}

	if len(r.Val)-r.offset > len(p) {
		n = len(p)
	} else {
		n = len(r.Val) - r.offset
	}

	for i, v := range r.Val[r.offset : r.offset+n] {
		p[i] = byte(v)
		r.offset++
	}

	return n, nil
}

// Reset - resets current offset.
func (r *StringReader) Reset() {
	r.offset = 0
}
