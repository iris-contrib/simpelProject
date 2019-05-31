package model

import (
	"time"

	"github.com/shopspring/decimal"

	"gopkg.in/guregu/null.v3"
)

//User model ....
type User struct {
	ID           int32     `db:"id"`
	FirstName    string    `db:"firstName"`
	LastName     string    `db:"lastName"`
	Email        string    `db:"email"`
	UserName     string    `db:"userName"`
	Password     string    `db:"password"`
	Mobile       string    `db:"mobile"`
	NationalCode string    `db:"nationalCode"`
	CreateDate   time.Time `db:"createDate"`
	UpdateDate   time.Time `db:"updateDate"`
	IsDelete     bool      `db:"isDelete"`
}

//Address model ....
type Address struct {
	ID         int32     `db:"id"`
	PostalCode string    `db:"postalCode"`
	Address    string    `db:"address"`
	Phone      string    `db:"phone"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//Category model ....
type Category struct {
	ID         int32     `db:"id"`
	Name       string    `db:"name"`
	ParentID   int32     `db:"parentId"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//CategoryProduct model ....
type CategoryProduct struct {
	ID         int32     `db:"id"`
	ProductID  int32     `db:"productId"`
	CategoryID int32     `db:"categoryId"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//Feature model ....
type Feature struct {
	ID         int32     `db:"id"`
	Name       string    `db:"name"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//FeatureValue model ....
type FeatureValue struct {
	ID         int32     `db:"id"`
	Value      string    `db:"value"`
	FeatureID  int32     `db:"featureId"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//Order model ....
type Order struct {
	ID                int32           `db:"id"`
	PaymentType       int32           `db:"paymentType"`
	AddressID         int32           `db:"addressId"`
	TotalPricePayable decimal.Decimal `db:"totalPricePayable"`
	TtalPrice         decimal.Decimal `db:"totalPrice"`
	Amount            int32           `db:"amount"`
	Workflow          int32           `db:"workflow"`
	SerialNumber      string          `db:"serialNumber"`
	PaymentStatus     bool            `db:"paymentStatus"`
	CreateDate        time.Time       `db:"createDate"`
	UpdateDate        time.Time       `db:"updateDate"`
	IsDelete          bool            `db:"isDelete"`
}

//OrderProduct model ....
type OrderProduct struct {
	ID         int32     `db:"id"`
	OrderID    int32     `db:"orderId"`
	ProductID  int32     `db:"productId"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//Product model ....
type Product struct {
	ID          int32           `db:"id"`
	FullName    string          `db:"fullName"`
	Price       decimal.Decimal `db:"price"`
	Summery     string          `db:"summery"`
	Description string          `db:"description"`
	Status      bool            `db:"status"`
	MainImage   string          `db:"mainImage"`
	CreateDate  time.Time       `db:"createDate"`
	UpdateDate  time.Time       `db:"updateDate"`
	IsDelete    bool            `db:"isDelete"`
}

//ProductComment model ....
type ProductComment struct {
	ID         int32     `db:"id"`
	UserID     int32     `db:"userId"`
	ProductID  int32     `db:"productId"`
	Comment    string    `db:"comment"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//ProductFeature model ....
type ProductFeature struct {
	ID         int32     `db:"id"`
	FeatureID  int32     `db:"featureId"`
	ProductID  int32     `db:"productId"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//ProductFeatureValue model ....
type ProductFeatureValue struct {
	ID             int32     `db:"id"`
	FeatureValueID int32     `db:"featureValueId"`
	ProductID      int32     `db:"productId"`
	CreateDate     time.Time `db:"createDate"`
	UpdateDate     time.Time `db:"updateDate"`
	IsDelete       bool      `db:"isDelete"`
}

//ProductImage model ....
type ProductImage struct {
	ID         int32     `db:"id"`
	Image      string    `db:"image"`
	ProductID  int32     `db:"productId"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//ProductPoint model ....
type ProductPoint struct {
	ID         int32     `db:"id"`
	Point      int32     `db:"point"`
	UserID     null.Int  `db:"userId"`
	ProductID  int32     `db:"productId"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//ProductPriceHistory model ....
type ProductPriceHistory struct {
	ID         int32           `db:"id"`
	Price      decimal.Decimal `db:"price"`
	ProductID  int32           `db:"productId"`
	CreateDate time.Time       `db:"createDate"`
	UpdateDate time.Time       `db:"updateDate"`
	IsDelete   bool            `db:"isDelete"`
}

//ProductPromotion model ....
type ProductPromotion struct {
	ID              int32           `db:"id"`
	Name            string          `db:"name"`
	ProductID       int32           `db:"productId"`
	DisCountPercent int32           `db:"disCountPercent"`
	Price           decimal.Decimal `db:"price"`
	Using           int32           `db:"using"`
	Status          bool            `db:"status"`
	CreateDate      time.Time       `db:"createDate"`
	UpdateDate      time.Time       `db:"updateDate"`
	IsDelete        bool            `db:"isDelete"`
}

//Role model ....
type Role struct {
	ID         int32     `db:"id"`
	Name       string    `db:"name"`
	Value      int32     `db:"value"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//Slider model ....
type Slider struct {
	ID         int32     `db:"id"`
	Image      string    `db:"image"`
	Status     bool      `db:"status"`
	URL        string    `db:"url"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//Transcation model ....
type Transcation struct {
	ID                int32           `db:"id"`
	TranscationNumber string          `db:"transcationNumber"`
	StatusCode        int32           `db:"statusCode"`
	Price             decimal.Decimal `db:"price"`
	OrderID           int32           `db:"orderId"`
	CreateDate        time.Time       `db:"createDate"`
	UpdateDate        time.Time       `db:"updateDate"`
	IsDelete          bool            `db:"isDelete"`
}

//UserAddress model ....
type UserAddress struct {
	ID         int32     `db:"id"`
	UserID     int32     `db:"userId"`
	AddressID  int32     `db:"addressId"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//UserFavorateProduct model ....
type UserFavorateProduct struct {
	ID         int32     `db:"id"`
	UserID     int32     `db:"userId"`
	ProductID  int32     `db:"productId"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}

//UserRole model ....
type UserRole struct {
	ID         int32     `db:"id"`
	UserID     int32     `db:"userId"`
	RoleID     int32     `db:"roleId"`
	CreateDate time.Time `db:"createDate"`
	UpdateDate time.Time `db:"updateDate"`
	IsDelete   bool      `db:"isDelete"`
}
