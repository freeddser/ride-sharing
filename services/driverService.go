package services

import "github.com/freeddser/ride-sharing/repository"

func GetDriverByID(ID int64) (*GenericResponse, error) {
	if ID == 0 {
		return NewGenericResponse(-1001), nil
	}
	repo := repository.GetDriverRepository()
	data, err := repo.GetDriverByID(ID)
	if err != nil {
		return NewGenericResponse(-1002), nil
	}
	if data.DriverID == 0 {
		return NewGenericResponse(-1002), nil
	}

	genResp := NewGenericResponse(0)
	(*genResp).Data = data
	return genResp, nil
}
