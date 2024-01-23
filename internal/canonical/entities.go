package canonical

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrorNotFound = fmt.Errorf("entity not found")
)

type BaseStatus int

const (
	STATUS_ACTIVE BaseStatus = iota
	STATUS_INACTIVE
)

var MapBaseStatus = map[string]BaseStatus{ //ajustar chaves
	"ACTIVE":   STATUS_ACTIVE,
	"INACTIVE": STATUS_INACTIVE,
}

type Product struct {
	ID          string     `bson:"_id"`
	Name        string     `bson:"name"`
	Description string     `bson:"description"`
	Price       float64    `bson:"price"`
	Category    string     `bson:"category"`
	Status      BaseStatus `bson:"status"`
	ImagePath   string     `bson:"image_path"`
}

func NewUUID() string {
	return uuid.New().String()
}

func HandleError(err error) error {
	if errors.Is(err, ErrorNotFound) {
		return err
	}
	return fmt.Errorf("unexpected error occurred %w", err)

}
