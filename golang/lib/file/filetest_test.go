package file

import (
	"path/filepath"
	"testing"
)

func TestExists(t *testing.T) {
	if !Exists(path(".")) {
		t.Errorf("TestExists error")
	}
}

func TestNotExists(t *testing.T) {
	if Exists(path("hoge")) {
		t.Errorf("TestNotExists error")
	}
}

func TestIsFile(t *testing.T) {
	if !Exists(path("test.md")) {
		t.Errorf("TestIsFile error")
	}
}
func TestIsNotFile(t *testing.T) {
	if Exists(path("nottest.md")) {
		t.Errorf("TestIsNotFile error")
	}
}

func TestIsDir(t *testing.T) {
	if !Exists(path("../filetest")) {
		t.Errorf("TestIsDir error")
	}
}

func TestIsNotDir(t *testing.T) {
	if Exists(path("newdir")) {
		t.Errorf("TestIsFile error")
	}
}

func path(name string) string {
	base := "../test/file/filetest"
	return filepath.Join(base, name)
}
