package image_manager

import (
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestLocalMachineImpl_Copy(t *testing.T) {
	src := "it is a test string for copy"
	f, err := os.Create("testFile.txt")
	assert.Equal(t, err, nil)
	n, err := f.Write([]byte(src))
	assert.Equal(t, err, nil)
	assert.Equal(t, n, len([]byte(src)))

	impl := localMachineImpl{}
	err = impl.Copy("newFile.txt", "testFile.txt")
	assert.Equal(t, err, nil)
	err = f.Close()
	assert.Equal(t, err, nil)
}

func TestLocalMachineImpl_Delete(t *testing.T) {
	impl := localMachineImpl{}
	err := impl.Delete("testFile.txt")
	assert.Equal(t, err, nil)
	err = impl.Delete("newFile.txt")
	assert.Equal(t, err, nil)
}
