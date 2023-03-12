// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"tophub/app/topdata/dal/model"
)

func newTopdatum(db *gorm.DB, opts ...gen.DOOption) topdatum {
	_topdatum := topdatum{}

	_topdatum.topdatumDo.UseDB(db, opts...)
	_topdatum.topdatumDo.UseModel(&model.Topdatum{})

	tableName := _topdatum.topdatumDo.TableName()
	_topdatum.ALL = field.NewAsterisk(tableName)
	_topdatum.ID = field.NewInt64(tableName, "id")
	_topdatum.Tab = field.NewString(tableName, "tab")
	_topdatum.Position = field.NewInt64(tableName, "position")
	_topdatum.Title = field.NewString(tableName, "title")
	_topdatum.URL = field.NewString(tableName, "url")
	_topdatum.Extra = field.NewString(tableName, "extra")
	_topdatum.SpiderTime = field.NewTime(tableName, "spider_time")

	_topdatum.fillFieldMap()

	return _topdatum
}

type topdatum struct {
	topdatumDo topdatumDo

	ALL        field.Asterisk
	ID         field.Int64
	Tab        field.String // 标识
	Position   field.Int64  // 资讯位置
	Title      field.String // 资讯标题
	URL        field.String // 资讯地址
	Extra      field.String // 拓展字段
	SpiderTime field.Time   // 抓取时间

	fieldMap map[string]field.Expr
}

func (t topdatum) Table(newTableName string) *topdatum {
	t.topdatumDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t topdatum) As(alias string) *topdatum {
	t.topdatumDo.DO = *(t.topdatumDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *topdatum) updateTableName(table string) *topdatum {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.Tab = field.NewString(table, "tab")
	t.Position = field.NewInt64(table, "position")
	t.Title = field.NewString(table, "title")
	t.URL = field.NewString(table, "url")
	t.Extra = field.NewString(table, "extra")
	t.SpiderTime = field.NewTime(table, "spider_time")

	t.fillFieldMap()

	return t
}

func (t *topdatum) WithContext(ctx context.Context) ITopdatumDo { return t.topdatumDo.WithContext(ctx) }

func (t topdatum) TableName() string { return t.topdatumDo.TableName() }

func (t topdatum) Alias() string { return t.topdatumDo.Alias() }

func (t *topdatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *topdatum) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["tab"] = t.Tab
	t.fieldMap["position"] = t.Position
	t.fieldMap["title"] = t.Title
	t.fieldMap["url"] = t.URL
	t.fieldMap["extra"] = t.Extra
	t.fieldMap["spider_time"] = t.SpiderTime
}

func (t topdatum) clone(db *gorm.DB) topdatum {
	t.topdatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t topdatum) replaceDB(db *gorm.DB) topdatum {
	t.topdatumDo.ReplaceDB(db)
	return t
}

type topdatumDo struct{ gen.DO }

