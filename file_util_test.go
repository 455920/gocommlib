package file

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	tmpDir := os.TempDir()
	path := tmpDir + "/_TestCreate_/dir/file"
	f, err := Create(path)
	if err != nil {
		t.Fatalf("create file err:%q", path)
	}
	if f == nil {
		t.Fatal("create file failed")
	}
}

func TestIsExist(t *testing.T) {
	tmpDir := os.TempDir()
	path := tmpDir + "/_TestIsExist_/dir/not_exist"
	if IsExist(path) {
		t.Fatalf("file should be not exist:%q", path)
	}
	path = tmpDir + "/_TestCreate_/dir/file_exist"
	f, err := Create(path)
	if err != nil {
		t.Fatalf("create file err:%s", err)
	}
	defer f.Close()
	if !IsExist(path) {
		t.Fatalf("file should be exist:%q", path)
	}
}

func TestMd5Sum(t *testing.T) {
	tmpDir := os.TempDir()
	path := tmpDir + "/_TestMd5Sum_/dir/md5sum"
	f, err := Create(path)
	if err != nil {
		t.Fatalf("create file err:%s", err)
	}
	f.WriteString("1")
	f.Close()
	md5sum, err := Md5Sum(path)
	if err != nil {
		t.Fatalf("file md5sum err:%s", err)
	}

	hash := md5.Sum([]byte("1"))
	expected := hex.EncodeToString(hash[:])
	if md5sum != expected {
		t.Fatalf("check md5sum err, actual:%s, expected:%s", md5sum, expected)
	}
}

func TestRemove(t *testing.T) {
	tmpDir := os.TempDir()
	path := tmpDir + "/_TestRemove_/dir/not_exist"
	err := Remove(path)
	if err != nil {
		t.Fatalf("remove file err:%s", err)
	}

	path = tmpDir + "/_TestRemove_/dir/file_exist"
	f, err := Create(path)
	if err != nil {
		t.Fatalf("create file err:%q", path)
	}
	f.Close()
	err = Remove(path)
	if err != nil {
		t.Fatalf("remove file err:%s", err)
	}
	if IsExist(path) {
		t.Fatalf("file should be not exist:%q", path)
	}
}
