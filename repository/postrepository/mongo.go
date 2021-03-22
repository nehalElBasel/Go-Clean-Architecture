package postrepository

import (
	"context"
	"log"

	"github.com/nehal1992/Go-Clean-Architecture/models"
	"github.com/nehal1992/Go-Clean-Architecture/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoPostRepo struct {
	Conn *mongo.Client
}

var db_name, collection_name = "go_crud", "posts"

func NewMongoPost(conn *mongo.Client) repository.PostRepo {
	return &mongoPostRepo{
		Conn: conn,
	}
}

func ConnectCollection(client *mongo.Client) (collection *mongo.Collection) {
	collection = client.Database(db_name).Collection(collection_name)
	return
}
func (mongo *mongoPostRepo) List(ctx context.Context) (posts []models.Post, err error) {
	collection := ConnectCollection(mongo.Conn)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return posts, err
	}
	for cursor.Next(ctx) {
		var post models.Post
		err = cursor.Decode(&post)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)

	}
	return posts, nil
}

func (mongo *mongoPostRepo) Create(ctx context.Context, e models.Post) (interface{}, error) {
	collection := ConnectCollection(mongo.Conn)
	insertResult, err := collection.InsertOne(ctx, e)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult.InsertedID, err
}

func (mongo *mongoPostRepo) Get(ctx context.Context, id int) (models.Post, error) {
	collection := ConnectCollection(mongo.Conn)
	var post models.Post
	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	return post, err
}

func (mongo *mongoPostRepo) Update(ctx context.Context, id int, p models.Post) error {
	collection := ConnectCollection(mongo.Conn)
	filter := bson.M{"id": id}
	change := bson.M{
		"$set": bson.M{
			"title":  p.Title,
			"author": p.Author,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, change)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (mysql *mongoPostRepo) Delete(ctx context.Context, id int) (bool, error) {
	collection := ConnectCollection(mysql.Conn)
	filter := bson.M{"id": id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
