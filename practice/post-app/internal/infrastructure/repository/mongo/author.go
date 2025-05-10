package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	dom "post-app/internal/domain/author"
	"post-app/internal/domain/vo"
	"post-app/internal/infrastructure/repository/mongo/mapper"
	"time"
)

var _ dom.Repository = (*AuthorRepository)(nil)

// AuthorRepository представляет собой репозиторий для работы с авторами в MongoDB.
type AuthorRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

// NewAuthorRepository создаёт новый Mongo-репозиторий для авторов.
func NewAuthorRepository(db *mongo.Database, timeout time.Duration) *AuthorRepository {
	return &AuthorRepository{
		db:         db,
		collection: db.Collection("authors"),
		timeout:    timeout,
	}
}

// Create добавляет нового автора в БД.
func (r *AuthorRepository) Create(ctx context.Context, author *dom.Author) (vo.AuthorID, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	id, err := r.getNextID(ctx)
	if err != nil {
		return vo.AuthorID{}, fmt.Errorf("AuthorRepository.Create: %w", err)
	}

	authorID, err := vo.NewAuthorID(id)
	if err != nil {
		return vo.AuthorID{}, fmt.Errorf("AuthorRepository.Create: %w", err)
	}

	author.SetID(authorID)

	doc := mapper.FromAuthorToDoc(author)

	_, err = r.collection.InsertOne(ctx, doc)
	if err != nil {
		return vo.AuthorID{}, fmt.Errorf("AuthorRepository.Create: %w", err)
	}

	return authorID, nil
}

// FindByID находит автора по его ID.
func (r *AuthorRepository) FindByID(ctx context.Context, id vo.AuthorID) (*dom.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var doc mapper.AuthorDocument

	err := r.collection.FindOne(ctx, bson.M{"_id": id.Value()}).Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("AuthorRepository.FindByID: %w", err)
	}

	return mapper.MapDocToAuthor(doc)
}

// FindByIDs находит авторов по их ID.
func (r *AuthorRepository) FindByIDs(ctx context.Context, ids []vo.AuthorID) ([]*dom.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var authorsIDs []int32
	for _, id := range ids {
		authorsIDs = append(authorsIDs, id.Value())
	}

	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": authorsIDs}})
	if err != nil {
		return nil, fmt.Errorf("AuthorRepository.FindAll: %w", err)
	}
	defer cursor.Close(ctx)

	return r.decodeManyAuthors(ctx, cursor)
}

// Save сохраняет изменения в существующем авторе.
func (r *AuthorRepository) Save(ctx context.Context, author *dom.Author) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	doc := mapper.FromAuthorToDoc(author)

	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": doc.ID}, doc)
	if err != nil {
		return fmt.Errorf("AuthorRepository.Save: %w", err)
	}

	return nil
}

// getNextID возвращает следующее значение идентификатора.
func (r *AuthorRepository) getNextID(ctx context.Context) (int32, error) {
	filter := bson.M{"_id": "authors"}
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

// decodeManyAuthors декодирует курсор в массив авторов.
func (r *AuthorRepository) decodeManyAuthors(ctx context.Context, cursor *mongo.Cursor) ([]*dom.Author, error) {
	var authors []*dom.Author

	for cursor.Next(ctx) {
		var doc mapper.AuthorDocument
		if err := cursor.Decode(&doc); err != nil {
			return nil, fmt.Errorf("AuthorRepository.decodeManyPosts: %w", err)
		}

		author, err := mapper.MapDocToAuthor(doc)
		if err != nil {
			return nil, fmt.Errorf("AuthorRepository.decodeManyPosts: %w", err)
		}

		authors = append(authors, author)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("AuthorRepository.decodeManyPosts: %w", err)
	}

	return authors, nil
}
