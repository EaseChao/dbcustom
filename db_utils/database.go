package db_utils

import (
	"gorm.io/gorm"
)

/**
	不同框架随意亲自搭建db
	需要用上封装方法，可以自己重写当前go
*/

type Db = gorm.DB //别名类型(类型还是gorm.DB)，控制"gorm.io/gorm"的注入
// if no modal ,DB func must return
type DB struct {
	*Db
}

// you can update this func
func NewDB(db ...*gorm.DB) *DB {
	if len(db) > 0 {
		return &DB{db[0]}
	}
	return &DB{}	//此处可以设置默认db
}

// you can update this func or remove
func (db *DB) Model(value interface{}) *DB {
	return &DB{db.Db.Model(value)}
}

// do transaction
func (db *DB) Transaction(fc func(tx *Db) error) error {
	return db.Db.Transaction(func(tx *gorm.DB) error{
		return fc(tx)
	})
}