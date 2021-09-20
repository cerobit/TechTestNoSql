package handlers

import (
	"TechTestNoSql/domain/entities"
	"TechTestNoSql/domain/usecases"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type MainHandler struct {
	bvUseCases usecases.VoteUseCases
}

func (mh MainHandler)  Start(bvalidateExistsGT entities.BirdValidateExistsGateway,
	birdVoteGateway entities.BirdVoteGateway){

	mh.bvUseCases = usecases.VoteUseCases{bvalidateExistsGT, birdVoteGateway}

	http.HandleFunc("/api/v1/voteup", mh.voteUp )
}


func (mh MainHandler) voteUp(w http.ResponseWriter, r *http.Request)  {

	var birdRequest entities.BirdVoteRequest

	if r.Method == "POST" {
	  body,err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad Request Info", http.StatusBadRequest)
			json.NewEncoder(w).Encode(string("Bad request"))
		}

		err = json.Unmarshal(body, &birdRequest)
		if err != nil {
			http.Error(w, "Bad Request Info", http.StatusBadRequest)
			json.NewEncoder(w).Encode(string("Bad request"))
		}

		fmt.Println(birdRequest)
		 // ID bird info obtained correctly from request body
		 // Proceed with voting process
		birdId, votes, error :=  mh.bvUseCases.CaseBirdVoteUp(birdRequest.ID)
		birdResponse := entities.BirdVoteResponse{
			birdId,
			votes,
			error.Error()}

		json.NewEncoder(w).Encode(birdResponse)

	} else {
		 http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

