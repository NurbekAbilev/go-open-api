package fs

import (
	"path"
	"runtime"
)

func RootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(path.Dir(path.Dir(b))))
}
