package storage

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MongoDbStorage struct {
	db         *mongo.Database
	collection *mongo.Collection
}

// NewMongoDb
// @param db
// @param collection
// @return *MongoDbStorage
func NewMongoDb(db *mongo.Database, collection string) *MongoDbStorage {

	return &MongoDbStorage{db, db.Collection(collection)}
}

// jsonStr2Bson
// @receiver m
// @param str
// @return interface{}
// @return error
func (m *MongoDbStorage) jsonStr2Bson(str string) (interface{}, error) {
	var want interface{}
	err := bson.UnmarshalJSON([]byte(str), &want)
	if err != nil {
		return nil, err
	}
	return want, nil
}

// InsertToDb 直接插入Json字段
// @receiver m
// @param wantStr
// @return string
// @return error
func (m *MongoDbStorage) InsertToDb(wantStr string) (string, error) {
	if wantStr == "" {
		return "", errors.New("转换的字符串为空")
	}
	want, err := m.jsonStr2Bson(wantStr)
	if err != nil {
		return "", err
	}
	res, err := m.collection.InsertOne(context.TODO(), want)
	if err != nil {
		return "", err
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("断言错误")
	}
	return id.Hex(), nil
}

// FindInfoByField 通过字段与KEY进行查询
// @receiver m
// @param field
// @param want
// @return string
// @return error
func (m *MongoDbStorage) FindInfoByField(field, want string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{field: want}
	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		return "", err
	}
	defer cursor.Close(ctx)
	var temp []bson.M
	if err = cursor.All(context.Background(), &temp); err != nil {
		return "", err
	}
	if len(temp) == 0 {
		return "", nil
	}
	jsonInfo, err := json.Marshal(temp)
	if err != nil {
		return "", err
	}
	return string(jsonInfo), nil
}

// FindInfoById 通过Mongo自己的ID进行查询
// @receiver m
// @param id
// @return string
// @return error
func (m *MongoDbStorage) FindInfoById(id string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		return "", err
	}
	defer cursor.Close(ctx)
	var temp []bson.M
	if err = cursor.All(context.Background(), &temp); err != nil {
		return "", err
	}
	if len(temp) == 0 {
		return "", nil
	}
	jsonInfo, err := json.Marshal(temp[0])
	if err != nil {
		return "", err
	}
	return string(jsonInfo), nil
}
