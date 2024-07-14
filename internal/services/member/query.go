package member

import (
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm/clause"
)

/**
Converter HTTP request query to repository structure
*/

func ConvertQueryToLimit(r *http.Request) int {
	var limit = 10

	limitQuery := r.URL.Query().Get("limit")

	if r.URL.Query().Get("limit") != "" {
		limitQueryInt, err := strconv.Atoi(limitQuery)
		if err != nil {
			return limit
		}

		return limitQueryInt
	}

	return limit
}

func ConvertQueryToFilter(r *http.Request) map[string]string {
	var filter = make(map[string]string)

	for key, val := range r.URL.Query() {
		for _, allowKey := range AllowedFilterQuery {
			if allowKey == key {
				filter[key+" LIKE ?"] = "%" + val[0] + "%"
			}
		}
	}

	return filter
}

func ConvertQueryToOrder(r *http.Request) clause.OrderByColumn {
	var isDescOrder = false
	var orderBy = "id"

	order := strings.ToLower(r.URL.Query().Get("order"))
	by := strings.ToLower(r.URL.Query().Get("by"))

	if order == "desc" {
		isDescOrder = true
	}

	if by != "" {
		orderBy = by
	}

	return clause.OrderByColumn{Column: clause.Column{Name: orderBy}, Desc: isDescOrder}
}
