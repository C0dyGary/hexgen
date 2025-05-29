package configpage

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
    DB *mongo.Database
}