package services

import (
	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/db/reports"
	"github.com/batthanhvan/src/lib"
	"golang.org/x/xerrors"
)

func GetAllReports(req *pb.GetRequest) (*pb.ReportGetResponse_Data, error) {

	limit := lib.ParseInt32Val(req.Limit)
	offset := lib.ParseInt32Val(req.Offset)

	reportNum, err := reports.ReportCount(&db.Search{
		Limit: int(limit),
		Skip:  int(offset),
		Query: req.Query,
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	result, err := reports.AllReport(&db.Search{
		Limit: int(limit),
		Skip:  int(offset),
		Query: req.Query,
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return &pb.ReportGetResponse_Data{
		Result:     result,
		Pagination: lib.Pagination(offset, limit, reportNum),
	}, nil
}

func GetReportByUserName(req *pb.GetRequest) (*pb.ReportGetResponse_Data, error) {

	limit := lib.ParseInt32Val(req.Limit)
	offset := lib.ParseInt32Val(req.Offset)

	total, err := reports.ReportCount(&db.Search{
		Limit: int(limit),
		Skip:  int(offset),
		Query: req.Query,
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	res, err := reports.ListReports(&db.Search{
		Limit: int(limit),
		Skip:  int(offset),
		Query: req.Query,
	})

	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return &pb.ReportGetResponse_Data{
		Result:     res,
		Pagination: lib.Pagination(offset, limit, total),
	}, nil
}
