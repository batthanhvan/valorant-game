package services

import (
	"database/sql"
	"net/http"

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

func PostNewReport(username string, reportCategory string, matchID string, reportDetail string) (string, int) {
	db, err := sql.Open(lib.DRIVER_NAME, db.ConStr)
	if err != nil {
		return "We've encountered an error, please refresh the page.", http.StatusBadRequest
	}
	defer db.Close()

	checkDuplicateReport, err := db.Query("SELECT 1 FROM reports WHERE username=? AND reportCategory=? AND matchID=?",
		username, reportCategory, matchID)
	if err != nil {
		return "We've encountered an error, please refresh the page.", http.StatusBadRequest
	}

	_, err = db.Query(reports.PostNewReport, username, reportCategory, matchID, reportDetail)
	if err != nil {
		return "Please check for the correct username and match ID.", http.StatusOK
	}

	if checkDuplicateReport.Next() {
		return "This player's behaviour in this match has already been reported.", http.StatusOK
	} else {

		checkPlayerInMatch, err := db.Query(reports.CheckPlayerInMatch, username, matchID)
		if err != nil {
			return "We've encountered an error, please refresh the page.", http.StatusBadRequest
		}

		if !checkPlayerInMatch.Next() {
			return "The player you report was not present in this match.", http.StatusOK
		}

		return "Report sent. We will review this soon. Thank you for your contribution", http.StatusOK
	}
}
