package filesys_test

import (
	"testing"

	"github.com/roman-kartash/gocommons/filesys"
	"github.com/stretchr/testify/require"
)

func TestIsDirectory(t *testing.T) {
	t.Parallel()

	require.NoError(t, filesys.IsDirectory("test_data/dir"))
	require.ErrorIs(t, filesys.IsDirectory("test_data/dir/empty_file"), filesys.ErrPathIsNotDirectory)
	require.ErrorIs(t, filesys.IsDirectory("test_data/none"), filesys.ErrPathNotExists)
}
