// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: revenue.sql

package sqlc

import (
	"context"
)

const getAllRevenues = `-- name: GetAllRevenues :many
SELECT month, revenue FROM revenue
`

func (q *Queries) GetAllRevenues(ctx context.Context) ([]Revenue, error) {
	rows, err := q.db.Query(ctx, getAllRevenues)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Revenue
	for rows.Next() {
		var i Revenue
		if err := rows.Scan(&i.Month, &i.Revenue); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}