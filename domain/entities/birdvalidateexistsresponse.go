package entities

type BirdValidateExistsResponse struct{
	ID int  `json:"id"`
	Specie string `json:"specie"`
	Name string `json:"name"`
	Characteristics string `json:"characteristics"`
}
