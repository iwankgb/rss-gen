//parse friendy yaml into not so friendly RSS feed!
package main

import y "gopkg.in/yaml.v2"
import "flag"
import "os"
import "bufio"
import "github.com/iwankgb/rss-gen/yaml"
import "github.com/iwankgb/rss-gen/rss"
import "encoding/xml"
import "io/ioutil"
import "fmt"

const worldReadable = 0644

var yamlFilePath = flag.String("yaml", "", "Path to yaml file")
var existingRssFilePath = flag.String("rss", "", "Path to RSS file to be updated")
var itemsCount = flag.Int("count", 10, "Number of items to be included in the feed")

func main() {
	flag.Parse()
	if *yamlFilePath == "" {
		fmt.Errorf("You have to provide -yaml parameter!")
		panic("No yaml file path specified")
	}
	if *existingRssFilePath == "" {
		fmt.Errorf("You have to provide -rss parameter!")
		panic("No RSS file path specified")
	}
	existingRss := new(rss.Channel)
	rssFile, _ := os.Open(*existingRssFilePath)
	defer rssFile.Close()
	rssFileInfo, isRssFileValid := rssFile.Stat()
	fmt.Printf("%+v\n", isRssFileValid)
	if isRssFileValid == nil {
		rssFileSize := rssFileInfo.Size()
		rssContent := make([]byte, rssFileSize)
		rssReader := bufio.NewReader(rssFile)
		rssReader.Read(rssContent)
		existingRss = new(rss.Channel)
		xml.Unmarshal(rssContent, existingRss)
		fmt.Printf("%+v\n", existingRss)
	}
	fmt.Printf("%+v\n", existingRss.Items)
	datesMap := rss.NewDates(existingRss)
	fmt.Printf("%+v\n", datesMap)
	yamlFile, _ := os.Open(*yamlFilePath)
	defer yamlFile.Close()
	yamlFileInfo, _ := yamlFile.Stat()
	var yamlContent = make([]byte, yamlFileInfo.Size())
	yamlReader := bufio.NewReader(yamlFile)
	yamlReader.Read(yamlContent)
	var yamlObject yaml.Channel
	y.Unmarshal(yamlContent, &yamlObject)
	rssObject := rss.BuildRssFromYaml(yamlObject, *itemsCount, datesMap)
	//	fmt.Printf("%+v", rssObject)
	rssXml, _ := xml.Marshal(&rssObject)
	ioutil.WriteFile(*existingRssFilePath, rssXml, 0644)
}
