package b28

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

//https://storj.io/blog/a-tale-of-two-copies

const numEntries = 100_000

type Entry struct {
	b [28]byte
}

type Buffer struct {
	fh      *os.File
	flushes int
	n       int
	buf     [numEntries]Entry
}

func (b *Buffer) Append(ent Entry) error {
	if b.n < numEntries-1 {
		b.buf[b.n] = ent
		b.n++
		return nil
	}
	return b.appendSlow(ent)
}

func (b *Buffer) appendSlow(ent Entry) error {
	if err := b.Flush(); err != nil {
		return err
	}
	b.buf[0] = ent
	b.n = 1
	return nil
}

func (b *Buffer) Flush() error {
	b.flushes++
	for i := 0; i < b.n; i++ {
		_, err := b.fh.Write(b.buf[i].b[:])
		if err != nil {
			return err
		}
	}
	return nil
}

func BenchmarkBuffer(b *testing.B) {
	fh, err := ioutil.TempFile("", "buf")
	if err != nil {
		b.Fatal(err)
	}
	defer func() {
		_ = fh.Close()
	}()

	buf := &Buffer{fh: fh}
	now := time.Now()
	ent := Entry{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := fh.Seek(0, io.SeekStart)
		if err != nil {
			b.Fatal(err)
		}

		for i := 0; i < 1e5; i++ {
			_ = buf.Append(ent)
		}
		_ = buf.Flush()
	}

	b.ReportMetric(float64(time.Since(now).Nanoseconds())/float64(b.N)/1e5, "ns/key")
	b.ReportMetric(float64(buf.flushes)/float64(b.N), "flushes")
}
