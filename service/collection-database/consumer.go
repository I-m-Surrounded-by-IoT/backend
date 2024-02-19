package collection_database

import (
	"strconv"
	"time"

	database "github.com/I-m-Surrounded-by-IoT/backend/api/collection-database"
	"github.com/I-m-Surrounded-by-IoT/backend/service/collection-database/model"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var _ sarama.ConsumerGroupHandler = (*CollectionConsumer)(nil)

type CollectionConsumer struct {
	db *dbUtils
}

func NewCollectionConsumer(dbs *CollectionDatabaseService) *CollectionConsumer {
	return &CollectionConsumer{
		db: dbs.db,
	}
}

func (s *CollectionConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *CollectionConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *CollectionConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Infof("start consume device report...")
	msgCh := claim.Messages()
	var record database.CollectionRecord
	for {
		select {
		case msg := <-msgCh:
			log.Infof("receive device log message: key: %v, value: %v", string(msg.Key), string(msg.Value))
			_, err := strconv.ParseUint(string(msg.Key), 10, 64)
			if err != nil {
				log.Errorf("failed to parse device id (%s): %v", msg.Key, err)
				continue
			}
			err = proto.Unmarshal(msg.Value, &record)
			if err != nil {
				log.Errorf("failed to unmarshal device report (%s): %v", msg.Value, err)
				continue
			}
			err = s.db.CreateCollection(&model.CollectionRecord{
				DeviceID:  record.DeviceId,
				Timestamp: time.UnixMicro(record.Timestamp),
				GeoPoint: model.GeoPoint{
					Lat: record.GeoPoint.Lat,
					Lng: record.GeoPoint.Lng,
				},
				Temperature: record.Temperature,
			})
			if err != nil {
				log.Errorf("failed to create collection record: %v", err)
				continue
			}
		case <-session.Context().Done():
			log.Infof("stop consume device report...")
			return nil
		}
	}
}
