package main

import (
	"fmt"

	"github.com/twpayne/go-geom/encoding/wkb"
	"github.com/twpayne/go-geom/encoding/wkt"
	"github.com/volcengine/volc-sdk-golang/example/dts/data-subscription-demo/avro/avro"
)

func printRecord(record *avro.Record) {
	fmt.Printf("Operation[%s] ObjectName[%s] SourceTimestamp[%v] Id[%v]\n",
		record.Operation.String(), record.ObjectName.String, record.SourceTimestamp, record.Id)
	switch record.Operation {
	case avro.OperationDDL:
		fmt.Printf("DDL[%s] \n", record.AfterImages.String)
	case avro.OperationINSERT:
		for i := 0; i < len(record.AfterImages.ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject); i++ {
			after, err := getValue(record.AfterImages.ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject[i])
			if err != nil {
				panic(err)
			}
			fmt.Printf("Field[%s] After[%s]\n", record.Fields.ArrayField[i].Name, after)
		}
	case avro.OperationUPDATE:
		for i := 0; i < len(record.AfterImages.ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject); i++ {
			before, err := getValue(record.BeforeImages.ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject[i])
			if err != nil {
				panic(err)
			}
			after, err := getValue(record.AfterImages.ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject[i])
			if err != nil {
				panic(err)
			}
			fmt.Printf("Field[%s] Before[%s] After[%s]\n", record.Fields.ArrayField[i].Name, before, after)
		}
	case avro.OperationDELETE:
		for i := 0; i < len(record.BeforeImages.ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject); i++ {
			before, err := getValue(record.BeforeImages.ArrayUnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject[i])
			if err != nil {
				panic(err)
			}
			fmt.Printf("Field[%s] Before[%s]\n", record.Fields.ArrayField[i].Name, before)
		}
	}
}

func getValue(in *avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject) (string, error) {
	if in == nil {
		return "null", nil
	}

	switch in.UnionType {
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumInteger:
		return in.Integer.Value, nil
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumCharacter:
		return string(in.Character.Value), nil
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumDecimal:
		return fmt.Sprintf("%s", in.Decimal.Value), nil
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumFloat:
		return fmt.Sprintf("%v", in.Float.Value), nil
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumTimestamp:
		return fmt.Sprintf("%v.%v", in.Timestamp.Timestamp, in.Timestamp.Millis), nil
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumDateTime:
		var year, month, day, hour, minute, second int32
		if in.DateTime.Year != nil {
			year = in.DateTime.Year.Int
		}
		if in.DateTime.Month != nil {
			month = in.DateTime.Month.Int
		}
		if in.DateTime.Day != nil {
			day = in.DateTime.Day.Int
		}
		if in.DateTime.Hour != nil {
			hour = in.DateTime.Hour.Int
		}
		if in.DateTime.Minute != nil {
			minute = in.DateTime.Minute.Int
		}
		if in.DateTime.Second != nil {
			second = in.DateTime.Second.Int
		}
		if year != 0 {
			return fmt.Sprintf("%v-%v-%v %v:%v:%v", year, month, day, hour, minute, second), nil
		}
		return fmt.Sprintf("%v:%v:%v", hour, minute, second), nil
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumTimestampWithTimeZone:
		var year, month, day, hour, minute, second int32
		if in.TimestampWithTimeZone.Value.Year != nil {
			year = in.TimestampWithTimeZone.Value.Year.Int
		}
		if in.TimestampWithTimeZone.Value.Month != nil {
			month = in.TimestampWithTimeZone.Value.Month.Int
		}
		if in.TimestampWithTimeZone.Value.Day != nil {
			day = in.TimestampWithTimeZone.Value.Day.Int
		}
		if in.TimestampWithTimeZone.Value.Hour != nil {
			hour = in.TimestampWithTimeZone.Value.Hour.Int
		}
		if in.TimestampWithTimeZone.Value.Minute != nil {
			minute = in.TimestampWithTimeZone.Value.Minute.Int
		}
		if in.TimestampWithTimeZone.Value.Second != nil {
			second = in.TimestampWithTimeZone.Value.Second.Int
		}
		if year != 0 {
			return fmt.Sprintf("%v-%v-%v %v:%v:%v+%v", year, month, day, hour, minute, second, in.TimestampWithTimeZone.Timezone), nil
		}
		return fmt.Sprintf("%v:%v:%v", hour, minute, second), nil

	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumBinaryGeometry:
		return getReadableGeometry(in.BinaryGeometry.Value)
		//return fmt.Sprintf("%s", string(in.BinaryGeometry.Value)), nil
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumTextGeometry:
		return getReadableGeometry([]byte(in.TextGeometry.Value))
		//return fmt.Sprintf("%s", in.TextGeometry.Value), nil
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumBinaryObject:
		return fmt.Sprintf("%s", string(in.BinaryObject.Value)), nil
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumTextObject:
		return fmt.Sprintf("%s", in.TextObject.Value), nil
	case avro.UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObjectTypeEnumEmptyObject:
		return fmt.Sprintf("%s", in.EmptyObject.String()), nil
	}
	return "", fmt.Errorf("invalid value for *UnionNullIntegerCharacterDecimalFloatTimestampDateTimeTimestampWithTimeZoneBinaryGeometryTextGeometryBinaryObjectTextObjectEmptyObject")
}

func getReadableGeometry(raw []byte) (string, error) {
	if len(raw) < 4 {
		return "", fmt.Errorf("not a valid geometry %v", raw)
	}
	geomObj, err := wkb.Unmarshal(raw[4:])
	if err != nil {
		return "", err
	}

	wktObj, err := wkt.Marshal(geomObj)
	if err != nil {
		return "", err
	}
	return wktObj, nil
}
