package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ranxx/wire-in-action/foobarbaz"
)

func main() {
	fmt.Println("Hello World !")

	start := time.Now()
	slice1, err := recursion("./")
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Now().Sub(start).String())

	start = time.Now()
	slice2, err := findAllSubDir([]string{"./"})
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Now().Sub(start).String())

	fmt.Println(len(slice1) == len(slice2))
	fmt.Println(foobarbaz.InitializeBaz(context.TODO()))
}

func recursion(d string) ([]string, error) {
	dir, err := os.Open(d)
	if err != nil {
		panic(err)
	}
	stat, err := dir.Stat()
	if err != nil {
		panic(err)
	}
	if !stat.IsDir() {
		return nil, err
	}
	subdirs, err := dir.ReadDir(0)
	if err != nil {
		panic(err)
	}

	ret := make([]string, 0, len(subdirs))
	has := false
	for _, v := range subdirs {
		if !v.IsDir() {
			has = true
		}
	}
	if has {
		ret = append(ret, d)
	}
	for _, v := range subdirs {
		if !v.IsDir() {
			continue
		}
		tmp := strings.TrimSuffix(d, "/") + "/" + v.Name()
		subb, err := recursion(tmp)
		if err != nil {
			panic(err)
		}
		ret = append(ret, subb...)
	}
	return ret, nil
}

func findAllSubDir(dirs []string) ([]string, error) {
	retDirs := make([]string, len(dirs), len(dirs)*1024)
	copy(retDirs, dirs)
	realDirs := make([]string, 0, len(dirs)*1024)
	for index := 0; index < len(retDirs); index++ {
		dir, err := os.Open(retDirs[index])
		if err != nil {
			return nil, err
		}
		stat, err := dir.Stat()
		if err != nil {
			return nil, err
		}
		if !stat.IsDir() {
			realDirs = append(realDirs, retDirs[index])
			continue
		}

		subdirs, err := dir.ReadDir(0)
		if err != nil {
			return nil, err
		}

		for _, v := range subdirs {
			if !v.IsDir() {
				realDirs = append(realDirs, retDirs[index])
				break
			}
		}

		dirName := strings.TrimSuffix(retDirs[index], "/") + "/"
		for _, v := range subdirs {
			if !v.IsDir() {
				continue
			}
			retDirs = append(retDirs, dirName+v.Name())
		}
	}
	return realDirs, nil
}
