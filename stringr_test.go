package stringr

import (
	"io"
	"io/ioutil"
	"testing"
)

func TestStringReader(t *testing.T) {
	str := "Some nice string example!"
	r := &StringReader{Val: str}

	toRead := 10
	sub := make([]byte, toRead, toRead)
	n, err := r.Read(sub)
	if err != nil {
		t.Errorf("Can't read first %d bytes. Err: %v", toRead, err)
	}

	if n != toRead {
		t.Errorf("Read wrong number of bytes. Got: %d, Expected: %d", n, toRead)
	}

	if string(sub) != str[0:toRead] {
		t.Errorf("Read wrong bytes! Got: %s, Expected: %s", string(sub), str[0:toRead])
	}
}

func TestShortString(t *testing.T) {
	str := "example"
	r := &StringReader{Val: str}

	toRead := 10
	sub := make([]byte, toRead, toRead)
	n, err := r.Read(sub)
	if err != nil {
		t.Errorf("Can't read first %d bytes. Err: %v", toRead, err)
	}

	if n != len(str) {
		t.Errorf("Read wrong number of bytes. Got: %d, Expected: %d", n, len(str))
	}

	if string(sub[:n]) != str {
		t.Errorf("Read wrong bytes! Got: %s, Expected: %s", string(sub), str[:n])
	}
}

func TestEmptyString(t *testing.T) {
	str := ""
	r := &StringReader{Val: str}

	toRead := 10
	sub := make([]byte, toRead, toRead)
	_, err := r.Read(sub)
	if err != io.EOF {
		t.Errorf("Error! Got: %v, expected: %v", err, io.EOF)
	}
}

func TestReadAll(t *testing.T) {
	str := "Some nice string example!"
	r := &StringReader{Val: str}

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		t.Errorf("Can't read all bytes. Err: %v", err)
	}

	if str != string(bytes) {
		t.Errorf("Read wrong bytes! Got: %s, Expected: %s", string(bytes), str)
	}
}
