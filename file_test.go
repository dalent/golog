package golog

import (
	"os"
	"testing"
)

//func TestCreateDir(t *testing.T) {
//	err := NewFileWriter().CreateDir()
//	if err != nil {
//		t.Fatal(err)
//	}
//}
//func TestCreateFile(t *testing.T) {
//	file := NewFileWriter()
//	err := file.CreateDir()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	err = file.OpenFile()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//}

//func TestStart(t *testing.T) {
//	file := NewFileWriter()
//	file.StartLogger()
//}

func TestLog(t *testing.T) {
	file := NewFileWriter()
	file.StartLogger()
	file.MaxDays(1)
	file.WriteMsg("test")
	if file.writer.fd == nil {
		t.Fatal("fd nil")
	}

	fi, _ := os.Open(file.writer.fd.Name())
	defer fi.Close()
	var bytes [4]byte
	fi.Read(bytes[:])
	if string(bytes[:]) != "test" {
		t.Fatal("write byte not equal")
	}

	os.RemoveAll(file.dir)
}

//func TestRo(t *testing.T) {
//	file := NewFileWriter()
//	file.StartLogger()
//}

func BenchmarkLog(B *testing.B) {
	file := NewFileWriter()
	file.StartLogger()
	file.MaxDays(1)
	file.MaxLogSize(1000 * KB)
	for i := 0; i < B.N; i++ {
		file.WriteMsg("test")
	}

	os.RemoveAll(file.dir)
}
