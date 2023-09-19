package models

import "time"

type State string

const (
	Unknown     State = "unknown"
	Stale       State = "stale"
	Downloading State = "downloading"
	Seeding     State = "seeding"
	Finished    State = "finished"
	Paused      State = "paused"
	Errored     State = "errored"
)

type Priority string

const (
	None   Priority = "none"
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type File struct {
	Index         uint     `json:"index"`
	Name          string   `json:"name"`
	Priority      Priority `json:"priority"`
	Size          uint64   `json:"size"`
	AvailableSize uint64   `json:"availableSize"`
	Path          string   `json:"path"`
}

type Item struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	AddedDate time.Time `json:"addedDate"`
	Files     []File    `json:"files"`

	State         State  `json:"state"`
	Size          uint64 `json:"size"`
	AvailableSize uint64 `json:"availableSize"`
	Percent       uint   `json:"percent"`
	UploadedSize  uint64 `json:"uploadedSize"`
	// Hexadecimal representation of the download progress.
	// The highest bit corresponds to the piece at index 0. Any set bits indicate loaded pieces,
	// while unset bits indicate not yet loaded and/or missing pieces.
	// Any overflow bits at the end are set to zero.
	// When the download was not started yet, this will be an empty string.
	BitField      string `json:"bitfield"`
	DownloadSpeed uint   `json:"downloadSpeed"`
	UploadSpeed   uint   `json:"uploadSpeed"`
	SeedCount     uint   `json:"seedCount"`
	Connections   uint   `json:"connections"`

	ErrorMessage *string `json:"errorMessage"`
}
