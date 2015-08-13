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

const worldReadable = 0644

var yamlFilePath = flag.String("yaml", "", "Path to yaml file")

func main() {
	flag.Parse()
	yamlFile, _ := os.Open(*yamlFilePath)
	defer yamlFile.Close()
	yamlFileInfo, _ := yamlFile.Stat()
	var yamlContent = make([]byte, yamlFileInfo.Size())
	yamlReader := bufio.NewReader(yamlFile)
	yamlReader.Read(yamlContent)
	var yamlObject yaml.Channel
	y.Unmarshal(yamlContent, &yamlObject)
	rssObject := rss.BuildRssFromYaml(yamlObject, 10)
	rssXml, _ := xml.Marshal(&rssObject)
	ioutil.WriteFile("/tmp/test.xml", rssXml, os.ModeExclusive)
}
