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

	"github.com/feihua/zero-admin/rpc/sys/gen/model"
)

func newSysOperateLog(db *gorm.DB, opts ...gen.DOOption) sysOperateLog {
	_sysOperateLog := sysOperateLog{}

	_sysOperateLog.sysOperateLogDo.UseDB(db, opts...)
	_sysOperateLog.sysOperateLogDo.UseModel(&model.SysOperateLog{})

	tableName := _sysOperateLog.sysOperateLogDo.TableName()
	_sysOperateLog.ALL = field.NewAsterisk(tableName)
	_sysOperateLog.ID = field.NewInt64(tableName, "id")
	_sysOperateLog.Title = field.NewString(tableName, "title")
	_sysOperateLog.OperationType = field.NewString(tableName, "operation_type")
	_sysOperateLog.OperationName = field.NewString(tableName, "operation_name")
	_sysOperateLog.RequestMethod = field.NewString(tableName, "request_method")
	_sysOperateLog.OperationURL = field.NewString(tableName, "operation_url")
	_sysOperateLog.OperationParams = field.NewString(tableName, "operation_params")
	_sysOperateLog.OperationResponse = field.NewString(tableName, "operation_response")
	_sysOperateLog.OperationStatus = field.NewInt32(tableName, "operation_status")
	_sysOperateLog.DeptName = field.NewString(tableName, "dept_name")
	_sysOperateLog.UseTime = field.NewInt64(tableName, "use_time")
	_sysOperateLog.Browser = field.NewString(tableName, "browser")
	_sysOperateLog.Os = field.NewString(tableName, "os")
	_sysOperateLog.OperationIP = field.NewString(tableName, "operation_ip")
	_sysOperateLog.OperationTime = field.NewTime(tableName, "operation_time")

	_sysOperateLog.fillFieldMap()

	return _sysOperateLog
}

// sysOperateLog 系统操作日志表
type sysOperateLog struct {
	sysOperateLogDo sysOperateLogDo

	ALL               field.Asterisk
	ID                field.Int64  // 编号
	Title             field.String // 系统模块
	OperationType     field.String // 操作类型
	OperationName     field.String // 操作人员
	RequestMethod     field.String // 请求方式
	OperationURL      field.String // 操作方法
	OperationParams   field.String // 请求参数
	OperationResponse field.String // 响应参数
	OperationStatus   field.Int32  // 操作状态
	DeptName          field.String // 部门名称
	UseTime           field.Int64  // 执行时长(毫秒)
	Browser           field.String // 浏览器
	Os                field.String // 操作信息
	OperationIP       field.String // 操作地址
	OperationTime     field.Time   // 操作时间

	fieldMap map[string]field.Expr
}

