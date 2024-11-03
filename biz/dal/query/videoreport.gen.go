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

func newVideoReport(db *gorm.DB, opts ...gen.DOOption) videoReport {
	_videoReport := videoReport{}

	_videoReport.videoReportDo.UseDB(db, opts...)
	_videoReport.videoReportDo.UseModel(&model.VideoReport{})

	tableName := _videoReport.videoReportDo.TableName()
	_videoReport.ALL = field.NewAsterisk(tableName)
	_videoReport.ID = field.NewInt64(tableName, "id")
	_videoReport.UserID = field.NewInt64(tableName, "user_id")
	_videoReport.VideoID = field.NewInt64(tableName, "video_id")
	_videoReport.Reason = field.NewString(tableName, "reason")
	_videoReport.Label = field.NewString(tableName, "label")
	_videoReport.CreatedAt = field.NewInt64(tableName, "created_at")
	_videoReport.Status = field.NewString(tableName, "status")
	_videoReport.ResolvedAt = field.NewInt64(tableName, "resolved_at")
	_videoReport.AdminID = field.NewInt64(tableName, "admin_id")

	_videoReport.fillFieldMap()

	return _videoReport
}

// videoReport 视频举报表
type videoReport struct {
	videoReportDo videoReportDo

	ALL        field.Asterisk
	ID         field.Int64  // 举报ID
	UserID     field.Int64  // 用户ID
	VideoID    field.Int64  // 视频ID
	Reason     field.String // 举报原因
	Label      field.String // 举报标签
	CreatedAt  field.Int64  // 创建时间
	Status     field.String // 举报状态
	ResolvedAt field.Int64  // 解决时间
	AdminID    field.Int64  // 管理员ID

	fieldMap map[string]field.Expr
}

func (v videoReport) Table(newTableName string) *videoReport {
	v.videoReportDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v videoReport) As(alias string) *videoReport {
	v.videoReportDo.DO = *(v.videoReportDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *videoReport) updateTableName(table string) *videoReport {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewInt64(table, "id")
	v.UserID = field.NewInt64(table, "user_id")
	v.VideoID = field.NewInt64(table, "video_id")
	v.Reason = field.NewString(table, "reason")
	v.Label = field.NewString(table, "label")
	v.CreatedAt = field.NewInt64(table, "created_at")
	v.Status = field.NewString(table, "status")
	v.ResolvedAt = field.NewInt64(table, "resolved_at")
	v.AdminID = field.NewInt64(table, "admin_id")

	v.fillFieldMap()

	return v
}

func (v *videoReport) WithContext(ctx context.Context) *videoReportDo {
	return v.videoReportDo.WithContext(ctx)
}

func (v videoReport) TableName() string { return v.videoReportDo.TableName() }

func (v videoReport) Alias() string { return v.videoReportDo.Alias() }

func (v videoReport) Columns(cols ...field.Expr) gen.Columns { return v.videoReportDo.Columns(cols...) }

func (v *videoReport) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *videoReport) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 9)
	v.fieldMap["id"] = v.ID
	v.fieldMap["user_id"] = v.UserID
	v.fieldMap["video_id"] = v.VideoID
	v.fieldMap["reason"] = v.Reason
	v.fieldMap["label"] = v.Label
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["status"] = v.Status
	v.fieldMap["resolved_at"] = v.ResolvedAt
	v.fieldMap["admin_id"] = v.AdminID
}

func (v videoReport) clone(db *gorm.DB) videoReport {
	v.videoReportDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v videoReport) replaceDB(db *gorm.DB) videoReport {
	v.videoReportDo.ReplaceDB(db)
	return v
}

type videoReportDo struct{ gen.DO }

func (v videoReportDo) Debug() *videoReportDo {
	return v.withDO(v.DO.Debug())
}

func (v videoReportDo) WithContext(ctx context.Context) *videoReportDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v videoReportDo) ReadDB() *videoReportDo {
	return v.Clauses(dbresolver.Read)
}

func (v videoReportDo) WriteDB() *videoReportDo {
	return v.Clauses(dbresolver.Write)
}

func (v videoReportDo) Session(config *gorm.Session) *videoReportDo {
	return v.withDO(v.DO.Session(config))
}

func (v videoReportDo) Clauses(conds ...clause.Expression) *videoReportDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v videoReportDo) Returning(value interface{}, columns ...string) *videoReportDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v videoReportDo) Not(conds ...gen.Condition) *videoReportDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v videoReportDo) Or(conds ...gen.Condition) *videoReportDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v videoReportDo) Select(conds ...field.Expr) *videoReportDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v videoReportDo) Where(conds ...gen.Condition) *videoReportDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v videoReportDo) Order(conds ...field.Expr) *videoReportDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v videoReportDo) Distinct(cols ...field.Expr) *videoReportDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v videoReportDo) Omit(cols ...field.Expr) *videoReportDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v videoReportDo) Join(table schema.Tabler, on ...field.Expr) *videoReportDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v videoReportDo) LeftJoin(table schema.Tabler, on ...field.Expr) *videoReportDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v videoReportDo) RightJoin(table schema.Tabler, on ...field.Expr) *videoReportDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v videoReportDo) Group(cols ...field.Expr) *videoReportDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v videoReportDo) Having(conds ...gen.Condition) *videoReportDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v videoReportDo) Limit(limit int) *videoReportDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v videoReportDo) Offset(offset int) *videoReportDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v videoReportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *videoReportDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v videoReportDo) Unscoped() *videoReportDo {
	return v.withDO(v.DO.Unscoped())
}

func (v videoReportDo) Create(values ...*model.VideoReport) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v videoReportDo) CreateInBatches(values []*model.VideoReport, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v videoReportDo) Save(values ...*model.VideoReport) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v videoReportDo) First() (*model.VideoReport, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoReport), nil
	}
}

func (v videoReportDo) Take() (*model.VideoReport, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoReport), nil
	}
}

func (v videoReportDo) Last() (*model.VideoReport, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoReport), nil
	}
}

func (v videoReportDo) Find() ([]*model.VideoReport, error) {
	result, err := v.DO.Find()
	return result.([]*model.VideoReport), err
}

func (v videoReportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VideoReport, err error) {
	buf := make([]*model.VideoReport, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v videoReportDo) FindInBatches(result *[]*model.VideoReport, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v videoReportDo) Attrs(attrs ...field.AssignExpr) *videoReportDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v videoReportDo) Assign(attrs ...field.AssignExpr) *videoReportDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v videoReportDo) Joins(fields ...field.RelationField) *videoReportDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v videoReportDo) Preload(fields ...field.RelationField) *videoReportDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v videoReportDo) FirstOrInit() (*model.VideoReport, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoReport), nil
	}
}

func (v videoReportDo) FirstOrCreate() (*model.VideoReport, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoReport), nil
	}
}

func (v videoReportDo) FindByPage(offset int, limit int) (result []*model.VideoReport, count int64, err error) {
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

func (v videoReportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v videoReportDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v videoReportDo) Delete(models ...*model.VideoReport) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *videoReportDo) withDO(do gen.Dao) *videoReportDo {
	v.DO = *do.(*gen.DO)
	return v
}
