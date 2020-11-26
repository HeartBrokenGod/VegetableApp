package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ProductCategoriesMgr struct {
	*_BaseMgr
}

// ProductCategoriesMgr open func
func ProductCategoriesMgr(db *gorm.DB) *_ProductCategoriesMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCategoriesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCategoriesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_categories"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductCategoriesMgr) GetTableName() string {
	return "product_categories"
}

// Get 获取
func (obj *_ProductCategoriesMgr) Get() (result ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductCategoriesMgr) Gets() (results []*ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ProductCategoriesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_ProductCategoriesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取
func (obj *_ProductCategoriesMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithCreatedOn created_on获取
func (obj *_ProductCategoriesMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_ProductCategoriesMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_ProductCategoriesMgr) GetByOption(opts ...Option) (result ProductCategories, err error) {
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
func (obj *_ProductCategoriesMgr) GetByOptions(opts ...Option) (results []*ProductCategories, err error) {
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
func (obj *_ProductCategoriesMgr) GetFromID(id int) (result ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ProductCategoriesMgr) GetBatchFromID(ids []int) (results []*ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ProductCategoriesMgr) GetFromName(name string) (results []*ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ProductCategoriesMgr) GetBatchFromName(names []string) (results []*ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_ProductCategoriesMgr) GetFromDescription(description string) (results []*ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量唯一主键查找
func (obj *_ProductCategoriesMgr) GetBatchFromDescription(descriptions []string) (results []*ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_ProductCategoriesMgr) GetFromCreatedOn(createdOn time.Time) (results []*ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error

	return
}

// GetBatchFromCreatedOn 批量唯一主键查找
func (obj *_ProductCategoriesMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error

	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_ProductCategoriesMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error

	return
}

// GetBatchFromUpdatedOn 批量唯一主键查找
func (obj *_ProductCategoriesMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductCategoriesMgr) FetchByPrimaryKey(id int) (result ProductCategories, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
