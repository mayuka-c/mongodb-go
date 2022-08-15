package main

import (
	"context"

	"github.com/mayuka-c/mongodb-go/internal/config"
	"github.com/mayuka-c/mongodb-go/internal/models"
	"github.com/mayuka-c/mongodb-go/internal/service"
)

func main() {
	ctx := context.Background()

	item := models.CreateContentDataReq{
		Title:       "t1",
		Description: "d1",
		Services:    []string{"s1", "s2"},
		Roles:       []string{"r1", "r2"},
		ContentType: "video",
		LocationURL: "url-1",
	}

	dbConfig := config.GetDBConfig(ctx)
	svc := service.NewService(ctx, dbConfig)
	svc.AddSingleItem(ctx, item)
}
