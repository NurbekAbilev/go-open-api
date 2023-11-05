package pagination

import (
	"context"
	"database/sql"
	"fmt"
)

const LimitPerPage = 10

type PaginatedData[T any] struct {
	CurentPage    int
	AmountOfPages int
	Data          []T
}

type PaginationRequest struct {
	PerPageAmount int
	CurrentPage   int
}

func PaginateQuery(ctx context.Context, db *sql.DB, query string, pg PaginationRequest, args ...any) (pageCount int, rows *sql.Rows, err error) {
	countQuery := "select count(1) as count from (" + query + ") t"

	var allCount int
	err = db.QueryRowContext(ctx, countQuery).Scan(&allCount)
	if err != nil {
		return 0, nil, err
	}

	offset := pg.PerPageAmount * (pg.CurrentPage - 1)
	mainQuery := fmt.Sprintf("select * from (%s) t limit %d offset %d", query, pg.PerPageAmount, offset)

	rows, err = db.QueryContext(ctx, mainQuery, args...)
	if err != nil {
		return 0, nil, err
	}

	pageCount = (allCount / pg.PerPageAmount) + 1

	return pageCount, rows, nil
}
