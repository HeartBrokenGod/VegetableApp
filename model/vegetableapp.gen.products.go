package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ProductsMgr struct {
	*_BaseMgr
}

// ProductsMgr open func
func ProductsMgr(db *gorm.DB) *_ProductsMgr {
	if db == nil {
		panic(fmt.Errorf("ProductsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("products"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductsMgr) GetTableName() string {
	return "products"
}

// Get 获取
func (obj *_ProductsMgr) Get() (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_categories").Where("id = ?", result.CategoryID).Find(&result.ProductCategories).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("product_statuses").Where("id = ?", result.StatusID).Find(&result.ProductStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_ProductsMgr) Gets() (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ProductsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_ProductsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithColor color获取
func (obj *_ProductsMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}

// WithDescription description获取
func (obj *_ProductsMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithRatePerKg rate_per_kg获取
func (obj *_ProductsMgr) WithRatePerKg(ratePerKg int) Option {
	return optionFunc(func(o *options) { o.query["rate_per_kg"] = ratePerKg })
}

// WithStockQty stock_qty获取
func (obj *_ProductsMgr) WithStockQty(stockQty int) Option {
	return optionFunc(func(o *options) { o.query["stock_qty"] = stockQty })
}

// WithCategoryID category_id获取
func (obj *_ProductsMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithImageURI image_uri获取
func (obj *_ProductsMgr) WithImageURI(imageURI string) Option {
	return optionFunc(func(o *options) { o.query["image_uri"] = imageURI })
}

// WithUserID user_id获取
func (obj *_ProductsMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithStatusID status_id获取
func (obj *_ProductsMgr) WithStatusID(statusID int) Option {
	return optionFunc(func(o *options) { o.query["status_id"] = statusID })
}

// WithCreatedOn created_on获取
func (obj *_ProductsMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_ProductsMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_ProductsMgr) GetByOption(opts ...Option) (result Products, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_categories").Where("id = ?", result.CategoryID).Find(&result.ProductCategories).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("product_statuses").Where("id = ?", result.StatusID).Find(&result.ProductStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductsMgr) GetByOptions(opts ...Option) (results []*Products, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_ProductsMgr) GetFromID(id int) (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_categories").Where("id = ?", result.CategoryID).Find(&result.ProductCategories).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("product_statuses").Where("id = ?", result.StatusID).Find(&result.ProductStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromID(ids []int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_ProductsMgr) GetFromName(name string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromName(names []string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromColor 通过color获取内容
func (obj *_ProductsMgr) GetFromColor(color string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromColor 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromColor(colors []string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDescription 通过description获取内容
func (obj *_ProductsMgr) GetFromDescription(description string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDescription 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromDescription(descriptions []string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRatePerKg 通过rate_per_kg获取内容
func (obj *_ProductsMgr) GetFromRatePerKg(ratePerKg int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate_per_kg = ?", ratePerKg).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRatePerKg 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromRatePerKg(ratePerKgs []int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate_per_kg IN (?)", ratePerKgs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStockQty 通过stock_qty获取内容
func (obj *_ProductsMgr) GetFromStockQty(stockQty int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stock_qty = ?", stockQty).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStockQty 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromStockQty(stockQtys []int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stock_qty IN (?)", stockQtys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCategoryID 通过category_id获取内容
func (obj *_ProductsMgr) GetFromCategoryID(categoryID int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_id = ?", categoryID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCategoryID 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_id IN (?)", categoryIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromImageURI 通过image_uri获取内容
func (obj *_ProductsMgr) GetFromImageURI(imageURI string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("image_uri = ?", imageURI).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromImageURI 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromImageURI(imageURIs []string) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("image_uri IN (?)", imageURIs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_ProductsMgr) GetFromUserID(userID int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromUserID(userIDs []int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatusID 通过status_id获取内容
func (obj *_ProductsMgr) GetFromStatusID(statusID int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id = ?", statusID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatusID 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromStatusID(statusIDs []int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id IN (?)", statusIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_ProductsMgr) GetFromCreatedOn(createdOn time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedOn 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_ProductsMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedOn 批量唯一主键查找
func (obj *_ProductsMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductsMgr) FetchByPrimaryKey(id int) (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_categories").Where("id = ?", result.CategoryID).Find(&result.ProductCategories).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("product_statuses").Where("id = ?", result.StatusID).Find(&result.ProductStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByName primay or index 获取唯一内容
func (obj *_ProductsMgr) FetchUniqueIndexByName(name string, userID int) (result Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ? AND user_id = ?", name, userID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("product_categories").Where("id = ?", result.CategoryID).Find(&result.ProductCategories).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("product_statuses").Where("id = ?", result.StatusID).Find(&result.ProductStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkProductCategoryID  获取多个内容
func (obj *_ProductsMgr) FetchIndexByFkProductCategoryID(categoryID int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("category_id = ?", categoryID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkProductUserID  获取多个内容
func (obj *_ProductsMgr) FetchIndexByFkProductUserID(userID int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkProductStatusID  获取多个内容
func (obj *_ProductsMgr) FetchIndexByFkProductStatusID(statusID int) (results []*Products, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id = ?", statusID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("product_categories").Where("id = ?", results[i].CategoryID).Find(&results[i].ProductCategories).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("product_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].ProductStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
