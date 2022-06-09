package services

import (
	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/db/matches"
	"github.com/batthanhvan/src/lib"
	"golang.org/x/xerrors"
)

func GetMatchByUsername(req *pb.GetRequest) (*pb.MatchGetResponse_Data, error) {

	limit := lib.ParseInt32Val(req.Limit)
	offset := lib.ParseInt32Val(req.Offset)

	total, err := matches.MatchCount(&db.Search{
		Query: req.Query,
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	res, err := matches.MatchesByUsername(&db.Search{
		Limit: int(limit),
		Skip:  int(offset),
		Query: req.Query,
	})

	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return &pb.MatchGetResponse_Data{
		Result:     res,
		Pagination: lib.Pagination(offset, limit, total),
	}, nil
}
