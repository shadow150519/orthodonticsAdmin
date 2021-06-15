package my_errors

const (
	// 系统配置问题
	ErrorsBasePath                  string = "初始化项目根目录失败"
	ErrorsConfigYamlNotExists       string = "config.yaml 文件不存在"
	ErrorsStorageLogsNotExists      string = "storage/logs 目录不存在"
	ErrorsConfigInitFail            string = "初始化配置文件发生错误"
	ErrorsZapLoggerInitFail         string = "初始化配置文件发生错误"
	ErrorsGormInitFail              string = "Gorm 数据库驱动、连接初始化失败"
)

// 数据库问题
const (
	ErrorsUnknownError 				string = "未知的错误"

	ErrorsQueryPatient				string = "查询患者用户失败"
	ErrorsCreatePatient 			string = "创建患者用户失败"
	ErrorsModifyPatient 			string = "修改患者用户失败"
	ErrorsDeletePatient             string = "删除患者用户失败"
	ErrorsPatientRecordNotFound 	string = "该患者用户不存在"

	ErrorsQueryReservation			string = "查询预约失败"
	ErrorsCreateReservation			string = "创建预约失败"
	ErrorsConfilictReservation 		string = "该时段已被预约"
	ErrorsModifyReservation			string = "修改预约失败"
	ErrorsCancelReservation			string = "取消预约失败"
	ErrorsReservationRecordNotFound string = "该预约不存在"

	ErrorsQueryDoctor				string = "查询医生用户失败"
	ErrorsCreateDoctor				string = "创建医生用户失败"
	ErrorsModifyDoctor 				string = "修改医生用户失败"
	ErrorsDeleteDoctor            	string = "删除医生用户失败"
	ErrorsDoctorRecordNotFound 		string = "该医生用户不存在"

	ErrorsQueryTreatmentRecord		string = "查询治疗记录失败"
	ErrorsCreateTreatmentRecord 	string = "创建治疗记录失败"
	ErrorsModifyTreatmentRecord		string = "修改治疗记录失败"
	ErrorsDeleteTreatmentRecord		string = "删除治疗记录失败"
	ErrorsTreatmentRecordNotFound	string = "该治疗记录不存在"
)

//token部分
const (

	ErrorsTokenInvalid      string = "无效的token"
	ErrorsTokenNotActiveYet string = "jwttoken 尚未激活"
	ErrorsTokenMalFormed    string = "jwttoken 格式不正确"
)

