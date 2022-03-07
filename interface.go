package gorose

import "constraints"

type IQuery[T constraints.Ordered] interface {
	Query(sqlString string, args ...T)
}

type IExecute[T constraints.Ordered] interface {
	Execute(sqlString string, args ...T)
}

type IOrm[T any] interface {
	OrderBy(field string, sort string) IOrm
	OrderByRaw(orderString string) IOrm
	Select(object T) IOrm
}

type IConvert[T IOrm] interface {
	ToJSON() string
}
