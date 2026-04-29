package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"
	"strings"
	"log"
	"github.com/Sebastien-Johnson/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
    	return errors.New("Not enough arguments submitted")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	//create feed with params struct
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("Could not create feed: %w\n", err)
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("User unable to follow feed: %w\n", err)
	}
	fmt.Println("Feed created successfully:")
	printFeed(feed, user)

	fmt.Println("Feed followed successfully:")
    printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	//iterate through slice of feeds
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w\n", err)
		}
		printFeed(feed, user)
	}
	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to locate next feed: %w\n", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("Unable to mark feed as fetch: %w\n", err)
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)

	if err != nil {
		return fmt.Errorf("Unable to retrieve next feed: %w\n", err)
	}

	for _,item := range rssFeed.Channel.Item {
		publishDate := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
    		publishDate = sql.NullTime{Time: t, Valid: true}
		}
		_, err = s.db.CreatPost(context.Background(), database.CreatPostParams{
			ID:        		uuid.New(),
			CreatedAt: 		time.Now().UTC(),
			UpdatedAt: 		time.Now().UTC(),
			Title:       	item.Title,
			Url:         	item.Link,
			Description: 	sql.NullString{String: item.Description, Valid: true},
			PublishedAt: 	publishDate,
			FeedID:      	nextFeed.ID,
	})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
    		continue
		}
	}
	return nil
}

func handlerBrowse(s *state, cmd command, user database.User) error{
	limit := 0

	if len(cmd.args) == 0 {
		limit = 2
	} else {
		i, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("Argument must be number: %w", err)
		}
		limit = i
	}
	posts, err := s.db.GetPostForUser(context.Background(), int32(limit))
	if err != nil {
		return fmt.Errorf("Unable to retrieve post: %w", err)
	}
	for _, post := range posts {
		fmt.Println(post)
	}
	return nil
}