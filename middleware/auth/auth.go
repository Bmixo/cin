package auth

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
)

type Auth struct {
	Api               func(s *grpc.Server)
	NoVerifyMethodMap map[string]bool
	ParseSessionFunc  func(ctx context.Context, md *metadata.MD) (newCtx context.Context, err error)
}

func (self *Auth) AddNoVerifyMethods(methods []string) {
	for i := 0; i < len(methods); i++ {
		self.NoVerifyMethodMap[methods[i]] = true
		log.Println("Method:", methods[i], "public")
	}
}

func NewAuth() *Auth {
	return &Auth{
		NoVerifyMethodMap: map[string]bool{},
	}
}

// AuthInterceptor 认证拦截器，对以authorization为头部，形式为`bearer token`的Token进行验证
func (self *Auth) AuthInterceptor(ctx context.Context) (context.Context, error) {
	methodString, has := grpc.Method(ctx)
	if has == true {
		if v, ok := self.NoVerifyMethodMap[methodString]; ok && v == true {
			return ctx, nil
		}
	}
	//token, err := grpc_auth.AuthFromMD(ctx, "authorization")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}

	if self.ParseSessionFunc == nil {
		return nil, grpc.Errorf(codes.Unavailable, " %v", errors.New("error ParseSessionFunc null "))
	}
	newCtx, err := self.ParseSessionFunc(ctx, &md)
	if err != nil {
		return nil, err
	}
	if newCtx == nil {
		return nil, grpc.Errorf(codes.Unavailable, " %v", errors.New("system error"))
	}
	return newCtx, nil
}
