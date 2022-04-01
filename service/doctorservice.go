package service

import "github.com/Jestinepa/doctor/domain"

type DoctorService interface {
	GetAllDoctor() ([]domain.Doctor, error)
}

type DefaultDoctorService struct {
	repo domain.DoctorRepository
}

func (s DefaultDoctorService) GetAllDoctor() ([]domain.Doctor, error) {
	return s.repo.FindAll()
}

func NewDoctorService(repository domain.DoctorRepository) DefaultDoctorService {
	return DefaultDoctorService{repository}
}
