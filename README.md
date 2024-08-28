# gormdt

`gormdt` is a Golang package that simplifies data filtering and pagination when working with GORM. It provides two main functions, `FilterAndPaginate` and `FilterAndPaginateCustomQuery`, which allow you to easily filter, search, and paginate database records.

## Installation

To install `gormdt`, run:

```bash
go get github.com/sanda0/gormdt
```

## Usage

### Basic Usage: `FilterAndPaginate`

This function filters and paginates data from a specific table based on the search criteria provided. 

#### Example

```go
package main

import (
	"fmt"
	"github.com/yourusername/gormdt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize GORM DB connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Define the filter and paginate parameters
	params := gormdt.FilterAndPaginateParam{
		DB:        db,
		TableName: "users",
		Fields:    []string{"name", "email"},
		Request: gormdt.DataTableRequest{
			Draw:   1,
			Start:  0,
			Length: 10,
			Search: "john",
		},
	}

	// Call the FilterAndPaginate function
	response, err := gormdt.FilterAndPaginate(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output the results
	fmt.Printf("Total Records: %d\n", response.RecordsTotal)
	fmt.Printf("Filtered Records: %d\n", response.RecordsFiltered)
	fmt.Printf("Data: %+v\n", response.Data)
}
```

### Advanced Usage: `FilterAndPaginateCustomQuery`

This function allows for more advanced filtering and pagination using a custom SQL query.

#### Example

```go
package main

import (
	"fmt"
	"github.com/yourusername/gormdt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize GORM DB connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Define the custom filter and paginate parameters
	params := gormdt.FilterAndPaginateCustomQueryParam{
		DB:        db,
		BaseQuery: "SELECT name, email FROM users",
		Where:     "status = 'active'",
		Fields:    []string{"name", "email"},
		Request: gormdt.DataTableRequest{
			Draw:   1,
			Start:  0,
			Length: 10,
			Search: "john",
		},
	}

	// Call the FilterAndPaginateCustomQuery function
	response, err := gormdt.FilterAndPaginateCustomQuery(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output the results
	fmt.Printf("Total Records: %d\n", response.RecordsTotal)
	fmt.Printf("Filtered Records: %d\n", response.RecordsFiltered)
	fmt.Printf("Data: %+v\n", response.Data)
}
```

## Data Structures

### `DataTableRequest`

This struct represents the request for data filtering and pagination.

```go
type DataTableRequest struct {
	Draw   int    `json:"draw" form:"draw"`
	Start  int    `json:"start" form:"start"`
	Length int    `json:"length" form:"length"`
	Search string `json:"search" form:"search[value]"`
}
```

### `DataTableResponse`

This struct represents the response after filtering and pagination.

```go
type DataTableResponse struct {
	Draw            int         `json:"draw"`
	RecordsTotal    int         `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
}
```

### Parameters

#### `FilterAndPaginateParam`

Parameters for the `FilterAndPaginate` function.

```go
type FilterAndPaginateParam struct {
	DB        *gorm.DB
	TableName string
	Fields    []string
	Request   DataTableRequest
}
```

#### `FilterAndPaginateCustomQueryParam`

Parameters for the `FilterAndPaginateCustomQuery` function.

```go
type FilterAndPaginateCustomQueryParam struct {
	DB        *gorm.DB
	BaseQuery string
	Where     string
	GroupBy   string
	Having    string
	Fields    []string
	Request   DataTableRequest
}
```
