package entity

type PythonStat struct {
	PythonID string
	Hp       int
	Attack   int
	Defense  int
	Speed    int
	Skill1   *int
	Skill2   *int
	Skill3   *int
	Skill4   *int
}

func (p *PythonStat) SkillIDs() []int {
	skillIDs := make([]int, 0, 4)
	if p.Skill1 != nil {
		skillIDs = append(skillIDs, *p.Skill1)
	}
	if p.Skill2 != nil {
		skillIDs = append(skillIDs, *p.Skill2)
	}
	if p.Skill3 != nil {
		skillIDs = append(skillIDs, *p.Skill3)
	}
	if p.Skill4 != nil {
		skillIDs = append(skillIDs, *p.Skill4)
	}
	return skillIDs
}
