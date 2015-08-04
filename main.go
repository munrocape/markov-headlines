package main

import (
	"fmt"
	hn "github.com/munrocape/hn/hnclient"
)

func main() {
	hnClient := hn.NewClient()
	stories, _ := hnClient.GetTopStories(500)
	fmt.Println(stories)
}