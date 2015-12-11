package global

import (
	"github.com/DigitalInnovation/vitamns/data"
	"log"
	"os"
	"strconv"
)

var (
	Config ConfigStruct
	Dal    data.DAL
)

type ConfigStruct struct {
	Port     int
	MongoURI string
	DBName   string
	CName    string
	APIKey   string
}

func loadConfig() {

	Config = ConfigStruct{}

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	Config.Port = port
	Config.MongoURI = os.Getenv("MONGOURI")
	Config.DBName = os.Getenv("DBNAME")
	Config.CName = os.Getenv("CNAME")
	Config.APIKey = os.Getenv("APIKEY")
}

func Setup() {
	loadConfig()

	Dal = new(data.Mongodal)

	err := Dal.CreateDataConnection(Config.MongoURI, Config.DBName, Config.CName)
	if err != nil {
		log.Printf("Error connecting to Mongo on cache update: %v", err)
		panic(err)
	}
}
