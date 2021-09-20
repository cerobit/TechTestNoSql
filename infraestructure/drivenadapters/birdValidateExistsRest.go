package drivenadapters

import (
	"TechTestNoSql/domain/entities"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)


type BirdValidateExistsRest struct{
	urlBaseFindBird string
}

// Rest client, get bird info validating if exist on  Birds Microservice
// Todo inject config variable url properly
func (b BirdValidateExistsRest) Validate(i int) (bool, error) {
	b.urlBaseFindBird = "http://localhost:8080/api/v1/byid?id="
	var myClient = &http.Client{Timeout: 10 * time.Second}
    r, err := myClient.Get(b.urlBaseFindBird+strconv.Itoa(i))
    if err != nil {
        return false, err
    }
    defer r.Body.Close()
	var birdInfo entities.BirdValidateExistsResponse
	err =  json.NewDecoder(r.Body).Decode(&birdInfo)
	if err != nil {
		return false,err
	}
	return true, nil
}
