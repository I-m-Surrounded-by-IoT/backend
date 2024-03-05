package collection

import (
	"strconv"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/service/collection/model"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/zijiren233/stream"
	"google.golang.org/protobuf/proto"
)

var _ sarama.ConsumerGroupHandler = (*CollectionConsumer)(nil)

type CollectionConsumer struct {
	db *dbUtils
}

func NewCollectionConsumer(dbs *CollectionService) *CollectionConsumer {
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
	for {
		select {
		case msg := <-msgCh:
			id, err := strconv.ParseUint(stream.BytesToString(msg.Key), 10, 64)
			if err != nil {
				log.Errorf("failed to parse device id (%s): %v", stream.BytesToString(msg.Key), err)
				continue
			}
			data := &collection.CollectionData{}
			err = proto.Unmarshal(msg.Value, data)
			if err != nil {
				log.Errorf("failed to unmarshal device report (%s): %v", msg.Value, err)
				continue
			}
			err = s.db.CreateCollectionRecord(&model.CollectionRecord{
				DeviceID:  id,
				Timestamp: time.UnixMilli(data.Timestamp),
				GeoPoint: model.GeoPoint{
					Lat: data.GeoPoint.Lat,
					Lon: data.GeoPoint.Lon,
				},
				Temperature: data.Temperature,
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
