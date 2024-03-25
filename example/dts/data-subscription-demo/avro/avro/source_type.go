// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     record.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type SourceType int32

const (
	SourceTypeMySQL      SourceType = 0
	SourceTypeOracle     SourceType = 1
	SourceTypeSQLServer  SourceType = 2
	SourceTypePostgreSQL SourceType = 3
	SourceTypeMongoDB    SourceType = 4
	SourceTypeRedis      SourceType = 5
	SourceTypeDB2        SourceType = 6
	SourceTypePPAS       SourceType = 7
	SourceTypeDRDS       SourceType = 8
	SourceTypeHBASE      SourceType = 9
	SourceTypeHDFS       SourceType = 10
	SourceTypeFILE       SourceType = 11
	SourceTypeOTHER      SourceType = 12
)

func (e SourceType) String() string {
	switch e {
	case SourceTypeMySQL:
		return "MySQL"
	case SourceTypeOracle:
		return "Oracle"
	case SourceTypeSQLServer:
		return "SQLServer"
	case SourceTypePostgreSQL:
		return "PostgreSQL"
	case SourceTypeMongoDB:
		return "MongoDB"
	case SourceTypeRedis:
		return "Redis"
	case SourceTypeDB2:
		return "DB2"
	case SourceTypePPAS:
		return "PPAS"
	case SourceTypeDRDS:
		return "DRDS"
	case SourceTypeHBASE:
		return "HBASE"
	case SourceTypeHDFS:
		return "HDFS"
	case SourceTypeFILE:
		return "FILE"
	case SourceTypeOTHER:
		return "OTHER"
	}
	return "unknown"
}

func writeSourceType(r SourceType, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewSourceTypeValue(raw string) (r SourceType, err error) {
	switch raw {
	case "MySQL":
		return SourceTypeMySQL, nil
	case "Oracle":
		return SourceTypeOracle, nil
	case "SQLServer":
		return SourceTypeSQLServer, nil
	case "PostgreSQL":
		return SourceTypePostgreSQL, nil
	case "MongoDB":
		return SourceTypeMongoDB, nil
	case "Redis":
		return SourceTypeRedis, nil
	case "DB2":
		return SourceTypeDB2, nil
	case "PPAS":
		return SourceTypePPAS, nil
	case "DRDS":
		return SourceTypeDRDS, nil
	case "HBASE":
		return SourceTypeHBASE, nil
	case "HDFS":
		return SourceTypeHDFS, nil
	case "FILE":
		return SourceTypeFILE, nil
	case "OTHER":
		return SourceTypeOTHER, nil
	}

	return -1, fmt.Errorf("invalid value for SourceType: '%s'", raw)

}

func (b SourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *SourceType) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewSourceTypeValue(stringVal)
	*b = val
	return err
}

type SourceTypeWrapper struct {
	Target *SourceType
}

func (b SourceTypeWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b SourceTypeWrapper) SetInt(v int32) {
	*(b.Target) = SourceType(v)
}

func (b SourceTypeWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b SourceTypeWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b SourceTypeWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b SourceTypeWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b SourceTypeWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b SourceTypeWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b SourceTypeWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b SourceTypeWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b SourceTypeWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b SourceTypeWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b SourceTypeWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b SourceTypeWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b SourceTypeWrapper) Finalize() {}