package utils

import (
	"testing"
)

func TestReadFileIntoSlice(t *testing.T) {
	file := ReadFileIntoSlice("input_file_test")
	if len(file) != 3 {
		t.Fatal("Slice length should be 3.")
	}
	if file[0] != "ab" && file[1] != "cd" && file[2] != "12" {
		t.Fatal("File contents have not been read correctly.")
	}
}
