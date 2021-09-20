package entities

type BirdVoteGateway interface {
	AddVote(int) (int, error)
}