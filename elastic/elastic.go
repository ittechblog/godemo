package main

import(
	"fmt"
	"context"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	// Create a context
	ctx := context.Background()

	// Create a client
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL("192.168.1.41:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	// Search with a term query
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index("dataflow").   // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Number of hits
	if searchResult.Hits != nil {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		//for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			//var t Tweet
			//err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			//fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		//}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}
}
