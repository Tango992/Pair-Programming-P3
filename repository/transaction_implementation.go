package repository

import (
	"context"
	"errors"
	"pair-programming/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionImplementation struct {
	Collection *mongo.Collection
}

func NewTransactionRepository(collection *mongo.Collection) TransactionImplementation {
	return TransactionImplementation{
		Collection: collection,
	}
}

func (t TransactionImplementation) Post(data *models.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	res, err := t.Collection.InsertOne(ctx, *data)
	if err != nil {
		return err
	}
	
	data.Id =  res.InsertedID.(primitive.ObjectID)
	return nil
}

func (t TransactionImplementation) GetAll() ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	
	cursor, err := t.Collection.Find(ctx, bson.M{})
	if err != nil {
		return []models.Transaction{}, err
	}
	
	var transactions []models.Transaction
	if err := cursor.All(ctx, &transactions); err != nil {
		return []models.Transaction{}, err
	}
	
	return transactions, nil
}

func (t TransactionImplementation) GetById(transactionId string) (models.Transaction, error) {
	objectID, err := primitive.ObjectIDFromHex(transactionId)
	if err != nil {
		return models.Transaction{}, nil
	}
	filter := bson.M{"_id ": objectID}

	var transaction models.Transaction

	
	err = t.Collection.FindOne(context.TODO(), filter).Decode(&transaction)
	if err != nil {
		
		if err == mongo.ErrNoDocuments {
			
			return models.Transaction{}, errors.New("transaction not found")
		}
		return models.Transaction{}, err
	}

	
	return transaction, nil
	
}



func (t TransactionImplementation) PutById(transactionId string, updatedTransaction models.Transaction) error {
	objectID, err := primitive.ObjectIDFromHex(transactionId)
	if err != nil {
		return err
	}

		filter := bson.M{"_id": objectID}

		update := bson.M{"$set": bson.M{
		"description": updatedTransaction.Description,
		"amount":      updatedTransaction.Amount,
		}}

	
	result, err :=t.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
	return err
	}

	if result.ModifiedCount == 0 {
		
	return errors.New("no document updated")
	}

	return nil
}



func (t TransactionImplementation) DeleteById(transactionId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(transactionId)
	if err != nil {
		return err
	}
	
	res, err := t.Collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	
	if res.DeletedCount == 0 {
		return errors.New("data not found")
	}
	
	return nil
}