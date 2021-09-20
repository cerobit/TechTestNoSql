package entities

import (
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type BirdValidateExistsGateway interface {
	Validate(int) (bool, error)
}
