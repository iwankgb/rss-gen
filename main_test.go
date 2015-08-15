package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestRssGeneration(t *testing.T) {
	tmpFile := createTemporaryFile()
	defer tmpFile.Close()
	defer os.Remove("./example.xml")
	commandOutputString := runCommand(t)
	assertCommandOutput(t, &commandOutputString)
	assertNewFile(t, tmpFile)
}

func createTemporaryFile() *os.File {
	testRssFile, _ := os.Open("resources/example.xml")
	defer testRssFile.Close()
	tmpFile, _ := os.Create("./example.xml")
	io.Copy(bufio.NewWriter(tmpFile), bufio.NewReader(testRssFile))
	tmpFile.Sync()
	return tmpFile
}

func runCommand(t *testing.T) string {
	cmd := exec.Command("./rss-gen", "-yaml=resources/example.yaml", "-rss=example.xml")
	commandOutput, _ := cmd.Output()
	return fmt.Sprintf("%s", commandOutput)

}

func assertCommandOutput(t *testing.T, commandOutput *string) {
	if *commandOutput != "RSS updated successfully\n" {
		t.Errorf("Wrong command output: '%v'\n", commandOutput)
	}
}

func assertNewFile(t *testing.T, tmpFile *os.File) {
	tmpFile.Sync()
	rssReader := bufio.NewReader(tmpFile)
	rssBytes := make([]byte, 1178)
	rssReader.Read(rssBytes)
	rssString := fmt.Sprintf("%s", rssBytes)
	if !strings.Contains(rssString, "a135f7d0341d8444adc19b20f330fe280a2c950b81f9068d1a1139cf5ebd37791cb5e5245d7ba43ff9971756f6fd19d523029a7ca73a561defcfae8a31eba796") {
		t.Error("File does not containt second item guid")
	}
}
