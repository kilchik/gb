package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)


type recorder struct {
	counter int
	writer func(bb []byte) (n int, err error)
}

func (r recorder) Write(bb []byte) (n int, err error) {
	return r.writer(bb)
}

func TestDumpFileInfo(t *testing.T) {
	content := []byte("SOME EXIF HEADER")
	buf := bytes.NewBuffer(content)
	rec := &recorder{}
	rec.writer = func(bb []byte) (n int, err error) {
		assert.Equal(t, "type:JPEG,orientation:LEFT-TOP", string(bb))
		rec.counter++
		return len(bb), nil
	}

	dumpFileInfo(rec, buf)

	require.Equal(t, 1, rec.counter)
}
