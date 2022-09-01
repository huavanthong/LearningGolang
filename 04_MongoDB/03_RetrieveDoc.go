/***********

Refer:
	https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-read-documents
***********/
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")
	episodesCollection := quickstartDatabase.Collection("episodes")

	/*
		Solution 1: Find all
			Reading all data from a collection consists of making the request,
			then working with the results cursor.
			Knowing what fields exist on each of the documents isn't too important, only knowing the collection name itself.
		Disadvantage:
			However, because we're trying to return all documents, there aren't any fields in our query.
	*/
	fmt.Println("################## Example 1 #####################################\n")
	cursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)

	/*
		Solution 2: Find all iterate each document
			If your expected result set is large, using the *mongo.Cursor.All function might not be the best idea.
			Instead, you can iterate over your cursor and have it retrieve your data in batches.
			To do this, our code might change to the following:
	*/
	fmt.Println("\n################## Example 2 #####################################\n")
	cursor2, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor2.Close(ctx)
	for cursor2.Next(ctx) {
		var episode bson.M
		if err = cursor2.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		fmt.Println(episode)
	}

	/*
		Solution 3: Find One
	*/
	fmt.Println("\n################## Example 3 #####################################\n")
	var podcast bson.M
	if err = podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcast)

	/*
		Solution 4: Find with filter
	*/
	fmt.Println("\n################## Example 4 #####################################\n")

	filterCursor, err := episodesCollection.Find(ctx, bson.M{"duration": 25})
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesFiltered)

	/*
		Solution 5: Find with Sorting Documents in a Query
	*/
	fmt.Println("\n################## Example 5 #####################################\n")

	opts := options.Find()
	opts.SetSort(bson.D{{"duration", -1}})
	sortCursor, err := episodesCollection.Find(ctx, bson.D{{"duration", bson.D{{"$gt", 24}}}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var episodesSorted []bson.M
	if err = sortCursor.All(ctx, &episodesSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesSorted)
}
