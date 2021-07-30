package dbcustom_db

type Conditions struct {
	Params     []ParamPair  	// 参数
	OrParams   []ParamPair  	// or参数
}

func (c *Conditions) Eq(column string, args ...interface{}) *Conditions {
	c.Where(column+" = ?", args)
	return c
}

func (c *Conditions) EqOr(column string, args ...interface{}) *Conditions {
	c.Or(column+" = ?", args)
	return c
}

func (c *Conditions) NotEq(column string, args ...interface{}) *Conditions {
	c.Where(column+" <> ?", args)
	return c
}

func (c *Conditions) NotEqOr(column string, args ...interface{}) *Conditions {
	c.Or(column+" <> ?", args)
	return c
}

func (c *Conditions) Gt(column string, args ...interface{}) *Conditions {
	c.Where(column+" > ?", args)
	return c
}

func (c *Conditions) GtOr(column string, args ...interface{}) *Conditions {
	c.Or(column+" > ?", args)
	return c
}

func (c *Conditions) Gte(column string, args ...interface{}) *Conditions {
	c.Where(column+" >= ?", args)
	return c
}

func (c *Conditions) GteOr(column string, args ...interface{}) *Conditions {
	c.Or(column+" >= ?", args)
	return c
}

func (c *Conditions) Lt(column string, args ...interface{}) *Conditions {
	c.Where(column+" < ?", args)
	return c
}

func (c *Conditions) LtOr(column string, args ...interface{}) *Conditions {
	c.Or(column+" < ?", args)
	return c
}

func (c *Conditions) Lte(column string, args ...interface{}) *Conditions {
	c.Where(column+" <= ?", args)
	return c
}

func (c *Conditions) LteOr(column string, args ...interface{}) *Conditions {
	c.Or(column+" <= ?", args)
	return c
}

func (c *Conditions) Like(column string, str string) *Conditions {
	c.Where(column+" LIKE ?", "%"+str+"%")
	return c
}

func (c *Conditions) LikeOr(column string, str string) *Conditions {
	c.Or(column+" LIKE ?", "%"+str+"%")
	return c
}

func (c *Conditions) Starting(column string, str string) *Conditions {
	c.Where(column+" LIKE ?", str+"%")
	return c
}

func (c *Conditions) StartingOr(column string, str string) *Conditions {
	c.Or(column+" LIKE ?", str+"%")
	return c
}

func (c *Conditions) Ending(column string, str string) *Conditions {
	c.Where(column+" LIKE ?", "%"+str)
	return c
}

func (c *Conditions) EndingOr(column string, str string) *Conditions {
	c.Or(column+" LIKE ?", "%"+str)
	return c
}

func (c *Conditions) In(column string, params interface{}) *Conditions {
	c.Where(column+" in (?) ", params)
	return c
}

func (c *Conditions) InOr(column string, params interface{}) *Conditions {
	c.Or(column+" in (?) ", params)
	return c
}

func (c *Conditions) NotIn(column string, params interface{}) *Conditions {
	c.Where(column+" not in (?) ", params)
	return c
}

func (c *Conditions) NotInOr(column string, params interface{}) *Conditions {
	c.Or(column+" not in (?) ", params)
	return c
}

func (c *Conditions) Where(query string, args ...interface{}) *Conditions {
	c.Params = append(c.Params, ParamPair{query, args})
	return c
}

func (c *Conditions) Or(query string, args ...interface{}) *Conditions {
	c.OrParams = append(c.OrParams, ParamPair{query, args})
	return c
}


func (c *Conditions) Build(db *DB) *DB {
	ret := db
	// where
	if len(c.Params) > 0 {
		for _, param := range c.Params {
			ret.Db = ret.Where(param.Query, param.Args...)
		}
	}

	// or
	if len(c.OrParams) > 0 {
		for _, param := range c.OrParams {
			ret.Db = ret.Or(param.Query, param.Args...)
		}
	}

	return ret
}