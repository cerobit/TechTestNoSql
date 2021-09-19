package entities

type BirdVoteGateway interface {
	addVote(BirdVote) (int, error)
}