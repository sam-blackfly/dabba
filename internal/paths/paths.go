package paths

import (
	"os"
	"path"
)

var cwd, _ = os.Getwd()
var TempPath = path.Join(cwd, "tmp")
var FileSystemsPath = path.Join(cwd, "fs")
