package circle

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

const pathMembers = "community_members"

// Member represents a member in the Circle.so system.
type Member struct {
	ID                 int       `json:"id"`
	Email              string    `json:"email"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	Headline           string    `json:"headline"`
	Bio                string    `json:"bio"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	CommunityId        int       `json:"community_id"`
	LastSeenAt         time.Time `json:"last_seen_at"`
	ProfileUrl         string    `json:"profile_url"`
	PublicUid          string    `json:"public_uid"`
	TopicsCount        int       `json:"topics_count"`
	PostsCount         int       `json:"posts_count"`
	CommentsCount      int       `json:"comments_count"`
	Location           string    `json:"location"`
	WebsiteUrl         string    `json:"website_url"`
	InstagramUrl       string    `json:"instagram_url"`
	TwitterUrl         string    `json:"twitter_url"`
	LinkedinUrl        string    `json:"linkedin_url"`
	FacebookUrl        string    `json:"facebook_url"`
	AcceptedInvitation string    `json:"accepted_invitation"`
	Active             bool      `json:"active"`
}

// MembersOptions provides optional parameters to the Members method.
type MembersOptions struct {
	Sort    string `url:"sort,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
	Page    int    `url:"page,omitempty"`
	Status  string `url:"status,omitempty"`
	Email   string `url:"email,omitempty"`
}

// GetMembers retrieve a set of members from your community. Uses per_page and page params to paginate results.
// See: https://api.circle.so/#b5d3063a-1b8d-4310-944c-60483515e8f5
func (c *Client) GetMembers(ctx context.Context, opts *MembersOptions) ([]Member, error) {
	req, err := c.newRequest(ctx, http.MethodGet, pathMembers, nil, opts)
	if err != nil {
		return nil, err
	}

	var resp []Member
	return resp, c.do(req, &resp)
}

// GetMemberById retrieve a set of members from your community. Uses per_page and page params to paginate results.
// See: https://api.circle.so/#234e7550-4a81-41a0-b9e4-5d1b4d9bd246
func (c *Client) GetMemberById(ctx context.Context, id int) (Member, error) {
	req, err := c.newRequest(ctx, http.MethodGet, pathMembers+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return Member{}, err
	}

	var resp Member
	return resp, c.do(req, &resp)
}

type searchByEmail struct {
	Email string `url:"email"`
}

// GetMemberByEmail retrieve a set of members from your community. Uses per_page and page params to paginate results.
// See: https://api.circle.so/#234e7550-4a81-41a0-b9e4-5d1b4d9bd246
func (c *Client) GetMemberByEmail(ctx context.Context, email string) (Member, error) {
	req, err := c.newRequest(ctx, http.MethodGet, pathMembers+"/search", nil, &searchByEmail{Email: email})
	if err != nil {
		return Member{}, err
	}

	var resp Member
	return resp, c.do(req, &resp)
}
