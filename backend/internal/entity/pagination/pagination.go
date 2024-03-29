package pagination

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"go.uber.org/zap"
)

type Pagination struct {
	HasNext     bool   `json:"hasNext"`
	HasPrevious bool   `json:"hasPrevious"`
	CountPages  uint64 `json:"countPages"`
	Limit       uint64 `json:"limit"`
	Page        uint64 `json:"page"`
	TotalItems  uint64 `json:"totalItems"`
}

func NewPagination(pag *Pagination) *Pagination {
	return &Pagination{
		HasNext:     pag.HasNext,
		HasPrevious: pag.HasPrevious,
		CountPages:  pag.CountPages,
		Limit:       pag.Limit,
		Page:        pag.Page,
		TotalItems:  pag.TotalItems,
	}
}

func ApplyPagination(sqlQuery string, page uint64, limit uint64) string {
	offset := (page - 1) * limit
	sqlQuery += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
	return sqlQuery
}

func GetTotalItems(ctx context.Context, db *sql.DB, sqlQuery string) (uint64, error) {
	var totalItems uint64
	err := db.QueryRowContext(ctx, sqlQuery).Scan(&totalItems)
	if err != nil {
		logger.Log.Debug(
			"error func GetTotalItems, method QueryRowContext by path internal/entity/pagination/pagination.go",
			zap.Error(err))
		return 0, err
	}
	return totalItems, nil
}

func getCountPages(limit uint64, totalItems uint64) uint64 {
	return (totalItems + limit - 1) / limit
}

func GetPagination(limit uint64, page uint64, totalItems uint64) *Pagination {
	hasPrevious := page > 1
	hasNext := (page * limit) < totalItems
	constPages := getCountPages(limit, totalItems)
	paging := NewPagination(&Pagination{
		HasNext:     hasNext,
		HasPrevious: hasPrevious,
		CountPages:  constPages,
		Limit:       limit,
		Page:        page,
		TotalItems:  totalItems,
	})
	return paging
}
