package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnector struct {
	Client *mongo.Client
}

type MongoDBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (cfg *MongoDBConfig) GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
}

// URI format: mongodb://username:password@host:port
func NewMongoDBConnector(ctx context.Context, cfg *MongoDBConfig) (*MongoDBConnector, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.GetURI()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}
	return &MongoDBConnector{Client: client}, nil
}

func (conn *MongoDBConnector) Close(ctx context.Context) error {
	err := conn.Client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("failed to disconnect from mongodb: %w", err)
	}
	return nil
}
