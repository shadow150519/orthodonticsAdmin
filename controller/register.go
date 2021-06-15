package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hello/orthodonticsAdmin/global/my_errors"
	"hello/orthodonticsAdmin/global/response"
	"hello/orthodonticsAdmin/global/variable"
	"hello/orthodonticsAdmin/model"
	"hello/orthodonticsAdmin/utils/myencrypt"
	"hello/orthodonticsAdmin/utils/zapLogger"
	"net/http"
	"time"
)

func PatientRegister(c *gin.Context) {
	registerData := struct {
		Username  string    `json:"username" form:"username" binding:"required"`
		Password  string    `json:"password" form:"password" binding:"required"`
		FirstName string    `json:"first_name" form:"first_name" binding:"required"`
		LastName  string    `json:"last_name" form:"last_name" binding:"required"`
		Email     string    `json:"email" form:"email"`
		Phone     string    `json:"phone" form:"phone"`
		Birthdate time.Time `json:"birthdate" form:"birthdate" binding:"required" time_format:"2006-01-02"`
		Identity  string    `json:"identity" form:"identity" binding:"required"`
	}{}

	c.ShouldBind(&registerData)
	result := variable.GormDbMysql.Table("patients").Where("identity = ?", registerData.Identity).
		First(&model.Patient{})
	if !errors.Is(result.Error,gorm.ErrRecordNotFound){
		c.AbortWithStatusJSON(http.StatusBadRequest,
			response.Response{Msg: "用户已经存在",Code: response.ErrorsPatientAlreadyExist } )
		return
	}
	hash, err := myencrypt.EncryptPassWord(registerData.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Response{
			Msg:  "创建hash错误",
			Code: -1,
		})
		return
	}
	p := model.Patient{
		Model:     gorm.Model{},
		Username:  registerData.Username,
		Password:  hash,
		Identity:  registerData.Identity,
		FirstName: registerData.FirstName,
		LastName:  registerData.LastName,
		Email:     registerData.Email,
		Birthdate: registerData.Birthdate,
		Phone:     registerData.Phone,
	}
	err = p.AddPatient()
	if err != nil {
		zapLogger.ZapSugarLogger().Warnf("%s, error:%s", my_errors.ErrorsCreatePatient, err.Error())
		defer zapLogger.ZapSugarLogger().Sync()
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Response{Msg: "创建用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": response.Success,
		"data": gin.H{
			"patientData": gin.H{
				"ID":     	 p.ID,
				"username":  p.Username,
				"identity":  p.Identity,
				"firstName": p.FirstName,
				"lastName":  p.LastName,
				"email":     p.Email,
				"birthdate": p.Birthdate,
				"phone":     p.Phone,
			},
		},
	})
}

func DoctorRegister(c *gin.Context) {
	registerData := struct {
		WorkerId  string    `json:"worker_id" form:"worker_id" binding:"required"`
		Password  string    `json:"password" form:"password" binding:"required"`
		FirstName string    `json:"first_name" form:"first_name" binding:"required"`
		LastName  string    `json:"last_name" form:"last_name" binding:"required"`
		Email     string    `json:"email" form:"email"`
		Phone     string    `json:"phone" form:"phone"`
		Birthdate time.Time `json:"birthdate" form:"birthdate" binding:"required" time_format:"2006-01-02"`
		JobTitle  model.Title `json:"job_title" form:"job_title" binding:"required" gorm:"type:int"`
	}{}

	err := c.ShouldBind(&registerData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{Msg: "绑定参数失败" + err.Error(),Code: response.ErrorsBindingData})
		return
	}
	result := variable.GormDbMysql.Table("doctors").Where("worker_id = ?", registerData.WorkerId).
		First(&model.Doctor{})
	if !errors.Is(result.Error,gorm.ErrRecordNotFound){
		c.AbortWithStatusJSON(http.StatusBadRequest,
			response.Response{Msg: "用户已经存在",Code: response.ErrorsDoctorAlreadyExist } )
		return
	}
	hash, err := myencrypt.EncryptPassWord(registerData.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Response{
			Msg:  "创建hash错误",
			Code: -1,
		})
		return
	}
	d := model.Doctor{
		Model:     gorm.Model{},
		WorkerId:  registerData.WorkerId,
		Password:  hash,
		JobTitle:  registerData.JobTitle,
		FirstName: registerData.FirstName,
		LastName:  registerData.LastName,
		Email:     registerData.Email,
		Birthdate: registerData.Birthdate,
		Phone:     registerData.Phone,
	}
	err = model.AddDoctor(&d)
	if err != nil {
		zapLogger.ZapSugarLogger().Warnf("%s, error:%s", my_errors.ErrorsCreateDoctor, err.Error())
		defer zapLogger.ZapSugarLogger().Sync()
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Response{Msg: "创建用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": response.Success,
		"data": gin.H{
			"patientData": gin.H{
				"ID":     d.ID,
				"workerId":  d.WorkerId,
				"jobTitle":  d.JobTitle,
				"firstName": d.FirstName,
				"lastName":  d.LastName,
				"email":     d.Email,
				"birthdate": d.Birthdate,
				"phone":     d.Phone,
			},
		},
	})
}

