package main

import (
	"errors"
	"io"
	"os"
	"path"
	"regexp"
)

const filePerm = 0666

type Fs struct {
	root string
}

func NewFs(root string) (*Fs, error) {
	fs := &Fs{root: root}

	err := os.MkdirAll(root, filePerm)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(fs.CurrPath(), filePerm)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(fs.SnapPath(), filePerm)
	if err != nil {
		return nil, err
	}

	return fs, nil
}

func (f Fs) CurrPath() string {
	return path.Join(f.root, "curr")
}

func (f Fs) SnapPath() string {
	return path.Join(f.root, "snap")
}

func (f Fs) GetCurrFilePath(filepath string) string {
	return path.Join(f.CurrPath(), filepath)
}

var snap_validate = regexp.MustCompile("$[a-f0-9]{64}")
var InvalidSnapErr = errors.New("snap is not valid")

func (f Fs) GetSnapPath(snap string) (string, error) {
	if !snap_validate.Match([]byte(snap)) {
		return "", InvalidSnapErr
	}
	path := path.Join(f.SnapPath(), snap)

	return path, nil
}

func (f Fs) Upload(r io.Reader, pathname string) error {
	p := path.Join(f.CurrPath(), pathname)

	err := os.MkdirAll(path.Dir(p), filePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(p)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, r)
	if err != nil {
		os.Remove(p)
		return err
	}

	return nil
}
