package model

type Python struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Exp        int32        `json:"exp"`
	Repository string       `json:"repository"`
	UserID     string       `json:"userId"`
	User       *User        `json:"user"`
	Stats      *PythonStats `json:"stats"`
}
