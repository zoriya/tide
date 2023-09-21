package controllers

import (
	models "github.com/zoriya/tide/api/models"
	services "github.com/zoriya/tide/api/services"
)

type Controller struct {
	Database *services.Database
	Qbittorent    *services.Qbittorent
}

func NewController(db *services.Database, aria2 *services.Qbittorent) *Controller {
	c := new(Controller)
	c.Database = db
	c.Qbittorent = aria2
	return c
}

func (c *Controller) NewItem(newItem models.NewItem) (*models.Item, error) {
	item, err := c.Qbittorent.AddItem(newItem.Uri)
	if err != nil {
		return nil, err
	}
	if newItem.Path != nil {
		item.Path = *newItem.Path
	}
	return c.Database.AddItem(item)
}
