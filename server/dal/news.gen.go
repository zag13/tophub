// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/zag13/tophub/server/dal/model"
)

func newNews(db *gorm.DB, opts ...gen.DOOption) news {
	_news := news{}

	_news.newsDo.UseDB(db, opts...)
	_news.newsDo.UseModel(&model.News{})

	tableName := _news.newsDo.TableName()
	_news.ALL = field.NewAsterisk(tableName)
	_news.ID = field.NewInt64(tableName, "id")
	_news.SpiderTime = field.NewTime(tableName, "spider_time")
	_news.Site = field.NewString(tableName, "site")
	_news.Ranking = field.NewInt32(tableName, "ranking")
	_news.Title = field.NewString(tableName, "title")
	_news.URL = field.NewString(tableName, "url")

	_news.fillFieldMap()

	return _news
}

type news struct {
	newsDo newsDo

	ALL        field.Asterisk
	ID         field.Int64
	SpiderTime field.Time
	Site       field.String
	Ranking    field.Int32
	Title      field.String
	URL        field.String

	fieldMap map[string]field.Expr
}

func (n news) Table(newTableName string) *news {
	n.newsDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n news) As(alias string) *news {
	n.newsDo.DO = *(n.newsDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *news) updateTableName(table string) *news {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewInt64(table, "id")
	n.SpiderTime = field.NewTime(table, "spider_time")
	n.Site = field.NewString(table, "site")
	n.Ranking = field.NewInt32(table, "ranking")
	n.Title = field.NewString(table, "title")
	n.URL = field.NewString(table, "url")

	n.fillFieldMap()

	return n
}

func (n *news) WithContext(ctx context.Context) *newsDo { return n.newsDo.WithContext(ctx) }

func (n news) TableName() string { return n.newsDo.TableName() }

func (n news) Alias() string { return n.newsDo.Alias() }

func (n news) Columns(cols ...field.Expr) gen.Columns { return n.newsDo.Columns(cols...) }

func (n *news) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *news) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 6)
	n.fieldMap["id"] = n.ID
	n.fieldMap["spider_time"] = n.SpiderTime
	n.fieldMap["site"] = n.Site
	n.fieldMap["ranking"] = n.Ranking
	n.fieldMap["title"] = n.Title
	n.fieldMap["url"] = n.URL
}

func (n news) clone(db *gorm.DB) news {
	n.newsDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n news) replaceDB(db *gorm.DB) news {
	n.newsDo.ReplaceDB(db)
	return n
}

type newsDo struct{ gen.DO }

// SELECT * FROM @@table WHERE id = @id
func (n newsDo) FindOne(id int64) (result *model.News, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM news WHERE id = ? ")

	var executeSQL *gorm.DB
	executeSQL = n.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (n newsDo) Debug() *newsDo {
	return n.withDO(n.DO.Debug())
}

func (n newsDo) WithContext(ctx context.Context) *newsDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n newsDo) ReadDB() *newsDo {
	return n.Clauses(dbresolver.Read)
}

func (n newsDo) WriteDB() *newsDo {
	return n.Clauses(dbresolver.Write)
}

func (n newsDo) Session(config *gorm.Session) *newsDo {
	return n.withDO(n.DO.Session(config))
}

func (n newsDo) Clauses(conds ...clause.Expression) *newsDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n newsDo) Returning(value interface{}, columns ...string) *newsDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n newsDo) Not(conds ...gen.Condition) *newsDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n newsDo) Or(conds ...gen.Condition) *newsDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n newsDo) Select(conds ...field.Expr) *newsDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n newsDo) Where(conds ...gen.Condition) *newsDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n newsDo) Order(conds ...field.Expr) *newsDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n newsDo) Distinct(cols ...field.Expr) *newsDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n newsDo) Omit(cols ...field.Expr) *newsDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n newsDo) Join(table schema.Tabler, on ...field.Expr) *newsDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n newsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *newsDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n newsDo) RightJoin(table schema.Tabler, on ...field.Expr) *newsDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n newsDo) Group(cols ...field.Expr) *newsDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n newsDo) Having(conds ...gen.Condition) *newsDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n newsDo) Limit(limit int) *newsDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n newsDo) Offset(offset int) *newsDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n newsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *newsDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n newsDo) Unscoped() *newsDo {
	return n.withDO(n.DO.Unscoped())
}

func (n newsDo) Create(values ...*model.News) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n newsDo) CreateInBatches(values []*model.News, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n newsDo) Save(values ...*model.News) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n newsDo) First() (*model.News, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.News), nil
	}
}

func (n newsDo) Take() (*model.News, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.News), nil
	}
}

func (n newsDo) Last() (*model.News, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.News), nil
	}
}

func (n newsDo) Find() ([]*model.News, error) {
	result, err := n.DO.Find()
	return result.([]*model.News), err
}

func (n newsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.News, err error) {
	buf := make([]*model.News, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n newsDo) FindInBatches(result *[]*model.News, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n newsDo) Attrs(attrs ...field.AssignExpr) *newsDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n newsDo) Assign(attrs ...field.AssignExpr) *newsDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n newsDo) Joins(fields ...field.RelationField) *newsDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n newsDo) Preload(fields ...field.RelationField) *newsDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n newsDo) FirstOrInit() (*model.News, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.News), nil
	}
}

func (n newsDo) FirstOrCreate() (*model.News, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.News), nil
	}
}

func (n newsDo) FindByPage(offset int, limit int) (result []*model.News, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n newsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n newsDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n newsDo) Delete(models ...*model.News) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *newsDo) withDO(do gen.Dao) *newsDo {
	n.DO = *do.(*gen.DO)
	return n
}
