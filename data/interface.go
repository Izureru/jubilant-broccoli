package data

import (
	"github.com/DigitalInnovation/vitamns/entities"
)

type DAL interface {
	CreateDataConnection(url, database, collection string) error
	GetAllItems() ([]entities.Item, error)
	GetItemInfo(upc string) (entities.Item, error)
}
