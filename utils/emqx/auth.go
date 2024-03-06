package emqx

import (
	"context"
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
)

type Authenticator = string

var (
	PasswordBased_BuildInDatabase Authenticator = `password_based:built_in_database`
)

func (c *Client) GetAuthUsers(ctx context.Context, id Authenticator, query *GetAuthUsersReq) (*GetAuthUsersResp, error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/authentication/%s/users", id), nil)
	if err != nil {
		return nil, err
	}

	if query != nil {
		q := req.URL.Query()
		if query.Page != 0 {
			q.Add("page", fmt.Sprint(query.Page))
		}
		if query.Limit != 0 {
			q.Add("limit", fmt.Sprint(query.Limit))
		}
		if query.LikeUserID != "" {
			q.Add("like_user_id", query.LikeUserID)
		}
		if query.IsSuperUser {
			q.Add("is_super_user", "true")
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.hcli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var au GetAuthUsersResp
	if err := json.NewDecoder(resp.Body).Decode(&au); err != nil {
		return nil, err
	}
	return &au, nil
}

func (c *Client) GetAuthUser(ctx context.Context, id Authenticator, uid string) (*GetAuthUserResp, error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/authentication/%s/users/%s", id, uid), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.hcli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var au GetAuthUserResp
	if err := json.NewDecoder(resp.Body).Decode(&au); err != nil {
		return nil, err
	}
	return &au, nil
}

func (c *Client) CreateUsername(ctx context.Context, id Authenticator, uid string, password string) error {
	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("/authentication/%s/users", id), &CreateUserReq{
		UserID:   uid,
		Password: password,
	})
	if err != nil {
		return err
	}

	resp, err := c.hcli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var cu CreateUserResp
	if err := json.NewDecoder(resp.Body).Decode(&cu); err != nil {
		return err
	}

	if cu.Code != "" {
		return fmt.Errorf("create user failed: %s", cu.Code)
	}

	return nil
}

func (c *Client) SetUserPassword(ctx context.Context, id Authenticator, uid string, password string) error {
	req, err := c.newRequest(ctx, http.MethodPut, fmt.Sprintf("/authentication/%s/users/%s", id, uid), &CreateUserReq{
		Password: password,
	})
	if err != nil {
		return err
	}

	resp, err := c.hcli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var cu CodeResp
	if err := json.NewDecoder(resp.Body).Decode(&cu); err != nil {
		return err
	}

	if cu.Code != "" {
		return fmt.Errorf("set user password failed code: %s", cu.Code)
	}

	return nil
}

func (c *Client) DeleteUser(ctx context.Context, id Authenticator, uid string) error {
	req, err := c.newRequest(ctx, http.MethodDelete, fmt.Sprintf("/authentication/%s/users/%s", id, uid), nil)
	if err != nil {
		return err
	}

	resp, err := c.hcli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	var cu CodeResp
	if err := json.NewDecoder(resp.Body).Decode(&cu); err != nil {
		return err
	}

	if cu.Code != "" {
		return fmt.Errorf("delete uid failed code: %s", cu.Code)
	}

	return nil
}
