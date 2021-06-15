package response

// 业务返回状态码
const (
	Success               			int = 200
	ErrorsNoAuthorization 			int = -1000 // 未获得授权
	ErrorsInvalidData     			int = -1001 // 用户输入了非法的数据
	ErrorsBindingData 				int = -1002 // 绑定参数失败
	ErrorsTokenGenerate 			int = -1003 // token生成失败

	ErrorsPatientLogin  			int	= -4001 // 登录时内部处理错误
	ErrorPatientRegister 			int	= -4002 // 注册时内部处理错误
	ErrorsPatientNotFound			int = -4003 // 数据库不存在该用户
	ErrorsPatientPassword			int = -4004 // 用户密码错误
	ErrorsPatientAlreadyExist 		int = -4006 // 患者账户已存在

	ErrorsAddReservation			int = -4008 // 添加预约失败
	ErrorsModifyReservation			int = -4009 // 修改预约失败
	ErrorsGetReservation 			int = -4010 // 获取预约失败
	ErrorsCancelReservation			int = -4011 // 取消预约失败

	ErrorsDoctorAlreadyExist		int = -4012 // 医生账户已存在
	ErrorsDoctorLogin 				int = -4013 // 医生登陆失败
	ErrorsDoctorNotFound			int = -4014 // 数据库不存在该用户
	ErrorsDoctorPassword			int = -4015 // 用户密码错误

	ErrorsAddTreatmentRecord		int = -4016 // 添加治疗记录失败
	ErrorsModifyTreatmentRecord		int = -4017 // 修改治疗记录失败
	ErrorsGetTreatmentRecord		int = -4018 // 获取治疗记录失败
	ErrorsDeleteTreatmentRecord		int = -4019 // 删除治疗记录失败


)
