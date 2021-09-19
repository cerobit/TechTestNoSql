package entities

type BirdVote struct {
	ID int `json:"id"`
	BirdID int `json:"bird_id"`
	VoteUp int `json:"vote_up"`
	VoteDown int `json:"vote_down"`
}
