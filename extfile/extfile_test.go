package extfile

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestAppendToFile(t *testing.T) {
	file := filepath.Join(t.TempDir(), "test.txt")

	err := AppendToFile(file, []string{"Hello World!\n"})
	require.NoError(t, err)

	err = AppendToFile(file, []string{"H", "e", "l", "l", "o", " ", "Daniel!"})
	require.NoError(t, err)

	content, err := os.ReadFile(file)
	assert.Contains(t, string(content), "Hello World!\nHello Daniel!")
}

func TestFile2Base64(t *testing.T) {
	file := filepath.Join(t.TempDir(), "test.txt")

	err := AppendToFile(file, []string{"Hello World!\n"})
	require.NoError(t, err)

	content, err := File2Base64(file)
	assert.Contains(t, string(content), "SGVsbG8gV29ybGQhCg==")
}
