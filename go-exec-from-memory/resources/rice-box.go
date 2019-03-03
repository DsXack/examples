package resources

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    ".gitignore",
		FileModTime: time.Unix(1542921683, 0),
		Content:     string("*\n!.gitignore\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "embeded.gz",
		FileModTime: time.Unix(1542924826, 0),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1542924827, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // ".gitignore"
			file3, // "embeded.gz"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`data`, &embedded.EmbeddedBox{
		Name: `data`,
		Time: time.Unix(1542924827, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			".gitignore": file2,
			"embeded.gz": file3,
		},
	})
}