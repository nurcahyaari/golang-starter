package router

import (
	"golang-starter/internal/utils/auth"
	"golang-starter/internal/web"
)

type RouterStruct struct {
	web.RouterStruct
	jwtAuth auth.JwtTokenInterface
}
