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

func (x *Aria2) AddItem(item string) {
	x.client.AddURI([]string{item}, nil)
}

func (x *Aria2) List() []models.Item {
	return make([]models.Item, 0)
}
