package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _UserStatusesMgr struct {
	*_BaseMgr
}

// UserStatusesMgr open func
func UserStatusesMgr(db *gorm.DB) *_UserStatusesMgr {
	if db == nil {
		panic(fmt.Errorf("UserStatusesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserStatusesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_statuses"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserStatusesMgr) GetTableName() string {
	return "user_statuses"
}

// Get 获取
func (obj *_UserStatusesMgr) Get() (result UserStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserStatusesMgr) Gets() (results []*UserStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserStatusesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStatus status获取
func (obj *_UserStatusesMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_UserStatusesMgr) GetByOption(opts ...Option) (result UserStatuses, err error) {
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
func (obj *_UserStatusesMgr) GetByOptions(opts ...Option) (results []*UserStatuses, err error) {
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
func (obj *_UserStatusesMgr) GetFromID(id int) (result UserStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserStatusesMgr) GetBatchFromID(ids []int) (results []*UserStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_UserStatusesMgr) GetFromStatus(status string) (results []*UserStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_UserStatusesMgr) GetBatchFromStatus(statuss []string) (results []*UserStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserStatusesMgr) FetchByPrimaryKey(id int) (result UserStatuses, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
