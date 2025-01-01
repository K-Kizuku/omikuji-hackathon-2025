package entity

type PythonSkill struct {
	ID          int
	Name        string
	Description string
	Pp          int
	Attack      int
	HitRate     int
	MinVersion  string
	MaxVersion  string
}

type PythonSkills []*PythonSkill
