package db

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/sync/errgroup"
	"os"
	"warehouse/simulation"
	"warehouse/utils"
)

type MongoDB struct {
	conn *mongo.Client
}

func (m *MongoDB) Connect() error {
	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGODB_ADDRESS")))
	if err != nil {
		return err
	}
	m.conn = client
	return nil
}

func (m *MongoDB) Close() error {
	err := m.conn.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) InsertSimulationsInDB(sims []simulation.SimResults) error {
	simAmount := len(sims)
	batchSize := utils.CalculateOptimalBatchSizeForCPUTasks(simAmount)

	collection := m.conn.Database("testing").Collection("simulation")
	// The speed is improved by disabling the insert order
	opts := options.InsertMany().SetOrdered(false)

	// Uses an error group as a way to return errors from a go function
	var eg errgroup.Group

	for batchStart := 0; batchStart < simAmount; batchStart += batchSize {
		// check if the last batchSize is bigger than the remaining simulation amount
		batchEnd := batchStart + batchSize
		if batchEnd > simAmount {
			batchEnd = simAmount
		}

		// starts a go routine that can return errors
		eg.Go(func() error {
			_, err := collection.InsertMany(context.TODO(), sims[batchStart:batchEnd], opts)
			if err != nil {
				return err
			}
			return nil
		})
	}
	return eg.Wait()
}
