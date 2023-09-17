package controllers

import (
	models "tide/api/models"
	services "tide/api/services"
)

type Controller struct {
	Database *services.Database
	Aria2    *services.Aria2
}

func NewController(db *services.Database, aria2 *services.Aria2) *Controller {
	c := new(Controller)
	c.Database = db
	c.Aria2 = aria2
	return c
}

func (c *Controller) NewItem(newItem models.NewItem) (*models.Item, error) {
	item, err := c.Aria2.AddItem(newItem.Uri)
	if err != nil {
		return nil, err
	}
	if newItem.Path != nil {
		item.Path = *newItem.Path
	}
	return c.Database.AddItem(item)
}
