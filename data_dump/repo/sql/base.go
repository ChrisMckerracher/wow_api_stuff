package sql

import "reflect"

type RepoSqlBase struct {
}

func (r *RepoSqlBase) ExtractBase(data interface{}) error {
	dataReflection := reflect.ValueOf(data)
	tableMap := re

	for i := 0; i < dataReflection.NumField(); i++ {

	}
}
