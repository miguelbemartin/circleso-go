package circle

import (
	"context"
	"net/http"
)

const pathMembers = "community_members"

// Member represents a label in the Nylas system.
type Member struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Headline  string `json:"headline"`
}

// MembersOptions provides optional parameters to the Members method.
type MembersOptions struct {
	Sort    string `url:"sort,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
	Page    int    `url:"page,omitempty"`
	Status  string `url:"status,omitempty"`
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
