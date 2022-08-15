package service

import (
	"context"
	"fmt"

	"github.com/mayuka-c/mongodb-go/internal/config"
	"github.com/mayuka-c/mongodb-go/internal/db"
	"github.com/mayuka-c/mongodb-go/internal/models"
)

type service struct {
	db db.InterfaceDB
}

func NewService(ctx context.Context, dbConfig config.DBConfiguration) *service {
	return &service{
		db: db.NewDBConnection(ctx, dbConfig),
	}
}

func (s *service) AddSingleItem(ctx context.Context, input models.CreateContentDataReq) {
	insertOneResult, err := s.db.InsertOne(ctx, input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted one item: %v\n", insertOneResult)
}

func (s *service) GetItem() {

}
