package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name   string `json:"name"`
	Values []struct {
		Value    int `json:"value,omitempty"`
		Comments int `json:"comments,omitempty"`
		Likes    int `json:"likes,omitempty"`
		Shares   int `json:"shares,omitempty"`
	} `json:"values"`
}

func main() {
	input := []byte(`
[{
    "name": "organic_impressions_unique",
    "values": [{
        "value": 8288
    }]
	}, 
	{
        "name": "post_story_actions_by_type",
        "values": [{
            "shares": 234,
            "comments": 838,
            "likes": 8768
        }]
    }]
`)

	messages := []Message{} // Slice of Message instances
	json.Unmarshal(input, &messages)
	fmt.Println(messages)
}
