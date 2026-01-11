package query

import go_kit "github.com/lubosgarancovsky/go-kit"

type FindProjectLabelByIDQuery struct {
	ID        string
	UserID    string
	ProjectID string
}

type ListProjectLabelsQuery struct {
	UserID    string
	ProjectID string
	go_kit.ListingQuery
}
