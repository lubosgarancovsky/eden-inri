package dto

import "time"

type ListProjectUsersReq struct {
	UserID  string `header:"X-User-ID"`
	ScopeID string `uri:"projectId"`
}

type DeleteProjectUserReq struct {
	UserID  string `header:"X-User-ID"`
	ScopeID string `uri:"projectId"`
	ID      string `uri:"memberId"`
}

type ChangeProjectUserRoleReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	MemberID  string `uri:"memberId"`
	Role      string `json:"role"`
}

type ProjectUserRes struct {
	Role      string    `json:"role"`
	IsStarred bool      `json:"isStarred"`
	JoinedAt  time.Time `json:"joinedAt"`
	User      *UserRes  `json:"user"`
}
