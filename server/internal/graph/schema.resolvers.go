package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.62

import (
	"context"
	"fmt"
	"time"

	"github.com/K-Kizuku/pymon-graphql/internal/graph/model"
)

// User is the resolver for the user field.
func (r *pythonResolver) User(ctx context.Context, obj *model.Python) (*model.User, error) {
	user, err := r.UserRepository.GetUserByID(ctx, obj.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
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
	if stat == nil {
		return nil, nil
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
	if user == nil {
		return nil, nil
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
	if python == nil {
		return nil, nil
	}
	return &model.Python{
		ID:         python.ID,
		UserID:     python.UserID,
		Name:       python.Name,
		Exp:        python.Exp,
		Repository: python.Repository,
	}, nil
}

// CurrentTime is the resolver for the currentTime field.
func (r *subscriptionResolver) CurrentTime(ctx context.Context) (<-chan *model.Time, error) {
	ch := make(chan *model.Time)

	go func() {
		defer close(ch)
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("Tick")
			currentTime := time.Now()
			t := &model.Time{
				UnixTime:  int(currentTime.Unix()),
				TimeStamp: currentTime.Format(time.RFC3339),
			}

			select {
			case <-ctx.Done():
				fmt.Println("Subscription is closed")
				return
			case ch <- t:
			}
		}
	}()

	return ch, nil
}

// Python returns PythonResolver implementation.
func (r *Resolver) Python() PythonResolver { return &pythonResolver{r} }

// PythonStats returns PythonStatsResolver implementation.
func (r *Resolver) PythonStats() PythonStatsResolver { return &pythonStatsResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type pythonResolver struct{ *Resolver }
type pythonStatsResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *queryResolver) Placeholder(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Placeholder - placeholder"))
}
*/
