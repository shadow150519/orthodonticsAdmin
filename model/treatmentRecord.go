package model

import (
	"gorm.io/gorm"
	"hello/orthodonticsAdmin/global/my_errors"
	"hello/orthodonticsAdmin/global/variable"
	"hello/orthodonticsAdmin/utils/zapLogger"
	"time"
)

type TreatmentRecord struct {
	gorm.Model
	TreatmentTime time.Time `json:"treatment_time" form:"treatment_time" time_format:"2006-01-02 15:04:05"`
	Remark string `json:"remark" form:"remark" binding:"required"`
	DoctorId string `json:"doctor_id" form:"doctor_id" binding:"required"`
	PatientId int `json:"patient_id" form:"patient_id" binding:"required"`
	State int `json:"state" form:"state" gorm:"default:0"`
}
// AddTreatmentRecord 添加治疗记录
func (tr *TreatmentRecord)AddTreatmentRecord()error{
	if result := variable.GormDbMysql.Table("treatment_records").Create(tr); result.Error != nil{
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsCreateTreatmentRecord,result.Error.Error())
		return result.Error
	}
	return nil
}

// ModifyTreatmentRecord 修改治疗记录
func (tr *TreatmentRecord)ModifyTreatmentRecord()error{
	if result := variable.GormDbMysql.Table("treatment_records").Save(tr); result.Error != nil{
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsModifyTreatmentRecord,result.Error.Error())
		return result.Error
	}
	return nil
}

// DeleteTreatmentRecord 删除治疗记录
func DeleteTreatmentRecord (id int)error{
	if result := variable.GormDbMysql.Table("treatment_records").Delete(&TreatmentRecord{},id); result.Error != nil{
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsDeleteTreatmentRecord,result.Error.Error())
		return result.Error
	}
	return nil
}

// GetAllTreatment 获取所有的治疗记录
func GetAllTreatment()[]*TreatmentRecord  {
	treatmentRecords := make([]*TreatmentRecord,1)
	if result := variable.GormDbMysql.Table("treatment_records").Find(&treatmentRecords); result.Error != nil{
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsQueryTreatmentRecord,result.Error.Error())
		return nil
	}
	return treatmentRecords
}

// GetTreatmentRecordById 根据id找到一条治疗记录
func GetTreatmentRecordById (id int)(*TreatmentRecord, error){
	tr := &TreatmentRecord{}
	if result := variable.GormDbMysql.Table("treatment_records").Where("id = ?",id).First(tr); result.Error != nil{
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsQueryTreatmentRecord,result.Error.Error())
		return nil, result.Error
	}
	return tr, nil
}

// GetTreatmentRecordsByDoctorId 根据DoctorId找到治疗记录
func GetTreatmentRecordsByDoctorId(doctorId string)(trs []*TreatmentRecord,err error)  {
	if result := variable.GormDbMysql.Table("treatment_records").Where("doctor_id = ?",doctorId).Find(&trs);
	result.Error != nil{
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsQueryTreatmentRecord,result.Error.Error())
		return nil, result.Error
	}
	return trs, nil
}

// GetTreatmentRecordsByPatientId 根据PatientId找到治疗记录
func GetTreatmentRecordsByPatientId(patientId int)(trs []*TreatmentRecord,err error)  {
	if result := variable.GormDbMysql.Table("treatment_records").Where("patient_id = ?",patientId).Find(&trs);
		result.Error != nil{
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsQueryTreatmentRecord,result.Error.Error())
		return nil, result.Error
	}
	return trs, nil
}
