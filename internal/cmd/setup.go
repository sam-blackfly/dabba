package cmd

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/codeclysm/extract/v3"
	"github.com/sam-blackfly/dabba/internal/colors"
	"github.com/sam-blackfly/dabba/internal/paths"
	"github.com/spf13/cobra"
)

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the environment",
	Run: func(cmd *cobra.Command, args []string) {
		setup(args)
	},
}

const downloadURL = "https://dl-cdn.alpinelinux.org/alpine/v3.14/releases/x86_64/alpine-minirootfs-3.14.2-x86_64.tar.gz"

func setup(args []string) {
	fsDirPath := path.Join(paths.FileSystemsPath, "alpine")

	if dirExists(fsDirPath) {
		log.Printf("skipping download")
		return
	}

	if !dirExists(paths.TempPath) {
		os.MkdirAll(paths.TempPath, 0755)
	}

	var filePath = path.Join(paths.TempPath, "alpine.tar.gz")
	if !fileExists(filePath) {
		downloadFile(downloadURL, filePath)
		log.Printf("downloaded filesystem archive to %v\n", colors.Info(filePath))
	}

	data, _ := ioutil.ReadFile(filePath)
	buffer := bytes.NewBuffer(data)

	log.Printf("extracting archive to %v\n", colors.Info(fsDirPath))
	extract.Gz(context.Background(), buffer, fsDirPath, nil)

	log.Printf("extracted archive to %v\n", colors.Info(fsDirPath))
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return !os.IsNotExist(err) && !info.IsDir()
}

// dirExists checks if a directory exists and is not a directory before we
// try using it to prevent further errors.
func dirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	return !os.IsNotExist(err) && info.IsDir()
}

func downloadFile(url string, dest string) {
	out, err := os.Create(dest)
	if err != nil {
		log.Fatalf("could not create file at location %v\n", dest)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("could not download file %v\n", url)
	}
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("could not download file %v\n", url)
	}

	log.Printf("wrote %v bytes to %v", n, dest)
}
