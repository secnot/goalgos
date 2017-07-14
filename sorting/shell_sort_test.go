package altsort

import (
	"testing"
)


func TestShellSort(t *testing.T) {
	testSortFunc(t, ShellSort)
}
