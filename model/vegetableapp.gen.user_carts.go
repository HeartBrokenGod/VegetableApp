package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _UserCartsMgr struct {
	*_BaseMgr
}

// UserCartsMgr open func
func UserCartsMgr(db *gorm.DB) *_UserCartsMgr {
	if db == nil {
		panic(fmt.Errorf("UserCartsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserCartsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_carts"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserCartsMgr) GetTableName() string {
	return "user_carts"
}

// Get 获取
func (obj *_UserCartsMgr) Get() (result UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("products").Where("id = ?", result.ProductID).Find(&result.Products).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UserCartsMgr) Gets() (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_UserCartsMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithProductID product_id获取
func (obj *_UserCartsMgr) WithProductID(productID int) Option {
	return optionFunc(func(o *options) { o.query["product_id"] = productID })
}

// WithQtyInKg qty_in_kg获取
func (obj *_UserCartsMgr) WithQtyInKg(qtyInKg int) Option {
	return optionFunc(func(o *options) { o.query["qty_in_kg"] = qtyInKg })
}

// WithRatePerKg rate_per_kg获取
func (obj *_UserCartsMgr) WithRatePerKg(ratePerKg int) Option {
	return optionFunc(func(o *options) { o.query["rate_per_kg"] = ratePerKg })
}

// WithCreatedOn created_on获取
func (obj *_UserCartsMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// GetByOption 功能选项模式获取
func (obj *_UserCartsMgr) GetByOption(opts ...Option) (result UserCarts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("products").Where("id = ?", result.ProductID).Find(&result.Products).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserCartsMgr) GetByOptions(opts ...Option) (results []*UserCarts, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUserID 通过user_id获取内容
func (obj *_UserCartsMgr) GetFromUserID(userID int) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_UserCartsMgr) GetBatchFromUserID(userIDs []int) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProductID 通过product_id获取内容
func (obj *_UserCartsMgr) GetFromProductID(productID int) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_id = ?", productID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProductID 批量唯一主键查找
func (obj *_UserCartsMgr) GetBatchFromProductID(productIDs []int) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_id IN (?)", productIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromQtyInKg 通过qty_in_kg获取内容
func (obj *_UserCartsMgr) GetFromQtyInKg(qtyInKg int) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("qty_in_kg = ?", qtyInKg).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromQtyInKg 批量唯一主键查找
func (obj *_UserCartsMgr) GetBatchFromQtyInKg(qtyInKgs []int) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("qty_in_kg IN (?)", qtyInKgs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRatePerKg 通过rate_per_kg获取内容
func (obj *_UserCartsMgr) GetFromRatePerKg(ratePerKg int) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate_per_kg = ?", ratePerKg).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRatePerKg 批量唯一主键查找
func (obj *_UserCartsMgr) GetBatchFromRatePerKg(ratePerKgs []int) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate_per_kg IN (?)", ratePerKgs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_UserCartsMgr) GetFromCreatedOn(createdOn time.Time) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedOn 批量唯一主键查找
func (obj *_UserCartsMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchUniqueIndexByUserID primay or index 获取唯一内容
func (obj *_UserCartsMgr) FetchUniqueIndexByUserID(userID int, productID int) (result UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ? AND product_id = ?", userID, productID).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("products").Where("id = ?", result.ProductID).Find(&result.Products).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkUserCartProductID  获取多个内容
func (obj *_UserCartsMgr) FetchIndexByFkUserCartProductID(productID int) (results []*UserCarts, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_id = ?", productID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("products").Where("id = ?", results[i].ProductID).Find(&results[i].Products).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
