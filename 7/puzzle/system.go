package puzzle

import (
	"regexp"
)

var command = regexp.MustCompile("^\\$ (\\w+)\\s?(.*)")
var dir = regexp.MustCompile("^dir (\\w+)")
var file = regexp.MustCompile("^(\\d+)(.*\\w?\\.?\\w{3}?)")

func ParseCommand(s string) ([]string, bool) {
	matches := command.FindStringSubmatch(s)
	return matches, len(matches) > 0
}

func ParseDir(s string) ([]string, bool) {
	matches := dir.FindStringSubmatch(s)
	return matches, len(matches) > 0
}

func ParseFile(s string) ([]string, bool) {
	matches := file.FindStringSubmatch(s)
	return matches, len(matches) > 0
}

func NewDirectory(name string, directory *Directory) *Directory {
	return &Directory{
		name:            name,
		parentDirectory: directory,
		subdirectories:  make(map[string]*Directory),
	}
}

type Directory struct {
	name            string
	parentDirectory *Directory
	subdirectories  map[string]*Directory
	files           []File
}

func (d *Directory) Root() *Directory {
	rootDir := d
	for rootDir.parentDirectory != nil {
		rootDir = rootDir.parentDirectory
	}
	return rootDir
}

func (d *Directory) Solve(maxSize int64, candidateDirs *[]Directory) int64 {
	filesInCurrentDir := d.Size()

	subtreeSize := int64(0)
	for _, directory := range d.subdirectories {
		var childDirSize = int64(0)
		childDirSize = directory.Solve(maxSize, candidateDirs)
		subtreeSize += childDirSize
	}

	totalSize := filesInCurrentDir + subtreeSize

	if totalSize < maxSize {
		if candidateDirs != nil {
			*candidateDirs = append(*candidateDirs, *d)
		}
	}
	return totalSize
}

func (d *Directory) Size() int64 {
	total := int64(0)
	for _, f := range d.files {
		total += f.size
	}
	return total
}

type File struct {
	name string
	size int64
}

func NewFile(name string, size int64) File {
	return File{
		name: name,
		size: size,
	}
}
