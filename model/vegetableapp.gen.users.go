package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _UsersMgr struct {
	*_BaseMgr
}

// UsersMgr open func
func UsersMgr(db *gorm.DB) *_UsersMgr {
	if db == nil {
		panic(fmt.Errorf("UsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("users"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UsersMgr) GetTableName() string {
	return "users"
}

// Get 获取
func (obj *_UsersMgr) Get() (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user_types").Where("id = ?", result.TypeID).Find(&result.UserTypes).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_statuses").Where("id = ?", result.StatusID).Find(&result.UserStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UsersMgr) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
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
func (obj *_UsersMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFirstName first_name获取
func (obj *_UsersMgr) WithFirstName(firstName string) Option {
	return optionFunc(func(o *options) { o.query["first_name"] = firstName })
}

// WithLastName last_name获取
func (obj *_UsersMgr) WithLastName(lastName string) Option {
	return optionFunc(func(o *options) { o.query["last_name"] = lastName })
}

// WithEmail email获取
func (obj *_UsersMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithMobile mobile获取
func (obj *_UsersMgr) WithMobile(mobile int64) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithAddress address获取
func (obj *_UsersMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithPassword password获取
func (obj *_UsersMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithTypeID type_id获取
func (obj *_UsersMgr) WithTypeID(typeID int) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithStatusID status_id获取
func (obj *_UsersMgr) WithStatusID(statusID int) Option {
	return optionFunc(func(o *options) { o.query["status_id"] = statusID })
}

// WithCreatedOn created_on获取
func (obj *_UsersMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_UsersMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_UsersMgr) GetByOption(opts ...Option) (result Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user_types").Where("id = ?", result.TypeID).Find(&result.UserTypes).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_statuses").Where("id = ?", result.StatusID).Find(&result.UserStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UsersMgr) GetByOptions(opts ...Option) (results []*Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
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
func (obj *_UsersMgr) GetFromID(id int) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user_types").Where("id = ?", result.TypeID).Find(&result.UserTypes).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_statuses").Where("id = ?", result.StatusID).Find(&result.UserStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromID(ids []int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromFirstName 通过first_name获取内容
func (obj *_UsersMgr) GetFromFirstName(firstName string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("first_name = ?", firstName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromFirstName 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromFirstName(firstNames []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("first_name IN (?)", firstNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromLastName 通过last_name获取内容
func (obj *_UsersMgr) GetFromLastName(lastName string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_name = ?", lastName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromLastName 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromLastName(lastNames []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_name IN (?)", lastNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEmail 通过email获取内容
func (obj *_UsersMgr) GetFromEmail(email string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user_types").Where("id = ?", result.TypeID).Find(&result.UserTypes).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_statuses").Where("id = ?", result.StatusID).Find(&result.UserStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromEmail 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromEmail(emails []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMobile 通过mobile获取内容
func (obj *_UsersMgr) GetFromMobile(mobile int64) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user_types").Where("id = ?", result.TypeID).Find(&result.UserTypes).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_statuses").Where("id = ?", result.StatusID).Find(&result.UserStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromMobile 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromMobile(mobiles []int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAddress 通过address获取内容
func (obj *_UsersMgr) GetFromAddress(address string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address = ?", address).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAddress 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromAddress(addresss []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address IN (?)", addresss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPassword 通过password获取内容
func (obj *_UsersMgr) GetFromPassword(password string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password = ?", password).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPassword 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromPassword(passwords []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password IN (?)", passwords).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_UsersMgr) GetFromTypeID(typeID int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_id = ?", typeID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTypeID 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromTypeID(typeIDs []int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_id IN (?)", typeIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatusID 通过status_id获取内容
func (obj *_UsersMgr) GetFromStatusID(statusID int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id = ?", statusID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatusID 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromStatusID(statusIDs []int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id IN (?)", statusIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_UsersMgr) GetFromCreatedOn(createdOn time.Time) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedOn 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_UsersMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedOn 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
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
func (obj *_UsersMgr) FetchByPrimaryKey(id int) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user_types").Where("id = ?", result.TypeID).Find(&result.UserTypes).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_statuses").Where("id = ?", result.StatusID).Find(&result.UserStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByEmail primay or index 获取唯一内容
func (obj *_UsersMgr) FetchUniqueByEmail(email string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user_types").Where("id = ?", result.TypeID).Find(&result.UserTypes).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_statuses").Where("id = ?", result.StatusID).Find(&result.UserStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByMobile primay or index 获取唯一内容
func (obj *_UsersMgr) FetchUniqueByMobile(mobile int64) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("user_types").Where("id = ?", result.TypeID).Find(&result.UserTypes).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_statuses").Where("id = ?", result.StatusID).Find(&result.UserStatuses).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkUserTypeID  获取多个内容
func (obj *_UsersMgr) FetchIndexByFkUserTypeID(typeID int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_id = ?", typeID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkUserStatusID  获取多个内容
func (obj *_UsersMgr) FetchIndexByFkUserStatusID(statusID int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status_id = ?", statusID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("user_types").Where("id = ?", results[i].TypeID).Find(&results[i].UserTypes).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_statuses").Where("id = ?", results[i].StatusID).Find(&results[i].UserStatuses).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
