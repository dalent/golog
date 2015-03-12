package golog

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

const (
	DEFAULT_MAX_DAYS = 7
	//DEFAULT_FILE_SIZE  = 500 * MB
	DEFAULT_FILE_SIZE  = 200 * MB
	DEDAULT_LINE_WIDTH = 120 * B
	DEFAULT_DIR        = "logs"
	SUFFIX             = ".log"
	PREFIX             = "undefined_"
)

type FileWriter struct {
	maxLogSize   int64 //最大的单个log的size
	maxDays      int64 //保存日志天数
	maxLineWidth int64
	dir          string
	prefix       string
	suffix       string
	writer       *FileMux
}

func NewFileWriter() *FileWriter {
	return &FileWriter{DEFAULT_FILE_SIZE, DEFAULT_MAX_DAYS, DEDAULT_LINE_WIDTH, DEFAULT_DIR, PREFIX, SUFFIX, new(FileMux)}
}

func (p *FileWriter) createDir() error {
	return os.MkdirAll(p.dir, os.ModePerm)
}

func (p *FileWriter) Dir(dir string) {
	p.dir = dir
}
func (p *FileWriter) Prefix(prefix string) {
	p.prefix = prefix
}

func (p *FileWriter) Suffix(suffix string) {
	p.suffix = suffix
}

func (p *FileWriter) MaxLogSize(size int64) {
	p.maxLogSize = size
}

func (p *FileWriter) MaxDays(days int64) {
	p.maxDays = days
}

func (p *FileWriter) flush() {
	p.writer.flush()
}

func (p *FileWriter) openFile() error {
	fileNamePre := fmt.Sprintf("%s/%s%s%s",
		p.dir,
		p.prefix,
		time.Now().Format("2006-01-02"),
		p.suffix,
	)

	var tmpName string
	var index = 0
	for {
		tmpName = fmt.Sprintf("%s_%03d", fileNamePre, index)
		info, err := os.Stat(tmpName)
		if err == nil {
			if info.Size() > p.maxLogSize {
				index++
				if index > 100 {
					return errors.New("index more")
				}
				continue
			}
		}

		fd, err := os.OpenFile(tmpName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil {
			return err
		}

		p.writer.initFd(fd)

		break
	}

	p.deleteOldFile()
	return nil
}

func (p *FileWriter) deleteOldFile() {
	filepath.Walk(p.dir, func(path string, info os.FileInfo, err error) (returnErr error) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Unable to delete old log '%s', error: %+v\n", path, r)
			}
		}()

		if !info.IsDir() && info.ModTime().Unix() < (time.Now().Unix()-60*60*24*p.maxDays) {
			if strings.HasPrefix(info.Name(), p.prefix) {
				os.Remove(path)
			}
		}
		return
	})
}

func (p *FileWriter) startLogger() {
	err := p.createDir()
	if err != nil {
		panic("dir create failed: " + err.Error())
	}

	err = p.openFile()
	if err != nil {
		panic("open file failed: " + err.Error())
	}
}

func (p *FileWriter) checkFile() {
	size, err := p.writer.size()
	if err != nil {
		fmt.Println("get log file size error:", err)
		return
	}

	if size > p.maxLogSize {
		err = p.openFile()
		if err != nil {
			fmt.Println("rotate log file  error:", err)
			return
		}
	}
}

func (p *FileWriter) color(_ int) string {
	return ""
}
func (p *FileWriter) writeMsg(msg string) {
	p.writer.writeString(msg)
	p.checkFile()
}

type FileMux struct {
	sync.Mutex
	fd *os.File
}

func (p *FileMux) initFd(fd *os.File) {
	p.Lock()
	defer p.Unlock()
	if p.fd != nil {
		p.fd.Sync()
		p.fd.Close()
	}

	p.fd = fd
}
func (p *FileMux) flush() {
	if p.fd != nil {
		p.Lock()
		defer p.Unlock()
		p.fd.Sync()
	}
}

func (p *FileMux) size() (int64, error) {
	p.Lock()
	defer p.Unlock()
	if p.fd != nil {
		fi, err := p.fd.Stat()
		return fi.Size(), err
	}

	return 0, errors.New("fd not exist")
}

func (p *FileMux) writeString(msg string) {
	p.Lock()
	defer p.Unlock()
	if p.fd != nil {
		if _, err := p.fd.WriteString(msg); err != nil {
			fmt.Println("write message failed:", err)
		}
	}
}
