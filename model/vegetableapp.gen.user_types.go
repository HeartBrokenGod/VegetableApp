package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _UserTypesMgr struct {
	*_BaseMgr
}

// UserTypesMgr open func
func UserTypesMgr(db *gorm.DB) *_UserTypesMgr {
	if db == nil {
		panic(fmt.Errorf("UserTypesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserTypesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_types"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserTypesMgr) GetTableName() string {
	return "user_types"
}

// Get 获取
func (obj *_UserTypesMgr) Get() (result UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserTypesMgr) Gets() (results []*UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserTypesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithType type获取
func (obj *_UserTypesMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithDescription description获取
func (obj *_UserTypesMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithCreatedOn created_on获取
func (obj *_UserTypesMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_UserTypesMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_UserTypesMgr) GetByOption(opts ...Option) (result UserTypes, err error) {
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
func (obj *_UserTypesMgr) GetByOptions(opts ...Option) (results []*UserTypes, err error) {
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
func (obj *_UserTypesMgr) GetFromID(id int) (result UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserTypesMgr) GetBatchFromID(ids []int) (results []*UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_UserTypesMgr) GetFromType(_type string) (results []*UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_UserTypesMgr) GetBatchFromType(_types []string) (results []*UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_UserTypesMgr) GetFromDescription(description string) (results []*UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量唯一主键查找
func (obj *_UserTypesMgr) GetBatchFromDescription(descriptions []string) (results []*UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_UserTypesMgr) GetFromCreatedOn(createdOn time.Time) (results []*UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error

	return
}

// GetBatchFromCreatedOn 批量唯一主键查找
func (obj *_UserTypesMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error

	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_UserTypesMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error

	return
}

// GetBatchFromUpdatedOn 批量唯一主键查找
func (obj *_UserTypesMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserTypesMgr) FetchByPrimaryKey(id int) (result UserTypes, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
