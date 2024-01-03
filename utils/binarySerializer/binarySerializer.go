package fastjsonserializer

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"reflect"

	"gorm.io/gorm/schema"
)

type BinarySerializer struct {
	order binary.ByteOrder
}

func (b *BinarySerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue any) (err error) {
	fieldValue := reflect.New(field.FieldType)

	if dbValue != nil {
		var _bytes []byte
		switch v := dbValue.(type) {
		case []byte:
			_bytes = v
		case string:
			_bytes = []byte(v)
		default:
			return fmt.Errorf("failed to unmarshal JSONB value: %#v", dbValue)
		}

		err = binary.Read(bytes.NewReader(_bytes), b.order, fieldValue.Interface())
	}

	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

func (b *BinarySerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error) {
	buf := bytes.Buffer{}
	err := binary.Write(&buf, b.order, fieldValue)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func init() {
	schema.RegisterSerializer("binary-b", &BinarySerializer{
		order: binary.BigEndian,
	})
	schema.RegisterSerializer("binary-l", &BinarySerializer{
		order: binary.LittleEndian,
	})
}
