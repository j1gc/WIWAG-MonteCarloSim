package db

import "warehouse/simulation"

type Database interface {
	Connect() error
	Close() error
	InsertSimulationsInDB(sims []simulation.SimResults) error
}
