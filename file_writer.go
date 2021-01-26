package logger

import (
	"os"
	"strings"
)

type fileWriter struct {
	file *os.File
}

func NewFileWriter(addr string) (*fileWriter, error) {
	p := strings.Split(addr, "/")
	path := strings.Join(p[:len(p)-1], "/")
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}

	f, err := os.Create(p[len(p)-1])
	if err != nil {
		return nil, err
	}
	return &fileWriter{
		file: f,
	}, nil
}

func (f *fileWriter) Emerg(m string) error   { _, err := f.Write([]byte(m)); return err }
func (f *fileWriter) Alert(m string) error   { _, err := f.Write([]byte(m)); return err }
func (f *fileWriter) Crit(m string) error    { _, err := f.Write([]byte(m)); return err }
func (f *fileWriter) Err(m string) error     { _, err := f.Write([]byte(m)); return err }
func (f *fileWriter) Warning(m string) error { _, err := f.Write([]byte(m)); return err }
func (f *fileWriter) Notice(m string) error  { _, err := f.Write([]byte(m)); return err }
func (f *fileWriter) Info(m string) error    { _, err := f.Write([]byte(m)); return err }
func (f *fileWriter) Debug(m string) error   { _, err := f.Write([]byte(m)); return err }

func (f *fileWriter) Write(ba []byte) (int, error) {
	return f.file.Write(ba)
}

func (f *fileWriter) Close() error {
	return f.file.Close()
}
