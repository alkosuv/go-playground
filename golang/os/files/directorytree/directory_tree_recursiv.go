package directorytree

import (
	"os"
	"path/filepath"
)

func DirectoryTreeRecursiv(dir string) ([]string, error) {
	dirEntreis, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	if len(dirEntreis) == 0 {
		return []string{}, nil
	}

	paths := []string{}
	for i := range dirEntreis {
		path := filepath.Join(dir, dirEntreis[i].Name())
		paths = append(paths, path)

		if dirEntreis[i].IsDir() {
			subFiles, err := DirectoryTreeRecursiv(path)
			if err != nil {
				return nil, err
			}
			paths = append(paths, subFiles...)
		}
	}

	return paths, nil
}

// /var/folders/sc/vcztlzh94y90wx_vgglfbn3w0000gn/T/TestDirectoryTreeRecursiv2721829220/001/dir1/file2.txt
// /var/folders/sc/vcztlzh94y90wx_vgglfbn3w0000gn/T/TestDirectoryTreeRecursiv2721829220/001/dir2/dir3/file3.txt
// /var/folders/sc/vcztlzh94y90wx_vgglfbn3w0000gn/T/TestDirectoryTreeRecursiv2721829220/001/file1.txt

// /var/folders/sc/vcztlzh94y90wx_vgglfbn3w0000gn/T/TestDirectoryTreeRecursiv2721829220/001/dir1
// /var/folders/sc/vcztlzh94y90wx_vgglfbn3w0000gn/T/TestDirectoryTreeRecursiv2721829220/001/dir1/file2.txt
// /var/folders/sc/vcztlzh94y90wx_vgglfbn3w0000gn/T/TestDirectoryTreeRecursiv2721829220/001/dir2
// /var/folders/sc/vcztlzh94y90wx_vgglfbn3w0000gn/T/TestDirectoryTreeRecursiv2721829220/001/dir2/dir3
// /var/folders/sc/vcztlzh94y90wx_vgglfbn3w0000gn/T/TestDirectoryTreeRecursiv2721829220/001/dir2/dir3/file3.txt
// /var/folders/sc/vcztlzh94y90wx_vgglfbn3w0000gn/T/TestDirectoryTreeRecursiv2721829220/001/file1.txt
