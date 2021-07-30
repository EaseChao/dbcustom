package db_utils

// params[where]
type QueryPw string
const (
	EQ       QueryPw = "eq"
	NEQ      QueryPw = "neq"
	GT       QueryPw = "gt"
	GTE      QueryPw = "gte"
	LT       QueryPw = "lt"
	LTE      QueryPw = "lte"
	LIKE     QueryPw = "like"
	STARTING QueryPw = "starting"
	ENDING   QueryPw = "ending"
	IN       QueryPw = "in"
	NIN      QueryPw = "nin"
)
// params[where]
type QueryPor string
const (
	EQ_OR       QueryPor = "eq_or"
	NEQ_OR      QueryPor = "neq_or"
	GT_OR       QueryPor = "gt_or"
	GTE_OR      QueryPor = "gte_or"
	LT_OR       QueryPor = "lt_or"
	LTE_OR      QueryPor = "lte_or"
	LIKE_OR     QueryPor = "like_or"
	STARTING_OR QueryPor = "starting_or"
	ENDING_OR   QueryPor = "ending_or"
	IN_OR       QueryPor = "in_or"
	NIN_OR      QueryPor = "nin_or"
)
// params[order/limit]
type QueryPol string
const (
	ASC      QueryPol = "asc"
	DESC     QueryPol = "desc"
	LIMIT    QueryPol = "limit"
	PAGE     QueryPol = "page"
)

// 分页请求数据
type Paging struct {
	Page  int   `json:"page" form:"page"`  // 页码
	Limit int   `json:"limit" form:"limit"` // 每页条数
	Total int64 `json:"total" form:"total"` // 总数据条数
}

func (p *Paging) Offset() int {
	offset := 0
	if p.Page > 0 {
		offset = (p.Page - 1) * p.Limit
	}
	return offset
}

func (p *Paging) TotalPage() int {
	if p.Total == 0 || p.Limit == 0 {
		return 0
	}
	totalPage := int(p.Total) / p.Limit
	if int(p.Total)%p.Limit > 0 {
		totalPage = totalPage + 1
	}
	return totalPage
}

type ParamPair struct {
	Query string        // 查询
	Args  []interface{} // 参数
}

// 排序信息
type OrderByCol struct {
	Column string // 排序字段
	Asc    bool   // 是否正序
}

type PreloadPair struct {
	Args		string	 	// 参数
	Model		string		// 连接表
}

type SqlPair struct {
	Sql		string			// sql
	Args	[]interface{}	// args
}