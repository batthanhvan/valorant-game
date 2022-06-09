package lib

import (
	"golang.org/x/xerrors"
	// "google.golang.org/grpc/status"
)

// var (
// 	ErrParseError = xerrors.New("parse grpc error fail")
// )

// func ParseGrpcError(err error) error {
// 	s, ok := status.FromError(err)
// 	if !ok {
// 		return ErrParseError
// 	}

// 	return xerrors.New(s.Message())
// }

func Unwrap(err error) error {
	curerror := err
	for {
		err, ok := curerror.(xerrors.Wrapper)
		if !ok {
			return curerror
		}
		curerror = err.Unwrap()
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
