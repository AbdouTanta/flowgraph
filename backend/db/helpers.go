package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Global variables
var (
	DefaultTimeout = 10 * time.Second
)

// FindOptions defines the options for the Find operations
type FindOptions struct {
	Filter     interface{} // Query filter (e.g., bson.M{"active": true})
	Sort       interface{} // e.g., bson.D{{"fieldName", 1}} for ascending, -1 for descending
	Skip       int64       // Number of documents to skip
	Limit      int64       // Maximum number of documents to return
	Projection interface{} // Fields to include/exclude in the results
}

// DefaultFindOptions returns the default options for FindManyDocuments
func DefaultFindOptions() *FindOptions {
	return &FindOptions{
		Filter: bson.M{}, // Default to empty filter
	}
}

// WithFilter sets the query filter criteria
func (o *FindOptions) WithFilter(filter interface{}) *FindOptions {
	o.Filter = filter
	return o
}

// WithSort adds sorting to the find options
func (o *FindOptions) WithSort(sort interface{}) *FindOptions {
	o.Sort = sort
	return o
}

// WithSkip adds skip to the find options for pagination
func (o *FindOptions) WithSkip(skip int64) *FindOptions {
	o.Skip = skip
	return o
}

// WithLimit adds limit to the find options
func (o *FindOptions) WithLimit(limit int64) *FindOptions {
	o.Limit = limit
	return o
}

// WithProjection adds projection to the find options
func (o *FindOptions) WithProjection(projection interface{}) *FindOptions {
	o.Projection = projection
	return o
}

// CreateDocument inserts a new document into the collection
func CreateDocument[T any](db *mongo.Database, collectionName string, document T) (*bson.ObjectID, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)
	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, fmt.Errorf("insert failed: %v", err)
	}

	id := result.InsertedID.(bson.ObjectID)
	return &id, nil
}

// CreateManyDocuments inserts multiple documents into the collection
func CreateManyDocuments[T any](db *mongo.Database, collectionName string, documents []T) ([]bson.ObjectID, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized, call Initialize first")
	}

	if len(documents) == 0 {
		return []bson.ObjectID{}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	// Convert to []interface{}
	docs := make([]interface{}, len(documents))
	for i, doc := range documents {
		docs[i] = doc
	}

	collection := db.Collection(collectionName)
	result, err := collection.InsertMany(ctx, docs)
	if err != nil {
		return nil, fmt.Errorf("insert many failed: %v", err)
	}

	ids := make([]bson.ObjectID, len(result.InsertedIDs))
	for i, id := range result.InsertedIDs {
		ids[i] = id.(bson.ObjectID)
	}

	return ids, nil
}

// FindOneDocument finds a single document by filter
func FindOneDocument[T any](db *mongo.Database, collectionName string, filter interface{}) (*T, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)

	if filter == nil {
		filter = bson.M{}
	}

	var result T
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil for not found
		}
		return nil, fmt.Errorf("find one failed: %v", err)
	}

	return &result, nil
}

// FindDocumentByID finds a document by its string ID
func FindDocumentByID[T any](db *mongo.Database, collectionName string, idStr string) (*T, error) {
	id, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	return FindOneDocument[T](db, collectionName, bson.M{"_id": id})
}

// FindManyDocuments retrieves multiple documents matching the options
func FindManyDocuments[T any](db *mongo.Database, collectionName string, opts ...*FindOptions) ([]T, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)

	// Set default options if none provided
	var opt *FindOptions
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	} else {
		opt = DefaultFindOptions()
	}

	// Ensure filter is not nil
	if opt.Filter == nil {
		opt.Filter = bson.M{}
	}

	// Build MongoDB options
	findOpts := options.Find()
	if opt.Sort != nil {
		findOpts.SetSort(opt.Sort)
	}
	if opt.Skip > 0 {
		findOpts.SetSkip(opt.Skip)
	}
	if opt.Limit > 0 {
		findOpts.SetLimit(opt.Limit)
	}
	if opt.Projection != nil {
		findOpts.SetProjection(opt.Projection)
	}

	cursor, err := collection.Find(ctx, opt.Filter, findOpts)
	if err != nil {
		return nil, fmt.Errorf("find failed: %v", err)
	}
	defer cursor.Close(ctx)

	var results []T
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("cursor iteration failed: %v", err)
	}

	return results, nil
}

