package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductStatusesMgr struct {
	*_BaseMgr
}

// ProductStatusesMgr open func
func ProductStatusesMgr(db *gorm.DB) *_ProductStatusesMgr {
	if db == nil {
		panic(fmt.Errorf("ProductStatusesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductStatusesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_statuses"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductStatusesMgr) GetTableName() string {
	return "product_statuses"
}

// Get 获取
func (obj *_ProductStatusesMgr) Get() (result ProductStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductStatusesMgr) Gets() (results []*ProductStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ProductStatusesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStatus status获取
func (obj *_ProductStatusesMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_ProductStatusesMgr) GetByOption(opts ...Option) (result ProductStatuses, err error) {
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
func (obj *_ProductStatusesMgr) GetByOptions(opts ...Option) (results []*ProductStatuses, err error) {
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
func (obj *_ProductStatusesMgr) GetFromID(id int) (result ProductStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ProductStatusesMgr) GetBatchFromID(ids []int) (results []*ProductStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ProductStatusesMgr) GetFromStatus(status string) (results []*ProductStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_ProductStatusesMgr) GetBatchFromStatus(statuss []string) (results []*ProductStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductStatusesMgr) FetchByPrimaryKey(id int) (result ProductStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
