package rc

import (
	"encoding/json"
	"github.com/SherClockHolmes/webpush-go"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/mail"
	"github.com/spo-iitk/ras-backend/middleware"
	"github.com/spo-iitk/ras-backend/util"
)

func getAllNoticesHandler(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var notices []Notice

	err := fetchAllNotices(ctx, rid, &notices)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notices)
}

func postNoticeHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var notice Notice
	err = ctx.ShouldBindJSON(&notice)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = CreateNotice(ctx, rid, &notice)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var subscriptions []Subscription
	if err := getSubscriptions(ctx, &subscriptions); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	go func() {
		for _, sub := range subscriptions {
			subData, _ := json.Marshal(sub)
			s := webpush.Subscription{}
			json.Unmarshal(subData, &s)
			go webpush.SendNotification([]byte(notice.Title), &s, &webpush.Options{
				VAPIDPublicKey:  vapidPublicKey,
				VAPIDPrivateKey: vapidPrivateKey,
				TTL:             30,
			})
		}
	}()
	ctx.JSON(http.StatusOK, gin.H{"status": "notice created", "subscriptions": subscriptions})
}

func CreateNotice(ctx *gin.Context, id uint, notice *Notice) error {
	notice.RecruitmentCycleID = uint(id)
	notice.LastReminderAt = 0
	notice.CreatedBy = middleware.GetUserID(ctx)

	return createNotice(ctx, notice)
}

func putNoticeHandler(ctx *gin.Context) {
	var editNoticeRequest Notice

	err := ctx.ShouldBindJSON(&editNoticeRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if editNoticeRequest.RecruitmentCycleID != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Recruitment cycle id is not allowed"})
		return
	}

	if editNoticeRequest.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	err = updateNotice(ctx, &editNoticeRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, editNoticeRequest)
}

func deleteNoticeHandler(ctx *gin.Context) {
	nid := ctx.Param("nid")

	err := removeNotice(ctx, nid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "status"})
}

func postReminderHandler(mail_channel chan mail.Mail) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid, err := util.ParseUint(ctx.Param("rid"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		nid := ctx.Param("nid")

		var notice Notice
		err = fetchNotice(ctx, nid, &notice)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if notice.LastReminderAt > time.Now().Add(-6*time.Hour).UnixMilli() {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Reminder already sent"})
			return
		}

		notice.LastReminderAt = time.Now().UnixMilli()
		err = updateNotice(ctx, &notice)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		emails, err := fetchAllUnfrozenEmails(ctx, rid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mail_channel <- mail.GenerateMails(emails, "Notice: "+notice.Title, notice.Description)

		ctx.JSON(http.StatusOK, gin.H{"status": "mail sent"})
	}
}
