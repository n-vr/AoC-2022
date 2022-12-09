// Package filesystem implements a filesystem representation for the Advent of Code challenge day 7.
package filesystem

import (
	"fmt"
	"strings"
)

// Operation is a type of operation that can be performed on a filesystem.
type Operation string

// Constants for the different types of operations.
const (
	OperationChangeDir Operation = "cd"
	OperationList      Operation = "ls"
)

// Command is an operation from the Advent of Code puzzle input.
type Command struct {
	Operation
	Args []string
}

// Parse puzzle input and return slice of commands.
func ParseCommands(input []string) []Command {
	var cmds []Command
	var currentCommand Command

	for _, line := range input {
		if isCommand(line) {
			// Append the last finished command to cmds before starting the next.
			cmds = append(cmds, currentCommand)

			currentCommand = Command{}
			currentCommand.Operation = parseOperation(line)
			if currentCommand.Operation == OperationChangeDir {
				currentCommand.Args = append(currentCommand.Args, parseDirName(line))
			}
			continue
		}
		currentCommand.Args = append(currentCommand.Args, line)
	}
	cmds = append(cmds, currentCommand)

	// Dont return the first two, since the first is empty, and the seconds one is cd /.
	return cmds[2:]
}

// Returns if a line of puzzle input is a command.
// If it is false, it is not a command, but an output.
func isCommand(input string) bool {
	return input[0] == '$'
}

// Returns what operation a line of puzzle input is.
func parseOperation(input string) Operation {
	if strings.Contains(input, "$ cd") {
		return OperationChangeDir
	} else if strings.Contains(input, "$ ls") {
		return OperationList
	}
	return ""
}

// Returns the directory name from a line of puzzle input.
func parseDirName(input string) string {
	return strings.Split(input, " ")[2]
}

// Parse args into a slice of file pointers.
func (c *Command) parseFiles() (files []*File) {
	for _, fileString := range c.Args {
		if strings.Contains(fileString, "dir ") {
			continue
		}
		files = append(files, NewFileFromString(fileString))
	}
	return
}

// Execute command on the [Filesystem].
func (c *Command) Execute(fs *Filesystem) {
	switch c.Operation {
	case OperationChangeDir:
		fs.ChangeDir(c.Args[0])
	case OperationList:
		fs.AddFiles(c.parseFiles())
	}
}

// Filesystem is a structure that allows you to move around a filesystem.
type Filesystem struct {
	Root       *Directory
	WorkingDir *Directory
}

// Create a new [Filesystem] with an empty root folder
// The working directory also defaults to `/`.
func NewFilesystem() *Filesystem {
	rootDir := &Directory{
		Name:     "/",
		Size:     0,
		Children: make(map[string]*Directory),
	}

	return &Filesystem{
		Root:       rootDir,
		WorkingDir: rootDir,
	}
}

// Change directory in the filesystem.
// The directory will be created if it does not exist.
func (f *Filesystem) ChangeDir(name string) {
	if _, ok := f.WorkingDir.Children[name]; !ok {
		f.WorkingDir.NewSubdir(name)
	}
	f.WorkingDir = f.WorkingDir.Children[name]
}

// Add a file to the current working directory.
func (f *Filesystem) AddFile(file *File) {
	f.WorkingDir.Files = append(f.WorkingDir.Files, file)
}

// Add multiple files to the current working directory.
func (f *Filesystem) AddFiles(files []*File) {
	for _, file := range files {
		f.AddFile(file)
	}
}

// Calculate sizes of all directories in the filesystem.
func (f *Filesystem) GetDirectorySizes() (sizes []int) {
	f.Root.CalculateSize()

	sizes = append(sizes, f.Root.Size)

	f.Root.WalkTree(func(dir *Directory) {
		sizes = append(sizes, dir.Size)
	})

	return
}

// Execute [Command] on the filesystem.
func (f *Filesystem) Execute(cmd Command) {
	cmd.Execute(f)
}

// Directory is a folder in a filesystem.
type Directory struct {
	Name     string
	Size     int
	Parent   *Directory
	Children map[string]*Directory
	Files    []*File
}

// Create a new subdirectory.
// If the subdirectory already exists, no error is thrown.
func (d *Directory) NewSubdir(name string) {
	if _, exists := d.Children[name]; exists {
		return
	}

	d.Children[name] = &Directory{
		Name:   name,
		Parent: d,
		Children: map[string]*Directory{
			"..": d,
		},
	}
}

// Walk the tree from the current directory and call fn on all directories.
func (d *Directory) WalkTree(fn func(dir *Directory)) {
	for name, child := range d.Children {
		if name == ".." {
			continue
		}

		fn(child)
		child.WalkTree(fn)
	}
}

// Calculate the size of the directory, including subdirectories.
// This will not update the size of the parent directories, so it is
// best to manually call this method on the root directory only.
func (d *Directory) CalculateSize() int {
	d.Size = d.totalFileSize()

	for name, child := range d.Children {
		if name == ".." {
			continue
		}

		d.Size += child.CalculateSize()
	}

	return d.Size
}

// Calculate the total size of all files in directory,
// not including subdirectories.
func (d *Directory) totalFileSize() (sum int) {
	for _, file := range d.Files {
		sum += file.Size
	}
	return
}

// A file in the [Filesystem].
type File struct {
	Name string
	Size int
}

// Create a new file from a string representation in the puzzle input.
func NewFileFromString(input string) *File {
	file := &File{}
	_, err := fmt.Sscanf(input, "%d %s", &file.Size, &file.Name)
	if err != nil {
		panic(err)
	}
	return file
}