// UpdateDocumentByID updates a document by its string ID
func UpdateDocumentByID(db *mongo.Database, collectionName string, idStr string, update interface{}) error {
	id, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	if db == nil {
		return fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return fmt.Errorf("update failed: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID %s", idStr)
	}

	return nil
}

// UpdateOneDocument updates a single document matching the filter
func UpdateOneDocument(db *mongo.Database, collectionName string, filter interface{}, update interface{}) error {
	if db == nil {
		return fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)

	if filter == nil {
		return fmt.Errorf("filter cannot be nil")
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("update failed: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found matching filter")
	}

	return nil
}

// UpdateManyDocuments updates multiple documents matching the filter
func UpdateManyDocuments(db *mongo.Database, collectionName string, filter interface{}, update interface{}) (int64, error) {
	if db == nil {
		return 0, fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)

	if filter == nil {
		return 0, fmt.Errorf("filter cannot be nil")
	}

	result, err := collection.UpdateMany(ctx, filter, update)
	if err != nil {
		return 0, fmt.Errorf("update many failed: %v", err)
	}

	return result.ModifiedCount, nil
}

// ReplaceDocumentByID replaces a document by its string ID
func ReplaceDocumentByID[T any](db *mongo.Database, collectionName string, idStr string, replacement T) error {
	id, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	if db == nil {
		return fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)
	result, err := collection.ReplaceOne(ctx, bson.M{"_id": id}, replacement)
	if err != nil {
		return fmt.Errorf("replace failed: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID %s", idStr)
	}

	return nil
}

// DeleteDocumentByID deletes a document by its string ID
func DeleteDocumentByID(db *mongo.Database, collectionName string, idStr string) error {
	id, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	if db == nil {
		return fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("delete failed: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with ID %s", idStr)
	}

	return nil
}

// DeleteOneDocument deletes a single document matching the filter
func DeleteOneDocument(db *mongo.Database, collectionName string, filter interface{}) error {
	if db == nil {
		return fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)

	if filter == nil {
		return fmt.Errorf("filter cannot be nil")
	}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete failed: %v", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found matching filter")
	}

	return nil
}

// DeleteManyDocuments deletes multiple documents matching the filter
func DeleteManyDocuments(db *mongo.Database, collectionName string, filter interface{}) (int64, error) {
	if db == nil {
		return 0, fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)

	if filter == nil {
		return 0, fmt.Errorf("filter cannot be nil")
	}

	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("delete many failed: %v", err)
	}

	return result.DeletedCount, nil
}

// CountDocuments counts documents matching the filter
func CountDocuments(db *mongo.Database, collectionName string, filter interface{}) (int64, error) {
	if db == nil {
		return 0, fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)

	if filter == nil {
		filter = bson.M{}
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("count failed: %v", err)
	}

	return count, nil
}

// DocumentExists checks if a document exists matching the filter
func DocumentExists(db *mongo.Database, collectionName string, filter interface{}) (bool, error) {
	count, err := CountDocuments(db, collectionName, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// DocumentExistsByID checks if a document exists by its string ID
func DocumentExistsByID(db *mongo.Database, collectionName string, idStr string) (bool, error) {
	id, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		return false, fmt.Errorf("invalid ID format: %v", err)
	}

	return DocumentExists(db, collectionName, bson.M{"_id": id})
}

// UpsertDocument performs an upsert operation (update or insert)
func UpsertDocument[T any](db *mongo.Database, collectionName string, filter interface{}, replacement T) (*bson.ObjectID, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)

	if filter == nil {
		return nil, fmt.Errorf("filter cannot be nil")
	}

	opts := options.Replace().SetUpsert(true)
	result, err := collection.ReplaceOne(ctx, filter, replacement, opts)
	if err != nil {
		return nil, fmt.Errorf("upsert failed: %v", err)
	}

	if result.UpsertedID != nil {
		id := result.UpsertedID.(bson.ObjectID)
		return &id, nil
	}

	return nil, nil // Document was updated, not inserted
}

// DropCollection drops (deletes) an entire collection
func DropCollection(db *mongo.Database, collectionName string) error {
	if db == nil {
		return fmt.Errorf("database not initialized, call Initialize first")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	collection := db.Collection(collectionName)
	return collection.Drop(ctx)
}
