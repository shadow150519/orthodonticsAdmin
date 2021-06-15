package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello/orthodonticsAdmin/global/response"
	"hello/orthodonticsAdmin/model"
	"net/http"
	"strconv"
)

// AddReservation 添加预约
func AddReservation(c *gin.Context)  {
	r := &model.Reservation{}
	err := c.ShouldBind(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  fmt.Sprintf("绑定预约参数失败" + err.Error()),
			Code: response.ErrorsBindingData,
		})
		return
	}
	err = model.AddReservation(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{
			Msg:  err.Error(),
			Code: response.ErrorsAddReservation,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": response.Success,
		"data": gin.H{
			"reservationData": *r,
		},
	})
}

// ModifyReservation 修改预约
func ModifyReservation(c *gin.Context)  {
	r := &model.Reservation{}
	err := c.ShouldBind(r)
	id, _ := strconv.Atoi(c.Param("id"))
	r.ID = uint(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  "绑定预约参数失败",
			Code: response.ErrorsBindingData,
		})
		return
	}
	err = model.ModifyReservation(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{
			Msg:  err.Error(),
			Code: response.ErrorsModifyReservation,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": response.Success,
		"data": gin.H{
			"reservationData": *r,
		},
	})
}

// GetAllReservationByPatientId 根据患者id查看他的预约
func GetAllReservationByPatientId(c *gin.Context) {
	patientId, err := strconv.Atoi(c.Param("patientId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  "非法的patientId",
			Code: response.ErrorsInvalidData,
		})
		return
	}
	reservations, err :=  model.GetAllReservationByPatientId(patientId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{
			Msg:  "查询预约失败",
			Code: response.ErrorsGetReservation,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": response.Success,
		"data": reservations,
	})
}

// CancelReservationById 根据Id删除预约
func CancelReservationById(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  "非法的id",
			Code: response.ErrorsInvalidData,
		})
		return
	}
	err = model.CancelReservation(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{
			Msg:  "取消预约失败",
			Code: response.ErrorsCancelReservation,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code" : response.Success,
		"data" : "",
	})

}

// GetAllReservationByDoctorId 根据医生doctorId查看他的预约
func GetAllReservationByDoctorId(c *gin.Context) {
	doctorId, err := strconv.Atoi(c.Param("doctorId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  "非法的doctorId",
			Code: response.ErrorsInvalidData,
		})
		return
	}
	reservations, err :=  model.GetAllReservationsByWorkerId(doctorId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{
			Msg:  "查询预约失败",
			Code: response.ErrorsGetReservation,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": response.Success,
		"data": reservations,
	})
}
