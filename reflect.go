package dbcustom

import (
	"reflect"
)

type OfReflect struct {
	Vof []reflect.Value
	Tof []reflect.StructField
}
type OfReflectName 		map[string]int
type OfReflectValue 	map[string]interface{}
type OfReflectTag 		map[string]string
type OfReflectTags 		map[string]map[string]string

type Kind = reflect.Kind
const (
	STRING		= reflect.String
	SLICE		= reflect.Slice
	STRUCT		= reflect.Struct
	MAP			= reflect.Map
	PTR			= reflect.Ptr
)

const (
	JSON 	=	"json"
	FORM 	=	"form"
	QUERY 	=	"query"
	JOIN 	=	"join"
)
var Tags = []string{JSON, FORM, QUERY, JOIN}

// 反射字段构造器，不反射字段里struct，struct变量
//反射获取所有字段(包含镶嵌) []reflect.StructField,[]reflect.Value
// struct中 &struct,[]struct 无法反射字段
func NewOfReflect(i interface{}) *OfReflect {
	structFields, values := toReflect(i,false)
	return &OfReflect{
		Vof: values,
		Tof: structFields,
	}
}

// 反射字段全参构造器，反射字段里struct，struct变量
func NewAllOfReflect(i interface{}) *OfReflect {
	structFields, values := toReflect(i,true)
	return &OfReflect{
		Vof: values,
		Tof: structFields,
	}
}

// 反射获取接口类型
func GetKind(i interface{}) Kind {
	return reflect.TypeOf(i).Kind()
}

// 反射获取信息
func getReflect(i interface{}) (reflect.Type, reflect.Value) {
	tof := reflect.TypeOf(i)
	var vof reflect.Value
	switch tof.Kind() {
	case STRUCT: vof = reflect.ValueOf(i)
	case PTR:
		if tof.Elem().Kind() == STRUCT {
			vof = reflect.ValueOf(i).Elem()
		}

	}
	return tof,vof
}

// 反射字段,接口类型是指针或则实体
func toReflect(i interface{},isAll bool) (rts []reflect.StructField,rvs []reflect.Value) {
	_, vof := getReflect(i)
	for k := 0; k<vof.NumField(); k++ {
		if vof.Field(k).Kind() == STRUCT && isAll {
			rt,rv := nReflect(vof.Field(k).Interface())
			if rt == nil && rv == nil{
				rvs = append(rvs,vof.Field(k))
				rts = append(rts,vof.Type().Field(k))
			}else {
				rvs = append(rvs,rv...)  //数组添加数组
				rts = append(rts,rt...)  //数组添加数组
			}

		} else{
			rvs = append(rvs,vof.Field(k))
			rts = append(rts,vof.Type().Field(k))
		}
	}
	return
}

// 获取非第一层多次反射struct，有json tag的字段
func nReflect(i interface{}) (rts []reflect.StructField,rvs []reflect.Value) {
	tof, vof := getReflect(i)
	count := 0
	for k := 0; k<vof.NumField(); k++ {
		if vof.Field(k).Kind() == STRUCT {
			if rt, rv := nReflect(vof.Field(k).Interface()); rt != nil && rv != nil{
				rts = append(rts, rt...)
				rvs = append(rvs, rv...)
			}
		}
		if _, ok := tof.Field(k).Tag.Lookup("json"); ok{
			rvs = append(rvs,vof.Field(k))
			rts = append(rts,vof.Type().Field(k))
		}else {
			count++
		}
	}
	if count == vof.NumField(){
		return nil,nil
	}
	return
}

// 取出字典，key is column，val is index
func (o *OfReflect) GetAllFieldName () OfReflectName {
	mp := make(OfReflectName)
	for v,k := range o.Tof{
		mp[k.Name] = v
	}
	return mp
}

// 取出字段值，key is column，val is column-value
// default value is all val
func (o *OfReflect) GetFieldValue (name ...string) OfReflectValue {
	mp := make(OfReflectValue)
	fieldName := o.GetAllFieldName()
	if len(name) == 0{
		for _,k := range name{
			if valInt,ok := fieldName[k];ok{
				mp[k] = o.Vof[valInt]
			}
		}
	}else{
		for v,k := range fieldName{
			mp[v] = o.Vof[k]
		}
	}
	return mp
}

// 取出字段tag，key is column，val is tag-value
func (o *OfReflect) GetKeyTags (tag string) OfReflectTag {
	mp := make(OfReflectTag)
	name := o.GetAllFieldName()
	for v,k := range name{
		if val,ok := o.Tof[k].Tag.Lookup(tag);ok{
			mp[v] = val
		}
	}
	return mp
}

// 取出字段tag，key is column，val is tag-value[]
func (o *OfReflect) GetAllKeyTags () OfReflectTags {
	mps := make(OfReflectTags)
	name := o.GetAllFieldName()
	for v,k := range name{
		mp := make(OfReflectTag)
		for _,tag := range Tags {
			if val,ok := o.Tof[k].Tag.Lookup(tag); ok {
				mp[tag] = val
			}
		}
		mps[v] = mp
	}
	return mps
}

func (o *OfReflect) GetHasField() []string {

	sts := make([]string,0)
	for i, v := range o.Vof {
		if ! IsNULL(v.Interface()) {
			sts = append(sts,o.Tof[i].Name)
		}

	}
	return sts
}