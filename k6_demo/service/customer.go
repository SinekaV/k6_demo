package service

import (
	"context"
	interfaces "k6_demo/Interfaces"
	"k6_demo/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Cust struct{
	ctx context.Context
	customercollection *mongo.Collection
}


func InitCustomer(collection *mongo.Collection, ctx context.Context) interfaces.Icustomer{
	return &Cust{ctx,collection}
}
func(c *Cust) CreateCustomer(user *models.Customer)(*mongo.InsertOneResult,error){
	indexModel := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "account_id", Value: 1},{Key: "customer_id", Value: 1}}, // 1 for ascending, -1 for descending
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := c.customercollection.Indexes().CreateMany(c.ctx, indexModel)
	if err != nil {
		return nil,err
	}
	// date := time.Now()
	//.Format("2006-01-02 15:04:05")
	for i:=0;i<len(user.Transaction);i++{
		user.Transaction[i].Date = time.Now()
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password),7)
	user.Password = string(hashedPassword)
	res,err := c.customercollection.InsertOne(c.ctx, &user)
	if err!=nil{
		if mongo.IsDuplicateKeyError(err){
			log.Fatal("Duplicate key error")
		}
		return nil,err
	}
	
	return res,nil
}
