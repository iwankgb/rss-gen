//parse friendy yaml into not so friendly RSS feed!
package main

import y "gopkg.in/yaml.v2"
import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/iwankgb/rss-gen/rss"
	"github.com/iwankgb/rss-gen/yaml"
	"io/ioutil"
	"os"
)

type arguments struct {
	yamlFilePath *string
	rssFilePath  *string
	itemsCount   *int
}

var args = arguments{
	flag.String("yaml", "", "Path to yaml file"),
	flag.String("rss", "", "Path to RSS file to be updated"),
	flag.Int("count", 10, "Number of items to be included in the feed"),
}

func main() {
	flag.Parse()
	validateArguments(&args)
	rssFile, _ := os.Open(*args.rssFilePath)
	defer rssFile.Close()
	existingRss := prepareExistingRssChannel(rssFile)
	dates := rss.NewDates(existingRss)
	yamlFile, _ := os.Open(*args.yamlFilePath)
	defer yamlFile.Close()
	yamlObject := prepareYaml(yamlFile)
	rssXml := prepareRss(yamlObject, &args, dates)
	ioutil.WriteFile(*args.rssFilePath, rssXml, 0644)
	fmt.Println("RSS updated successfully")
}

func validateArguments(argument *arguments) {
	if *argument.yamlFilePath == "" {
		fmt.Println("You have to provide -yaml parameter!")
		os.Exit(1)
	}
	if *argument.rssFilePath == "" {
		fmt.Println("You have to provide -rss parameter!")
		os.Exit(2)
	}
}

func prepareExistingRssChannel(rssFile *os.File) *rss.Channel {
	_, isRssFileValid := rssFile.Stat()
	existingRss := new(rss.Channel)
	if isRssFileValid == nil {
		rssContent := readFile(rssFile)
		existingRss = new(rss.Channel)
		xml.Unmarshal(rssContent, existingRss)
	}
	return existingRss
}

func readFile(file *os.File) []byte {
	fileInfo, _ := file.Stat()
	var fileContent = make([]byte, fileInfo.Size())
	fileReader := bufio.NewReader(file)
	fileReader.Read(fileContent)
	return fileContent
}

func prepareYaml(file *os.File) *yaml.Channel {
	yamlContent := readFile(file)
	return getYamlObject(&yamlContent)
}

func getYamlObject(yamlContent *[]byte) *yaml.Channel {
	yamlObject := new(yaml.Channel)
	y.Unmarshal(*yamlContent, yamlObject)
	return yamlObject
}

func prepareRss(yamlObject *yaml.Channel, args *arguments, dates rss.DateResolver) []byte {
	rssObject := getRssObject(yamlObject, args, dates)
	rssXml, _ := xml.Marshal(&rssObject)
	return rssXml
}

func getRssObject(yamlObject *yaml.Channel, args *arguments, dates rss.DateResolver) *rss.Channel {
	rssBuilder := rss.NewBuilder(dates, args.itemsCount)
	rssBuilder.BuildRssFromYaml(yamlObject)
	return rssBuilder.GetRss()
}
