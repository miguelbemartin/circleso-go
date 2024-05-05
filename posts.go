package circle

import (
	"context"
	"net/http"
)

const pathPosts = "posts"

// Post represents a post in the Circle.so system.
type Post struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Url    string `json:"url"`
	UserId int    `json:"user_id"`
}

// PostsOptions provides optional parameters to the Posts method.
type PostsOptions struct {
	SpaceId    int    `url:"space_id,omitempty"`
	Sort       string `url:"sort,omitempty"`
	SearchText string `url:"search_text,omitempty"`
}

// GetPosts retrieve all the posts in a given space.
// See: https://api.circle.so/#42cf7913-312a-4bb5-a1a4-e6697a072291
func (c *Client) GetPosts(ctx context.Context, opts *PostsOptions) ([]Post, error) {
	req, err := c.newRequest(ctx, http.MethodGet, pathPosts, nil, opts)
	if err != nil {
		return nil, err
	}

	var resp []Post
	return resp, c.do(req, &resp)
}
