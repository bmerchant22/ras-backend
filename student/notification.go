package student

import (
	"encoding/json"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	subscription    = `{"endpoint":"https://fcm.googleapis.com/fcm/send/e2RG-M7Tf74:APA91bFYoDpd--uoErVmSaI8b0Pc4ywplrfER2BHFuJ2eGHPcEkXp54jzqIowMCyECJl7oi7qyXT1rk2dpXwdLlDuFt1XfvyodbDKiY2JlFxa0xk5GUohO37rH1IlyIL9ixZd4siPR1e","expirationTime":null,"keys":{"p256dh":"BDjj-Ti1Hw80I5H3THnxgKRj1Lqn6oSGleCZNGRRDjdRKhfsFZUJee7Nypo8KT_O9CTjSQVZv5zTwjPR1sVxO5w","auth":"sAtl5Hv2nqCVmrk1P22FOw"}}`
	vapidPublicKey  = "BD02hnS3Y1WU--EHZ8LTqs19uUf2Jh3_rV-ROU55d0lQkev8P-2g_EZEbFcIN32eiuYRtrSNS9d94sBnlNPvjkw"
	vapidPrivateKey = "Qo4vPl8D77SL4NSXKR6o4QnoV18P79WuFOJVCA9GEjw"
)

func sendNotificationHandler(ctx *gin.Context) {
	s := &webpush.Subscription{}
	json.Unmarshal([]byte(subscription), s)

	// Send Notification
	resp, err := webpush.SendNotification([]byte("Take my penis"), s, &webpush.Options{
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
		return
	}
	defer resp.Body.Close()
}
