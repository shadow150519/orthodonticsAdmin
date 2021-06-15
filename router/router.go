package router

import (
	"github.com/gin-gonic/gin"
	"hello/orthodonticsAdmin/controller"
	middleware "hello/orthodonticsAdmin/middleware/cors"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	// 测试路由
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200,"hello world")
	})
	router.StaticFS("/static/img",http.Dir("./storage/img"))

	 patientGroup := router.Group("/patient")
	 {
	 	patientGroup.POST("/login",controller.PatientLogin)
	 	patientGroup.POST("/register",controller.PatientRegister)
	 }

	 reservationGroup := router.Group("/reservation")
	 {
		 reservationGroup.POST("/create",controller.AddReservation)
		 reservationGroup.PUT("/modify/:id",controller.ModifyReservation)
		 reservationGroup.DELETE("cancel/:id",controller.CancelReservationById)
		 reservationGroup.GET("/list/patient/:patientId",controller.GetAllReservationByPatientId)
		 reservationGroup.GET("/list/doctor/:doctorId",controller.GetAllReservationByDoctorId)
	 }

	 doctorGroup := router.Group("/doctor")
	 {
		 doctorGroup.POST("/login",controller.DoctorLogin)
		 doctorGroup.POST("/register",controller.DoctorRegister)
	 }

	 treatmentGroup := router.Group("/treatment")
	 {
	 	treatmentGroup.POST("/create",controller.AddTreatmentRecord)
	 	treatmentGroup.PUT("/modify",controller.ModifyTreatmentRecord)
	 	treatmentGroup.DELETE("/delete/:id",controller.DeleteTreatmentRecord)
	 	treatmentGroup.GET("/record/list",controller.GetTreatmentRecords)
	 	treatmentGroup.GET("/record/id/:id",controller.GetTreatmentRecordById)
	 	treatmentGroup.GET("/record/doctor/:doctorId",controller.GetTreatmentRecordsByDoctorId)
	 	treatmentGroup.GET("/record/patient/:patientId",controller.GetTreatmentRecordsByPatientId)
	 }

	return router
}
