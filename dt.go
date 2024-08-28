package gormdt

import (
	"fmt"
	"strings"
)

func FilterAndPaginate(param FilterAndPaginateParam) (DataTableResponse, error) {
	var totalRecords int64
	var filteredRecords int64
	var data []map[string]interface{}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", param.TableName)
	err := param.DB.Raw(countQuery).Scan(&totalRecords).Error
	if err != nil {
		return DataTableResponse{}, err
	}

	searchCondition := fmt.Sprintf(" WHERE %s.deleted_at IS NULL", param.TableName)
	if param.Request.Search != "" {
		conditions := []string{}
		for _, field := range param.Fields {
			conditions = append(conditions, fmt.Sprintf("LOWER(%s) LIKE LOWER('%%%s%%')", field, strings.ToLower(param.Request.Search)))
		}
		searchCondition = searchCondition + " AND ( " + strings.Join(conditions, " OR ") + ")"
	}

	filteredCountQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s%s", param.TableName, searchCondition)
	err = param.DB.Raw(filteredCountQuery).Scan(&filteredRecords).Error
	if err != nil {
		return DataTableResponse{}, err
	}

	dataQuery := fmt.Sprintf("SELECT * FROM %s%s LIMIT %d OFFSET %d", param.TableName, searchCondition, param.Request.Length, param.Request.Start)
	err = param.DB.Raw(dataQuery).Scan(&data).Error
	if err != nil {
		return DataTableResponse{}, err
	}

	return DataTableResponse{
		Draw:            param.Request.Draw,
		RecordsTotal:    int(totalRecords),
		RecordsFiltered: int(filteredRecords),
		Data:            convertMapToInterfaceSlice(data),
	}, nil
}

func convertMapToInterfaceSlice(data []map[string]interface{}) []interface{} {
	result := make([]interface{}, len(data))
	for i, d := range data {
		result[i] = d
	}
	return result
}

func FilterAndPaginateCustomQuery(param FilterAndPaginateCustomQueryParam) (DataTableResponse, error) {
	var totalRecords int64
	var filteredRecords int64
	var data []map[string]interface{}

	searchCondition := ""
	if param.Where != "" {
		searchCondition = " WHERE " + param.Where
	}

	if param.Request.Search != "" {
		conditions := []string{}
		for _, field := range param.Fields {
			conditions = append(conditions, fmt.Sprintf("LOWER(%s) LIKE LOWER('%%%s%%')", field, strings.ToLower(param.Request.Search)))
		}
		if searchCondition != "" {
			searchCondition += " AND (" + strings.Join(conditions, " OR ") + ")"
		} else {
			searchCondition = " WHERE " + strings.Join(conditions, " OR ")
		}
	}

	if param.GroupBy != "" {
		searchCondition += " GROUP BY " + param.GroupBy
	}

	if param.Having != "" {
		searchCondition += " HAVING " + param.Having
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s%s) AS total", param.BaseQuery, searchCondition)
	err := param.DB.Raw(countQuery).Scan(&totalRecords).Error
	if err != nil {
		return DataTableResponse{}, err
	}

	filteredCountQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s%s) AS filtered", param.BaseQuery, searchCondition)
	err = param.DB.Raw(filteredCountQuery).Scan(&filteredRecords).Error
	if err != nil {
		return DataTableResponse{}, err
	}

	dataQuery := fmt.Sprintf("%s%s LIMIT %d OFFSET %d", param.BaseQuery, searchCondition, param.Request.Length, param.Request.Start)
	fmt.Println("\n\n\n", dataQuery)
	err = param.DB.Raw(dataQuery).Scan(&data).Error
	if err != nil {
		return DataTableResponse{}, err
	}

	return DataTableResponse{
		Draw:            param.Request.Draw,
		RecordsTotal:    int(totalRecords),
		RecordsFiltered: int(filteredRecords),
		Data:            convertMapToInterfaceSlice(data),
	}, nil
}
