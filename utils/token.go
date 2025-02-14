package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 安全密钥
var TokenKey string = "6151cab95d38889b4b970372c451b24cdabd9b3b"
var TokenExpiresIn time.Duration = 30 * 24 * time.Hour

type UserClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
	IssuedAt    string `json:"issued_at"`
	ExpiresIn   int64  `json:"expires_in"`
}

func GenerateToken(userID int64, username string) (result *TokenResult, err error) {
	claims := UserClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(TokenKey)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		ZapLog().Error("token", "生成token失败")
		return
	}

	result = &TokenResult{
		AccessToken: tokenString,
		ExpiresAt:   claims.ExpiresAt.Format("2006-01-02 15:04:05"),
		IssuedAt:    claims.IssuedAt.Format("2006-01-02 15:04:05"),
		ExpiresIn:   int64(TokenExpiresIn.Seconds()),
	}

	return
}

// 解析token
func ParseToken(tokenString string) (user UserClaims, err error) {
	// var (
	// 	ErrInvalidKey                = errors.New("key is invalid")
	// 	ErrInvalidKeyType            = errors.New("key is of invalid type")
	// 	ErrHashUnavailable           = errors.New("the requested hash function is unavailable")
	// 	ErrTokenMalformed            = errors.New("token is malformed")
	// 	ErrTokenUnverifiable         = errors.New("token is unverifiable")
	// 	ErrTokenSignatureInvalid     = errors.New("token signature is invalid")
	// 	ErrTokenRequiredClaimMissing = errors.New("token is missing required claim")
	// 	ErrTokenInvalidAudience      = errors.New("token has invalid audience")
	// 	ErrTokenExpired              = errors.New("token is expired")
	// 	ErrTokenUsedBeforeIssued     = errors.New("token used before issued")
	// 	ErrTokenInvalidIssuer        = errors.New("token has invalid issuer")
	// 	ErrTokenInvalidSubject       = errors.New("token has invalid subject")
	// 	ErrTokenNotValidYet          = errors.New("token is not valid yet")
	// 	ErrTokenInvalidId            = errors.New("token has invalid id")
	// 	ErrTokenInvalidClaims        = errors.New("token has invalid claims")
	// 	ErrInvalidType               = errors.New("invalid type for claim")
	// )

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenKey), nil
	})

	if err != nil {
		ZapLog().Error("token", "解析token失败")
		return user, err
	}

	// 验证token
	switch {
	case token.Valid:
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			issuedAt, _ := claims.GetIssuedAt()
			expiresAt, _ := claims.GetExpirationTime()
			issuer, _ := claims.GetIssuer()
			user = UserClaims{
				UserID: int64(claims["user_id"].(float64)),
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: expiresAt,
					IssuedAt:  issuedAt,
					Issuer:    issuer,
				},
			}
			return user, nil
		} else {
			return user, errors.New("invalid type for claim")
		}
	case errors.Is(err, jwt.ErrTokenMalformed):
		return user, errors.New("token is malformed")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return user, errors.New("token signature is invalid")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return user, errors.New("token is expired")
	default:
		return user, err
	}
}
