package build_sql

import (
	"github.com/super-npc/bronya-go/src/commons/db"
)

func ExecSql(dbProvider db.DBProvider, sql string) []map[string]interface{} {
	rows, err := dbProvider.GetDb().Raw(sql).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// 4. 构造结果：[]map[string]interface{}
	cols, _ := rows.Columns()
	result := make([]map[string]interface{}, 0)

	for rows.Next() {
		// 为每一列建一个 interface{} 容器
		columns := make([]interface{}, len(cols))
		columnPtrs := make([]interface{}, len(cols))
		for i := range columns {
			columnPtrs[i] = &columns[i]
		}

		if err := rows.Scan(columnPtrs...); err != nil {
			panic(err)
		}

		rowMap := make(map[string]interface{})
		for i, colName := range cols {
			val := *(columnPtrs[i].(*interface{}))
			// 把 []byte 转成 string，更直观
			if b, ok := val.([]byte); ok {
				val = string(b)
			}
			rowMap[colName] = val
		}
		result = append(result, rowMap)
	}

	return result
}
