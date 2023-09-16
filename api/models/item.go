package models

import "time";

type State string
const (
	Stale State = "stale"
	Downloading State = "downloading"
	Seeding State = "seeding"
	Finished State = "finished"
)

type File struct {
	Name string
	Priority int
	Size uint64
	AvailableSize uint64
	Path string
}

type Item struct {
	Id string
	Name string
	State State
	Size uint64
	AvailableSize uint64
	Path string
	AddedDate time.Time
	Files []File
}
