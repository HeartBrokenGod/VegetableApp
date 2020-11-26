package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderStatusesMgr struct {
	*_BaseMgr
}

// OrderStatusesMgr open func
func OrderStatusesMgr(db *gorm.DB) *_OrderStatusesMgr {
	if db == nil {
		panic(fmt.Errorf("OrderStatusesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderStatusesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("order_statuses"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderStatusesMgr) GetTableName() string {
	return "order_statuses"
}

// Get 获取
func (obj *_OrderStatusesMgr) Get() (result OrderStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderStatusesMgr) Gets() (results []*OrderStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_OrderStatusesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStatus status获取
func (obj *_OrderStatusesMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_OrderStatusesMgr) GetByOption(opts ...Option) (result OrderStatuses, err error) {
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
func (obj *_OrderStatusesMgr) GetByOptions(opts ...Option) (results []*OrderStatuses, err error) {
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
func (obj *_OrderStatusesMgr) GetFromID(id int) (result OrderStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_OrderStatusesMgr) GetBatchFromID(ids []int) (results []*OrderStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_OrderStatusesMgr) GetFromStatus(status string) (results []*OrderStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_OrderStatusesMgr) GetBatchFromStatus(statuss []string) (results []*OrderStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderStatusesMgr) FetchByPrimaryKey(id int) (result OrderStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
