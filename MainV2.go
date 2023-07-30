package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"sync"
)

func main() {
	
	filePath := "/home/bombchu/mnt-points/1/bak/master0/no-clobber/dta/"
	
	var result0, result1 []string
	
	files, _ := ioutil.ReadDir(filePath)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".mp4" {
			result0 = append(result0, filePath + f.Name())
		}
	}
	
	files, _ = ioutil.ReadDir(filePath)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".rar" {
			result1 = append(result1, filePath + f.Name())
		}
	}
	
	result0 = append(result0, result1...)
	
	var result2 []string
	
	files, _ = ioutil.ReadDir(filePath)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".sha256" {
			result2 = append(result2, filePath + f.Name())
		}
	}
	
	var finalList []string
	var wg sync.WaitGroup
	for _, item0 := range result0 {
		wg.Add(1)
		go func(item0 string) {
			defer wg.Done()
			match := false
			for _, item2 := range result2 {
				if filepath.Base(item0) == filepath.Base(item2) {
					match = true
					break
				}
			}
			if !match {
				r := regexp.MustCompile(`/.+/`)
				matches := r.FindAllString(item0, -1)
				for _, m := range matches {
					finalList = append(finalList, m)
				}
			}
		}(item0)
	}
	wg.Wait()
	
	for _, x := range finalList {
		fmt.Println(x)
	}
}
