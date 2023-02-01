package templates

import (
	"embed"
	"errors"
	staticF "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

//go:embed dashboard/* statics/*
var TemplatesEmbed embed.FS

func StaticServer() gin.HandlerFunc {
	static1 := &ServeFileSystem{
		TemplatesEmbed,
		"statics",
	}
	return staticF.Serve("/css", static1)
}

type ServeFileSystem struct {
	e    embed.FS
	path string
}

type File struct {
	name string
	fs.File
}

func (f *File) Readdir(count int) ([]fs.FileInfo, error) {
	ff, ok := f.File.(fs.ReadDirFile)
	if !ok {
		return nil, &fs.PathError{Op: "readdir", Path: f.name, Err: errors.New("not implemented")}
	}
	fileList, err := ff.ReadDir(count)
	if err != nil {
		return nil, err
	}
	rspList := []fs.FileInfo{}
	for _, v := range fileList {
		temp, err := v.Info()
		if err != nil {
			return nil, err
		}
		rspList = append(rspList, temp)
	}
	return rspList, nil
}

func (f *File) Seek(offset int64, whence int) (int64, error) {
	ff, ok := f.File.(io.Seeker)
	if !ok {
		return 0, &fs.PathError{Op: "Seek", Path: f.name, Err: errors.New("not implemented")}
	}
	return ff.Seek(offset, whence)
}

func (c *ServeFileSystem) Open(name string) (http.File, error) {
	name = path.Join(c.path, name)
	f, err := c.e.Open(name)
	if err != nil {
		return nil, err
	}
	ff := File{
		name: name,
		File: f,
	}
	return &ff, nil
}

func (c *ServeFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {

		p = path.Join(c.path, p)
		f, err := c.e.Open(p)
		if err != nil {
			return false
		}
		err = f.Close()
		return err == nil
	}
	return false
}
