package service

import "github.com/cloudwego/hertz/pkg/app"

var relationService = NewRelationService(nil, new(app.RequestContext))
