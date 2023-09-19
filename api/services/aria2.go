package services

import (
	"fmt"
	"log"
	"os"
	"time"

	models "tide/api/models"

	"github.com/siku2/arigo"
)

type Service interface {
	AddItem(item string)
	List() []models.Item
}

type Aria2 struct {
	client *arigo.Client
	token  string
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

func toState(status arigo.DownloadStatus, percent uint) models.State {
	switch status {
	case arigo.StatusActive:
		if percent == 100 {
			return models.Seeding
		}
		return models.Downloading
	case arigo.StatusWaiting:
		return models.Stale
	case arigo.StatusPaused:
		return models.Paused
	case arigo.StatusError:
		return models.Errored
	case arigo.StatusCompleted:
		return models.Finished
	case arigo.StatusRemoved:
		fallthrough
	default:
		return models.Unknown
	}
}

func (x *Aria2) AddItem(uri string) (*models.Item, error) {
	id, err := x.client.AddURI([]string{uri}, nil)
	if err != nil {
		return nil, err
	}
	status, err := id.TellStatus()
	if err != nil {
		log.Println("Tell status error")
		return nil, err
	}

	percent := status.CompletedLength / status.TotalLength
	files := make([]models.File, 0, len(status.Files))
	for _, file := range status.Files {
		priority := models.Medium
		if !file.Selected {
			priority = models.None
		}

		files = append(files, models.File{
			Index:         uint(file.Index),
			Name:          "???",
			Path:          "???",
			Priority:      priority,
			Size:          uint64(file.Length),
			AvailableSize: uint64(file.CompletedLength),
		})
	}
	return &models.Item{
		Id: id.GID,

		Name:      "??",
		Path:      "??",
		AddedDate: time.Now(),
		Files:     files,

		State:         toState(status.Status, percent),
		Size:          uint64(status.TotalLength),
		AvailableSize: uint64(status.CompletedLength),
		Percent:       percent,
		UploadedSize:  uint64(status.UploadLength),
		BitField:      status.BitField,
		DownloadSpeed: status.DownloadSpeed,
		UploadSpeed:   status.UploadSpeed,
		SeedCount:     status.NumSeeders,
		Connections:   status.Connections,
		ErrorMessage:  &status.ErrorMessage,
	}, nil
}

func (x *Aria2) List() []models.Item {
	return make([]models.Item, 0)
}
