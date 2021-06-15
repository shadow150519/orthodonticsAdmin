package model

import (
	"errors"
	"gorm.io/gorm"
	"hello/orthodonticsAdmin/global/my_errors"
	"hello/orthodonticsAdmin/global/variable"
	"hello/orthodonticsAdmin/utils/zapLogger"
	"time"
)



type Reservation struct {
	gorm.Model
	ReservationTime time.Time `json:"reservation_time" form:"reservation_time" binding:"required" time_format:"2006-01-02 15:04"`
	EndTime time.Time `json:"end_time" form:"end_time" binding:"required" time_format:"2006-01-02 15:04"`
	PatientId int `json:"patient_id" form:"patient_id" binding:"required"`
	State ReservationState `json:"state" form:"state" binding:"required" gorm:"type:int"`
	Remark string `json:"remark" form:"remark"`
}



// AddReservation 添加预约
func AddReservation(r *Reservation)error  {
	if isConfilict := VerifyReservationConfilct(r); isConfilict{
		return errors.New(my_errors.ErrorsConfilictReservation)
	}

	result := variable.GormDbMysql.Table("reservations").Create(r)
	if err := result.Error; err != nil {
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsCreateReservation, err.Error())
		return err
	}
	return nil
}

// VerifyReservationConfilct 确定预约是否冲突
func VerifyReservationConfilct(r *Reservation)bool{
	var count int64
	result := variable.GormDbMysql.Raw("select count(*) from reservations " +
		"where reservation_time < ? and end_time > ?", r.EndTime,r.ReservationTime).Count(&count)
	if err := result.Error; err != nil {
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsUnknownError,err.Error())
		return true
	}
	return count != 0
}

// ModifyReservation  修改预约
func ModifyReservation(r *Reservation)error{
	if isConfilict := VerifyModifyReservationConflict(r); isConfilict{
		return errors.New(my_errors.ErrorsConfilictReservation)
	}

	result := variable.GormDbMysql.Updates(r)
	if err := result.Error; err != nil {
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsModifyReservation,err.Error())
		return err
	}
	return nil
}

// IsReservationExist 根据ID判断预约是否存在
func IsReservationExist(id int)bool{
	var count int64
	result := variable.GormDbMysql.Table("reservations").Where("id = ?",id).Count(&count)
	if err := result.Error; err != nil{
		if errors.Is(result.Error,gorm.ErrRecordNotFound) {
			return false
		}
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsQueryReservation,err.Error())
		return false
	}
	return true
}


// CancelReservation 取消预约
func CancelReservation(id int)error{
	if exist := IsReservationExist(id); !exist{
		return errors.New(my_errors.ErrorsReservationRecordNotFound)
	}
	if result := variable.GormDbMysql.Table("reservations").Delete(&Reservation{},id);result.Error != nil{
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsCancelReservation, result.Error.Error())
		return result.Error
	}
	return nil
}

// GetAllReservationByPatientId 根据id获取到所有的预约
func GetAllReservationByPatientId(patientId int)( []Reservation,  error){
	var reservations  = make([]Reservation,1)
	if result := variable.GormDbMysql.Where("patient_id = ?",patientId).Find(&reservations); result.Error != nil{
		if errors.Is(result.Error,gorm.ErrRecordNotFound) {
			return reservations, nil
		} else {
			zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsQueryPatient,result.Error.Error())
			return reservations, result.Error
		}
	}
	return reservations, nil
}

// VerifyModifyReservationConflict 确定预约是否冲突
func VerifyModifyReservationConflict(r *Reservation)bool{
	var count int64
	result := variable.GormDbMysql.Raw("select count(*) from reservations " +
		"where reservation_time < ? and end_time > ? and id <> ? ", r.EndTime,r.ReservationTime,r.ID).Count(&count)
	if err := result.Error; err != nil {
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsUnknownError,err.Error())
		return true
	}
	return count != 0
}

// GetAllReservationsByWorkerId 根据 WorkerId来获取所有预约
func GetAllReservationsByWorkerId(workerId int)([]Reservation,  error){
	var reservations  = make([]Reservation,1)
	if result := variable.GormDbMysql.Where("worker_id = ?",workerId).Find(&reservations); result.Error != nil{
		if errors.Is(result.Error,gorm.ErrRecordNotFound) {
			return reservations, nil
		} else {
			zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsQueryDoctor,result.Error.Error())
			return reservations, result.Error
		}
	}
	return reservations, nil
}

// GetAllReservationsByPatientId 根据 WorkerId来获取所有预约
func GetAllReservationsByPatientId(patientId int)([]Reservation,  error){
	var reservations  = make([]Reservation,1)
	if result := variable.GormDbMysql.Where("worker_id = ?",patientId).Find(&reservations); result.Error != nil{
		if errors.Is(result.Error,gorm.ErrRecordNotFound) {
			return reservations, nil
		} else {
			zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsQueryDoctor,result.Error.Error())
			return reservations, result.Error
		}
	}
	return reservations, nil
}
