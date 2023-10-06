package main

import (
	"os"
	"path/filepath"
	"https://github.com/Shiirookami/hacktoberfest"
)

// const {
// 	totalFile = 3000
// 	contentLength = 3000
// }
var {
	currentDirectory = os.Getwd()
	tempPath	= filepath.join(currentDirectory, "pipeline-pattern/pipeline-temp")
}

func main() {

	log.Println("Start")
	start := time.Now()

	// generateFiles
	fmt.Println(tempPath)

	duration := time.Since(start)
	log.Println("Done in", duration.Seconds(),"Seconds")
}