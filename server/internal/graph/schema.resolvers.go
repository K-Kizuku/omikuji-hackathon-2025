package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"

	"github.com/K-Kizuku/pymon-graphql/internal/graph/model"
)

// User is the resolver for the user field.
func (r *pythonResolver) User(ctx context.Context, obj *model.Python) (*model.User, error) {
	user, err := r.UserRepository.GetUserByID(ctx, obj.UserID)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

// Stats is the resolver for the stats field.
func (r *pythonResolver) Stats(ctx context.Context, obj *model.Python) (*model.PythonStats, error) {
	stat, err := r.PythonStatRepository.GetPythonStatByID(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return &model.PythonStats{
		Hp:       stat.Hp,
		Attack:   stat.Attack,
		Defense:  stat.Defense,
		Speed:    stat.Speed,
		SkillIDs: stat.SkillIDs(),
	}, nil
}

// Skills is the resolver for the skills field.
func (r *pythonStatsResolver) Skills(ctx context.Context, obj *model.PythonStats) ([]*model.PythonSkill, error) {
	skills, err := r.PythonSkillRepository.SelectByIDs(ctx, obj.SkillIDs)
	if err != nil {
		return nil, err
	}
	resp := make([]*model.PythonSkill, 0, len(skills))
	for _, skill := range skills {
		resp = append(resp, &model.PythonSkill{
			ID:          skill.ID,
			Name:        skill.Name,
			Description: skill.Description,
			Pp:          skill.Pp,
			Attack:      skill.Attack,
			HitRate:     skill.HitRate,
			MinVersion:  skill.MinVersion,
			MaxVersion:  skill.MaxVersion,
		})
	}
	return resp, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	// TODO: 認証からuserIDを取得する
	user, err := r.UserRepository.GetUserByID(ctx, "1")
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

// Python is the resolver for the python field.
func (r *queryResolver) Python(ctx context.Context) (*model.Python, error) {
	// TODO: 認証からuserIDを取得する
	python, err := r.PythonRepository.GetPythonByUserID(ctx, "1")
	if err != nil {
		return nil, err
	}
	return &model.Python{
		ID:         python.ID,
		UserID:     python.UserID,
		Name:       python.Name,
		Exp:        python.Exp,
		Repository: python.Repository,
	}, nil
}

// Python returns PythonResolver implementation.
func (r *Resolver) Python() PythonResolver { return &pythonResolver{r} }

// PythonStats returns PythonStatsResolver implementation.
func (r *Resolver) PythonStats() PythonStatsResolver { return &pythonStatsResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type pythonResolver struct{ *Resolver }
type pythonStatsResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
