package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/WadeGulbrandsen/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		new_limit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("Usage: %s [num_posts]", cmd.name)
		}
		limit = new_limit
	}
	posts, err := s.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{UserID: user.ID, Limit: int32(limit)},
	)
	if err != nil {
		return fmt.Errorf("Couldn't get posts: %w", err)
	}
	if len(posts) == 0 {
		fmt.Println("No posts found")
		return nil
	}
	for i, post := range posts {
		fmt.Printf("==== Post %3d ====\n", i+1)
		printPost(post)
		fmt.Println()
	}
	fmt.Printf("Found %d posts:\n", len(posts))
	return nil
}

func printPost(post database.Post) {
	fmt.Printf(" * ID:           %v\n", post.ID)
	fmt.Printf(" * Title:        %v\n", post.Title)
	fmt.Printf(" * Published At: %v\n", post.PublishedAt)
	fmt.Printf(" * Description:  %v\n", post.Description)
	fmt.Printf(" * Link:         %v\n", post.Url)
}
