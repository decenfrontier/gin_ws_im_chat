package service

import (
	"chat/conf"
	"chat/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sort"
	"time"
)

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     uint   `json:"read"`
	CreateAt int64  `json:"create_at"`
}

func InsertMsg(database string, id string, content string, read uint, expire int64) (err error) {
	collection := conf.MongoDBClient.Database(database).Collection(id)
	comment := model.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err = collection.InsertOne(context.TODO(), comment)
	return
}

func FindMany(database string, sendId string, id string, time int64, pageSize int) (results []model.Result, err error) {
	var resultsMe []model.Trainer  // id
	var resultsYou []model.Trainer // sendId
	db := conf.MongoDBClient.Database(database)
	sendIdCollection := db.Collection(sendId)
	idCollection := db.Collection(id)
	sendIdTimeCursor, err := sendIdCollection.Find(context.TODO(),
		options.Find().SetSort(bson.D{{"startTime", -1}}),
		options.Find().SetLimit(int64(pageSize)))
	idTimeCursor, err := idCollection.Find(context.TODO(),
		options.Find().SetSort(bson.D{{"startTime", -1}}),
		options.Find().SetLimit(int64(pageSize)))
	err = sendIdTimeCursor.All(context.TODO(), &resultsYou) // sendId 对面发过来的
	err = idTimeCursor.All(context.TODO(), &resultsMe)      // Id 发给对面的
	results, _ = AppendAndSort(resultsMe, resultsYou)
	return
}

func AppendAndSort(resultsMe, resultsYou []model.Trainer) (results []model.Result, err error) {
	for _, r := range resultsMe {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := model.Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "me",
		}
		results = append(results, result)
	}
	for _, r := range resultsYou {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := model.Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "you",
		}
		results = append(results, result)
	}
	// 最后进行排序
	sort.Slice(results, func(i, j int) bool {
		return results[i].StartTime < results[j].StartTime
	})
	return results, nil
}
