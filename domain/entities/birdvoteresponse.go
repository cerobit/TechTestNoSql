package entities

type BirdVoteResponse struct {
	BirdId int `json:"bird_id"`
	Votes int  `json:"votes"`
	Error string `json:"error"`
}