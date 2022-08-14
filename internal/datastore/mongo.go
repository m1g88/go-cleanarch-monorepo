package datastore

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDatastore struct {
	*mongo.Database
}

func (m *MongoDatastore) Insert(ctx context.Context, v interface{}) error {
	_, err := m.Database.Collection("hello").InsertOne(ctx, v)

	return err
}
