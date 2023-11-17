package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Scheduler struct {
	Cron       *cron.Cron
	Collection *mongo.Collection
}

func NewScheduler(collection *mongo.Collection) Scheduler {
	return Scheduler{
		Cron: cron.New(),
		Collection: collection,
	}
}

func (s Scheduler) ClearDatabase() {
	s.Cron.AddFunc("0 0 * * *", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()
		
		result, err := s.Collection.DeleteMany(ctx, bson.M{})
		if err != nil {
			fmt.Println("Cron job failed:", err.Error())
		}
		fmt.Printf("Cron scheduler deleted %v entries\n", result.DeletedCount)
	})
}

func (s Scheduler) StartCronJob() {
	s.ClearDatabase()
	s.Cron.Start()
}