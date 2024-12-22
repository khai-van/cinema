package cinemaservice

import (
	"context"

	pb "cinema/api/gen/cinema"
	"cinema/internal/config"
)

type Service struct {
	pb.UnimplementedCinemaServiceServer
	cinemaRepo Repository
}

func New() *Service {
	return &Service{
		cinemaRepo: newCinemaRepository(),
	}
}

func (s *Service) ConfigureCinema(ctx context.Context, req *pb.CinemaConfig) (*pb.StatusResponse, error) {
	_, err := s.cinemaRepo.CreateCinema(req.Name, int(req.Rows), int(req.Columns))
	if err != nil {
		return &pb.StatusResponse{Status: pb.Status_FAILURE, Message: err.Error()}, nil
	}
	return &pb.StatusResponse{Status: pb.Status_SUCCESS, Message: "Cinema created successfully"}, nil
}

func (s *Service) QueryReservedSeats(ctx context.Context, req *pb.QueryReservedSeatsRequest) (*pb.QueryReservedSeatsResponse, error) {
	cinema, err := s.cinemaRepo.GetCinemaConfig(int(req.CinemaId))
	if err != nil {
		return nil, err
	}

	seatGroups, err := s.cinemaRepo.GetReservedSeats(int(req.CinemaId))
	if err != nil {
		return nil, err
	}

	var protoSeatGroups []*pb.Seat
	for _, seat := range seatGroups {
		protoSeatGroups = append(protoSeatGroups, &pb.Seat{
			Row:    int32(seat.Row),
			Column: int32(seat.Column),
		})
	}

	return &pb.QueryReservedSeatsResponse{
		Seats:       protoSeatGroups,
		MaxRows:     int32(cinema.Rows),
		MaxColumns:  int32(cinema.Columns),
		MinDistance: int32(config.Get().Service.MinDistance),
	}, nil
}

func (s *Service) ReserveSeats(ctx context.Context, req *pb.SeatReservationRequest) (*pb.StatusResponse, error) {
	var seats []Seat
	for _, protoSeat := range req.Seats {
		seats = append(seats, Seat{
			Row:    int(protoSeat.Row),
			Column: int(protoSeat.Column),
		})
	}

	err := s.cinemaRepo.ReserveSeats(int(req.CinemaId), seats, config.Get().Service.MinDistance)
	if err != nil {
		return &pb.StatusResponse{Status: pb.Status_FAILURE, Message: err.Error()}, nil
	}

	return &pb.StatusResponse{Status: pb.Status_SUCCESS, Message: "Seats reserved successfully"}, nil
}

func (s *Service) CancelSeats(ctx context.Context, req *pb.SeatCancellationRequest) (*pb.StatusResponse, error) {
	var seats []Seat
	for _, protoSeat := range req.Seats {
		seats = append(seats, Seat{
			Row:    int(protoSeat.Row),
			Column: int(protoSeat.Column),
		})
	}

	err := s.cinemaRepo.CancelSeats(int(req.CinemaId), seats)
	if err != nil {
		return &pb.StatusResponse{Status: pb.Status_FAILURE, Message: err.Error()}, nil
	}

	return &pb.StatusResponse{Status: pb.Status_SUCCESS, Message: "Seats cancellation successful"}, nil
}
