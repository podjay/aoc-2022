package puzzle

import (
	"errors"
	"fmt"
)

type Line interface {
	ApplyInputToDirectory(*Directory) (*Directory, error)
}

type ChangeDirectory struct {
	target string
}

func (c ChangeDirectory) ApplyInputToDirectory(directory *Directory) (*Directory, error) {
	target := c.target
	if target == "/" {
		rootDir := directory
		for rootDir.parentDirectory != nil {
			rootDir = rootDir.parentDirectory
		}
		return rootDir, nil
	}

	if target == ".." {
		return directory.parentDirectory, nil
	}

	if directory, ok := directory.subdirectories[target]; !ok {
		return nil, errors.New(fmt.Sprintf("no directory with name %s in %s", target, directory.parentDirectory.name))
	} else {
		return directory, nil
	}
}

func NewChangeDirectory(target string) ChangeDirectory {
	return ChangeDirectory{
		target: target,
	}
}

type ListFiles struct {
}

func (l ListFiles) ApplyInputToDirectory(directory *Directory) (*Directory, error) {
	// Does nothing
	return directory, nil
}

func NewListFiles() ListFiles {
	return ListFiles{}
}

type FileInput struct {
	name string
	size int64
}

func (f FileInput) ApplyInputToDirectory(directory *Directory) (*Directory, error) {
	directory.files = append(directory.files, NewFile(f.name, f.size))
	return directory, nil
}

func NewFileInput(name string, size int64) FileInput {
	return FileInput{
		name: name,
		size: size,
	}
}

func (d DirInput) ApplyInputToDirectory(directory *Directory) (*Directory, error) {
	directory.subdirectories[d.name] = NewDirectory(d.name, directory)
	return directory, nil
}

type DirInput struct {
	name string
}

func NewDirInput(name string) DirInput {
	return DirInput{name: name}
}
