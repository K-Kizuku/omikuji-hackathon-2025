package model

type PythonStats struct {
	Hp       int            `json:"hp"`
	Attack   int            `json:"attack"`
	Defense  int            `json:"defense"`
	Speed    int            `json:"speed"`
	SkillIDs []int          `json:"skillIds"`
	Skills   []*PythonSkill `json:"skills"`
}
