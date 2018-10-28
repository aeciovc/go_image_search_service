package imagesearch

import (
	"log"
	"flag"
	
	"github.com/aeciovc/gonf"
)

var config Configuration

type Configuration struct {
	Service  Service `json:"Service"`
}

type Service struct {
	Name string `json:"Name"`
	Queue string `json:"Queue"`
	Broker string `json:"Broker"`
}

var configPath *string

func init(){
	configPath = flag.String("config", "./config.json", "path of config file")
}

func LoadConfigs() Configuration{

	//File Config
	flag.Parse()

	log.Println("[config] Loading from ", *configPath)

	// Load the configuration file
	err := gonf.Load(*configPath, &config)
	if err != nil {
		log.Fatal("[config] Error loading configs: ", err)
	}

	return Config()
}

func Config() Configuration{
	return config
}