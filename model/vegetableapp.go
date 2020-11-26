package model

import (
	"time"
)

// OrderStatuses [...]
type OrderStatuses struct {
	ID     int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Status string `gorm:"column:status;type:varchar(20);not null" json:"status"`
}

// Orders [...]
type Orders struct {
	ID              int             `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	UserID          int             `gorm:"index:Fk_orders_user_id;column:user_id;type:int(11);not null" json:"user_id"`
	Users           Users           `gorm:"joinForeignKey:user_id;foreignKey:id" json:"users_list"`
	ProductID       int             `gorm:"index:Fk_orders_product_id;column:product_id;type:int(11);not null" json:"product_id"`
	Products        Products        `gorm:"joinForeignKey:product_id;foreignKey:id" json:"products_list"`
	RatePerKg       int             `gorm:"column:rate_per_kg;type:int(11);not null" json:"rate_per_kg"`
	QtyInKg         int             `gorm:"column:qty_in_kg;type:int(11);not null" json:"qty_in_kg"`
	TotalPrice      int             `gorm:"column:total_price;type:int(11);not null" json:"total_price"`
	PaymentStatusID int             `gorm:"index:Fk_orders_payment_status_id;column:payment_status_id;type:int(11);not null" json:"payment_status_id"`
	PaymentStatuses PaymentStatuses `gorm:"joinForeignKey:payment_status_id;foreignKey:id" json:"payment_statuses_list"`
	StatusID        int             `gorm:"index:Fk_orders_status_id;column:status_id;type:int(11);not null" json:"status_id"`
	OrderStatuses   OrderStatuses   `gorm:"joinForeignKey:status_id;foreignKey:id" json:"order_statuses_list"`
	CreatedOn       time.Time       `gorm:"column:created_on;type:timestamp;not null" json:"created_on"`
	UpdatedOn       time.Time       `gorm:"column:updated_on;type:timestamp;not null" json:"updated_on"`
}

// PaymentStatuses [...]
type PaymentStatuses struct {
	ID     int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Status string `gorm:"column:status;type:varchar(20);not null" json:"status"`
}

// ProductCategories [...]
type ProductCategories struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Name        string    `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Description string    `gorm:"column:description;type:varchar(500);not null" json:"description"`
	CreatedOn   time.Time `gorm:"column:created_on;type:timestamp;not null" json:"created_on"`
	UpdatedOn   time.Time `gorm:"column:updated_on;type:timestamp;not null" json:"updated_on"`
}

// ProductStatuses [...]
type ProductStatuses struct {
	ID     int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Status string `gorm:"column:status;type:varchar(20);not null" json:"status"`
}

