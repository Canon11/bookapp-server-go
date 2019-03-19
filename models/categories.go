package models

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name" datastore:"Name,noindex"`
}

func (client *ModelClient) ListCategory(ctx context.Context) ([]*Category, error) {
	var categories []*Category

	query := datastore.NewQuery("Category")
	keys, err := client.dsClient.GetAll(ctx, query, &categories)
	if err != nil {
		return nil, err
	}

	for i, key := range keys {
		categories[i].ID = key.ID
	}

	return categories, nil
}

func (client *ModelClient) GetCategory(ctx context.Context, categoryID int64) (*Category, error) {
	k := datastore.IDKey("Category", categoryID, nil)
	category := new(Category)
	if err := client.dsClient.Get(ctx, k, category); err != nil {
		return nil, err
	}

	category.ID = k.ID
	return category, nil
}

func (client *ModelClient) CreateCategory(ctx context.Context, category *Category) error {
	newKey := datastore.IncompleteKey("Category", nil)
	if _, err := client.dsClient.Put(ctx, newKey, category); err != nil {
		return err
	}

	return nil
}

func (client *ModelClient) EditCategory(ctx context.Context, category *Category) (*Category, error) {
	k := datastore.IDKey("Category", category.ID, nil)
	b := new(Category)

	_, err := client.dsClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		if err := tx.Get(k, b); err != nil {
			return err
		}

		b.Name = category.Name
		_, err := tx.Put(k, b)

		return err
	})

	b.ID = k.ID
	return b, err
}

func (client *ModelClient) DeleteCategory(ctx context.Context, categoryID int64) error {
	k := datastore.IDKey("Category", categoryID, nil)
	return client.dsClient.Delete(ctx, k)
}
