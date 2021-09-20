package usecases

import (
	"TechTestNoSql/domain/entities"
	"errors"
	"fmt"
)



type VoteUseCases struct {
	BirdValidateGateway entities.BirdValidateExistsGateway
	BirdVoteGateway entities.BirdVoteGateway
}

// Validates existence of bird on the external microservice ( Part One of TechTest )
// if Exists then sum votes
func (vuc  VoteUseCases) CaseBirdVoteUp(id int) (int, int, error){

	birdExists,err  := vuc.BirdValidateGateway.Validate(id)
	if err != nil {
	   return id, -1,  errors.New("Error validating  information on Bird System")
	} else{
		if  birdExists != false {
			 votes, err2 := vuc.BirdVoteGateway.AddVote( id )

			 fmt.Println(votes, err2)
			 // Vote complete
			if err2 != nil {
				return id, -1,  errors.New("Error saving vote")
			}

			return id, votes,nil


		} else {
			return id, -1,  errors.New("Bird does not exists on bird system")
		}
	}

	return id, -1,nil
}