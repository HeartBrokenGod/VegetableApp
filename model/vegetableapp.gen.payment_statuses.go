package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PaymentStatusesMgr struct {
	*_BaseMgr
}

// PaymentStatusesMgr open func
func PaymentStatusesMgr(db *gorm.DB) *_PaymentStatusesMgr {
	if db == nil {
		panic(fmt.Errorf("PaymentStatusesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PaymentStatusesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("payment_statuses"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PaymentStatusesMgr) GetTableName() string {
	return "payment_statuses"
}

// Get 获取
func (obj *_PaymentStatusesMgr) Get() (result PaymentStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PaymentStatusesMgr) Gets() (results []*PaymentStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PaymentStatusesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStatus status获取
func (obj *_PaymentStatusesMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_PaymentStatusesMgr) GetByOption(opts ...Option) (result PaymentStatuses, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PaymentStatusesMgr) GetByOptions(opts ...Option) (results []*PaymentStatuses, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_PaymentStatusesMgr) GetFromID(id int) (result PaymentStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_PaymentStatusesMgr) GetBatchFromID(ids []int) (results []*PaymentStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_PaymentStatusesMgr) GetFromStatus(status string) (results []*PaymentStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_PaymentStatusesMgr) GetBatchFromStatus(statuss []string) (results []*PaymentStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PaymentStatusesMgr) FetchByPrimaryKey(id int) (result PaymentStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
