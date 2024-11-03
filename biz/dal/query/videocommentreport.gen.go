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

func newVideoCommentReport(db *gorm.DB, opts ...gen.DOOption) videoCommentReport {
	_videoCommentReport := videoCommentReport{}

	_videoCommentReport.videoCommentReportDo.UseDB(db, opts...)
	_videoCommentReport.videoCommentReportDo.UseModel(&model.VideoCommentReport{})

	tableName := _videoCommentReport.videoCommentReportDo.TableName()
	_videoCommentReport.ALL = field.NewAsterisk(tableName)
	_videoCommentReport.ID = field.NewInt64(tableName, "id")
	_videoCommentReport.UserID = field.NewInt64(tableName, "user_id")
	_videoCommentReport.CommentID = field.NewInt64(tableName, "comment_id")
	_videoCommentReport.Reason = field.NewString(tableName, "reason")
	_videoCommentReport.Label = field.NewString(tableName, "label")
	_videoCommentReport.CreatedAt = field.NewInt64(tableName, "created_at")
	_videoCommentReport.Status = field.NewString(tableName, "status")
	_videoCommentReport.ResolvedAt = field.NewInt64(tableName, "resolved_at")
	_videoCommentReport.AdminID = field.NewInt64(tableName, "admin_id")

	_videoCommentReport.fillFieldMap()

	return _videoCommentReport
}

// videoCommentReport 视频评论举报表
type videoCommentReport struct {
	videoCommentReportDo videoCommentReportDo

	ALL        field.Asterisk
	ID         field.Int64  // 举报ID
	UserID     field.Int64  // 用户ID
	CommentID  field.Int64  // 视频评论ID
	Reason     field.String // 举报原因
	Label      field.String // 举报标签
	CreatedAt  field.Int64  // 创建时间
	Status     field.String // 举报状态
	ResolvedAt field.Int64  // 解决时间
	AdminID    field.Int64  // 管理员ID

	fieldMap map[string]field.Expr
}

func (v videoCommentReport) Table(newTableName string) *videoCommentReport {
	v.videoCommentReportDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v videoCommentReport) As(alias string) *videoCommentReport {
	v.videoCommentReportDo.DO = *(v.videoCommentReportDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *videoCommentReport) updateTableName(table string) *videoCommentReport {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewInt64(table, "id")
	v.UserID = field.NewInt64(table, "user_id")
	v.CommentID = field.NewInt64(table, "comment_id")
	v.Reason = field.NewString(table, "reason")
	v.Label = field.NewString(table, "label")
	v.CreatedAt = field.NewInt64(table, "created_at")
	v.Status = field.NewString(table, "status")
	v.ResolvedAt = field.NewInt64(table, "resolved_at")
	v.AdminID = field.NewInt64(table, "admin_id")

	v.fillFieldMap()

	return v
}

func (v *videoCommentReport) WithContext(ctx context.Context) *videoCommentReportDo {
	return v.videoCommentReportDo.WithContext(ctx)
}

func (v videoCommentReport) TableName() string { return v.videoCommentReportDo.TableName() }

func (v videoCommentReport) Alias() string { return v.videoCommentReportDo.Alias() }

func (v videoCommentReport) Columns(cols ...field.Expr) gen.Columns {
	return v.videoCommentReportDo.Columns(cols...)
}

func (v *videoCommentReport) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *videoCommentReport) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 9)
	v.fieldMap["id"] = v.ID
	v.fieldMap["user_id"] = v.UserID
	v.fieldMap["comment_id"] = v.CommentID
	v.fieldMap["reason"] = v.Reason
	v.fieldMap["label"] = v.Label
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["status"] = v.Status
	v.fieldMap["resolved_at"] = v.ResolvedAt
	v.fieldMap["admin_id"] = v.AdminID
}

func (v videoCommentReport) clone(db *gorm.DB) videoCommentReport {
	v.videoCommentReportDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v videoCommentReport) replaceDB(db *gorm.DB) videoCommentReport {
	v.videoCommentReportDo.ReplaceDB(db)
	return v
}

type videoCommentReportDo struct{ gen.DO }

func (v videoCommentReportDo) Debug() *videoCommentReportDo {
	return v.withDO(v.DO.Debug())
}

