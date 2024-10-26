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

	"sfw/biz/dal/model"
)

func newActivity(db *gorm.DB, opts ...gen.DOOption) activity {
	_activity := activity{}

	_activity.activityDo.UseDB(db, opts...)
	_activity.activityDo.UseModel(&model.Activity{})

	tableName := _activity.activityDo.TableName()
	_activity.ALL = field.NewAsterisk(tableName)
	_activity.ID = field.NewInt64(tableName, "id")
	_activity.UserID = field.NewInt64(tableName, "user_id")
	_activity.Content = field.NewString(tableName, "content")
	_activity.MediaURL = field.NewString(tableName, "media_url")
	_activity.VisitCount = field.NewInt64(tableName, "visit_count")
	_activity.CreatedAt = field.NewInt64(tableName, "created_at")
	_activity.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_activity.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_activity.fillFieldMap()

	return _activity
}

// activity 动态表
type activity struct {
	activityDo activityDo

	ALL        field.Asterisk
	ID         field.Int64  // 动态ID
	UserID     field.Int64  // 用户ID
	Content    field.String // 动态内容
	MediaURL   field.String // 媒体URL
	VisitCount field.Int64  // 浏览量
	CreatedAt  field.Int64  // 创建时间
	UpdatedAt  field.Int64  // 修改时间
	DeletedAt  field.Int64  // 删除时间

	fieldMap map[string]field.Expr
}

func (a activity) Table(newTableName string) *activity {
	a.activityDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a activity) As(alias string) *activity {
	a.activityDo.DO = *(a.activityDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *activity) updateTableName(table string) *activity {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.UserID = field.NewInt64(table, "user_id")
	a.Content = field.NewString(table, "content")
	a.MediaURL = field.NewString(table, "media_url")
	a.VisitCount = field.NewInt64(table, "visit_count")
	a.CreatedAt = field.NewInt64(table, "created_at")
	a.UpdatedAt = field.NewInt64(table, "updated_at")
	a.DeletedAt = field.NewInt64(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *activity) WithContext(ctx context.Context) *activityDo { return a.activityDo.WithContext(ctx) }

func (a activity) TableName() string { return a.activityDo.TableName() }

func (a activity) Alias() string { return a.activityDo.Alias() }

func (a activity) Columns(cols ...field.Expr) gen.Columns { return a.activityDo.Columns(cols...) }

func (a *activity) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *activity) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["content"] = a.Content
	a.fieldMap["media_url"] = a.MediaURL
	a.fieldMap["visit_count"] = a.VisitCount
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
}

func (a activity) clone(db *gorm.DB) activity {
	a.activityDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a activity) replaceDB(db *gorm.DB) activity {
	a.activityDo.ReplaceDB(db)
	return a
}

type activityDo struct{ gen.DO }

func (a activityDo) Debug() *activityDo {
	return a.withDO(a.DO.Debug())
}

func (a activityDo) WithContext(ctx context.Context) *activityDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a activityDo) ReadDB() *activityDo {
	return a.Clauses(dbresolver.Read)
}

func (a activityDo) WriteDB() *activityDo {
	return a.Clauses(dbresolver.Write)
}

func (a activityDo) Session(config *gorm.Session) *activityDo {
	return a.withDO(a.DO.Session(config))
}

func (a activityDo) Clauses(conds ...clause.Expression) *activityDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a activityDo) Returning(value interface{}, columns ...string) *activityDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a activityDo) Not(conds ...gen.Condition) *activityDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a activityDo) Or(conds ...gen.Condition) *activityDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a activityDo) Select(conds ...field.Expr) *activityDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a activityDo) Where(conds ...gen.Condition) *activityDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a activityDo) Order(conds ...field.Expr) *activityDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a activityDo) Distinct(cols ...field.Expr) *activityDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a activityDo) Omit(cols ...field.Expr) *activityDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a activityDo) Join(table schema.Tabler, on ...field.Expr) *activityDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a activityDo) LeftJoin(table schema.Tabler, on ...field.Expr) *activityDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a activityDo) RightJoin(table schema.Tabler, on ...field.Expr) *activityDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a activityDo) Group(cols ...field.Expr) *activityDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a activityDo) Having(conds ...gen.Condition) *activityDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a activityDo) Limit(limit int) *activityDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a activityDo) Offset(offset int) *activityDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a activityDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *activityDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a activityDo) Unscoped() *activityDo {
	return a.withDO(a.DO.Unscoped())
}

func (a activityDo) Create(values ...*model.Activity) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a activityDo) CreateInBatches(values []*model.Activity, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a activityDo) Save(values ...*model.Activity) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a activityDo) First() (*model.Activity, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) Take() (*model.Activity, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) Last() (*model.Activity, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) Find() ([]*model.Activity, error) {
	result, err := a.DO.Find()
	return result.([]*model.Activity), err
}

func (a activityDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Activity, err error) {
	buf := make([]*model.Activity, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a activityDo) FindInBatches(result *[]*model.Activity, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a activityDo) Attrs(attrs ...field.AssignExpr) *activityDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a activityDo) Assign(attrs ...field.AssignExpr) *activityDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a activityDo) Joins(fields ...field.RelationField) *activityDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a activityDo) Preload(fields ...field.RelationField) *activityDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a activityDo) FirstOrInit() (*model.Activity, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) FirstOrCreate() (*model.Activity, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Activity), nil
	}
}

func (a activityDo) FindByPage(offset int, limit int) (result []*model.Activity, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a activityDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a activityDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a activityDo) Delete(models ...*model.Activity) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *activityDo) withDO(do gen.Dao) *activityDo {
	a.DO = *do.(*gen.DO)
	return a
}
