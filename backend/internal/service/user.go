// file: backend/internal/service/user.go
package service

import (
	"errors" // 导入 config 包
	"seckill-project/backend/config"
	"seckill-project/backend/internal/dao"
	"seckill-project/backend/internal/model"
	"seckill-project/backend/internal/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Register 处理用户注册的业务逻辑
func Register(username, password string) error {
	// 1. 检查用户是否已存在
	userExist, err := dao.CheckUserExist(username)
	if err != nil {
		return err // 数据库查询出错
	}
	if userExist {
		return errors.New("用户已存在") // 用户名已被注册
	}

	// 2. 对密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err // 哈希处理出错
	}

	// 3. 创建用户实例
	user := &model.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}

	// 4. 将用户信息存入数据库
	return dao.CreateUser(user)
}

func Login(username, password string) (string, error) {
	// 1. 根据用户名从数据库中查找用户
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		return "", err // 数据库查询出错
	}
	if user == nil {
		return "", errors.New("用户不存在")
	}

	// 2. 验证密码
	// user.PasswordHash 是数据库里存的哈希值
	// []byte(password) 是用户本次输入的明文密码
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		// 密码不匹配
		return "", errors.New("密码错误")
	}

	// 3. 密码正确，生成 JWT
	token, err := generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateToken 生成 JWT
func generateToken(user *model.User) (string, error) {
	// 创建一个我们自己的声明
	claims := types.CustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用我们在配置中设置的 secret 来签名
	return token.SignedString([]byte(config.Conf.JWT.Secret))
}
