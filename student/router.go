package student

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/ras"
	"net/http"
)

func StudentRouter(r *gin.Engine) {
	student := r.Group("/api/student")
	{
		student.PUT("", updateStudentHandler)
		student.GET("", getStudentHandler)
	}
}

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/student")
	{
		admin.DELETE("/:sid", deleteStudentHandler)
		admin.GET("", getAllStudentsHandler)
		admin.GET("/limited", getLimitedStudentsHandler)
		admin.PUT("", updateStudentByIDHandler)
		admin.GET("/:sid", getStudentByIDHandler)
		admin.PUT("/:sid/verify", verifyStudentHandler)
		admin.GET("/:sid/history", ras.PlaceHolderController)
	}
}

func NotificationRouter(r *gin.Engine) {
	notif := r.Group("/api/notifications")
	{
		//notification.POST("/subscribe", subscribeHandler)
		notif.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello Bhaii")
		})
		notif.POST("/send-notifications", sendNotificationHandler)
	}
}
