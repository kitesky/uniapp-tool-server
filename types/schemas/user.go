package schemas

type UserPaymentAccountReq struct {
	UserID   int64  `json:"-" binding:"required"`
	PayType  string `json:"pay_type" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Account  string `json:"account" binding:"required"`
	BankName string `json:"bank_name"`
}

type UserInfoReq struct {
	UserID int64 `json:"-" binding:"required"`
}

type UserProfileReq struct {
	UserID  int64  `json:"-" binding:"required"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Gender  string `json:"gender"`
	Website string `json:"website"`
	Bio     string `json:"bio"`
}

type UserInfoRes struct {
	MMID      int64           `json:"mmid"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Balance   float64         `json:"balance"`
	Score     int64           `json:"score"`
	Reward    float64         `json:"reward"`
	Point     int64           `json:"point"`
	Gender    string          `json:"gender"`
	Avatar    string          `json:"avatar"`
	Bio       string          `json:"bio"`
	CreatedAt Datetime        `json:"created_at"`
	VIP       *UserInfoVIPRes `json:"vip"`
}

type SignUpReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type SignInReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRes struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
	IssuedAt    string `json:"issued_at"`
	ExpiresIn   int64  `json:"expires_in"`
}

type ChangePasswordReq struct {
	UserID          int64  `json:"-" binding:"required"`
	OldPassword     string `json:"old_password" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"comfirm_password" binding:"required"`
}

type WechatAppLoginReq struct {
	Code      string `json:"code"`
	AvatarUrl string `json:"avatar_url"`
	NickName  string `json:"nick_name"`
	InviteID  int64  `json:"invite_id"`
	OpenID    string `json:"open_id"`
}
