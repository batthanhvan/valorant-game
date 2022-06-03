package reports

import (
	"database/sql"

	pb "github.com/batthanhvan/proto/pb"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/lib"
	_ "github.com/go-sql-driver/mysql"
)

func ReportCount(search *db.Search) (*int64, error) {

	db, err := sql.Open(lib.DRIVER_NAME, db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	search.Query = "%" + search.Query + "%"
	total := db.QueryRow("select count(*) from reports WHERE username like ? LIMIT ?,?", search.Query, search.Skip, search.Limit)

	var r int64
	err = total.Scan(&r)
	if err != nil {
		panic(err.Error())
	}

	return &r, nil
}

func AllReport(search *db.Search) ([]*pb.Report, error) {

	db, err := sql.Open(lib.DRIVER_NAME, db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query(QueryAllString, search.Skip, search.Limit)

	if err != nil {
		panic(err.Error())
	}

	var report Report
	rr := make([]Report, 0)
	for results.Next() {

		err = results.Scan(&report.ReportCategory, &report.ReportDetail, &report.ReportDate, &report.RecordLink, &report.UserName)

		if err != nil {
			panic(err.Error())
		}
		rr = append(rr, report)
	}

	arr := make([]*pb.Report, 0)
	for _, report := range rr {
		arr = append(arr, ConvertReportToProto(report))
	}
	return arr, nil
}

func ListReports(search *db.Search) ([]*pb.Report, error) {

	db, err := sql.Open(lib.DRIVER_NAME, db.ConStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	search.Query = "%" + search.Query + "%"
	results, err := db.Query(QueryString, search.Query, search.Skip, search.Limit)

	if err != nil {
		panic(err.Error())
	}

	var report Report
	rr := make([]Report, 0)
	for results.Next() {

		err = results.Scan(&report.ReportCategory, &report.ReportDetail, &report.ReportDate, &report.RecordLink, &report.UserName)

		if err != nil {
			panic(err.Error())
		}
		rr = append(rr, report)
	}

	arr := make([]*pb.Report, 0)
	for _, v := range rr {
		arr = append(arr, ConvertReportToProto(v))
	}
	return arr, nil

}

func ConvertReportToProto(p Report) *pb.Report {
	ppb := &pb.Report{}

	ppb.ReportCategory = p.ReportCategory.String
	ppb.ReportDetail = p.ReportDetail.String
	ppb.ReportDate = p.ReportDate.String
	ppb.RecordLink = p.RecordLink.String
	ppb.UserName = p.UserName.String

	return ppb
}
