package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"post-app/internal/domain/author"
	dom "post-app/internal/domain/post"
	"post-app/internal/infrastructure/repository/mongo/mapper"
	"time"
)

var _ dom.Repository = (*PostRepository)(nil)

// PostRepository представляет собой репозиторий для работы с постами в MongoDB.
type PostRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

// NewPostRepository создаёт новый Mongo-репозиторий для постов.
func NewPostRepository(db *mongo.Database, timeout time.Duration) *PostRepository {
	return &PostRepository{
		db:         db,
		collection: db.Collection("posts"),
		timeout:    timeout,
	}
}

// Create добавляет новый пост в БД.
func (r *PostRepository) Create(ctx context.Context, post *dom.Post) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	id, err := r.getNextID(ctx)
	if err != nil {
		return fmt.Errorf("PostRepository.Create: %w", err)
	}

	postID, err := dom.NewPostID(id)
	if err != nil {
		return fmt.Errorf("PostRepository.Create: %w", err)
	}

	post.SetID(postID)

	doc := mapper.FromPostToDoc(post)

	_, err = r.collection.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("PostRepository.Create: %w", err)
	}

	return nil
}

// FindByID находит пост по его ID.
func (r *PostRepository) FindByID(ctx context.Context, id dom.PostID) (*dom.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var doc mapper.PostDocument

	err := r.collection.FindOne(ctx, bson.M{"_id": id.Value()}).Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("PostRepository.FindByID: %w", err)
	}

	return mapper.MapDocToPost(doc)
}

func (r *PostRepository) FindByAuthorID(ctx context.Context, authorID author.AuthorID) ([]*dom.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostRepository) FindAll(ctx context.Context) ([]*dom.Post, error) {
	//TODO implement me
	panic("implement me")
}

// Save сохраняет изменения в существующем посте.
func (r *PostRepository) Save(ctx context.Context, post *dom.Post) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	doc := mapper.FromPostToDoc(post)

	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": doc.ID}, doc)
	if err != nil {
		return fmt.Errorf("PostRepository.Save: %w", err)
	}

	return nil
}

func (r *PostRepository) DeleteByID(ctx context.Context, id dom.PostID) error {
	//TODO implement me
	panic("implement me")
}

// getNextID возвращает следующее значение идентификатора.
func (r *PostRepository) getNextID(ctx context.Context) (int32, error) {
	filter := bson.M{"_id": "posts"}
	update := bson.M{"$inc": bson.M{"seq": 1}}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var result struct {
		Seq int32 `bson:"seq"`
	}
	err := r.db.Collection("counters").FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return 0, fmt.Errorf("getNextID: %w", err)
	}
	return result.Seq, nil
}
