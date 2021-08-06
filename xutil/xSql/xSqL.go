package xSql

import "strings"

type XSQL struct {
	Where  *strings.Builder
	Updata *strings.Builder
	Value  []interface{}
	Len    int
}

func New() *XSQL {
	var a strings.Builder
	var b strings.Builder
	var v = make([]interface{}, 100)
	return &XSQL{&a, &b, v, 0}
}

//追加更新字段
//tmp.APPEND_UPDATA("aid","5").APPEND_UPDATA("bid","")
func (x *XSQL) APPEND_UPDATA(key, value string) *XSQL {
	if len(value) == 0 {
		return x
	}
	if x.Len != 0 {
		x.Updata.WriteString(",")
	}
	x.Updata.WriteString(key)
	x.Updata.WriteString("=?")
	x.Value[x.Len] = value
	x.Len++
	return x
}

//追加查询条件 准确条件查询
//tmp.APPEND_Where("aid","5").APPEND_Where("bid","")
func (x *XSQL) APPEND_Where(key, value string) *XSQL {
	if len(value) == 0 {
		return x
	}
	x.Where.WriteString(" and ")
	x.Where.WriteString(key)
	x.Where.WriteString("=?")
	x.Value[x.Len] = value
	x.Len++
	return x
}

//追加查询提交 比较条件
//tmp.APPEND_Where_1("aid<","5").APPEND_Where_1("bid>","")
func (x *XSQL) APPEND_Where_1(key, value string) *XSQL {
	if len(value) == 0 {
		return x
	}

	x.Where.WriteString(" and ")
	x.Where.WriteString(key)
	x.Where.WriteString("?")
	x.Value[x.Len] = value
	x.Len++
	return x
}
//追加查询提交 比较条件
//tmp.APPEND_Where_2("msg LIKE '%?%'","5")
func (x *XSQL) APPEND_Where_2(key, value string) *XSQL {
	if len(value) == 0 {
		return x
	}

	x.Where.WriteString(" and ")
	x.Where.WriteString(key)
	x.Value[x.Len] = value
	x.Len++
	return x
}
func (x *XSQL) Get_WhereSQL() string {
	return x.Where.String()
}
func (x *XSQL) Get_UpdataSQL() string {
	return x.Updata.String()
}
func (x *XSQL) Get_Value() []interface{} {
	return x.Value[0:x.Len]
}
