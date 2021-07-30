package dbcustom_db

type SqlCnd struct {
	SelectCols []string     	// 要查询的字段，如果为空，表示查询所有字段
	Conditions
	Orders     []OrderByCol 	// 排序
	Paging     *Paging      	// 分页
	Preload    []PreloadPair  	// 预加载
	SqlSet     *SqlPair  		// 自定义
}

func NewSqlCnd() *SqlCnd {
	return &SqlCnd{}
}

func (s *SqlCnd) Cols(selectCols ...string) *SqlCnd {
	if len(selectCols) > 0 {
		s.SelectCols = append(s.SelectCols, selectCols...)
	}
	return s
}

func (s *SqlCnd) Asc(column string) *SqlCnd {
	s.Orders = append(s.Orders, OrderByCol{Column: column, Asc: true})
	return s
}

func (s *SqlCnd) Desc(column string) *SqlCnd {
	s.Orders = append(s.Orders, OrderByCol{Column: column, Asc: false})
	return s
}

func (s *SqlCnd) Limit(limit int) *SqlCnd {
	s.Page(1, limit)
	return s
}

func (s *SqlCnd) Page(page, limit int) *SqlCnd {
	if s.Paging == nil {
		s.Paging = &Paging{Page: page, Limit: limit}
	} else {
		s.Paging.Page = page
		s.Paging.Limit = limit
	}
	return s
}



func (s *SqlCnd) Build(db *DB) *DB {
	ret := s.Conditions.Build(db)

	if len(s.SelectCols) > 0 {
		ret.Db = ret.Select(s.SelectCols)
	}

	// preload
	//preload.args = 1 : uuid in ()
	//preload.args = 2 : "uuid in ()" = 'uuid in ()'
	//preload.args = 3 : id in ('uuid in ()','uuid in ()' ...)
	if len(s.Preload) > 0 {
		for _, param := range s.Preload {
			ret.Db = ret.Preload(param.Model, param.Args)
		}
	}

	// order
	if len(s.Orders) > 0 {
		for _, order := range s.Orders {
			if order.Asc {
				ret.Db = ret.Order(order.Column + " ASC")
			} else {
				ret.Db = ret.Order(order.Column + " DESC")
			}
		}
	}

	// limit
	if s.Paging != nil && s.Paging.Limit > 0 {
		ret.Db = ret.Limit(s.Paging.Limit)
	}

	// offset
	if s.Paging != nil && s.Paging.Offset() > 0 {
		ret.Db = ret.Offset(s.Paging.Offset())
	}
	return ret
}

func (s *SqlCnd) Find(db *DB, out interface{}) error {
	return s.Build(db).Find(out).Error
}

func (s *SqlCnd) FindOne(db *DB, out interface{}) error {
	return s.Build(db).First(out).Error
}

func (s *SqlCnd) Custom(db *DB) *DB {
	ret := db
	if s.SqlSet != nil {
		if len(s.SqlSet.Args) > 0 {
			ret.Db = ret.Raw(s.SqlSet.Sql,s.SqlSet.Args)
		}else {
			ret.Db = ret.Raw(s.SqlSet.Sql)
		}
	}

	return ret
}

func (s *SqlCnd) FindSql(db *DB, out interface{}) error {
	return s.Custom(db).Scan(out).Error
}

// 查询count， isCount启动查询条件
func (s *SqlCnd) Count(db *DB, isCount bool) (int64, error) {
	ret := db

	if isCount {
		ret = s.Conditions.Build(ret)
		// preload
		//preload.args = 1 : uuid in ()
		//preload.args = 2 : "uuid in ()" = 'uuid in ()'
		//preload.args = 3 : id in ('uuid in ()','uuid in ()' ...)
		if len(s.Preload) > 0 {
			for _, param := range s.Preload {
				ret.Db = ret.Preload(param.Model, param.Args)
			}
		}
	}

	var count int64
	if err := ret.Count(&count).Error; err != nil {
		return 0, err
	}else {
		return count,nil
	}

}

