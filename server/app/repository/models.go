package repository

import (
	"context"
	"server-go/models"
)

type Model interface {
	Save(ctx context.Context, data models.CurrencyRate) error
}
