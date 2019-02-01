package util

import (
	"io/ioutil"
	"os"
	"testing"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("tear test case")
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
