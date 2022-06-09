package services

import (
	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/db/leaderboard"
	"github.com/batthanhvan/src/lib"

	"golang.org/x/xerrors"
)

func GetLeaderBoard(req *pb.GetRequest) (*pb.LeaderboardGetResponse_Data, error) {

	limit := lib.ParseInt32Val(req.Limit)
	offset := lib.ParseInt32Val(req.Offset)

	total, err := leaderboard.PlayerCount()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	res, err := leaderboard.GetLeaderboardDetail(&db.Search{
		Limit: int(limit),
		Skip:  int(offset),
		Query: req.Query,
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return &pb.LeaderboardGetResponse_Data{
		Result:     res,
		Pagination: lib.Pagination(offset, limit, total),
	}, nil
}
