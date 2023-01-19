package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func addHexLineBlock(curLine *[]byte, result *string) {
	for i := 0; i < len(*curLine); i++ {
		*result += fmt.Sprintf("%02x ", (*curLine)[i])
	}
	*result += "\r\n"
}

func convertFile(fpIn string) {
	// open file
	fileIn, err := os.Open(fpIn)
	if err != nil {
		// failed to open the file
		panic(err)
	}
	defer fileIn.Close()

	var result string
	var lineIdx int64 = 0

	// loop through each line
	fileScanner := bufio.NewScanner(fileIn)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		curLine := fileScanner.Bytes()
		addHexLineBlock(&curLine, &result)
		lineIdx++
		fmt.Printf("\rline = %d", lineIdx)

	}
	fmt.Println()

	// write to file
	var fpOut string = fpIn + ".FileToHexLines.txt"
	fileOut, err := os.Create(fpOut)
	if err != nil {
		// failed to create a file
		panic(err)
	}
	writeSize, err := fileOut.WriteString(result)
	if err != nil {
		// failed to write to the file
		panic(err)
	}
	fmt.Printf("%d bytes written to %s\n", writeSize, fpOut)
}

func main() {
	// Title
	fmt.Println("=== File To Hex Lines ===")

	// Expect argument count to be 2.
	if len(os.Args) != 2 {
		fmt.Printf("ERROR: len(os.Args)=%d. Expected 2.\n", len(os.Args))
		os.Exit(1)
	}

	// Get file or folder path as the first command line argument
	var pathIn string = os.Args[1]

	// Get file informaiton
	fileInfo, err := os.Stat(pathIn)
	if err != nil {
		// passed command line argument is not a file.
		panic(err)
	}

	if fileInfo.IsDir() {
		// case: folder
		files, err := os.ReadDir(pathIn)
		if err != nil {
			panic(err)
		}
		// loop through files
		for _, file := range files {
			// get file path
			fpIn := filepath.Join(pathIn, file.Name())
			convertFile(fpIn)
		}
	} else {
		// case: file
		convertFile(pathIn)
	}

}
