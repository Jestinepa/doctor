package domain

type Timings struct {
	From string
	To   string
}

type Doctor struct {
	DoctorId     string
	Fullname     string
	Availability []Timings
}

type DoctorRepository interface {
	FindAll() ([]Doctor, error)
}
