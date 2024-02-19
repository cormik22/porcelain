package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/cormik22/porcelain"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a search query: ")
	query, _ := reader.ReadString('\n')

	videos, err := porcelain.Search(strings.TrimSpace(query), 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(videos) == 0 {
		fmt.Println("No videos found")
		return
	}

	rand.Shuffle(len(videos), func(i, j int) {
		videos[i], videos[j] = videos[j], videos[i]
	})

	for i, video := range videos {
		if i >= 10 {
			break
		}

		fmt.Printf("%d. [%s] %s\n", i+1, video.Provider, video.Title)
	}

	fmt.Print("Enter the index of the video you want to qualify: ")
	indexRaw, _ := reader.ReadString('\n')
	index, _ := strconv.Atoi(strings.TrimSpace(indexRaw))

	video, err := porcelain.Qualify(videos[index-1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Title: %s\n", video.Title)
	fmt.Printf("URL: %s\n", video.Url)
	fmt.Printf("Duration: %d seconds\n", video.Duration)
	fmt.Printf("View count: %d\n", video.ViewCount)
	fmt.Printf("Uploader: %s\n", video.Uploader)
	fmt.Printf("Likes: %d\n", video.Likes)
	fmt.Printf("Dislikes: %d\n", video.Dislikes)
	fmt.Printf("Categories: %v\n", strings.Join(video.Categories, ", "))
	fmt.Printf("Tags: %v\n", strings.Join(video.Tags, ", "))
	fmt.Printf("Pornstars: %v\n", strings.Join(video.Pornstars, ", "))
}
