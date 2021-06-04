package server

import (
	"context"

	pb "github.com/pablocrivella/mancala/pkg/proto/mancala/v1"

	"github.com/pablocrivella/mancala/internal/engine"
)

type GameService struct {
	eng engine.Service
	pb.UnimplementedGameServiceServer
}

func NewGameService(svc engine.Service) *GameService {
	return &GameService{eng: svc}
}

func (s *GameService) Find(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	g, err := s.eng.FindGame(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res := &pb.FindResponse{
		Game: &pb.Game{
			Id: g.ID.String(),
		},
	}
	return res, nil
}

func (s *GameService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	g, err := s.eng.CreateGame(ctx, req.Player1, req.Player2)
	if err != nil {
		return nil, err
	}
	res := &pb.CreateResponse{
		Game: &pb.Game{
			Id: g.ID.String(),
		},
	}
	return res, nil
}

func (s *GameService) ExecutePlay(ctx context.Context, req *pb.ExecutePlayRequest) (*pb.ExecutePlayResponse, error) {
	g, err := s.eng.ExecutePlay(ctx, req.Id, req.PitIndex)
	if err != nil {
		return nil, err
	}
	res := &pb.ExecutePlayResponse{
		Game: &pb.Game{
			Id: g.ID.String(),
		},
	}
	return res, nil
}
