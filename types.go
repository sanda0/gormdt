package gormdt

import "gorm.io/gorm"

type DataTableRequest struct {
	Draw   int    `json:"draw" form:"draw"`
	Start  int    `json:"start" form:"start"`
	Length int    `json:"length" form:"length"`
	Search string `json:"search" form:"search[value]"`

	//TODO sorting
	// Order   []DataTableOrder  `json:"order" form:"order"`
	// Columns []DataTableColumn `json:"columns" form:"columns"`
}

type DataTableSearch struct {
	Value string `json:"value" form:"value"`
	Regex bool   `json:"regex" form:"regex"`
}

type DataTableOrder struct {
	Column int    `json:"column" form:"column"`
	Dir    string `json:"dir" form:"dir"`
}

type DataTableColumn struct {
	Data       string          `json:"data" form:"data"`
	Name       string          `json:"name" form:"name"`
	Searchable bool            `json:"searchable" form:"searchable"`
	Orderable  bool            `json:"orderable" form:"orderable"`
	Search     DataTableSearch `json:"search" form:"search"`
}

type DataTableResponse struct {
	Draw            int         `json:"draw"`
	RecordsTotal    int         `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
}

//params

type FilterAndPaginateParam struct {
	DB        *gorm.DB
	TableName string
	Fields    []string
	Request   DataTableRequest
}

type FilterAndPaginateCustomQueryParam struct {
	DB        *gorm.DB
	BaseQuery string
	Where     string
	GroupBy   string
	Having    string
	Fields    []string
	Request   DataTableRequest
}
