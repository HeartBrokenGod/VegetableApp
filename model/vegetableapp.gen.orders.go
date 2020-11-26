package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrdersMgr struct {
	*_BaseMgr
}

// OrdersMgr open func
func OrdersMgr(db *gorm.DB) *_OrdersMgr {
	if db == nil {
		panic(fmt.Errorf("OrdersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrdersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("orders"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrdersMgr) GetTableName() string {
	return "orders"
}

// Get 获取
func (obj *_OrdersMgr) Get() (result Orders, err error) {
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
		if err = obj.New().Table("payment_statuses").Where("id = ?", result.PaymentStatusID).Find(&result.PaymentStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("order_statuses").Where("id = ?", result.StatusID).Find(&result.OrderStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_OrdersMgr) Gets() (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
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
func (obj *_OrdersMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取
func (obj *_OrdersMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithProductID product_id获取
func (obj *_OrdersMgr) WithProductID(productID int) Option {
	return optionFunc(func(o *options) { o.query["product_id"] = productID })
}

// WithRatePerKg rate_per_kg获取
func (obj *_OrdersMgr) WithRatePerKg(ratePerKg int) Option {
	return optionFunc(func(o *options) { o.query["rate_per_kg"] = ratePerKg })
}

// WithQtyInKg qty_in_kg获取
func (obj *_OrdersMgr) WithQtyInKg(qtyInKg int) Option {
	return optionFunc(func(o *options) { o.query["qty_in_kg"] = qtyInKg })
}

// WithTotalPrice total_price获取
func (obj *_OrdersMgr) WithTotalPrice(totalPrice int) Option {
	return optionFunc(func(o *options) { o.query["total_price"] = totalPrice })
}

// WithPaymentStatusID payment_status_id获取
func (obj *_OrdersMgr) WithPaymentStatusID(paymentStatusID int) Option {
	return optionFunc(func(o *options) { o.query["payment_status_id"] = paymentStatusID })
}

// WithStatusID status_id获取
func (obj *_OrdersMgr) WithStatusID(statusID int) Option {
	return optionFunc(func(o *options) { o.query["status_id"] = statusID })
}

// WithCreatedOn created_on获取
func (obj *_OrdersMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_OrdersMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_OrdersMgr) GetByOption(opts ...Option) (result Orders, err error) {
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
		if err = obj.New().Table("payment_statuses").Where("id = ?", result.PaymentStatusID).Find(&result.PaymentStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("order_statuses").Where("id = ?", result.StatusID).Find(&result.OrderStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_OrdersMgr) GetByOptions(opts ...Option) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
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
func (obj *_OrdersMgr) GetFromID(id int) (result Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
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
		if err = obj.New().Table("payment_statuses").Where("id = ?", result.PaymentStatusID).Find(&result.PaymentStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("order_statuses").Where("id = ?", result.StatusID).Find(&result.OrderStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromID(ids []int) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_OrdersMgr) GetFromUserID(userID int) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromUserID(userIDs []int) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProductID 通过product_id获取内容
func (obj *_OrdersMgr) GetFromProductID(productID int) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProductID 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromProductID(productIDs []int) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRatePerKg 通过rate_per_kg获取内容
func (obj *_OrdersMgr) GetFromRatePerKg(ratePerKg int) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRatePerKg 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromRatePerKg(ratePerKgs []int) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromQtyInKg 通过qty_in_kg获取内容
func (obj *_OrdersMgr) GetFromQtyInKg(qtyInKg int) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromQtyInKg 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromQtyInKg(qtyInKgs []int) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTotalPrice 通过total_price获取内容
func (obj *_OrdersMgr) GetFromTotalPrice(totalPrice int) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price = ?", totalPrice).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTotalPrice 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromTotalPrice(totalPrices []int) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price IN (?)", totalPrices).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPaymentStatusID 通过payment_status_id获取内容
func (obj *_OrdersMgr) GetFromPaymentStatusID(paymentStatusID int) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment_status_id = ?", paymentStatusID).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPaymentStatusID 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromPaymentStatusID(paymentStatusIDs []int) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment_status_id IN (?)", paymentStatusIDs).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatusID 通过status_id获取内容
func (obj *_OrdersMgr) GetFromStatusID(statusID int) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id = ?", statusID).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatusID 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromStatusID(statusIDs []int) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id IN (?)", statusIDs).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_OrdersMgr) GetFromCreatedOn(createdOn time.Time) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedOn 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_OrdersMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedOn 批量唯一主键查找
func (obj *_OrdersMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
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
func (obj *_OrdersMgr) FetchByPrimaryKey(id int) (result Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
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
		if err = obj.New().Table("payment_statuses").Where("id = ?", result.PaymentStatusID).Find(&result.PaymentStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("order_statuses").Where("id = ?", result.StatusID).Find(&result.OrderStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkOrdersUserID  获取多个内容
func (obj *_OrdersMgr) FetchIndexByFkOrdersUserID(userID int) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkOrdersProductID  获取多个内容
func (obj *_OrdersMgr) FetchIndexByFkOrdersProductID(productID int) (results []*Orders, err error) {
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkOrdersPaymentStatusID  获取多个内容
func (obj *_OrdersMgr) FetchIndexByFkOrdersPaymentStatusID(paymentStatusID int) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment_status_id = ?", paymentStatusID).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkOrdersStatusID  获取多个内容
func (obj *_OrdersMgr) FetchIndexByFkOrdersStatusID(statusID int) (results []*Orders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id = ?", statusID).Find(&results).Error
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
			if err = obj.New().Table("payment_statuses").Where("id = ?", results[i].PaymentStatusID).Find(&results[i].PaymentStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("order_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].OrderStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
