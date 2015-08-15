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
var existingRssFilePath = flag.String("rss-existing", "", "Path to RSS file to be updated")

func main() {
	flag.Parse()
	if *yamlFilePath == "" {
		fmt.Errorf("You have to provide -yaml parameter!")
		panic("No yaml file path specified")
	}
	existingRss := new(rss.Channel)
	if *existingRssFilePath != "" {
		rssFile, _ := os.Open(*existingRssFilePath)
		defer rssFile.Close()
		rssFileInfo, _ := rssFile.Stat()
		rssContent := make([]byte, rssFileInfo.Size())
		rssReader := bufio.NewReader(rssFile)
		rssReader.Read(rssContent)
		existingRss := new(rss.Channel)
		xml.Unmarshal(rssContent, existingRss)
	}
	datesMap := rss.NewDates(existingRss)
	yamlFile, _ := os.Open(*yamlFilePath)
	defer yamlFile.Close()
	yamlFileInfo, _ := yamlFile.Stat()
	var yamlContent = make([]byte, yamlFileInfo.Size())
	yamlReader := bufio.NewReader(yamlFile)
	yamlReader.Read(yamlContent)
	var yamlObject yaml.Channel
	y.Unmarshal(yamlContent, &yamlObject)
	rssObject := rss.BuildRssFromYaml(yamlObject, 10, datesMap)
	//	fmt.Printf("%+v", rssObject)
	rssXml, _ := xml.Marshal(&rssObject)
	ioutil.WriteFile("/tmp/test.xml", rssXml, os.ModeExclusive)
}
