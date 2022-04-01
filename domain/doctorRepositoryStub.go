package domain

type DoctorRepositoryStub struct {
	doctors []Doctor
}

func (s DoctorRepositoryStub) FindAll() ([]Doctor, error) {
	return s.doctors, nil
}

func NewDoctorRespositoryStub() DoctorRepositoryStub {
	doctors := []Doctor{
		{
			DoctorId: "1",
			Fullname: "jackson",
			Availability: []Timings{
				{
					From: "10",
					To:   "12",
				},
				{
					From: "4",
					To:   "6",
				},
			},
		},
	}
	return DoctorRepositoryStub{doctors}
}
