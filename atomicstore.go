package atomicstore

import (
	"io/ioutil"
	"os"
)

// Prefix contains the prefix to be used for temporary files.
var Prefix = "astore"

// AtomicWrite writes the data to a temporary file, moving it to the
// filename on success.
func AtomicWrite(filename string, data []byte, perms os.FileMode) error {
	tmpFile, err := ioutil.TempFile("", Prefix)
	if err != nil {
		return err
	}

	tmpName := tmpFile.Name()
	tmpFile.Close()

	err = ioutil.WriteFile(tmpName, data, perms)
	if err != nil {
		os.Remove(tmpName)
		return err
	}
	return os.Rename(tmpName, filename)
}