func (v videoCommentReportDo) WithContext(ctx context.Context) *videoCommentReportDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v videoCommentReportDo) ReadDB() *videoCommentReportDo {
	return v.Clauses(dbresolver.Read)
}

func (v videoCommentReportDo) WriteDB() *videoCommentReportDo {
	return v.Clauses(dbresolver.Write)
}

func (v videoCommentReportDo) Session(config *gorm.Session) *videoCommentReportDo {
	return v.withDO(v.DO.Session(config))
}

func (v videoCommentReportDo) Clauses(conds ...clause.Expression) *videoCommentReportDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v videoCommentReportDo) Returning(value interface{}, columns ...string) *videoCommentReportDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v videoCommentReportDo) Not(conds ...gen.Condition) *videoCommentReportDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v videoCommentReportDo) Or(conds ...gen.Condition) *videoCommentReportDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v videoCommentReportDo) Select(conds ...field.Expr) *videoCommentReportDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v videoCommentReportDo) Where(conds ...gen.Condition) *videoCommentReportDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v videoCommentReportDo) Order(conds ...field.Expr) *videoCommentReportDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v videoCommentReportDo) Distinct(cols ...field.Expr) *videoCommentReportDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v videoCommentReportDo) Omit(cols ...field.Expr) *videoCommentReportDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v videoCommentReportDo) Join(table schema.Tabler, on ...field.Expr) *videoCommentReportDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v videoCommentReportDo) LeftJoin(table schema.Tabler, on ...field.Expr) *videoCommentReportDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v videoCommentReportDo) RightJoin(table schema.Tabler, on ...field.Expr) *videoCommentReportDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v videoCommentReportDo) Group(cols ...field.Expr) *videoCommentReportDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v videoCommentReportDo) Having(conds ...gen.Condition) *videoCommentReportDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v videoCommentReportDo) Limit(limit int) *videoCommentReportDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v videoCommentReportDo) Offset(offset int) *videoCommentReportDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v videoCommentReportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *videoCommentReportDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v videoCommentReportDo) Unscoped() *videoCommentReportDo {
	return v.withDO(v.DO.Unscoped())
}

func (v videoCommentReportDo) Create(values ...*model.VideoCommentReport) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v videoCommentReportDo) CreateInBatches(values []*model.VideoCommentReport, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v videoCommentReportDo) Save(values ...*model.VideoCommentReport) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v videoCommentReportDo) First() (*model.VideoCommentReport, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoCommentReport), nil
	}
}

func (v videoCommentReportDo) Take() (*model.VideoCommentReport, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoCommentReport), nil
	}
}

func (v videoCommentReportDo) Last() (*model.VideoCommentReport, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoCommentReport), nil
	}
}

func (v videoCommentReportDo) Find() ([]*model.VideoCommentReport, error) {
	result, err := v.DO.Find()
	return result.([]*model.VideoCommentReport), err
}

func (v videoCommentReportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VideoCommentReport, err error) {
	buf := make([]*model.VideoCommentReport, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v videoCommentReportDo) FindInBatches(result *[]*model.VideoCommentReport, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v videoCommentReportDo) Attrs(attrs ...field.AssignExpr) *videoCommentReportDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v videoCommentReportDo) Assign(attrs ...field.AssignExpr) *videoCommentReportDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v videoCommentReportDo) Joins(fields ...field.RelationField) *videoCommentReportDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v videoCommentReportDo) Preload(fields ...field.RelationField) *videoCommentReportDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v videoCommentReportDo) FirstOrInit() (*model.VideoCommentReport, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoCommentReport), nil
	}
}

func (v videoCommentReportDo) FirstOrCreate() (*model.VideoCommentReport, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoCommentReport), nil
	}
}

func (v videoCommentReportDo) FindByPage(offset int, limit int) (result []*model.VideoCommentReport, count int64, err error) {
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

func (v videoCommentReportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v videoCommentReportDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v videoCommentReportDo) Delete(models ...*model.VideoCommentReport) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *videoCommentReportDo) withDO(do gen.Dao) *videoCommentReportDo {
	v.DO = *do.(*gen.DO)
	return v
}
