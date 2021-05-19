package expr

type ExprItem struct {
	Id   int64  `json:"id"`   // 主键
	Code string `json:"code"` // 编码
	Name string `json:"name"` // 展示代码

	Value int `json:"value"` // 对应的值

	ProjectId int64 `json:"projectId"` // 区别项目
	ExpressId int64 `json:"expressId"` // 归属表达式 id = expressId 就是 表达式根

	ParentId int64 `json:"parentId"` // 递归结构树

	Kind int `json:"kind"` // -1-变量取反 0-变量 1-"|"或运算 2-"&"与运算

	Childrens []*ExprItem `json:"childrens"`
}

func (item *ExprItem) Clone() *ExprItem {
	return &ExprItem{
		Id:        item.Id,
		Code:      item.Code,
		Name:      item.Name,
		Value:     item.Value,
		ProjectId: item.ProjectId,
		ExpressId: item.ExpressId,
		ParentId:  item.ParentId,
		Kind:      item.Kind,
	}
}

func Flattening(item *ExprItem, result *[]*ExprItem, expId, parentId int64, idFunc func() int64) {
	item.Id = idFunc()
	if expId < 1 {
		expId = item.Id
	}
	item.ExpressId = expId
	item.ParentId = parentId
	*result = append(*result, item.Clone())
	for _, c := range item.Childrens {
		Flattening(c, result, expId, item.Id, idFunc)
	}
}

func Tree(result []*ExprItem) *ExprItem {
	m := make(map[int64]*ExprItem, 0)
	for _, v := range result {
		m[v.Id] = v
	}
	for _, v := range result {
		if val, ok := m[v.ParentId]; ok {
			if val.Childrens == nil {
				val.Childrens = make([]*ExprItem, 0)
			}
			val.Childrens = append(val.Childrens, v)
		}
	}
	return result[0]
}
