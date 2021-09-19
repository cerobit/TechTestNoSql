package drivenadapters

import "TechTestNoSql/model/entities"

type birdVoteRedis struct {
	// redis conn
}

func (b birdVoteRedis) addVote(birdVote entities.BirdVote) (int, error) {
	panic("implement me")
}
