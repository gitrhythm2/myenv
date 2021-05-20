// https://github.com/eihigh/filetestをベースとしています。
package file

import (
	"os"
)

func Exests(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsFile(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && !stat.IsDir()
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.IsDir()
}
