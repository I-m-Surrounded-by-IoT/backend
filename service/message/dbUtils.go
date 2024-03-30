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

func (u *dbUtils) Transaction(fn func(db *dbUtils) error) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		return fn(newDBUtils(tx))
	})
}

func (d *dbUtils) CreateMessage(m *model.Message) error {
	return d.db.Create(m).Error
}

func (d *dbUtils) CreateMessages(ms []*model.Message) error {
	return d.db.Create(ms).Error
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

func (d *dbUtils) GetMessageList(userID string, scopes ...utils.Scope) ([]*model.Message, error) {
	var ms []*model.Message
	err := d.db.
		Where("user_id = ?", userID).
		Scopes(scopes...).
		Find(&ms).
		Error
	return ms, err
}

func (d *dbUtils) GetMessageListCount(userID string, scopes ...utils.Scope) (int64, error) {
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

func (d *dbUtils) MarkAllRead(userID string, scopes ...utils.Scope) error {
	return d.db.
		Model(&model.Message{}).
		Where("user_id = ?", userID).
		Scopes(scopes...).
		Update("unread", false).
		Error
}

func (d *dbUtils) GetUnreadMessages(userID string, scopes ...utils.Scope) ([]*model.Message, error) {
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

func (d *dbUtils) GetUnreadMessagesCount(userID string, scopes ...utils.Scope) (int64, error) {
	var count int64
	err := d.db.
		Model(&model.Message{}).
		Where("user_id = ? AND unread = ?", userID, true).
		Scopes(scopes...).
		Count(&count).
		Error
	return count, err
}

func (d *dbUtils) GetUnreadMessagesCountGroupByMessageType(userID string, scopes ...utils.Scope) (map[int32]int64, error) {
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
