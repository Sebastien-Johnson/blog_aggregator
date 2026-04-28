package main

import (
	"context"
	"time"
	"fmt"
	"github.com/google/uuid"
	"github.com/Sebastien-Johnson/blog_aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return err
	}

	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}

	feed, err := s.db.GetFeedsByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:       	uuid.New(),
		CreatedAt: 	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		UserID: 	user.ID,
		FeedID: 	feed.ID, 
	})

	if err != nil {
		return err
	}

	fmt.Printf("Feed: %s\n User: %s\n", feedFollowRow.FeedName, feedFollowRow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return err
	}
	
	userFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(userFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, follow := range userFollows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	return nil
}