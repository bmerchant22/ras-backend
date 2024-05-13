package rc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

var vapidPublicKey = viper.GetString("NOTIFICATION.VAPIDPUBLICKEY")
var vapidPrivateKey = viper.GetString("NOTIFICATION.VAPIDPRIVATEKEY")

func subscribeNotificationHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": rid})
		return
	}

	email := middleware.GetUserID(ctx)
	var student StudentRecruitmentCycle

	err = fetchStudentByEmailAndRC(ctx, email, rid, &student)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int64

	if err := getDeviceNumberForUser(ctx, &student, count); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if count >= 5 {
		if err := deleteOldestSubscription(ctx, &student); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

	var subscription Subscription

	if err := ctx.ShouldBindJSON(&subscription); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := subscribeNotice(ctx, &subscription); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"error": "Subscription created"})

}
