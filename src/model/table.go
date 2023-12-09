package model

import "strings"

// 支持的数据类型
var TypeList = []string{"text", "integer"}

type Table struct {
	TableName string            `json:"tableName" form:"tableName"`
	Columns   map[string]string `json:"columns,omitempy" form:"columns,omitempy"` //Map为列名:类型，golang没有元组，这种写法可能还相对好一点
}

// 从 map[string]string -> 形如
//
//		  id bigserial not null primary key,
//	   uid uuid default gen_random_uuid() not null,
//
// 来插入到SQL语句中
func (table *Table) ExportDefs() string {
	ans := ""
	for columnName, columnType := range table.Columns {
		ans += columnName
		ans += " "
		ans += columnType
		ans += ","
	}
	ans = strings.TrimRight(ans, ",")
	return ans
}

// func (table *Table) ExportNames() string {
// 	ans := ""
// 	for columnName, _ := range table.Columns {
// 		ans += columnName
// 		ans += ","
// 	}
// 	ans = strings.TrimRight(ans, ",")
// 	return ans
// }
