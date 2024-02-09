package rc

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func fetchAllNotices(ctx *gin.Context, rid string, notices *[]Notice) error {
	tx := db.WithContext(ctx).Where("recruitment_cycle_id = ?", rid).Order("created_at desc").Find(notices)
	return tx.Error
}

func createNotice(ctx *gin.Context, notice *Notice) error {
	tx := db.WithContext(ctx).Create(notice)
	return tx.Error
}

func removeNotice(ctx *gin.Context, nid string) error {
	tx := db.WithContext(ctx).Where("id = ?", nid).Delete(&Notice{})
	if tx.RowsAffected == 0 {
		return errors.New("no notice found")
	}
	return tx.Error
}

func updateNotice(ctx *gin.Context, notice *Notice) error {
	tx := db.WithContext(ctx).Where("id = ?", notice.ID).Updates(notice)
	return tx.Error
}

func fetchNotice(ctx *gin.Context, nid string, notice *Notice) error {
	tx := db.WithContext(ctx).Where("id = ?", nid).First(notice)
	return tx.Error
}

func subscribeNotice(ctx *gin.Context, subscription *Subscription) error {
	tx := db.WithContext(ctx).Create(subscription)
	return tx.Error
}

func getSubscriptions(ctx *gin.Context, subscriptions *[]Subscription) error {
	tx := db.WithContext(ctx).Preload("Keys").Find(subscriptions)
	return tx.Error
}
