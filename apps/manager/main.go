package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type DeliverConfig struct {
	Filename string   `json:"filename"`
	Servers  []string `json:"servers"`
}

func (c *DeliverConfig) Init(f string) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Panic(err)
	}
	if err := json.Unmarshal(data, c); err != nil {
		log.Panic(err)
	}
}

var (
	configFile    string
	deliverConfig DeliverConfig
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.Ldate)
	deliverConfig.Init(configFile)

	dm := &DeliverManager{}
	dm.InitGroups(deliverConfig.Servers)
	dm.InitSessions(deliverConfig.Filename)
	dm.Start()
}

func init() {
	flag.StringVar(&configFile, "config", "deliver.json", "deliver config file")
}
