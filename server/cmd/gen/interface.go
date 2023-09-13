package main

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	FindOne(id int64) (*gen.T, error)
}
