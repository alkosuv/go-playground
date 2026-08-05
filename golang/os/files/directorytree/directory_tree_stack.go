package directorytree

import (
	"os"
	"path/filepath"

	"github.com/alkosuv/go-playground/structs-data/stack"
)

func DirectoryTreeStack(dir string) ([]string, error) {
	paths := []string{}
	roots := stack.NewGenericStack[string]()
	roots.Push(dir)

	for !roots.IsEmpty() {
		dir, ok := roots.Pop()
		if !ok {
			continue
		}

		dirEntries, err := os.ReadDir(dir)
		if err != nil {
			return nil, err
		}

		for i := range dirEntries {
			path := filepath.Join(dir, dirEntries[i].Name())
			paths = append(paths, path)

			if dirEntries[i].IsDir() {
				roots.Push(path)
			}
		}
	}

	return paths, nil
}
