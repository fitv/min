package service

import "github.com/fitv/min/core/app"

var _ app.Service = (*Service)(nil)

type Service struct{}

func (Service) Register(*app.Application) {}
func (Service) Boot(*app.Application)     {}
