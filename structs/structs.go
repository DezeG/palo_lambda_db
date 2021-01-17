package structs

type Patient struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Illness string `json:"illness"`
	Pain_level int `json:"painLevel"`
	Hospital string `json:"hospital"`
}