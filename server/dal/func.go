package dal

import "text/template"

// PreprocessForQuery 对查询参数进行预处理
func PreprocessForQuery(qry map[string]any) map[string]any {
	for k, v := range qry {
		switch v.(type) {
		case string:
			qry[k] = template.HTMLEscapeString(v.(string))
		case []string:
			for i, s := range v.([]string) {
				v.([]string)[i] = template.HTMLEscapeString(s)
			}
		}
	}
	return qry
}

// GetOffsetAndLimit 基于 page 和 pageSize 获取 offset 和 limit
func GetOffsetAndLimit(page, pageSize int) (offset, limit int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	return (page - 1) * pageSize, pageSize * 3
}
