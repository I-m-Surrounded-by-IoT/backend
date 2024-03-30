package message

import (
	"github.com/I-m-Surrounded-by-IoT/backend/service/message/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"gorm.io/gorm"
)

type dbUtils struct {
	db *gorm.DB
}

func newDBUtils(db *gorm.DB) *dbUtils {
	return &dbUtils{db: db}
}

func (d *dbUtils) CreateMessage(m *model.Message) error {
	return d.db.Create(m).Error
}

func (d *dbUtils) GetMessageByID(id uint64) (*model.Message, error) {
	var m model.Message
	err := d.db.First(&m, id).Error
	if err != nil {
		return nil, err
	}
	err = d.markMessageAsRead(id)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (d *dbUtils) GetMessageList(userID uint64, scopes ...utils.Scope) ([]*model.Message, error) {
	var ms []*model.Message
	err := d.db.
		Where("user_id = ?", userID).
		Scopes(scopes...).
		Find(&ms).
		Error
	return ms, err
}

func (d *dbUtils) GetMessageListCount(userID uint64, scopes ...utils.Scope) (int64, error) {
	var count int64
	err := d.db.
		Model(&model.Message{}).
		Where("user_id = ?", userID).
		Scopes(scopes...).
		Count(&count).
		Error
	return count, err
}

func (d *dbUtils) markMessageAsRead(id uint64) error {
	return d.db.Model(&model.Message{}).Where("id = ?", id).Update("unread", false).Error
}

func (d *dbUtils) MarkAllRead(userID uint64, scopes ...utils.Scope) error {
	return d.db.
		Model(&model.Message{}).
		Where("user_id = ?", userID).
		Scopes(scopes...).
		Update("unread", false).
		Error
}

func (d *dbUtils) GetUnreadMessages(userID uint64, scopes ...utils.Scope) ([]*model.Message, error) {
	var ms []*model.Message
	err := d.db.
		Where("user_id = ? AND unread = ?", userID, true).
		Scopes(scopes...).
		Find(&ms).
		Error
	if err != nil {
		return nil, err
	}
	for _, m := range ms {
		err = d.markMessageAsRead(m.ID)
		if err != nil {
			return nil, err
		}
	}
	return ms, nil
}

func (d *dbUtils) GetUnreadMessagesCount(userID uint64, scopes ...utils.Scope) (int64, error) {
	var count int64
	err := d.db.
		Model(&model.Message{}).
		Where("user_id = ? AND unread = ?", userID, true).
		Scopes(scopes...).
		Count(&count).
		Error
	return count, err
}

func (d *dbUtils) GetUnreadMessagesCountGroupByMessageType(userID uint64, scopes ...utils.Scope) (map[int32]int64, error) {
	var counts []*struct {
		MessageType int32
		Count       int64
	}
	err := d.db.
		Model(&model.Message{}).
		Select("message_type, COUNT(*) as count").
		Where("user_id = ? AND unread = ?", userID, true).
		Scopes(scopes...).
		Group("message_type").
		Scan(&counts).
		Error
	if err != nil {
		return nil, err
	}
	result := make(map[int32]int64, len(counts))
	for _, c := range counts {
		result[c.MessageType] = c.Count
	}
	return result, nil
}
