package emqx

type GetAuthUsersReq struct {
	Page        int
	Limit       int
	LikeUserID  string
	IsSuperUser bool
}

type GetAuthUsersResp struct {
	Meta struct {
		Page    int  `json:"page"`
		Limit   int  `json:"limit"`
		Count   int  `json:"count"`
		Hasnext bool `json:"hasnext"`
	} `json:"meta"`
	GetAuthUserResp `json:"data"`
	Code            int `json:"code"`
}

type GetAuthUserResp struct {
	UserID      string `json:"user_id"`
	IsSuperUser bool   `json:"is_superuser"`
}

type CreateUserReq struct {
	UserID   string `json:"user_id,omitempty"`
	Password string `json:"password"`
}

type CodeResp struct {
	Code string `json:"code"`
}

type CreateUserResp = CodeResp
