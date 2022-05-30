package service

import (
	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/db/mysql"
	"github.com/batthanhvan/src/lib"
	"golang.org/x/xerrors"
)

func GetByUserName(req *pb.GetRequest) (*pb.PlayerGetResponse_Data, error) {

	limit := lib.ParseInt32Val(req.Limit)
	offset := lib.ParseInt32Val(req.Offset)

	// total, err := mysql.PlayerCount(&db.Search{
	// 	Limit: int(limit),
	// 	Skip:  int(offset),
	// 	Query: req.Query,
	// })
	// if err != nil {
	// 	err = xerrors.Errorf("%w", err)
	// 	return nil, err
	// }

	res, err := mysql.GetPlayer(&db.Search{
		Limit: int(limit),
		Skip:  int(offset),
		Query: req.Query,
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return &pb.PlayerGetResponse_Data{
		Result: res,
		// Pagination: lib.Pagination(offset, limit, total),
	}, nil
}

func ModifyUser(req *pb.PostModifyUserReq) (*pb.PlayerGetResponse_Data, error) {

	res, status, err := mysql.ModifyPlayer(req.Username, req.Playername, req.Tagline)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return &pb.PlayerGetResponse_Data{
		Result:       res,
		ResultStatus: status,
	}, nil
}
