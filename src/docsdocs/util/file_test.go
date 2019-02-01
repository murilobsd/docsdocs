package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	filename := "notexist"
	contents, err := ReadFile(filename)
	if err == nil {
		t.Fatalf("Readfile %s: error expected, none found", filename)
	}
	filename = "file_test.go"
	contents, err = ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile %s: %v", filename, err)
	}
	dir, _ := os.Stat(filename)

	if dir.Size() != int64(len(contents)) {
		t.Errorf("Stat %q: size %d want %d", filename, dir.Size(), int64(len(contents)))
	}

}
func TestSaveFile(t *testing.T) {
	f, err := ioutil.TempFile("", "docsfile")
	if err != nil {
		t.Fatal(err)
	}
	filename := f.Name()
	data := "hi =]"

	if err := SaveFile(filename, []byte(data)); err != nil {
		t.Fatalf("SaveFile %s: %v", filename, err)
	}
	contents, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("Reafile %s: %v", filename, err)
	}

	if string(contents) != data {
		t.Fatalf("contents = %q\nexpected = %q", string(contents), data)
	}

	f.Close()
	os.Remove(filename)

}
