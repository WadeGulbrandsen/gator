package main

import (
	"context"
	"fmt"
	"time"

	"github.com/WadeGulbrandsen/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Usage: %s <feed_url>", cmd.name)
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Couldn't get feed: %w", err)
	}
	follow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		})
	if err != nil {
		return fmt.Errorf("Couldn't create feed follow: %w", err)
	}
	fmt.Println("Feed follow created:")
	printFeedFollow(follow.UserName, follow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Couldn't get feed follows: %w", err)
	}

	if len(following) == 0 {
		fmt.Println("No feed follows found")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for i, follow := range following {
		fmt.Printf("%d. %s\n", i+1, follow.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Usage: %s <feed_url>", cmd.name)
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Couldn't get feed: %w", err)
	}
	if err := s.db.DeleteFeedFollow(
		context.Background(),
		database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID},
	); err != nil {
		return fmt.Errorf("Couldn't delete feed follow: %w", err)
	}
	fmt.Printf("%s is no longer following %s\n", user.Name, feed.Name)
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
