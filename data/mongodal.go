package data

import (
	"log"

	"github.com/DigitalInnovation/vitamns/entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongodal struct {
	session        *mgo.Session
	databaseName   string
	collectionName string
}

func (m *Mongodal) CreateDataConnection(url, database, collection string) error {
	m.databaseName = database
	m.collectionName = collection

	s, err := mgo.Dial(url)
	if err != nil {
		log.Printf("There was an error connecting to Mongo: %v", err)
		return err
	}

	m.session = s
	return nil
}

func (m *Mongodal) close() {
	m.session.Close()
}

func (m *Mongodal) GetAllItems() ([]entities.Item, error) {
	var item []entities.Item

	ses := m.session.New()
	defer ses.Close()

	collection := ses.DB(m.databaseName).C(m.collectionName)
	err := collection.Find(nil).All(&item)
	if err != nil {
		log.Printf("Error finding user data in mongo: %v", err)
		return nil, err
	}

	return item, nil
}

func (m *Mongodal) GetItemInfo(upc string) (entities.Item, error) {

	var item entities.Item
	ses := m.session.New()
	defer ses.Close()

	collection := ses.DB(m.databaseName).C(m.collectionName)

	err := collection.Find(bson.M{"upc": upc}).One(&item)
	if err != nil {
		log.Printf("Error finding item info in mongo: %v", err)
		return entities.Item{}, err
	}
	return item, nil
}
