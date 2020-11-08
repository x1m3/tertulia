package zstrings

import (
	"bytes"
	"compress/gzip"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"sync"
)

// A pool of Gzip Writers. Creating a Gzip Writer is expensive so we will reuse the existing ones
// letting the GC to remove it when needed.
var zStringGzipWritersPool = sync.Pool{New: func() interface{} {
	w, err := gzip.NewWriterLevel(nil, gzip.BestSpeed)
	if err != nil {
		panic(err)
	}
	return w
}}

// A pool of Gzip Readers. Please, see comment for zStringGzipWritersPool
var zStringGzipReadersPool = sync.Pool{New: func() interface{} {
	var buff bytes.Buffer
	// We need to call gzip.NewReader with a valid gzipped stream. This is the value for ""
	gzipHeader := []byte{31, 139, 8, 0, 0, 9, 110, 136, 4, 255, 0, 0, 0, 255, 255}
	buff.Write(gzipHeader)
	r, err := gzip.NewReader(&buff)
	if err != nil {
		panic(err)
	}
	return r
}}

// ZString is a custom type to store compressed or uncompressed strings. The idea is to let the
// programmer chose between CPU usage (compressed) or memory usage.
type ZString struct {
	compressed bool
	value      []byte
}

// NewZstring creates a new ZString uncompressed. Almost as fast as a usual string, suitable for small and
// medium string sizes.
func NewZString(string string) *ZString {
	return &ZString{compressed: false, value: []byte(string)}
}

// NewZStringCompressed created a new ZString that stores the value in compressed format.
// Very convenient for in memory storage of very big strings. CPU usage will increase. Storing a compressed
// string is like 100 times slower than an uncompressed one
// It uses a pool of gzip writers to avoid the overhead of creating a new writer each time. Performance is like
// 10 times better than creating and destroying the writer on each call.
func NewZStringCompressed(string string) *ZString {
	var buff bytes.Buffer
	if len(string) < 256 {
		// Do not compress small strings
		return NewZString(string)
	}
	w := zStringGzipWritersPool.Get().(*gzip.Writer)
	w.Reset(&buff)
	w.Write([]byte(string))
	w.Flush()
	zStringGzipWritersPool.Put(w)
	return &ZString{compressed: true, value: buff.Bytes()}
}

// Returns the value of a string, whether it is compressed or not.
// It uses a pool of gzip readers to minimize the overhead of creating readers on each call.
func (s *ZString) Value() string {
	if s.compressed {
		r := zStringGzipReadersPool.Get().(*gzip.Reader)
		defer zStringGzipReadersPool.Put(r)
		err := r.Reset(bytes.NewReader(s.value))
		if err != nil {
			logrus.Errorf("Error decoding compressed ZString. Got <%s>", err)
		}
		value, _ := ioutil.ReadAll(r)
		return string(value)
	} else {
		return string(s.value)
	}
}
