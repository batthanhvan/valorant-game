package lib

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"reflect"
	"runtime"

	// "io/ioutil"
	// "mime"
	// "net/http"
	// "path"
	// "reflect"
	// "runtime"
	"strconv"
	// "strings"
	// "time"
	// "github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	// "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	// "golang.org/x/xerrors"
	"github.com/batthanhvan/proto/pb"
)

// func SetBody(c *gin.Context) error {
// 	if c.Request.Method == http.MethodGet {
// 		return nil
// 	}
// 	buf := &bytes.Buffer{}

// 	_, err := buf.ReadFrom(c.Request.Body)
// 	if err != nil {
// 		return xerrors.Errorf("%w", e.ErrMissingBody)
// 	}
// 	c.Set("body", buf.Bytes())
// 	c.Request.Body = ioutil.NopCloser(buf)

// 	return nil
// }

// func GetBody(g *gin.Context, target interface{}) error {
// 	var bbody []byte
// 	var err error
// 	ibody, ok := g.Get("body")
// 	if !ok {
// 		return xerrors.Errorf("%w", e.ErrMissingBody)
// 	} else {
// 		bbody = ibody.(json.RawMessage)
// 	}

// 	err = json.Unmarshal(bbody, target)
// 	if err != nil {
// 		return xerrors.Errorf("%w", e.ErrMissingBody)
// 	}
// 	return nil

// }

// func ToUUID(id interface{}) (uuid.UUID, error) {
// 	switch tmp := id.(type) {
// 	case uuid.UUID:
// 		return tmp, nil
// 	case string:
// 		tmp2, err := uuid.Parse(tmp)
// 		if err != nil {
// 			return uuid.Nil, e.ErrIdInvalidFormat
// 		}
// 		return tmp2, nil
// 	default:
// 		return uuid.Nil, e.ErrIdInvalidFormat
// 	}
// }

// func MustGetRole(g *gin.Context) c.ROLE {
// 	r, ok := g.Get("role")
// 	if !ok {
// 		panic("not role")
// 	}
// 	ro, ok := r.(c.ROLE)
// 	if !ok {
// 		panic("not role")
// 	}
// 	return ro
// }

// func GetContentType(name *string) (*string, error) {
// 	ext := path.Ext(*name)
// 	typ := mime.TypeByExtension(ext)

// 	if !strings.Contains(typ, "image") {
// 		return nil, e.ErrInvalidFile
// 	}
// 	return &typ, nil
// }

// func RandomKey(name *string) *string {
// 	ext := path.Ext(*name)
// 	newKey := fmt.Sprintf("%s%s", uuid.New().String(), ext)
// 	return &newKey
// }

func ParseInt32Val(a string) int32 {
	if a == "" {
		return 0
	}
	v, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}
	i32 := int32(v)
	return i32
}

func ParseInt32Ptr(a string) *int32 {
	if a == "" {
		return nil
	}
	v, err := strconv.Atoi(a)
	if err != nil {
		return nil
	}
	i32 := int32(v)
	return &i32
}

// func ParseIntPtr(a string) *int {
// 	if a == "" {
// 		return nil
// 	}
// 	v, err := strconv.Atoi(a)
// 	if err != nil {
// 		return nil
// 	}
// 	return &v
// }

// func ParseUUID(a string) uuid.UUID {
// 	uid, err := uuid.Parse(a)
// 	if err != nil {
// 		return uuid.Nil
// 	}
// 	return uid
// }

func Pagination(offset, limit int32, total *int64) *pb.Pagination {
	if offset == 0 && limit == 0 && total == nil {
		return nil
	}
	return &pb.Pagination{
		Offset: offset,
		Limit:  limit,
		Total:  Int64Val(total),
	}
}

func GetFunctionName(funcname interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(funcname).Pointer()).Name()
}

func RecordError(span trace.Span, err error) {
	span.RecordError(fmt.Errorf("%+v", err))
	span.SetStatus(codes.Error, err.Error())
}

// var (
// 	aDay  = time.Hour * 24
// 	aWeek = aDay * 7
// )

// func GetTimeRange(n int32) (left int64, right int64) {
// 	now := time.Now()
// 	weekday := int64(now.Weekday())
// 	left = now.UnixMilli() - now.UnixMilli()%aDay.Milliseconds() - weekday*aDay.Milliseconds()
// 	if n == 0 {
// 		return left, now.UnixMilli()
// 	}
// 	left -= int64(n) * aWeek.Milliseconds()
// 	return left, left + aWeek.Milliseconds()
// }

// func ParseInt64Ptr(a string) *int64 {
// 	if a == "" {
// 		return nil
// 	}
// 	v, err := strconv.Atoi(a)
// 	if err != nil {
// 		return nil
// 	}
// 	i64 := int64(v)
// 	return &i64
// }

// func ParseInt64Val(a string) int64 {
// 	if a == "" {
// 		return 0
// 	}
// 	v, err := strconv.Atoi(a)
// 	if err != nil {
// 		return 0
// 	}
// 	i64 := int64(v)
// 	return i64
// }
