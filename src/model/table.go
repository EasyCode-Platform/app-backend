package model

var TypeList = []string{"text", "integer"}

type Table struct {
	Columns map[string]string `json:"columns"` //Map为列名:类型，golang没有元组，这种写法可能还相对好一点
}
