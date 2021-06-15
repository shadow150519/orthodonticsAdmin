package model

import (
	"gorm.io/gorm"
	"hello/orthodonticsAdmin/global/my_errors"
	"hello/orthodonticsAdmin/global/variable"
	"hello/orthodonticsAdmin/utils/myencrypt"
	"hello/orthodonticsAdmin/utils/zapLogger"
	"time"
)

type Doctor struct {
	gorm.Model
	WorkerId string `json:"worker_id" form:"worker_id"`
	Email string `json:"email" form:"email"`
	JobTitle Title `json:"job_title" form:"job_title"`
	Phone string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName string `json:"last_name" form:"last_name"`
	Birthdate time.Time `json:"birthdate" form:"birthdate" time_format:"2006-01-02" binding:"required"'`

}


// FindDoctorByWorkerId 通过WorkId查找医生帐号
func FindDoctorByWorkerId(workerId string)(*Doctor, error){
	d := Doctor{}
	result := variable.GormDbMysql.Table("doctors").Where("worker_id = ?",workerId).First(&d)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsDoctorRecordNotFound,err.Error() )
			defer variable.ZapLogger.Sync()
			return nil, err
		}
	}
	return &d, nil
}

// FindDoctorById 通过Id查找医生帐号
func FindDoctorById(Id int)(*Doctor, error){
	d := Doctor{}
	result := variable.GormDbMysql.Table("doctors").Where("worker_id = ?",Id).First(&d)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsDoctorRecordNotFound,err.Error() )
			defer variable.ZapLogger.Sync()
			return nil, err
		}
	}
	return &d, nil
}

// VerifyDoctorExist 判断是否有这个医生帐号
func VerifyDoctorExist(workerId string)(exist bool){
	var count int64
	if result := variable.GormDbMysql.Table("doctors").Where("worker_id = ?", workerId).Count(&count); result.Error != nil{
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsCreateDoctor,result.Error.Error())
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

// AddDoctor 添加医生
func AddDoctor(d *Doctor)error{
	result := variable.GormDbMysql.Table("doctors").Create(d)
	if err := result.Error; err != nil {
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsCreateDoctor,err.Error())
		defer variable.ZapLogger.Sync()
		return err
	}
	return nil
}

// ModifyDoctor 修改医生帐号
func ModifyDoctor(d *Doctor)error  {
	result := variable.GormDbMysql.Table("doctors").Updates(*d)
	if err := result.Error; err != nil {
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsModifyDoctor,err.Error())
		defer variable.ZapLogger.Sync()
		return err
	}
	return nil
}

// GetDoctorList 获取所有的病人信息
func GetDoctorList()[]*Doctor {
	doctors := []*Doctor{}
	variable.GormDbMysql.Table("doctors").Select("id,worker_id,first_name,last_name,birthdate,phone,email,job_title").
		Find(doctors)
	return doctors
}

// VerifyDoctorInfo 验证用户名密码是否正确
func VerifyDoctorInfo(d Doctor, password string)bool{
	return myencrypt.CompareHashAndPassword(d.Password,password)
}

