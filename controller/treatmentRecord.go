package controller

import (
	"github.com/gin-gonic/gin"
	"hello/orthodonticsAdmin/global/response"
	"hello/orthodonticsAdmin/model"
	"net/http"
	"strconv"
)

// AddTreatmentRecord 添加治疗记录
func AddTreatmentRecord(c *gin.Context)  {
	tr := &model.TreatmentRecord{}
	err := c.ShouldBind(tr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{Msg: "绑定治疗记录参数失败",
			Code: response.ErrorsBindingData})
		return
	}
	err = tr.AddTreatmentRecord()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{Msg: "创建治疗记录失败",
			Code: response.ErrorsAddTreatmentRecord})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": response.Success,
		"data": gin.H{
			"reservationData": *tr,
		},
	})
}

// ModifyTreatmentRecord 修改治疗记录
func ModifyTreatmentRecord(c *gin.Context){
	tr := &model.TreatmentRecord{}
	err := c.ShouldBind(tr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{Msg: "绑定治疗记录参数失败",
			Code: response.ErrorsBindingData})
		return
	}
	err = tr.ModifyTreatmentRecord()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{Msg: "修改治疗记录失败",
			Code: response.ErrorsModifyTreatmentRecord})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": response.Success,
		"data": gin.H{
			"reservationData": *tr,
		},
	})
}

// DeleteTreatmentRecord 删除治疗记录
func DeleteTreatmentRecord(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  "非法的id",
			Code: response.ErrorsInvalidData,
		})
		return
	}
	err = model.DeleteTreatmentRecord(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{
			Msg:  "删除治疗记录失败",
			Code: response.ErrorsDeleteTreatmentRecord,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code" : response.Success,
		"data" : "",
	})
}

// GetTreatmentRecords 获取所有的治疗记录
func GetTreatmentRecords(c *gin.Context)  {
	treatmentRecords :=  model.GetAllTreatment()
	c.JSON(http.StatusOK, gin.H{
		"code": response.Success,
		"data": treatmentRecords,
	})
}

// GetTreatmentRecordsByPatientId 根据PatientId获取所有的治疗记录
func GetTreatmentRecordsByPatientId(c *gin.Context)  {
	patientId, err := strconv.Atoi(c.Param("patientId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  "非法的patientId",
			Code: response.ErrorsInvalidData,
		})
		return
	}
	treatmentRecords, err :=  model.GetTreatmentRecordsByPatientId(patientId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{
			Msg:  "查询预约失败",
			Code: response.ErrorsGetTreatmentRecord,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": response.Success,
		"data": treatmentRecords,
	})
}

// GetTreatmentRecordsByDoctorId 根据DoctorId获取所有的治疗记录
func GetTreatmentRecordsByDoctorId(c *gin.Context)  {
	doctorId:= c.Param("doctorId")
	if doctorId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  "为获取到doctorId",
			Code: response.ErrorsInvalidData,
		})
		return
	}
	treatmentRecords, err :=  model.GetTreatmentRecordsByDoctorId(doctorId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{
			Msg:  "查询预约失败",
			Code: response.ErrorsGetTreatmentRecord,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": response.Success,
		"data": treatmentRecords,
	})
}

// GetTreatmentRecordById 根据Id获取一条治疗记录
func GetTreatmentRecordById(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,response.Response{
			Msg:  "非法的id",
			Code: response.ErrorsInvalidData,
		})
		return
	}
	treatmentRecord, err :=  model.GetTreatmentRecordById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,response.Response{
			Msg:  "查询预约失败",
			Code: response.ErrorsGetTreatmentRecord,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": response.Success,
		"data": treatmentRecord,
	})
}