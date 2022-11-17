package models

type Book struct {
	BookId       string  `json:"bookid"`
	BookName     string  `json:"bookname,omitempty" validate:"required,min=2,max=100"`
	AuthorName   string  `json:"authorname"`
	Genre        string  `json:"genre" validate:"required"`
	BookRating   float64 `json:"bookrating" validate:"required,min=0,max=5"`
	Status       string  `json:"-" validate:"eq=ADDED|eq=BORROWED|eq=RETURNED"`
	Availability string  `json:"availability,omitempty" validate:"required,eq=YES|eq=NO"`
	CreatedAt    string  `json:"-"`
	UserId       string  `json:"-" bson:"userid"`
	BorrowedBy   string  `json:"-" bson:"borrowedby"`
	LenderName   string  `json:"lendername"`
	LenderRating float64 `json:"lenderrating"`
}

type User struct {
	UserId       string  `json:"userid,omitempty" bson:"userid"`
	UserName     string  `json:"username"`
	UserRating   float64 `json:"userrating" validate:"required,min=0,max=5"`
	Books        []Book  `json:"books"`
	Token        int64   `json:"token" gorm:"default:0"`
	CreationTime string  `json:"creationTime" bson:"creationTime"`
}

type ErrorBody struct {
	Error ErrorMsg `json:"error"`
}

type ErrorMsg struct {
	Message string `json:"message"`
	Type    string `json:"type,omitempty"`
	Code    int    `json:"code"`
}

type ResponseUserObject struct {
	Data User `json:"data"`
}

type ResponseBookObject struct {
	Message string `json:"message"`
	Data    Book   `json:"data"`
}

type ResponseDataSuccess struct {
	Data ResponseDataString `json:"data"`
}

type ResponseDataString struct {
	SuccessMsg string `json:"successMsg"`
}
