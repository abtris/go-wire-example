//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func Initialize() User {
	wire.Build(NewUser, NewUserName)
	return User{}
}