func (s sysOperateLog) Table(newTableName string) *sysOperateLog {
	s.sysOperateLogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysOperateLog) As(alias string) *sysOperateLog {
	s.sysOperateLogDo.DO = *(s.sysOperateLogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysOperateLog) updateTableName(table string) *sysOperateLog {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Title = field.NewString(table, "title")
	s.OperationType = field.NewString(table, "operation_type")
	s.OperationName = field.NewString(table, "operation_name")
	s.RequestMethod = field.NewString(table, "request_method")
	s.OperationURL = field.NewString(table, "operation_url")
	s.OperationParams = field.NewString(table, "operation_params")
	s.OperationResponse = field.NewString(table, "operation_response")
	s.OperationStatus = field.NewInt32(table, "operation_status")
	s.DeptName = field.NewString(table, "dept_name")
	s.UseTime = field.NewInt64(table, "use_time")
	s.Browser = field.NewString(table, "browser")
	s.Os = field.NewString(table, "os")
	s.OperationIP = field.NewString(table, "operation_ip")
	s.OperationTime = field.NewTime(table, "operation_time")

	s.fillFieldMap()

	return s
}

func (s *sysOperateLog) WithContext(ctx context.Context) ISysOperateLogDo {
	return s.sysOperateLogDo.WithContext(ctx)
}

func (s sysOperateLog) TableName() string { return s.sysOperateLogDo.TableName() }

func (s sysOperateLog) Alias() string { return s.sysOperateLogDo.Alias() }

func (s sysOperateLog) Columns(cols ...field.Expr) gen.Columns {
	return s.sysOperateLogDo.Columns(cols...)
}

func (s *sysOperateLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysOperateLog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["id"] = s.ID
	s.fieldMap["title"] = s.Title
	s.fieldMap["operation_type"] = s.OperationType
	s.fieldMap["operation_name"] = s.OperationName
	s.fieldMap["request_method"] = s.RequestMethod
	s.fieldMap["operation_url"] = s.OperationURL
	s.fieldMap["operation_params"] = s.OperationParams
	s.fieldMap["operation_response"] = s.OperationResponse
	s.fieldMap["operation_status"] = s.OperationStatus
	s.fieldMap["dept_name"] = s.DeptName
	s.fieldMap["use_time"] = s.UseTime
	s.fieldMap["browser"] = s.Browser
	s.fieldMap["os"] = s.Os
	s.fieldMap["operation_ip"] = s.OperationIP
	s.fieldMap["operation_time"] = s.OperationTime
}

func (s sysOperateLog) clone(db *gorm.DB) sysOperateLog {
	s.sysOperateLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysOperateLog) replaceDB(db *gorm.DB) sysOperateLog {
	s.sysOperateLogDo.ReplaceDB(db)
	return s
}

type sysOperateLogDo struct{ gen.DO }

type ISysOperateLogDo interface {
	gen.SubQuery
	Debug() ISysOperateLogDo
	WithContext(ctx context.Context) ISysOperateLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysOperateLogDo
	WriteDB() ISysOperateLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysOperateLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysOperateLogDo
	Not(conds ...gen.Condition) ISysOperateLogDo
	Or(conds ...gen.Condition) ISysOperateLogDo
	Select(conds ...field.Expr) ISysOperateLogDo
	Where(conds ...gen.Condition) ISysOperateLogDo
	Order(conds ...field.Expr) ISysOperateLogDo
	Distinct(cols ...field.Expr) ISysOperateLogDo
	Omit(cols ...field.Expr) ISysOperateLogDo
	Join(table schema.Tabler, on ...field.Expr) ISysOperateLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysOperateLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysOperateLogDo
	Group(cols ...field.Expr) ISysOperateLogDo
	Having(conds ...gen.Condition) ISysOperateLogDo
	Limit(limit int) ISysOperateLogDo
	Offset(offset int) ISysOperateLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysOperateLogDo
	Unscoped() ISysOperateLogDo
	Create(values ...*model.SysOperateLog) error
	CreateInBatches(values []*model.SysOperateLog, batchSize int) error
	Save(values ...*model.SysOperateLog) error
	First() (*model.SysOperateLog, error)
	Take() (*model.SysOperateLog, error)
	Last() (*model.SysOperateLog, error)
	Find() ([]*model.SysOperateLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysOperateLog, err error)
	FindInBatches(result *[]*model.SysOperateLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysOperateLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysOperateLogDo
	Assign(attrs ...field.AssignExpr) ISysOperateLogDo
	Joins(fields ...field.RelationField) ISysOperateLogDo
	Preload(fields ...field.RelationField) ISysOperateLogDo
	FirstOrInit() (*model.SysOperateLog, error)
	FirstOrCreate() (*model.SysOperateLog, error)
	FindByPage(offset int, limit int) (result []*model.SysOperateLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysOperateLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysOperateLogDo) Debug() ISysOperateLogDo {
	return s.withDO(s.DO.Debug())
}

func (s sysOperateLogDo) WithContext(ctx context.Context) ISysOperateLogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysOperateLogDo) ReadDB() ISysOperateLogDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysOperateLogDo) WriteDB() ISysOperateLogDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysOperateLogDo) Session(config *gorm.Session) ISysOperateLogDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysOperateLogDo) Clauses(conds ...clause.Expression) ISysOperateLogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysOperateLogDo) Returning(value interface{}, columns ...string) ISysOperateLogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysOperateLogDo) Not(conds ...gen.Condition) ISysOperateLogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysOperateLogDo) Or(conds ...gen.Condition) ISysOperateLogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysOperateLogDo) Select(conds ...field.Expr) ISysOperateLogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysOperateLogDo) Where(conds ...gen.Condition) ISysOperateLogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysOperateLogDo) Order(conds ...field.Expr) ISysOperateLogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysOperateLogDo) Distinct(cols ...field.Expr) ISysOperateLogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysOperateLogDo) Omit(cols ...field.Expr) ISysOperateLogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysOperateLogDo) Join(table schema.Tabler, on ...field.Expr) ISysOperateLogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysOperateLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysOperateLogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysOperateLogDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysOperateLogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysOperateLogDo) Group(cols ...field.Expr) ISysOperateLogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysOperateLogDo) Having(conds ...gen.Condition) ISysOperateLogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysOperateLogDo) Limit(limit int) ISysOperateLogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysOperateLogDo) Offset(offset int) ISysOperateLogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysOperateLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysOperateLogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysOperateLogDo) Unscoped() ISysOperateLogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysOperateLogDo) Create(values ...*model.SysOperateLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysOperateLogDo) CreateInBatches(values []*model.SysOperateLog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysOperateLogDo) Save(values ...*model.SysOperateLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysOperateLogDo) First() (*model.SysOperateLog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperateLog), nil
	}
}

func (s sysOperateLogDo) Take() (*model.SysOperateLog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperateLog), nil
	}
}

func (s sysOperateLogDo) Last() (*model.SysOperateLog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperateLog), nil
	}
}

func (s sysOperateLogDo) Find() ([]*model.SysOperateLog, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysOperateLog), err
}

func (s sysOperateLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysOperateLog, err error) {
	buf := make([]*model.SysOperateLog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysOperateLogDo) FindInBatches(result *[]*model.SysOperateLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysOperateLogDo) Attrs(attrs ...field.AssignExpr) ISysOperateLogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysOperateLogDo) Assign(attrs ...field.AssignExpr) ISysOperateLogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysOperateLogDo) Joins(fields ...field.RelationField) ISysOperateLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysOperateLogDo) Preload(fields ...field.RelationField) ISysOperateLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysOperateLogDo) FirstOrInit() (*model.SysOperateLog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperateLog), nil
	}
}

func (s sysOperateLogDo) FirstOrCreate() (*model.SysOperateLog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperateLog), nil
	}
}

func (s sysOperateLogDo) FindByPage(offset int, limit int) (result []*model.SysOperateLog, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysOperateLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysOperateLogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysOperateLogDo) Delete(models ...*model.SysOperateLog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysOperateLogDo) withDO(do gen.Dao) *sysOperateLogDo {
	s.DO = *do.(*gen.DO)
	return s
}
