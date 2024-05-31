package main

import (
"context"
"fmt"
"log"
"time"

"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
// Set MongoDB URI
uri := "mongodb://localhost:27017" // Change this to your MongoDB server URI

// Set up a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// Connect to MongoDB
client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
if err != nil {
log.Fatal(err)
}
defer func() {
if err = client.Disconnect(ctx); err != nil {
log.Fatal(err)
}
}()

// List databases
databases, err := client.ListDatabaseNames(ctx, nil)
if err != nil {
log.Fatal(err)
}
fmt.Println("Databases:")
for _, db := range databases {
fmt.Println(" -", db)
}

// Specify the database to list collections
dbName := "your_database_name" // Change this to the database you want to inspect

// List collections in the specified database
collections, err := client.Database(dbName).ListCollectionNames(ctx, nil)
if err != nil {
log.Fatal(err)
}
fmt.Println("Collections in database", dbName + ":")
for _, coll := range collections {
fmt.Println(" -", coll)
}
}
