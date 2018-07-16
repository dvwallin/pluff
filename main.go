package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mgutz/ansi"
	"github.com/mholt/archiver"
)

type (
	config struct {
		TargetFolder string
	}
)

var (
	Cfg               config
	fileList          = "pluff.conf"
	completeFileList  string
	fileListLocations = []string{"~/.", "~/", "/etc/", "./.", "./"}
	err               error
	timestamp         string
	t                 = time.Now()
	content           []string
)

func main() {
	completeFileList = locateFileList()

	flag.StringVar(&Cfg.TargetFolder, "targetFolder", "$HOME/.pluff", "the folder you want the archives to be stored in")
	flag.Parse()

	Cfg.TargetFolder = strings.TrimRight(Cfg.TargetFolder, "/")

	if !fileExists(Cfg.TargetFolder) {
		log.Println("aborted. could not find folder", Cfg.TargetFolder)
		os.Exit(1)
	}

	fmt.Printf("\n\nversion: 0.0\n¤-------------------------------¤\n¤\t\t%s\t\t¤\n¤\t\t\t\t¤\n¤\tDoesn't mean nuthin'\t¤\n¤-------------------------------¤\n\n\n", ansi.Color("PLUFF", "magenta+h"))

	fmt.Printf("%s:\t\t%s\n\n", ansi.Color("Using", "yellow+b"), ansi.Color(completeFileList, "green+b"))

	timestamp = fmt.Sprintf("%d-%02d-%02d_%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	if err != nil {
		log.Println("error getting timestamp", err.Error())
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("error getting hostname. using 'unknown' as name", err.Error())
		hostname = "unknown"
	}
	filename := fmt.Sprintf("%s/%s-%s.tar.xz", Cfg.TargetFolder, hostname, timestamp)
	createArchive(filename)
}

func createArchive(filename string) {
	content, err = readLines(fileList)
	if err != nil {
		log.Println("error fetching files/dirs to be archived", err.Error())
	}
	fmt.Printf("%s:\t", ansi.Color("Compressing", "yellow+b"))
	for _, v := range content {
		fmt.Printf("%s\n\t\t", ansi.Color(v, "green+b"))
	}
	fmt.Printf("\n%s:\t\t%s\n\n\n", ansi.Color("Output", "yellow+b"), ansi.Color(filename, "green+b"))
	err = archiver.TarXZ.Make(filename, content)
	if err != nil {
		log.Println("error compressing files", err.Error())
	}
}

func locateFileList() string {
	var completePaths []string
	for _, v := range fileListLocations {
		tmpFile := fmt.Sprintf("%s%s", v, fileList)
		completePaths = append(completePaths, tmpFile)
		if fileExists(tmpFile) {
			return tmpFile
		}
	}
	log.Println("aborting. Could not find any of these:")
	for _, v := range completePaths {
		fmt.Println(v)
	}
	os.Exit(1)
	return ""
}

func fileExists(file string) (exists bool) {
	if _, err := os.Stat(file); err == nil {
		exists = true
	}
	return
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
