package api

import (
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

func (s *Server) Login(ctx context.Context, in *LoginRequest) (*LoginReply, error) {
	if in.Username == "golang" && in.Password == "123456" {
		tokenString := CreateToken(in.Username)
		return &LoginReply{Status: "200", Token: tokenString}, nil
	} else {
		return &LoginReply{Status: "403", Token: ""}, nil
	}
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	msg := "message"
	userName := CheckAuth(ctx)
	// userName := fmt.Sprintf("%v", ctx.Value("username"))
	msg += " " + userName
	return &PingMessage{Greeting: msg}, nil
}
