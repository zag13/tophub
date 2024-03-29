// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/zag13/tophub/cli/dal/model"
)

func newTop(db *gorm.DB, opts ...gen.DOOption) top {
	_top := top{}

	_top.topDo.UseDB(db, opts...)
	_top.topDo.UseModel(&model.Top{})

	tableName := _top.topDo.TableName()
	_top.ALL = field.NewAsterisk(tableName)
	_top.ID = field.NewInt64(tableName, "id")
	_top.SpiderTime = field.NewTime(tableName, "spider_time")
	_top.Site = field.NewString(tableName, "site")
	_top.Rank = field.NewInt32(tableName, "rank")
	_top.Title = field.NewString(tableName, "title")
	_top.URL = field.NewString(tableName, "url")
	_top.Extra = field.NewString(tableName, "extra")

	_top.fillFieldMap()

	return _top
}

type top struct {
	topDo topDo

	ALL        field.Asterisk
	ID         field.Int64
	SpiderTime field.Time
	Site       field.String
	Rank       field.Int32
	Title      field.String
	URL        field.String
	Extra      field.String

	fieldMap map[string]field.Expr
}

func (t top) Table(newTableName string) *top {
	t.topDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t top) As(alias string) *top {
	t.topDo.DO = *(t.topDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *top) updateTableName(table string) *top {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.SpiderTime = field.NewTime(table, "spider_time")
	t.Site = field.NewString(table, "site")
	t.Rank = field.NewInt32(table, "rank")
	t.Title = field.NewString(table, "title")
	t.URL = field.NewString(table, "url")
	t.Extra = field.NewString(table, "extra")

	t.fillFieldMap()

	return t
}

func (t *top) WithContext(ctx context.Context) *topDo { return t.topDo.WithContext(ctx) }

func (t top) TableName() string { return t.topDo.TableName() }

func (t top) Alias() string { return t.topDo.Alias() }

func (t top) Columns(cols ...field.Expr) gen.Columns { return t.topDo.Columns(cols...) }

func (t *top) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *top) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["spider_time"] = t.SpiderTime
	t.fieldMap["site"] = t.Site
	t.fieldMap["rank"] = t.Rank
	t.fieldMap["title"] = t.Title
	t.fieldMap["url"] = t.URL
	t.fieldMap["extra"] = t.Extra
}

func (t top) clone(db *gorm.DB) top {
	t.topDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t top) replaceDB(db *gorm.DB) top {
	t.topDo.ReplaceDB(db)
	return t
}

type topDo struct{ gen.DO }

func (t topDo) Debug() *topDo {
	return t.withDO(t.DO.Debug())
}

func (t topDo) WithContext(ctx context.Context) *topDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t topDo) ReadDB() *topDo {
	return t.Clauses(dbresolver.Read)
}

func (t topDo) WriteDB() *topDo {
	return t.Clauses(dbresolver.Write)
}

func (t topDo) Session(config *gorm.Session) *topDo {
	return t.withDO(t.DO.Session(config))
}

func (t topDo) Clauses(conds ...clause.Expression) *topDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t topDo) Returning(value interface{}, columns ...string) *topDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t topDo) Not(conds ...gen.Condition) *topDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t topDo) Or(conds ...gen.Condition) *topDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t topDo) Select(conds ...field.Expr) *topDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t topDo) Where(conds ...gen.Condition) *topDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t topDo) Order(conds ...field.Expr) *topDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t topDo) Distinct(cols ...field.Expr) *topDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t topDo) Omit(cols ...field.Expr) *topDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t topDo) Join(table schema.Tabler, on ...field.Expr) *topDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t topDo) LeftJoin(table schema.Tabler, on ...field.Expr) *topDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t topDo) RightJoin(table schema.Tabler, on ...field.Expr) *topDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t topDo) Group(cols ...field.Expr) *topDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t topDo) Having(conds ...gen.Condition) *topDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t topDo) Limit(limit int) *topDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t topDo) Offset(offset int) *topDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t topDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *topDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t topDo) Unscoped() *topDo {
	return t.withDO(t.DO.Unscoped())
}

func (t topDo) Create(values ...*model.Top) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t topDo) CreateInBatches(values []*model.Top, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t topDo) Save(values ...*model.Top) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t topDo) First() (*model.Top, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Top), nil
	}
}

func (t topDo) Take() (*model.Top, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Top), nil
	}
}

func (t topDo) Last() (*model.Top, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Top), nil
	}
}

func (t topDo) Find() ([]*model.Top, error) {
	result, err := t.DO.Find()
	return result.([]*model.Top), err
}

func (t topDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Top, err error) {
	buf := make([]*model.Top, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t topDo) FindInBatches(result *[]*model.Top, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t topDo) Attrs(attrs ...field.AssignExpr) *topDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t topDo) Assign(attrs ...field.AssignExpr) *topDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t topDo) Joins(fields ...field.RelationField) *topDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t topDo) Preload(fields ...field.RelationField) *topDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t topDo) FirstOrInit() (*model.Top, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Top), nil
	}
}

func (t topDo) FirstOrCreate() (*model.Top, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Top), nil
	}
}

func (t topDo) FindByPage(offset int, limit int) (result []*model.Top, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t topDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t topDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t topDo) Delete(models ...*model.Top) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *topDo) withDO(do gen.Dao) *topDo {
	t.DO = *do.(*gen.DO)
	return t
}
