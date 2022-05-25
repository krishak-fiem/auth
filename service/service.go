package service

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/krishak-fiem/auth/constants"
	"github.com/krishak-fiem/auth/proto/pb"
	"github.com/krishak-fiem/auth/utils"
	consts "github.com/krishak-fiem/constants/go"
	kafka "github.com/krishak-fiem/kafka/go"
	authmodels "github.com/krishak-fiem/models/go/auth"
	kafkamodels "github.com/krishak-fiem/models/go/kafka"
	"github.com/krishak-fiem/utils/go/jwt"

	"github.com/krishak-fiem/utils/go/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

var validate = validator.New()

func (s *Server) Signup(ctx context.Context, message *pb.SignupMessage) (*pb.Response, error) {
	err := validate.Struct(message)
	if err != nil {
		return &pb.Response{}, status.Error(codes.Aborted, err.Error())
	}

	_, ok := utils.CheckUserExists(message.Email)
	if ok {
		return &pb.Response{}, status.Error(codes.Aborted, "User already exists.")
	}

	user := new(authmodels.User)
	user.Email = message.Email
	user.Password, err = bcrypt.HashPassword(message.Password)
	if err != nil {
		return &pb.Response{}, status.Error(codes.Internal, err.Error())
	}

	err = user.CreateUser()
	if err != nil {
		return &pb.Response{}, status.Error(codes.Internal, err.Error())
	}

	kafkaMessage := kafkamodels.UserCreatedMessage{Name: message.Name, Email: user.Email}
	m, err := json.Marshal(kafkaMessage)
	if err != nil {
		return &pb.Response{}, status.Error(codes.Internal, err.Error())
	}

	kafka.MessageWriter(constants.KAFKA_BROKER_URL, consts.USER_CREATED, consts.AUTH_USER, m)

	return &pb.Response{Payload: "User Created.", Status: true}, nil
}

func (s *Server) Signin(ctx context.Context, message *pb.SigninMessage) (*pb.Response, error) {
	err := validate.Struct(message)
	if err != nil {
		return &pb.Response{}, status.Error(codes.Aborted, err.Error())
	}

	password, ok := utils.CheckUserExists(message.Email)
	if !ok {
		return &pb.Response{}, status.Error(codes.PermissionDenied, "User does not exists.")
	}

	ok = bcrypt.CheckPasswordHash(message.Password, password)
	if !ok {
		return &pb.Response{}, status.Error(codes.PermissionDenied, "Password does not match.")
	}

	token, err := jwt.CreateToken(message.Email)
	if err != nil {
		return &pb.Response{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.Response{Payload: token, Status: true}, nil
}
