package byteexec

import (
	"os"
	"path"
)

func renameExecutable(orig string) string {
	return orig + ".exe"
}

func pathForRelativeFiles(homeDir string) string {
	return path.Join(os.Getenv("APPDATA"), "byteexec")
}
