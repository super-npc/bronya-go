package tool

import (
	"maps"
)

// Cell 对应 Guava 的 Table.Cell
type Cell[R, C comparable, V any] struct {
	Row R
	Col C
	Val V
}

// Table 等价于 Guava 的 Table<R,C,V>
type Table[R, C comparable, V any] struct {
	// 内部用两层 map 实现
	data map[R]map[C]V
}

// NewTable 工厂函数
func NewTable[R, C comparable, V any]() *Table[R, C, V] {
	return &Table[R, C, V]{data: make(map[R]map[C]V)}
}

// Put 写入/更新
func (t *Table[R, C, V]) Put(r R, c C, v V) {
	if t.data[r] == nil {
		t.data[r] = make(map[C]V)
	}
	t.data[r][c] = v
}

// Get 读取
func (t *Table[R, C, V]) Get(r R, c C) (v V, ok bool) {
	row, ok := t.data[r]
	if !ok {
		return v, false
	}
	v, ok = row[c]
	return
}

// Remove 删除
func (t *Table[R, C, V]) Remove(r R, c C) {
	if row, ok := t.data[r]; ok {
		delete(row, c)
		if len(row) == 0 {
			delete(t.data, r)
		}
	}
}

// Contains 判断是否存在
func (t *Table[R, C, V]) Contains(r R, c C) bool {
	_, ok := t.Get(r, c)
	return ok
}

// Row 返回整行数据（拷贝）
func (t *Table[R, C, V]) Row(r R) map[C]V {
	return maps.Clone(t.data[r])
}

// Column 返回整列数据（拷贝）
func (t *Table[R, C, V]) Column(c C) map[R]V {
	col := make(map[R]V)
	for r, row := range t.data {
		if v, ok := row[c]; ok {
			col[r] = v
		}
	}
	return col
}

// Rows 所有行键
func (t *Table[R, C, V]) Rows() []R {
	keys := make([]R, 0, len(t.data))
	for k := range t.data {
		keys = append(keys, k)
	}
	return keys
}

// Columns 所有列键（去重）
func (t *Table[R, C, V]) Columns() []C {
	set := make(map[C]struct{})
	for _, row := range t.data {
		for c := range row {
			set[c] = struct{}{}
		}
	}
	keys := make([]C, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	return keys
}

// CellSet 遍历所有单元格（等价 Guava cellSet）
func (t *Table[R, C, V]) CellSet() []Cell[R, C, V] {
	var cells []Cell[R, C, V]
	for r, row := range t.data {
		for c, v := range row {
			cells = append(cells, Cell[R, C, V]{r, c, v})
		}
	}
	return cells
}
