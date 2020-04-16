package token

import (
	"encoding/base64"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
)

type GfMallClaims struct {
	//【JWT ID】     该jwt的唯一ID编号
	//【issuer】     发布者的url地址
	//【issued at】  该jwt的发布时间；unix 时间戳
	//【subject】    该JWT所面向的用户，用于处理特定应用，不是常用的字段
	//【audience】   接受者的url地址
	//【expiration】 该jwt销毁的时间；unix时间戳
	//【not before】 该jwt的使用时间不能早于该时间；unix时间戳
	StandardClaims *jwt.StandardClaims
	RefreshTime    int64 //【The refresh time】 该jwt刷新的时间；unix时间戳
}

type Token struct {
	Claim    *GfMallClaims
	Token    string
	NewToken string
}

//创建Claims
func New(id string) *GfMallClaims {
	timeOut := g.Cfg().GetInt("api.jwt.timeout")

	if timeOut <= 0 {
		timeOut = 3600
	}

	refresh := g.Cfg().GetInt("api.jwt.refresh")

	if refresh <= 0 {
		refresh = timeOut / 2
	}

	var claims GfMallClaims
	standardClaims := new(jwt.StandardClaims)
	standardClaims.Id = id
	standardClaims.ExpiresAt = time.Now().Add(time.Second * time.Duration(timeOut)).Unix()
	standardClaims.IssuedAt = time.Now().Unix()

	claims.RefreshTime = time.Now().Add(time.Second * time.Duration(refresh)).Unix()
	claims.StandardClaims = standardClaims
	return &claims
}

func (c *GfMallClaims) SetIss(issuer string) *GfMallClaims {
	c.StandardClaims.Issuer = issuer
	return c
}

func (c *GfMallClaims) SetSub(subject string) *GfMallClaims {
	c.StandardClaims.Subject = subject
	return c
}

func (c *GfMallClaims) SetAud(audience string) *GfMallClaims {
	c.StandardClaims.Audience = audience
	return c
}

func (c *GfMallClaims) SetNbf(notBefore int64) *GfMallClaims {
	c.StandardClaims.NotBefore = notBefore
	return c
}

func (c *GfMallClaims) Valid() error {
	//标准验证
	return c.StandardClaims.Valid()
}

//创建token
func (c *GfMallClaims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	mySignKeyBytes, err := base64.URLEncoding.DecodeString(g.Cfg().GetString("api.jwt.encryptKey"))
	if err != nil {
		return "", err
	}
	return token.SignedString(mySignKeyBytes)
}

//验证token
func VerifyAuthToken(token string) (*Token, error) {
	mySignKey := g.Cfg().GetString("api.jwt.encryptKey")
	mySignKeyBytes, err := base64.URLEncoding.DecodeString(mySignKey) //需要用和加密时同样的方式转化成对应的字节数组
	if err != nil {
		return nil, err
	}
	var gfMallClaims GfMallClaims

	parseAuth, err := jwt.ParseWithClaims(token, &gfMallClaims, func(*jwt.Token) (interface{}, error) {
		return mySignKeyBytes, nil
	})

	if err != nil {
		return nil, err
	}

	//验证claims
	if err := parseAuth.Claims.Valid(); err != nil {
		return nil, err
	}

	rs := new(Token)
	rs.Claim = &gfMallClaims
	rs.Token = token
	//判断是否需要刷新
	if gfMallClaims.RefreshTime > time.Now().Unix() {
		//生成新token
		newToken, err := New(gfMallClaims.StandardClaims.Id).CreateToken()
		if err == nil {
			rs.NewToken = newToken
		}
	}
	return rs, nil
}