// Products [...]
type Products struct {
	ID                int               `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Name              string            `gorm:"uniqueIndex:name;column:name;type:varchar(50);not null" json:"name"`
	Color             string            `gorm:"column:color;type:varchar(10);not null" json:"color"`
	Description       string            `gorm:"column:description;type:varchar(500);not null" json:"description"`
	RatePerKg         int               `gorm:"column:rate_per_kg;type:int(11);not null" json:"rate_per_kg"`
	StockQty          int               `gorm:"column:stock_qty;type:int(11);not null" json:"stock_qty"`
	CategoryID        int               `gorm:"index:Fk_product_category_id;column:category_id;type:int(11);not null" json:"category_id"`
	ProductCategories ProductCategories `gorm:"joinForeignKey:category_id;foreignKey:id" json:"product_categories_list"`
	ImageURI          string            `gorm:"column:image_uri;type:varchar(256);not null" json:"image_uri"`
	UserID            int               `gorm:"uniqueIndex:name;index:Fk_product_user_id;column:user_id;type:int(11);not null" json:"user_id"`
	Users             Users             `gorm:"joinForeignKey:user_id;foreignKey:id" json:"users_list"`
	StatusID          int               `gorm:"index:Fk_product_status_id;column:status_id;type:int(11);not null" json:"status_id"`
	ProductStatuses   ProductStatuses   `gorm:"joinForeignKey:status_id;foreignKey:id" json:"product_statuses_list"`
	CreatedOn         time.Time         `gorm:"column:created_on;type:timestamp;not null" json:"created_on"`
	UpdatedOn         time.Time         `gorm:"column:updated_on;type:timestamp;not null" json:"updated_on"`
}

// TokenStatuses [...]
type TokenStatuses struct {
	ID     int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Status string `gorm:"column:status;type:varchar(20);not null" json:"status"`
}

// Tokens [...]
type Tokens struct {
	UserID        int           `gorm:"index:Fk_token_user_id;column:user_id;type:int(11);not null" json:"user_id"`
	Users         Users         `gorm:"joinForeignKey:user_id;foreignKey:id" json:"users_list"`
	Token         string        `gorm:"column:token;type:varchar(256);not null" json:"token"`
	StatusID      int           `gorm:"index:Fk_token_status_id;column:status_id;type:int(11);not null" json:"status_id"`
	TokenStatuses TokenStatuses `gorm:"joinForeignKey:status_id;foreignKey:id" json:"token_statuses_list"`
}

// UserCarts [...]
type UserCarts struct {
	UserID    int       `gorm:"uniqueIndex:user_id;column:user_id;type:int(11);not null" json:"user_id"`
	Users     Users     `gorm:"joinForeignKey:user_id;foreignKey:id" json:"users_list"`
	ProductID int       `gorm:"uniqueIndex:user_id;index:Fk_user_cart_product_id;column:product_id;type:int(11);not null" json:"product_id"`
	Products  Products  `gorm:"joinForeignKey:product_id;foreignKey:id" json:"products_list"`
	QtyInKg   int       `gorm:"column:qty_in_kg;type:int(11);not null" json:"qty_in_kg"`
	RatePerKg int       `gorm:"column:rate_per_kg;type:int(11);not null" json:"rate_per_kg"`
	CreatedOn time.Time `gorm:"column:created_on;type:timestamp;not null" json:"created_on"`
}

// UserStatuses [...]
type UserStatuses struct {
	ID     int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Status string `gorm:"column:status;type:varchar(20);not null" json:"status"`
}

// UserTypes [...]
type UserTypes struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	Type        string    `gorm:"column:type;type:varchar(50);not null" json:"type"`
	Description string    `gorm:"column:description;type:varchar(500);not null" json:"description"`
	CreatedOn   time.Time `gorm:"column:created_on;type:timestamp;not null" json:"created_on"`
	UpdatedOn   time.Time `gorm:"column:updated_on;type:timestamp;not null" json:"updated_on"`
}

// Users [...]
type Users struct {
	ID           int          `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	FirstName    string       `gorm:"column:first_name;type:varchar(50);not null" json:"first_name"`
	LastName     string       `gorm:"column:last_name;type:varchar(50);not null" json:"last_name"`
	Email        string       `gorm:"unique;column:email;type:varchar(200);not null" json:"email"`
	Mobile       int64        `gorm:"unique;column:mobile;type:bigint(10);not null" json:"mobile"`
	Address      string       `gorm:"column:address;type:varchar(500);not null" json:"address"`
	Password     string       `gorm:"column:password;type:varchar(256);not null" json:"password"`
	TypeID       int          `gorm:"index:Fk_user_type_id;column:type_id;type:int(11);not null" json:"type_id"`
	UserTypes    UserTypes    `gorm:"joinForeignKey:type_id;foreignKey:id" json:"user_types_list"`
	StatusID     int          `gorm:"index:Fk_user_status_id;column:status_id;type:int(11);not null" json:"status_id"`
	UserStatuses UserStatuses `gorm:"joinForeignKey:status_id;foreignKey:id" json:"user_statuses_list"`
	CreatedOn    time.Time    `gorm:"column:created_on;type:timestamp;not null" json:"created_on"`
	UpdatedOn    time.Time    `gorm:"column:updated_on;type:timestamp;not null" json:"updated_on"`
}
