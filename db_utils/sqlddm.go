package db_utils

// 包含事务tx
type SqlCud struct {
	SelectCols 	[]string     	// 要保存的字段，如果为空，表示保存所有字段
	Model 		interface{}		// 实体
	Conditions
}

func NewSqlCud() *SqlCud {
	return &SqlCud{}
}

func (s *SqlCud) Cols(selectCols ...string) *SqlCud {
	if len(selectCols) > 0 {
		s.SelectCols = append(s.SelectCols, selectCols...)
	}
	return s
}

func (s *SqlCud) Value(model interface{}) *SqlCud {
	s.Model = model
	return s
}


func (s *SqlCud) Id(id interface{}) *SqlCud {
	s.Eq("id",id)
	return s
}

func (s *SqlCud) Build(db *DB) *DB {
	ret := s.Conditions.Build(db)

	if len(s.SelectCols) > 0 {
		ret.Db = ret.Select(s.SelectCols)
	}

	return ret
}

func (s *SqlCud) Create(db *DB) error {
	return s.Build(db).Create(s.Model).Error
}

func (s *SqlCud) Update(db *DB) error {
	tx := s.Build(db)
	if err := tx.First(s.Model).Error; err != nil {
		return err
	}else {
		return tx.Updates(s.Model).Error
	}
}

func (s *SqlCud) Delete(db *DB) error {
	return s.Build(db).Delete(s.Model).Error
}

//func (s *SqlCud) Transaction(fc	func(tx *DB) error) error {
//	return NewDB().DB.Transaction(func(tx *gorm.DB) error{
//		return fc(NewDB(tx))
//	})
//}