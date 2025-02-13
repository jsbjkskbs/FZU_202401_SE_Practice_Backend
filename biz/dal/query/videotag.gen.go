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

func newVideoTag(db *gorm.DB, opts ...gen.DOOption) videoTag {
	_videoTag := videoTag{}

	_videoTag.videoTagDo.UseDB(db, opts...)
	_videoTag.videoTagDo.UseModel(&model.VideoTag{})

	tableName := _videoTag.videoTagDo.TableName()
	_videoTag.ALL = field.NewAsterisk(tableName)
	_videoTag.VideoID = field.NewInt64(tableName, "video_id")
	_videoTag.TagID = field.NewInt64(tableName, "tag_id")
	_videoTag.CreatedAt = field.NewInt64(tableName, "created_at")
	_videoTag.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_videoTag.fillFieldMap()

	return _videoTag
}

// videoTag 视频标签关联表
type videoTag struct {
	videoTagDo videoTagDo

	ALL       field.Asterisk
	VideoID   field.Int64 // 视频ID
	TagID     field.Int64 // 标签ID
	CreatedAt field.Int64 // 创建时间
	DeletedAt field.Int64 // 删除时间

	fieldMap map[string]field.Expr
}

func (v videoTag) Table(newTableName string) *videoTag {
	v.videoTagDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v videoTag) As(alias string) *videoTag {
	v.videoTagDo.DO = *(v.videoTagDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *videoTag) updateTableName(table string) *videoTag {
	v.ALL = field.NewAsterisk(table)
	v.VideoID = field.NewInt64(table, "video_id")
	v.TagID = field.NewInt64(table, "tag_id")
	v.CreatedAt = field.NewInt64(table, "created_at")
	v.DeletedAt = field.NewInt64(table, "deleted_at")

	v.fillFieldMap()

	return v
}

func (v *videoTag) WithContext(ctx context.Context) *videoTagDo { return v.videoTagDo.WithContext(ctx) }

func (v videoTag) TableName() string { return v.videoTagDo.TableName() }

func (v videoTag) Alias() string { return v.videoTagDo.Alias() }

func (v videoTag) Columns(cols ...field.Expr) gen.Columns { return v.videoTagDo.Columns(cols...) }

func (v *videoTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *videoTag) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 4)
	v.fieldMap["video_id"] = v.VideoID
	v.fieldMap["tag_id"] = v.TagID
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["deleted_at"] = v.DeletedAt
}

func (v videoTag) clone(db *gorm.DB) videoTag {
	v.videoTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v videoTag) replaceDB(db *gorm.DB) videoTag {
	v.videoTagDo.ReplaceDB(db)
	return v
}

type videoTagDo struct{ gen.DO }

func (v videoTagDo) Debug() *videoTagDo {
	return v.withDO(v.DO.Debug())
}

func (v videoTagDo) WithContext(ctx context.Context) *videoTagDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v videoTagDo) ReadDB() *videoTagDo {
	return v.Clauses(dbresolver.Read)
}

func (v videoTagDo) WriteDB() *videoTagDo {
	return v.Clauses(dbresolver.Write)
}

func (v videoTagDo) Session(config *gorm.Session) *videoTagDo {
	return v.withDO(v.DO.Session(config))
}

func (v videoTagDo) Clauses(conds ...clause.Expression) *videoTagDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v videoTagDo) Returning(value interface{}, columns ...string) *videoTagDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v videoTagDo) Not(conds ...gen.Condition) *videoTagDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v videoTagDo) Or(conds ...gen.Condition) *videoTagDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v videoTagDo) Select(conds ...field.Expr) *videoTagDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v videoTagDo) Where(conds ...gen.Condition) *videoTagDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v videoTagDo) Order(conds ...field.Expr) *videoTagDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v videoTagDo) Distinct(cols ...field.Expr) *videoTagDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v videoTagDo) Omit(cols ...field.Expr) *videoTagDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v videoTagDo) Join(table schema.Tabler, on ...field.Expr) *videoTagDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v videoTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) *videoTagDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v videoTagDo) RightJoin(table schema.Tabler, on ...field.Expr) *videoTagDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v videoTagDo) Group(cols ...field.Expr) *videoTagDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v videoTagDo) Having(conds ...gen.Condition) *videoTagDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v videoTagDo) Limit(limit int) *videoTagDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v videoTagDo) Offset(offset int) *videoTagDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v videoTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *videoTagDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v videoTagDo) Unscoped() *videoTagDo {
	return v.withDO(v.DO.Unscoped())
}

func (v videoTagDo) Create(values ...*model.VideoTag) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v videoTagDo) CreateInBatches(values []*model.VideoTag, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v videoTagDo) Save(values ...*model.VideoTag) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v videoTagDo) First() (*model.VideoTag, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoTag), nil
	}
}

func (v videoTagDo) Take() (*model.VideoTag, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoTag), nil
	}
}

func (v videoTagDo) Last() (*model.VideoTag, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoTag), nil
	}
}

func (v videoTagDo) Find() ([]*model.VideoTag, error) {
	result, err := v.DO.Find()
	return result.([]*model.VideoTag), err
}

func (v videoTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VideoTag, err error) {
	buf := make([]*model.VideoTag, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v videoTagDo) FindInBatches(result *[]*model.VideoTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v videoTagDo) Attrs(attrs ...field.AssignExpr) *videoTagDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v videoTagDo) Assign(attrs ...field.AssignExpr) *videoTagDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v videoTagDo) Joins(fields ...field.RelationField) *videoTagDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v videoTagDo) Preload(fields ...field.RelationField) *videoTagDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v videoTagDo) FirstOrInit() (*model.VideoTag, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoTag), nil
	}
}

func (v videoTagDo) FirstOrCreate() (*model.VideoTag, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoTag), nil
	}
}

func (v videoTagDo) FindByPage(offset int, limit int) (result []*model.VideoTag, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v videoTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v videoTagDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v videoTagDo) Delete(models ...*model.VideoTag) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *videoTagDo) withDO(do gen.Dao) *videoTagDo {
	v.DO = *do.(*gen.DO)
	return v
}
