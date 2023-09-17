package services

import (
	"fmt"
	"os"
	models "tide/api/models"

	"github.com/siku2/arigo"
)

type Service interface {
	AddItem(item string)
	List() []models.Item
}

type Aria2 struct {
	client *arigo.Client
	token string
}

func NewAria2() (*Aria2, error) {
	p := new(Aria2)
	p.token = os.Getenv("RPC_SECRET")
	c, err := arigo.Dial(fmt.Sprintf("ws://%v/jsonrpc", os.Getenv("ARIA_URI")), p.token)
	if err != nil {
		return nil, err
	}
	p.client = &c
	return p, nil
}

func (x *Aria2) AddItem(uri string) (*models.Item, error) {
	id, err := x.client.AddURI([]string{uri}, nil)
	if err != nil {
		return nil, err
	}
	item := new(models.Item)
	item.Id = id.GID
	// TODO: Download other datas
	return item, nil
}

func (x *Aria2) List() []models.Item {
	return make([]models.Item, 0)
}
