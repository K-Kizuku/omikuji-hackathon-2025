package model

type PythonStats struct {
	Hp      int32          `json:"hp"`
	Attack  int32          `json:"attack"`
	Defense int32          `json:"defense"`
	Speed   int32          `json:"speed"`
	Skills  []*PythonSkill `json:"skills"`
}
