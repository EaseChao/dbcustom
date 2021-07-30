package db_utils

import (
	"github.com/EaseChao/dbcustom"
	"github.com/EaseChao/dbcustom/string_utils"
	"strings"
)

type QueryParams struct {
	Context
	SqlCnd
}

func NewQueryParams(ctx Context) *QueryParams {
	return &QueryParams{Context: ctx}
}

func (q *QueryParams) getValue(name string) (str string, b bool) {
	value := q.GetValue(name)
	if len(value) > 0 {
		return value,true
	}
	return "",false
}

// 获取分页参数集
func (q *QueryParams) GetPgList(ita interface{}) *QueryParams {
	return q.PageByReq().GetList(ita)
}

// 获取队列参数集
func (q *QueryParams) GetList(ita interface{}) *QueryParams {
	tags := dbcustom.NewAllOfReflect(ita).GetAllKeyTags()
	for _,tag := range tags {
		if string_utils.IsNotBlank(tag[dbcustom.QUERY]){ // where
			querys := strings.Split(tag[dbcustom.QUERY], ",")

			if string_utils.IsNotBlank(tag[dbcustom.JOIN]){ // join
				//joins := strings.Split(tag[dbcustom.JOIN], ",")
				q.GetPj(QueryPw(querys[0]),querys[1],tag[dbcustom.JSON],tag[dbcustom.JOIN])
				continue
			}

			if strings.Contains(querys[0],"_") {
				q.GetPor(QueryPor(querys[0]),querys[1],tag[dbcustom.JSON])
			}

			q.GetPw(QueryPw(querys[0]),querys[1],tag[dbcustom.JSON])
		}else { // order &limit

			q.GetPol(QueryPol(tag[dbcustom.JSON]),tag[dbcustom.JSON])
		}
	}
	return q
}

// 组装where条件
func (q *QueryParams) GetPw(qp QueryPw, column, name string) *QueryParams {
	if value, b := q.getValue(name); string_utils.IsNotBlank(column) && b {
		switch qp {
			case EQ:		q.Eq(column,value)
			case NEQ:		q.NotEq(column,value)
			case GT:		q.Gt(column,value)
			case GTE:		q.Gte(column,value)
			case LT:		q.Lt(column,value)
			case LTE:		q.Lte(column,value)
			case LIKE:		q.Like(column,value)
			case STARTING:	q.Starting(column,value)
			case ENDING:	q.Ending(column,value)
			case IN:		q.In(column,strings.Split(value,","))
			case NIN:		q.NotIn(column,strings.Split(value,","))
		}
	}
	return q
}

// 组装or条件
func (q *QueryParams) GetPor(qp QueryPor, column, name string) *QueryParams {
	if value, b := q.getValue(name); string_utils.IsNotBlank(column) && b {
		switch qp {
			case EQ_OR:			q.EqOr(column,value)
			case NEQ_OR:		q.NotEqOr(column,value)
			case GT_OR:			q.GtOr(column,value)
			case GTE_OR:		q.GteOr(column,value)
			case LT_OR:			q.LtOr(column,value)
			case LTE_OR:		q.LteOr(column,value)
			case LIKE_OR:		q.LikeOr(column,value)
			case STARTING_OR:	q.StartingOr(column,value)
			case ENDING_OR:		q.EndingOr(column,value)
			case IN_OR:			q.InOr(column,strings.Split(value,","))
			case NIN_OR:		q.NotInOr(column,strings.Split(value,","))
		}
	}
	return q
}

// 组装order & limit 条件
func (q *QueryParams) GetPol(qp QueryPol, name string) *QueryParams {
	if value, b := q.getValue(name); b {
		switch qp {
			case ASC:		q.Asc(value)
			case DESC:		q.Desc(value)
			case LIMIT:
				if q.Paging == nil {
					q.Limit(dbcustom.ToInt(value))
				} else {
					q.Page(q.Paging.Page,dbcustom.ToInt(value))
				}
			case PAGE:
				if q.Paging == nil {
					q.Page(dbcustom.ToInt(value),1)
				} else {
					q.Page(dbcustom.ToInt(value),q.Paging.Limit)
				}
		}
	}
	return q
}

