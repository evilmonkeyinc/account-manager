package datastore

import "context"

type DataStore interface {
	Delete(ctx context.Context, collection, id string) error
	Retreive(ctx context.Context, collection, id string, ptr interface{}) error
	RetreiveMany(ctx context.Context, collection string, page, limit int, ptr interface{}) ([]interface{}, error)
	Create(ctx context.Context, collection, id string, ptr interface{}) error
	Update(ctx context.Context, collection, id string, ptr interface{}) error
}
