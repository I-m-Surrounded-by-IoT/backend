package message

import (
	"context"
	"time"

	messageApi "github.com/I-m-Surrounded-by-IoT/backend/api/message"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/message/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	log "github.com/sirupsen/logrus"
)

type MessageService struct {
	db *dbUtils
	messageApi.UnimplementedMessageServer
}

func NewMessageService(dc *conf.DatabaseServerConfig, lc *conf.MessageConfig, rc *conf.RedisConfig) *MessageService {
	d, err := dbdial.Dial(context.Background(), dc)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	if dc.AutoMigrate {
		log.Infof("auto migrate database...")
		err = d.AutoMigrate(
			new(model.Message),
		)
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	}

	return &MessageService{
		db: newDBUtils(d),
	}
}

func (ms *MessageService) GetUnreadNum(ctx context.Context, req *messageApi.GetUnreadNumReq) (*messageApi.GetUnreadNumResp, error) {
	c, err := ms.db.GetUnreadMessagesCountGroupByMessageType(req.UserId)
	if err != nil {
		return nil, err
	}
	return &messageApi.GetUnreadNumResp{Nums: c}, nil
}

func (ms *MessageService) MarkAllRead(ctx context.Context, req *messageApi.MarkAllReadReq) (*messageApi.Empty, error) {
	err := ms.db.MarkAllRead(req.UserId)
	if err != nil {
		return nil, err
	}
	return &messageApi.Empty{}, nil
}

func (ms *MessageService) SendMessage(ctx context.Context, req *messageApi.SendMessageReq) (*messageApi.Empty, error) {
	messages := make([]*model.Message, len(req.UserId))
	for i, userid := range req.UserId {
		m := model.Message{
			UserID:      userid,
			Timestamp:   time.UnixMilli(req.Payload.Timestamp),
			MessageType: req.Payload.MessageType,
			Title:       req.Payload.Title,
			Content:     req.Payload.Content,
		}
		messages[i] = &m
	}
	return &messageApi.Empty{}, ms.db.Transaction(func(db *dbUtils) error {
		return db.CreateMessages(messages)
	})
}

func messageRecordFromModel(m *model.Message) *messageApi.MessageRecord {
	return &messageApi.MessageRecord{
		Id:         m.ID,
		CreateTime: m.CreatedAt.UnixMilli(),
		UpdateTime: m.UpdateAt.UnixMilli(),
		Unread:     m.Unread.Bool,
		UserId:     m.UserID,
		Payload: &messageApi.MessagePayload{
			Timestamp:   m.Timestamp.UnixMilli(),
			MessageType: m.MessageType,
			Title:       m.Title,
			Content:     m.Content,
		},
	}
}

func (ms *MessageService) GetMessage(ctx context.Context, req *messageApi.GetMessageReq) (*messageApi.MessageRecord, error) {
	msg, err := ms.db.GetMessageByID(req.Id)
	if err != nil {
		return nil, err
	}
	return messageRecordFromModel(msg), nil
}

func (ms *MessageService) GetMessageList(ctx context.Context, req *messageApi.GetMessageListReq) (*messageApi.GetMessageListResp, error) {
	scopes := []utils.Scope{
		utils.WithUserIDEq(req.UserId),
	}
	if req.UnreadOnly {
		scopes = append(scopes, model.WithUnread())
	}
	if req.StartTime != 0 {
		scopes = append(scopes, utils.WithTimestampAfter(req.StartTime))
	}
	if req.EndTime != 0 {
		scopes = append(scopes, utils.WithTimestampBefore(req.EndTime))
	}
	msgs, err := ms.db.GetMessageList(req.UserId, scopes...)
	if err != nil {
		return nil, err
	}
	count, err := ms.db.GetMessageListCount(
		req.UserId,
		append(
			scopes,
			utils.WithPageAndPageSize(req.Page, req.Size),
		)...,
	)
	if err != nil {
		return nil, err
	}
	records := make([]*messageApi.MessageRecord, len(msgs))
	for i, m := range msgs {
		m.Content = ""
		records[i] = messageRecordFromModel(m)
	}
	return &messageApi.GetMessageListResp{Records: records, Total: count}, nil
}
