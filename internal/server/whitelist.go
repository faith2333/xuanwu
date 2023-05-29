package server

import (
	userPB "github.com/faith2333/xuanwu/api/user/v1"
)

var AuthenticationWhiteList = map[string]bool{
	userPB.OperationUserServerSignUp: true,
	userPB.OperationUserServerLogin:  true,
}
