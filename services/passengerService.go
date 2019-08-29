package services

import "github.com/freeddser/ride-sharing/repository"

func GetPassengerByID(ID int64) (*GenericResponse, error) {
	if ID == 0 {
		return NewGenericResponse(-1001), nil
	}
	repo := repository.GetPassengerRepository()
	data, err := repo.GetPassengerByID(ID)
	if err != nil {
		return NewGenericResponse(-1000), nil
	}
	if data.CustomerID == 0 {
		return NewGenericResponse(-1002), nil
	}

	genResp := NewGenericResponse(0)
	(*genResp).Data = data
	return genResp, nil
}
