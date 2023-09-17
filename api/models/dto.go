package models

type ItemType string

const (
	Torrent ItemType = "torrent"
	Magnet  ItemType = "magnet"
	Direct  ItemType = "direct"
)

type NewItem struct {
	Uri  string
	Type *ItemType
	Path *string
}
