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

var tag_validate = regexp.MustCompile(`^[\w\-. ]+$`)
var InvalidTagErr = errors.New("tag is not valid")

func ValidateTag(tag string) error {
	if !tag_validate.Match([]byte(tag)) {
		return InvalidTagErr
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

type Path struct {
	tag  string
	path string
}

func CurrPath(path string) Path {
	return Path{
		tag:  "",
		path: path,
	}
}

func TagPath(tag, path string) Path {
	return Path{
		tag:  tag,
		path: path,
	}
}

func (p *Path) Resolve(fs *Fs) string {
	if p.tag == "" {
		return filepath.Join(fs.root, "curr", p.path)
	}
	return filepath.Join(fs.root, "tag", p.tag, p.path)
}

func (p *Path) Stat(fs *Fs) (os.FileInfo, error) {
	info, err := os.Stat(p.Resolve(fs))
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (p *Path) Validate() error {
	var tagerr error = nil
	if p.tag != "" {
		tagerr = ValidateTag(p.tag)
	}
	patherr := ValidatePath(p.path)

	return errors.Join(tagerr, patherr)
}

var ResourceNotFoundErr = errors.New("resource does not exist")

type Fs struct {
	root string
}

// Creates new filesystem and creates basic dir structure
func NewFs(root string) (*Fs, error) {
	fs := &Fs{root: filepath.Join(root, "fs")}

	err := os.MkdirAll(root, filePerm)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(filepath.Join(root, "fs", "curr"), filePerm)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(filepath.Join(root, "fs", "tags"), filePerm)
	if err != nil {
		return nil, err
	}

	return fs, nil
}

func (f *Fs) Exists(path Path) bool {
	if _, err := os.Stat(path.Resolve(f)); err != nil {
		return false
	}
	return true
}

func (f *Fs) Upload(r io.Reader, path Path) error {
	if err := path.Validate(); err != nil {
		return err
	}

	p := path.Resolve(f)

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

func (f *Fs) Delete(path Path) error {
	if !f.Exists(path) {
		return ResourceNotFoundErr
	}
	if err := path.Validate(); err != nil {
		return err
	}
	if err := os.RemoveAll(path.Resolve(f)); err != nil {
		return err
	}

	return nil
}

func (f *Fs) MakeTag(name string, path Path) error {
	return nil
}

func (f *Fs) DeleteTag(name string) error {
	return nil
}
