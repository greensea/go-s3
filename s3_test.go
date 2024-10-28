package s3

import (
	"fmt"
	"io"
	"testing"
	"time"
)

var ts = fmt.Sprintf("%v", time.Now())

func TestPut1(t *testing.T) {

	err := Put("test", []byte(ts))
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	Put("test", []byte(ts))
	f, err := Get("test")
	if err != nil {
		t.Error(err)
	}

	defer f.Close()
	if raw, _ := io.ReadAll(f); string(raw) != ts {
		t.Error("wrong data")
	}

}

func TestHead(t *testing.T) {
	Put("test", []byte(ts))

	_, err := Head("test")

	if err != nil {
		t.Fatal(err)
	}

}

func TestDir(t *testing.T) {
	Put("testdir/testfile", []byte(ts))

	f, err := Get("testdir/testfile")
	if err != nil {
		t.Error(err)
	}

	defer f.Close()
	if raw, _ := io.ReadAll(f); string(raw) != ts {
		t.Error("wrong data")
	}

}
