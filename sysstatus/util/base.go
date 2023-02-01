package util

import "github.com/dustin/go-humanize"

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s int64) string {
	return humanize.IBytes(uint64(s))
}
