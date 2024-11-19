package dao

import (
	"IMProject/app/user/model"
	"IMProject/common/db"
	userPb "IMProject/pb/user"
	"IMProject/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"math/rand"
	"strings"
	"time"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{db.NewDBClient(ctx)}
}

// UserLogin 用户登录方法
func (dao *UserDao) UserLogin(req *userPb.UserRequest) (model.User, error) {
	user := model.User{}
	user.UserName = strings.TrimSpace(req.UserName) // 去除多余的空格
	password := req.PassWord

	// 检查用户名和密码是否为空
	if user.UserName == "" || password == "" {
		return user, errors.New("用户名或密码为空")
	}

	// 根据用户名查找用户
	err := dao.DB.Where("user_name = ?", user.UserName).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("用户不存在")
		}
		return user, fmt.Errorf("数据库查询出错: %v", err)
	}

	// 验证密码
	if !utils.ValidPassword(password, user.Salt, user.PassWord) {
		return user, errors.New("密码错误")
	}

	// 生成 token（identity 字段）
	str := fmt.Sprintf("%d", time.Now().Unix())
	identity := utils.MD5Encode(str)

	// 更新用户的 identity 字段
	if err := dao.DB.Model(&user).Update("identity", identity).Error; err != nil {
		return user, fmt.Errorf("更新identity字段失败: %v", err)
	}

	// 成功返回用户数据
	return user, nil
}

// CreateUser 用户创建方法
func (dao *UserDao) CreateUser(req *userPb.UserRequest) (err error) {
	var user model.User
	var count int64

	// 检查用户名和密码是否为空
	if req.UserName == "" || req.PassWord == "" {
		return errors.New("用户名或者密码不能为空")
	}

	// 检查两次密码是否一致
	if req.PassWord != req.PasswordConfirm {
		return errors.New("两次密码不一致")
	}

	// 检查用户名是否已存在
	dao.Model(&model.User{}).Where("user_name = ?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("用户名已经存在")
	}

	// 创建新用户
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.UserName = req.UserName
	user.PassWord = utils.MakePassword(req.PassWord, salt)
	user.Salt = salt
	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()

	// 插入用户
	if err = dao.Model(&model.User{}).Create(&user).Error; err != nil {
		return errors.New("插入用户失败")
	}
	return nil
}

// UpdateUser 用户更新
func (dao *UserDao) UpdateUser(req *userPb.UserRequest) (model.User, error) {
	user := model.User{}
	user.ID = uint(req.UserID)
	user.UserName = req.UserName
	user.PassWord = req.PassWord
	user.Phone = req.Phone
	user.Avatar = req.Icon
	user.Email = req.Email

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return user, errors.New("密码错误")
	} else {
		UpdateUser(user)
		return user, errors.New("密码错误")
	}
}

func UpdateUser(user model.User) *gorm.DB {
	return db.DB.Model(&user).Updates(model.User{UserName: user.UserName, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email, Avatar: user.Avatar})
}
