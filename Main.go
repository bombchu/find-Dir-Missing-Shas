package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

func main() {
	
	filePath := "/home/bombchu/mnt-points/1/bak/master0/no-clobber/dta/"
	
	var result0, result1 []string
	
	// ioutil.ReadDir instead of Java Files.walk to get file listings.
	files, _ := ioutil.ReadDir(filePath)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".mp4") {
			result0 = append(result0, filepath.Join(filePath, f.Name()))
		}
	}
	
	files, _ = ioutil.ReadDir(filePath)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".rar") {
			result1 = append(result1, filepath.Join(filePath, f.Name()))
		}
	}
	
	result0 = append(result0, result1...)
	
	var result2 []string
	
	files, _ = ioutil.ReadDir(filePath)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".sha256") {
			result2 = append(result2, filepath.Join(filePath, f.Name()))
		}
	}

	pattern := regexp.MustCompile(`/.+/`)

	var finalList []string

	// synchronize concurrent access with sync.WaitGroup instead of streams.
	var wg sync.WaitGroup
	for _, item0 := range result0 {
		wg.Add(1)
		go func(item0 string) {
			defer wg.Done()
			for _, item2 := range result2 {
				if !strings.Contains(item2, item0) {
					matches := pattern.FindAllString(item0, -1)
					for _, match := range matches {
						finalList = append(finalList, match) 
					}
				}
			}
		}(item0)
	}
	wg.Wait()
	
	// map instead of Java TreeSet for unique values.
	unique := make(map[string]bool)
	for _, entry := range finalList {
		unique[entry] = true
	}
	
	for key, _ := range unique {
		fmt.Println(key)
	}
}
