package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	errorString  = "error"
	tempFileName = "temp.txt"
)

var errorStringBytes = []byte("error")

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Make sure input is being piped to this program. STDIN should be a
	// named pipe.
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe != 0 {
		fmt.Println("Please pipe input to this program.")
		return
	}

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)

	efficientImplementation := flag.Bool("e", false, "Run efficient implementation")
	flag.Parse()

	setupCPUProfile()

	if !(*efficientImplementation) {
		// Naive Implementation 1
		err := implementationOne(reader)
		if err != nil {
			log.Fatalln("An unexpected error occured during Implementation 1: ", err)
		}
	} else {
		// Efficient Implementation 2
		err := implementationTwo(reader)
		if err != nil {
			log.Fatalln("An unexpected error occured during Implementation 2: ", err)
		}
	}

	setupMemoryProfile()
}

func setupCPUProfile() {
	cf, err := os.Create("cpu_profile.out")
	if err != nil {
		fmt.Printf("Error creating cpu profile: file %s\n", err.Error())
	}
	defer cf.Close()
	if err := pprof.StartCPUProfile(cf); err != nil {
		fmt.Printf("Error not start cpu profile: %s\n", err.Error())
	}
	defer pprof.StopCPUProfile()
}

func setupMemoryProfile() {
	f, err := os.Create("heap_profile.out")
	if err != nil {
		fmt.Printf("Error creating heap profile: file %s\n", err.Error())
	}
	defer f.Close()
	runtime.GC()

	err = pprof.WriteHeapProfile(f)
	if err != nil {
		fmt.Printf("Error writing heap profile: %s\n", err.Error())
	}
}

// Simple Quick & Dirty Approach
func implementationOne(reader *bufio.Reader) error {
	delimByte := byte('\n')
	for {
		input, err := reader.ReadBytes(delimByte)

		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		if bytes.Contains(input, []byte(errorString)) {
			fmt.Print(string(input))
		}
	}
	return nil
}

// Struct to hold information regarding temp file processing
// for large streams of data on stdin
type TempFileInfo struct {
	ErrorStringFound bool
	FileCreated      bool
	FilePointer      *os.File
}

// Efficient Approach -- If we have a contiguous stream of text
func implementationTwo(reader *bufio.Reader) error {
	newLineDelim := []byte("\n")
	tfi := &TempFileInfo{}
	buf := make([]byte, 1024)

	for {
		size, err := reader.Read(buf)
		buf = buf[:size]

		if err != nil {
			if err != io.EOF {
				return err
			}

			// A temp file was created but we are at end of stdin
			if tfi.FileCreated {
				if err := tfi.processTempFile([]byte("")); err != nil {
					return err
				}
			}
			break
		}

		// If no newline write to temp file and continue
		if !bytes.Contains(buf, newLineDelim) {
			if err := tfi.initTempFile(buf); err != nil {
				return err
			}
			continue
		}

		// I reorder the call to bytes.Split() after bytes.Contains() since
		// we were using much more memory on the operation below...
		// Split lines since we have at least 1 '\n' character
		lines := bytes.Split(buf, newLineDelim)

		// If the temp file is created then we check if the error string
		// is on the last line or not and then conditionally print the file
		if tfi.FileCreated {
			if err := tfi.processTempFile(lines[0]); err != nil {
				return err
			}

			lines = lines[1:]
			tfi.FileCreated = false
			tfi.ErrorStringFound = false

			// Check to see if string continues without a newline char and
			// setup the temp file for a continued string
			if len(lines) == 1 && (string(lines[0]) != "") {
				if err := tfi.initTempFile(lines[0]); err != nil {
					return err
				}
				continue
			}
		}

		// Check if last line isnt empty -- meaning the string continues so
		// we will setup the temp file for the continued string
		lastIndex := len(lines) - 1
		if string(lines[lastIndex]) != "" {
			if err := tfi.initTempFile(lines[lastIndex]); err != nil {
				return err
			}
			lines = lines[:lastIndex]
		}

		// Check the remaining lines that end with new line chars if they
		// contain the `error` keyword
		for _, line := range lines {
			if bytes.Contains(line, errorStringBytes) {
				fmt.Println(string(line))
			}
		}
	}

	return nil
}

// Initialize a temp file if the file pointer has not been initialized yet
// otherwise write to the temp file
func (tfi *TempFileInfo) initTempFile(currentLine []byte) error {
	tfi.FileCreated = true
	if !tfi.ErrorStringFound {
		tfi.ErrorStringFound = bytes.Contains(currentLine, errorStringBytes)
	}

	if tfi.FilePointer == nil {
		f, err := os.OpenFile(tempFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		tfi.FilePointer = f
	}

	if _, err := tfi.FilePointer.Write(currentLine); err != nil {
		return err
	}

	return nil
}

// Processes the temp file to check for the error string with a final line of
// input and prints and removes the file afterward
func (tfi *TempFileInfo) processTempFile(currentLine []byte) error {
	if !tfi.ErrorStringFound {
		tfi.ErrorStringFound = bytes.Contains(currentLine, errorStringBytes)
	}

	if tfi.ErrorStringFound {
		if _, err := tfi.FilePointer.Write(currentLine); err != nil {
			return err
		}

		if err := readTempFile(tfi.FilePointer); err != nil {
			return err
		}
	}

	if err := removeTempFile(tfi.FilePointer); err != nil {
		return err
	}

	tfi.FilePointer = nil
	return nil
}

//=====================
// Temp File operations
//===

// Prints the temp file to stdout. This is triggered when an instance of a new
// line has been found on a string > 1024 bytes and it contains the text `error`.
func readTempFile(f *os.File) error {
	fr, err := os.Open(f.Name())
	if err != nil {
		return err
	}

	reader := bufio.NewReader(fr)

	buf := make([]byte, 1024)
	for {
		size, err := reader.Read(buf)
		buf = buf[:size]

		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		fmt.Print(string(buf))
	}

	fmt.Println()
	return nil
}

// Closes and removes a temp file to prepare if we need to open a new one
// while processing the rest of the stdin
func removeTempFile(f *os.File) error {
	if err := f.Close(); err != nil {
		return err
	}

	if err := os.Remove(f.Name()); err != nil {
		return err
	}

	return nil
}
