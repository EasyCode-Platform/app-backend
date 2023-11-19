package request

type RetrieveAllAppRequest struct {
	TeamID int `json:"teamID" validate:"required"`
}

func NewRetriveAllAppRequest(teamId int) *RetrieveAllAppRequest {
	return &RetrieveAllAppRequest{
		TeamID: teamId,
	}
}

type RetrieveAppByNameRequest struct {
	TeamID int    `json:"teamID" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

func NewRetrieveAppByNameRequest(teamId int, name string) *RetrieveAppByNameRequest {
	return &RetrieveAppByNameRequest{
		TeamID: teamId,
		Name:   name,
	}
}

type RetrieveAppByIdRequest struct {
	UID int `json:"id" validate:"required"`
}

func NewRetrieveAppByIdRequest(id int) *RetrieveAppByIdRequest {
	return &RetrieveAppByIdRequest{
		UID: id,
	}
}
