package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"junsonjack.cn/go_vue/common"
	"junsonjack.cn/go_vue/dto"
	"junsonjack.cn/go_vue/model"
	"junsonjack.cn/go_vue/response"
	"junsonjack.cn/go_vue/util"
)

func Register (c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := c.PostForm("name")
	telephone  := c.PostForm("telephone")
	password := c.PostForm("password")

	// 校验参数
	if len(telephone) != 11 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位有效数字")
		return
	}
	if len(password) < 6 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	} 
	// 如果名称没有传，就给一个10位随机字符串
	if len(name) == 0 {
		name = util.RandomStr(10)
		
	}
	log.Println(name,telephone,password)

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone){
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户已经存在")
		return
	}
	// 用户不存在就新建用户
	// 密码加密
	hasedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		response.Response(c,http.StatusUnprocessableEntity,500,nil,"加密错误")
	}
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasedPassword),
	}
	DB.Create(&newUser)
	// 返回结果

	response.Success(c,"注册成功",nil)
}

func Login(c *gin.Context){
	DB := common.GetDB()
	// 获取参数
	telephone  := c.PostForm("telephone")
	password := c.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位有效数字")
		return
	}
	if len(password) < 6 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	} 
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?",telephone).First(&user)
	if user.ID == 0{
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户不存在")
		return 
	}

	// 判断密码是否正确
	// CompareHashAndPassword第一个参数是加密后的密码，第二参数是未加密的密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)) ; err != nil {
		response.Fail(c,"密码错误",nil)
		return 
	}

	// 发放token
	token , err:= common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"code": 500,
			"msg": "系统异常",
		})
		log.Printf("token generate error: %v",err)
		return
	}

	// 返回结果
	response.Success(c,"登录成功",gin.H{"token": token})
}

// 获取用户信息
func Info (c *gin.Context){
	user,_ := c.Get("user")
	response.Success(c,"获取用户信息成功",gin.H{"user": dto.ToUserDto(user.(model.User))})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false

}


