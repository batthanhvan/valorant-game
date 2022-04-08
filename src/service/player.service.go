package service

import (
	"github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/db/mysql"
	"github.com/batthanhvan/src/lib"
	"golang.org/x/xerrors"
)

func GetByUserName(req *pb.PlayerGetRequest) (*pb.PlayerGetResponse_Data, error) {

	limit := lib.ParseInt32Val(req.Limit)
	offset := lib.ParseInt32Val(req.Offset)

	total, err := mysql.TotalPlayer(&db.Search{
		Limit: int(limit),
		Skip:  int(offset),
		Query: req.Query,
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	res, err := mysql.ListPlayers(&db.Search{
		Limit: int(limit),
		Skip:  int(offset),
		Query: req.Query,
	})

	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return &pb.PlayerGetResponse_Data{
		Result:     res,
		Pagination: lib.Pagination(offset, limit, total),
	}, nil
}
