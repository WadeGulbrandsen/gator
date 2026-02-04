package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/WadeGulbrandsen/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Usage: %v <time_between_reqs>", cmd.name)
	}
	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Invalid duratin: %w", err)
	}
	log.Printf("Collecting feeds every %s...\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feed to fetch", err)
		return
	}
	if _, err := s.db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v\n", feed.Name, err)
		return
	}
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v\n", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		if err := savePost(s, feed, item); err != nil {
			log.Println(err)
		}
	}
	log.Printf("Feed %s collected, %v posts found\n", feed.Name, len(feedData.Channel.Item))
}

func savePost(s *state, feed database.Feed, item RSSItem) error {
	pub_time, err := parseTime(item.PubDate)
	if err != nil {
		return fmt.Errorf("Error parsing time: %w", err)
	}
	if _, err := s.db.CreatePost(
		context.Background(),
		database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pub_time,
			FeedID:      feed.ID,
		},
	); err != nil {
		if strings.Contains(err.Error(), "posts_url_key") {
			return nil
		}
		return fmt.Errorf("Could not create post: %w", err)
	}
	return nil
}

func parseTime(time_string string) (time.Time, error) {
	layouts := []string{time.RFC1123Z, time.RFC1123}
	for _, layout := range layouts {
		t, err := time.Parse(layout, time_string)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("Could not parse %q to a timestamp", time_string)
}
