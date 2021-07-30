package db_utils

import (
	"github.com/EaseChao/dbcustom"
)

type DdmParams struct {
	Context
	SqlCud
}

func NewDdmParams(ctx Context) *DdmParams {
	return &DdmParams{Context:ctx}
}

// 参数与实体不对称，出err，但是依然能赋值
func (d *DdmParams) postValue(pojo interface{}) error {
	return d.PostValue(pojo)
}

func (d *DdmParams) Set(pojo interface{}) *DdmParams {
	_ = d.postValue(pojo)
	d.SetValue(pojo)
	sts := dbcustom.NewAllOfReflect(pojo).GetHasField()
	if len(sts) > 0 {
		d.SetCols(sts...)
	}

	return d
}


func (d *DdmParams) SetCols(selectCols ...string) *DdmParams {
	d.Cols(selectCols...)
	return d
}


func (d *DdmParams) SetId(id interface{}) *DdmParams {
	d.Id(id)
	return d
}

func (d *DdmParams) SetValue(pojo interface{}) *DdmParams {
	d.Value(pojo)
	return d
}


func (d *DdmParams) Create() error {
	return d.SqlCud.Create(NewDB())
}

func (d *DdmParams) Update() error {
	return d.SqlCud.Update(NewDB())
}

func (d *DdmParams) Delete() error {
	return d.SqlCud.Delete(NewDB())
}
