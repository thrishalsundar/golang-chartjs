package services

import (
	"context"
	"dhlabs/backend/models"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactServicePasser struct {
	transactions *mongo.Collection
	ctx          context.Context
}

func TransConstruct(transactions *mongo.Collection, ctx context.Context) TransactService {
	return &TransactServicePasser{
		transactions: transactions,
		ctx:          ctx,
	}
}

func (t *TransactServicePasser) NewTrans(transaction *models.Transact) error {
	ts := time.Now()
	transaction.ID = primitive.NewObjectIDFromTimestamp(ts)
	_, err := t.transactions.InsertOne(t.ctx, transaction)
	return err
}

func (t *TransactServicePasser) GetTrans() ([]*models.Transact, error) {
	var transacs []*models.Transact
	cursor, err := t.transactions.Find(t.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(t.ctx) {
		var elem models.Transact
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		transacs = append(transacs, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(t.ctx)

	if len(transacs) == 0 {
		return nil, errors.New("no documents (I did this too)")
	}
	return transacs, nil
}
