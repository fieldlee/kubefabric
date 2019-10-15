package utils

import (
	"os"
	"path/filepath"
)

func GetCurrentPath()(string,error){
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	dir = filepath.Join(dir,"kube","config")
	return dir,nil
}
func Int32Ptr(i int32) *int32 { return &i }