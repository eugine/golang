package main

import (
    "bufio"
    "fmt"
    "os"
)

func main()  {
    keyProFileCount := make(map[string]map[string]int)

    files := os.Args[1:]
    if (len(files) == 0) {
        countLines(os.Stdin, keyProFileCount)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, keyProFileCount)
            f.Close()
        }
        for key, fileCount := range keyProFileCount {
            totalCount := 0
            files := make([]string, len(fileCount))
            for file, count := range fileCount {
                totalCount += count
                files = append(files, file)
            } 
            if totalCount > 1 {
                 fmt.Printf("%d\t%s\t%s\n", totalCount, key, files)
            }
        }
    }
}

func countLines(f *os.File, keyProFileCount map[string]map[string]int)  {
    input := bufio.NewScanner(f)
    for input.Scan() {
        fileCount := keyProFileCount[input.Text()]
        if (fileCount == nil) {
            fileCount = make(map[string]int)
            keyProFileCount[input.Text()] = fileCount
        }
        fileCount[f.Name()]++
    }
}