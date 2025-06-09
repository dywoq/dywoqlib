package filesystem

import (
	"bufio"
	"os"
)

// FileLine represents a line with its content of file.
type FileLine struct {
	content string
	line    int
}

// File represents a file with the content of lines.
type File struct {
	path    string
	content []FileLine
}

// Path returns the path of the file.
func (f File) Path() string {
	return f.path
}

// Content returns a slice of the file content.
func (f File) Content() []FileLine {
	return f.content
}

// NewFile creates new instance of File without pointer to it.
// The function can return the error.
func NewFile(path string) (f File, err error) {
	content, err := readFile(path)
	if err != nil {
		return
	}
	f = File{path, content}
	return
}

func readFile(path string) (result []FileLine, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	num := 0
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, FileLine{line, num})
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}
