package models

import (
	"context"

	"cloud.google.com/go/datastore"
)

type Book struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" datastore:"Name,noindex"`
	Category int    `json:"category" datastore:"Category"`
	ImageUrl string `json:"image_url" datastore:"ImageUrl,noindex"`
}

func (client *ModelClient) ListBook(ctx context.Context) ([]*Book, error) {
	var books []*Book

	query := datastore.NewQuery("Book")
	keys, err := client.dsClient.GetAll(ctx, query, &books)
	if err != nil {
		return nil, err
	}

	for i, key := range keys {
		books[i].ID = key.ID
	}

	return books, nil
}

func (client *ModelClient) GetBook(ctx context.Context, bookID int64) (*Book, error) {
	k := datastore.IDKey("Book", bookID, nil)
	book := new(Book)
	if err := client.dsClient.Get(ctx, k, book); err != nil {
		return nil, err
	}

	book.ID = k.ID
	return book, nil
}

func (client *ModelClient) CreateBook(ctx context.Context, book *Book) error {
	newKey := datastore.IncompleteKey("Book", nil)
	if _, err := client.dsClient.Put(ctx, newKey, book); err != nil {
		return err
	}

	return nil
}

func (client *ModelClient) EditBook(ctx context.Context, book *Book) (*Book, error) {
	k := datastore.IDKey("Book", book.ID, nil)
	b := new(Book)

	_, err := client.dsClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		if err := tx.Get(k, b); err != nil {
			return err
		}

		b.Name = book.Name
		b.Category = book.Category
		_, err := tx.Put(k, b)

		return err
	})

	b.ID = k.ID
	return b, err
}

func (client *ModelClient) DeleteBook(ctx context.Context, bookID int64) error {
	k := datastore.IDKey("Book", bookID, nil)
	return client.dsClient.Delete(ctx, k)
}
