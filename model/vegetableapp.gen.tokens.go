package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TokensMgr struct {
	*_BaseMgr
}

// TokensMgr open func
func TokensMgr(db *gorm.DB) *_TokensMgr {
	if db == nil {
		panic(fmt.Errorf("TokensMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TokensMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tokens"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TokensMgr) GetTableName() string {
	return "tokens"
}

// Get 获取
func (obj *_TokensMgr) Get() (result Tokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("users").Where("id = ?", result.UserID).Find(&result.Users).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("token_statuses").Where("id = ?", result.StatusID).Find(&result.TokenStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_TokensMgr) Gets() (results []*Tokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("token_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].TokenStatuses).Error; err != nil { //
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
func (obj *_TokensMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithToken token获取
func (obj *_TokensMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}

// WithStatusID status_id获取
func (obj *_TokensMgr) WithStatusID(statusID int) Option {
	return optionFunc(func(o *options) { o.query["status_id"] = statusID })
}

// GetByOption 功能选项模式获取
func (obj *_TokensMgr) GetByOption(opts ...Option) (result Tokens, err error) {
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
		if err = obj.New().Table("token_statuses").Where("id = ?", result.StatusID).Find(&result.TokenStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TokensMgr) GetByOptions(opts ...Option) (results []*Tokens, err error) {
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
			if err = obj.New().Table("token_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].TokenStatuses).Error; err != nil { //
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
func (obj *_TokensMgr) GetFromUserID(userID int) (results []*Tokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("token_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].TokenStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量唯一主键查找
func (obj *_TokensMgr) GetBatchFromUserID(userIDs []int) (results []*Tokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("token_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].TokenStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromToken 通过token获取内容
func (obj *_TokensMgr) GetFromToken(token string) (results []*Tokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token = ?", token).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("token_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].TokenStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromToken 批量唯一主键查找
func (obj *_TokensMgr) GetBatchFromToken(tokens []string) (results []*Tokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token IN (?)", tokens).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("token_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].TokenStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatusID 通过status_id获取内容
func (obj *_TokensMgr) GetFromStatusID(statusID int) (results []*Tokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id = ?", statusID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("token_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].TokenStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatusID 批量唯一主键查找
func (obj *_TokensMgr) GetBatchFromStatusID(statusIDs []int) (results []*Tokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id IN (?)", statusIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("token_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].TokenStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByFkTokenUserID  获取多个内容
func (obj *_TokensMgr) FetchIndexByFkTokenUserID(userID int) (results []*Tokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_id = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("token_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].TokenStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkTokenStatusID  获取多个内容
func (obj *_TokensMgr) FetchIndexByFkTokenStatusID(statusID int) (results []*Tokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id = ?", statusID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("users").Where("id = ?", results[i].UserID).Find(&results[i].Users).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("token_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].TokenStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