type ITopdatumDo interface {
	gen.SubQuery
	Debug() ITopdatumDo
	WithContext(ctx context.Context) ITopdatumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITopdatumDo
	WriteDB() ITopdatumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITopdatumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITopdatumDo
	Not(conds ...gen.Condition) ITopdatumDo
	Or(conds ...gen.Condition) ITopdatumDo
	Select(conds ...field.Expr) ITopdatumDo
	Where(conds ...gen.Condition) ITopdatumDo
	Order(conds ...field.Expr) ITopdatumDo
	Distinct(cols ...field.Expr) ITopdatumDo
	Omit(cols ...field.Expr) ITopdatumDo
	Join(table schema.Tabler, on ...field.Expr) ITopdatumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITopdatumDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITopdatumDo
	Group(cols ...field.Expr) ITopdatumDo
	Having(conds ...gen.Condition) ITopdatumDo
	Limit(limit int) ITopdatumDo
	Offset(offset int) ITopdatumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITopdatumDo
	Unscoped() ITopdatumDo
	Create(values ...*model.Topdatum) error
	CreateInBatches(values []*model.Topdatum, batchSize int) error
	Save(values ...*model.Topdatum) error
	First() (*model.Topdatum, error)
	Take() (*model.Topdatum, error)
	Last() (*model.Topdatum, error)
	Find() ([]*model.Topdatum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Topdatum, err error)
	FindInBatches(result *[]*model.Topdatum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Topdatum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITopdatumDo
	Assign(attrs ...field.AssignExpr) ITopdatumDo
	Joins(fields ...field.RelationField) ITopdatumDo
	Preload(fields ...field.RelationField) ITopdatumDo
	FirstOrInit() (*model.Topdatum, error)
	FirstOrCreate() (*model.Topdatum, error)
	FindByPage(offset int, limit int) (result []*model.Topdatum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITopdatumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t topdatumDo) Debug() ITopdatumDo {
	return t.withDO(t.DO.Debug())
}

func (t topdatumDo) WithContext(ctx context.Context) ITopdatumDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t topdatumDo) ReadDB() ITopdatumDo {
	return t.Clauses(dbresolver.Read)
}

func (t topdatumDo) WriteDB() ITopdatumDo {
	return t.Clauses(dbresolver.Write)
}

func (t topdatumDo) Session(config *gorm.Session) ITopdatumDo {
	return t.withDO(t.DO.Session(config))
}

func (t topdatumDo) Clauses(conds ...clause.Expression) ITopdatumDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t topdatumDo) Returning(value interface{}, columns ...string) ITopdatumDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t topdatumDo) Not(conds ...gen.Condition) ITopdatumDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t topdatumDo) Or(conds ...gen.Condition) ITopdatumDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t topdatumDo) Select(conds ...field.Expr) ITopdatumDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t topdatumDo) Where(conds ...gen.Condition) ITopdatumDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t topdatumDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITopdatumDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t topdatumDo) Order(conds ...field.Expr) ITopdatumDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t topdatumDo) Distinct(cols ...field.Expr) ITopdatumDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t topdatumDo) Omit(cols ...field.Expr) ITopdatumDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t topdatumDo) Join(table schema.Tabler, on ...field.Expr) ITopdatumDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t topdatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITopdatumDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t topdatumDo) RightJoin(table schema.Tabler, on ...field.Expr) ITopdatumDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t topdatumDo) Group(cols ...field.Expr) ITopdatumDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t topdatumDo) Having(conds ...gen.Condition) ITopdatumDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t topdatumDo) Limit(limit int) ITopdatumDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t topdatumDo) Offset(offset int) ITopdatumDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t topdatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITopdatumDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t topdatumDo) Unscoped() ITopdatumDo {
	return t.withDO(t.DO.Unscoped())
}

func (t topdatumDo) Create(values ...*model.Topdatum) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t topdatumDo) CreateInBatches(values []*model.Topdatum, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t topdatumDo) Save(values ...*model.Topdatum) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t topdatumDo) First() (*model.Topdatum, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Topdatum), nil
	}
}

func (t topdatumDo) Take() (*model.Topdatum, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Topdatum), nil
	}
}

func (t topdatumDo) Last() (*model.Topdatum, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Topdatum), nil
	}
}

func (t topdatumDo) Find() ([]*model.Topdatum, error) {
	result, err := t.DO.Find()
	return result.([]*model.Topdatum), err
}

func (t topdatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Topdatum, err error) {
	buf := make([]*model.Topdatum, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t topdatumDo) FindInBatches(result *[]*model.Topdatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t topdatumDo) Attrs(attrs ...field.AssignExpr) ITopdatumDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t topdatumDo) Assign(attrs ...field.AssignExpr) ITopdatumDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t topdatumDo) Joins(fields ...field.RelationField) ITopdatumDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t topdatumDo) Preload(fields ...field.RelationField) ITopdatumDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t topdatumDo) FirstOrInit() (*model.Topdatum, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Topdatum), nil
	}
}

func (t topdatumDo) FirstOrCreate() (*model.Topdatum, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Topdatum), nil
	}
}

func (t topdatumDo) FindByPage(offset int, limit int) (result []*model.Topdatum, count int64, err error) {
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

func (t topdatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t topdatumDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t topdatumDo) Delete(models ...*model.Topdatum) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *topdatumDo) withDO(do gen.Dao) *topdatumDo {
	t.DO = *do.(*gen.DO)
	return t
}