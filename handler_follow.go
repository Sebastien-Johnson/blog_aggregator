package main

import (
	"context"
	"fmt"
	"time"
	"github.com/Sebastien-Johnson/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
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

func handlerFollowing(s *state, cmd command, user database.User) error {
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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	feedUrl := cmd.args[0] 
	feed, err := s.db.GetFeedsByUrl(context.Background(), feedUrl)

	if err != nil {
		return fmt.Errorf("Unable to find feed: %w", err)
	}
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: feed.UserID,
    	FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Unable to find follow: %w", err)
	}
	fmt.Print("Unfollowed!")
	return nil
}