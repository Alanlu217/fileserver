package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const filePerm = 0666

var snap_validate = regexp.MustCompile("$[a-f0-9]{64}")
var InvalidSnapErr = errors.New("snap is not valid")

func ValidateSnap(snap string) error {
	if !snap_validate.Match([]byte(snap)) {
		return InvalidSnapErr
	}
	return nil
}

func ValidatePath(path string) error {
	clean := filepath.Clean(path)
	if filepath.IsAbs(clean) {
		return fmt.Errorf("absolute paths are not allowed: %s", path)
	}

	rel, err := filepath.Rel(".", clean)
	if err != nil {
		return fmt.Errorf("cannot evaluate path: %w", err)
	}

	if strings.HasPrefix(rel, "..") {
		return fmt.Errorf("path escapes base directory: %s", path)
	}

	return nil
}

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
	return filepath.Join(f.root, "curr")
}

func (f Fs) SnapPath() string {
	return filepath.Join(f.root, "snap")
}

func (f Fs) GetCurrPath(path string) string {
	return filepath.Join(f.CurrPath(), path)
}

func (f Fs) GetSnapPath(snap, path string) (string, error) {
	if err := ValidateSnap(snap); err != nil {
		return "", err
	}
	p := filepath.Join(f.SnapPath(), snap, path)

	return p, nil
}

func (f Fs) Upload(r io.Reader, path string) error {
	if err := ValidatePath(path); err != nil {
		return err
	}

	p := f.GetCurrPath(path)

	err := os.MkdirAll(filepath.Dir(p), filePerm)
	if err != nil {
		return err
	}

	f.Delete(path)
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

func (f Fs) Delete(path string) error {
	if err := ValidatePath(path); err != nil {
		return err
	}
	os.RemoveAll(f.GetCurrPath(path))
	return nil
}
