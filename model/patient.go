package model

import (
	"errors"
	"gorm.io/gorm"
	"hello/orthodonticsAdmin/global/my_errors"
	"hello/orthodonticsAdmin/global/variable"
	"hello/orthodonticsAdmin/utils/myencrypt"
	"hello/orthodonticsAdmin/utils/zapLogger"
	"time"
)

type Patient struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Identity string `json:"identity"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Birthdate time.Time `json:"birthdate" time_format:"2006-01-02"`
	Phone string `json:"phone"`
}

// AddPatient 创建病人帐号
func (p *Patient) AddPatient()error  {
	result := variable.GormDbMysql.Table("patients").Create(p)
	if err := result.Error; err != nil {
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsCreatePatient,err.Error())
		defer variable.ZapLogger.Sync()
		return err
	}
	return nil

}

// ModifyPatient 修改病人帐号
func (p *Patient) ModifyPatient()error  {
	result := variable.GormDbMysql.Table("patients").Updates(*p)
	if err := result.Error; err != nil {
		zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsModifyPatient,err.Error())
		defer variable.ZapLogger.Sync()
		return err
	}
	return nil
}

// FindPatientByID 通过id查找病人帐号
func FindPatientByID(id int)(*Patient, error){
	p := Patient{}
	result := variable.GormDbMysql.Table("patients").Where("id = ?",id).First(&p)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(my_errors.ErrorsPatientRecordNotFound)
		} else {
			zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsModifyPatient,err.Error() )
			defer variable.ZapLogger.Sync()
			return nil, err
		}
	}
	return &p, nil
}

// FindPatientByUserName 通过username查找病人帐号
func FindPatientByUserName(username string)(*Patient, error){
	p := Patient{}
	result := variable.GormDbMysql.Table("patients").Where("username = ?",username).First(&p)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			zapLogger.ZapSugarLogger().Errorf("%s, error:%s",my_errors.ErrorsModifyPatient,err.Error() )
			defer variable.ZapLogger.Sync()
			return nil, err
		}
	}
	return &p, nil
}


// DeletePatientByID 通过id删除病人帐号
func DeletePatientByID(id int)(bool, error){
	variable.GormDbMysql.Table("patients").Delete(&Patient{},id)
	return true, nil
}

// GetPatientList 获取所有的病人信息
func GetPatientList()[]*Patient {
	patients := []*Patient{}
	variable.GormDbMysql.Table("patients").Select("id,username,first_name,last_name,birthdate,phone,identity").Find(patients)
	return patients
}

// VerifyPatientExist 验证是否存在该用户
func VerifyPatientExist(username string)(exist bool, p *Patient){
	var err error
	p, err = FindPatientByUserName(username)
	if err != nil {
		return false, nil
	}
	return true, p
}

// VerifyPatientInfo 验证用户名密码是否正确
func VerifyPatientInfo(p Patient, password string)bool{
	return myencrypt.CompareHashAndPassword(p.Password,password)
}