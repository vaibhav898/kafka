package chat

import (
  "log"

  "golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Username(ctx context.Context, in *Message) (*Message, error) {
  log.Printf("Receive Message Body from client: %s", in.Body)
  return &Message{Body: "vaibhav From the Server!"}, nil
}

func (s *Server) Name(ctx context.Context, in *Message) (*Message, error) {
  log.Printf("Name from a client: %s", in.Body)
  return &Message{Body: "Hii vaibhav!"}, nil
}