// 组装join条件
func (q *QueryParams) GetPj(qp QueryPw, column, name, join string) *QueryParams {
	if string_utils.IsBlank(join) {
		return q
	}
	if value, b := q.getValue(name); string_utils.IsNotBlank(column) && b {
		preload := PreloadPair{Model: join}
		switch qp {
			case EQ:		preload.Args = column + " =' " + value + " ' "
			case NEQ:		preload.Args = column + " !=' " + value + " ' "
			case GT:		preload.Args = column + " >' " + value + " ' "
			case GTE:		preload.Args = column + " >=' " + value + " ' "
			case LT:		preload.Args = column + " <' " + value + " ' "
			case LTE:		preload.Args = column + " <=' " + value + " ' "
			case LIKE:		preload.Args = column + " like '%" + value + "%' "
			case STARTING:	preload.Args = column + " like '% " + value + " ' "
			case ENDING:	preload.Args = column + " like ' " + value + " %' "
			case IN:
				args := ""
				for _,val := range strings.Split(value, ",") {
					if args != ""{
						args += " , "
					}
					args = args + "'" + val + "'"
				}
				preload.Args = column + " in ( " + args + " ) "
			case NIN:
				args := ""
				for _,val := range strings.Split(value, ",") {
					if args != ""{
						args += " , "
					}
					args = args + "'" + val + "'"
				}
				preload.Args = column + " not in ( " + args + " ) "
		}
		q.Preload = append(q.Preload, preload)
	}else {
		q.Preload = append(q.Preload, PreloadPair{Model: join})
	}
	return q
}

//func (q *QueryParams) PjByReq(join string) *QueryParams {
//	return q.GetPj("","","",join)
//}

func (q *QueryParams) EqByReq(column string) *QueryParams {
	if value,ok := q.getValue(column); ok {
		q.Eq(column, value)
	}
	return q
}

func (q *QueryParams) NotEqByReq(column string) *QueryParams {
	if value,ok := q.getValue(column); ok {
		q.NotEq(column, value)
	}
	return q
}

func (q *QueryParams) GtByReq(column string) *QueryParams {
	if value,ok := q.getValue(column); ok {
		q.Gt(column, value)
	}
	return q
}

func (q *QueryParams) GteByReq(column string) *QueryParams {
	if value,ok := q.getValue(column); ok {
		q.Gte(column, value)
	}
	return q
}

func (q *QueryParams) LtByReq(column string) *QueryParams {
	if value,ok := q.getValue(column); ok {
		q.Lt(column, value)
	}
	return q
}

func (q *QueryParams) LteByReq(column string) *QueryParams {
	if value,ok := q.getValue(column); ok {
		q.Lte(column, value)
	}
	return q
}

func (q *QueryParams) LikeByReq(column string) *QueryParams {
	if value,ok := q.getValue(column); ok {
		q.Like(column, value)
	}
	return q
}

func (q *QueryParams) PageByReq() *QueryParams {
	if q.Paging == nil {
		return q.Page(1, 20)
	}
	return q
}

func (q *QueryParams) Asc(column string) *QueryParams {
	q.Orders = append(q.Orders, OrderByCol{Column: column, Asc: true})
	return q
}

func (q *QueryParams) Desc(column string) *QueryParams {
	q.Orders = append(q.Orders, OrderByCol{Column: column, Asc: false})
	return q
}

func (q *QueryParams) Limit(limit int) *QueryParams {
	q.Page(1, limit)
	return q
}

func (q *QueryParams) Page(page, limit int) *QueryParams {
	if q.Paging == nil {
		q.Paging = &Paging{Page: page, Limit: limit}
	} else {
		q.Paging.Page = page
		q.Paging.Limit = limit
	}
	return q
}

func (q *QueryParams) CreateSql(sql string, args ...interface{}) *QueryParams {
	q.SqlSet = &SqlPair{Sql: sql, Args: args}
	return q
}
