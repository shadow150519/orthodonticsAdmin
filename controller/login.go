package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hello/orthodonticsAdmin/global/response"
	"hello/orthodonticsAdmin/middleware/token"
	"hello/orthodonticsAdmin/model"
	"hello/orthodonticsAdmin/utils/jwttoken"
	"net/http"
)

func PatientLogin(c *gin.Context)  {
	loginData := struct{
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}{}
	err := c.ShouldBind(&loginData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  fmt.Sprintf("登陆失败, err:%v",err.Error()),
			Code: response.ErrorsPatientLogin,
		})
		return
	}

	// 验证该用户是否存在
	exists, p := model.VerifyPatientExist(loginData.Username)
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{Msg: "用户不存在",Code: response.ErrorsPatientNotFound})
		return
	}

	// 验证密码是否正确
	ok := model.VerifyPatientInfo(*p,loginData.Password)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{Msg: "密码错误", Code: response.ErrorsPatientPassword})
		return
	}

	customClaims := token.CustomClaims{
		Id:             int(p.ID),
		Username:       p.Username,
		Phone:          p.Phone,
		StandardClaims: jwt.StandardClaims{},
	}
	token, err := jwttoken.CreateToken(customClaims)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{Msg: "token生成失败",Code: response.ErrorsTokenGenerate})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"code" :response.Success,
		"data" : gin.H{
			"token" : token,
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



func DoctorLogin(c *gin.Context)  {
	loginData := struct{
		WorkerId string `json:"worker_id" form:"worker_id" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}{}
	err := c.ShouldBind(&loginData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  fmt.Sprintf("登陆失败, err:%v",err.Error()),
			Code: response.ErrorsDoctorLogin,
		})
		return
	}

	// 验证该用户是否存在
	exists := model.VerifyDoctorExist(loginData.WorkerId)
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{Msg: "工号不存在",Code: response.ErrorsDoctorNotFound})
		return
	}
	d, _ := model.FindDoctorByWorkerId(loginData.WorkerId)
	// 验证密码是否正确
	ok := model.VerifyDoctorInfo(*d,loginData.Password)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{Msg: "密码错误", Code: response.ErrorsDoctorPassword})
		return
	}

	customClaims := token.CustomClaims{
		Id:             int(d.ID),
		WorkerId:       d.WorkerId,
		Phone:          d.Phone,
		StandardClaims: jwt.StandardClaims{},
	}
	token, err := jwttoken.CreateToken(customClaims)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{Msg: "token生成失败",Code: response.ErrorsTokenGenerate})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"code" :response.Success,
		"data" : gin.H{
			"token" : token,
			"patientData": gin.H{
				"Id":     d.ID,
				"jobTitle":  d.JobTitle,
				"workerId":  d.WorkerId,
				"firstName": d.FirstName,
				"lastName":  d.LastName,
				"email":     d.Email,
				"birthdate": d.Birthdate,
				"phone":     d.Phone,
			},
		},
	})
}

