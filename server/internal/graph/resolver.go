package graph

import "github.com/K-Kizuku/pymon-graphql/internal/domain/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserRepository        repository.IUserRepository
	PythonRepository      repository.IPythonRepository
	PythonStatRepository  repository.IPythonStatRepository
	PythonSkillRepository repository.IPythonSkillRepository
}
