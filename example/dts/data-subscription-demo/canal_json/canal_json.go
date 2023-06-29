package canal_json

type FlatMessage struct {
	/* batch_id */
	Id       int64  `json:"id"`
	Database string `json:"database"`
	Table    string `json:"table"`
	/* 组成 primary key 的所有列名 */
	PkNames []string `json:"pkNames"`
	IsDdl   bool     `json:"isDdl"`
	/* 数据变更类型 */
	Type string `json:"type"`
	/* binlog 中的毫秒级时间戳,即数据原始变更的时间 */
	Es int64 `json:"es"`
	/* DTS 生成该条消息的毫秒级时间戳 */
	Ts int64 `json:"ts"`
	/* 当 isDdl 为 true 时，记录对应的 DDL 语句 */
	Sql string `json:"sql"`
	/* 当 isDdl 为 false 时，记录每一列数据类型在 Java 的 JDBC 中的类型表示 */
	SqlType map[string]int32 `json:"sqlType"`
	/* 当 isDdl 为 false 时，MySQL 中每一列数据类型的描述*/
	MysqlType map[string]string `json:"mysqlType"`
	/* DML 修改后的数据，包含每一个表结构字段的kv结构 */
	Data []map[string]string `json:"data"`
	/* 仅 UPDATE 操作时不为null，表示 UPDATE 语句变更的列，即变更前的列值 */
	Old []map[string]string `json:"old"`

	/* 新增字段，当 isDdl 为 false 时，PG 中每一列数据类型的描述*/
	PGType map[string]string `json:"pgType"`
}

//*数据库类型*
type Type int32

const (
	Type_TYPECOMPATIBLEPROTO2 Type = 0
	Type_ORACLE               Type = 1
	Type_MYSQL                Type = 2
	Type_PGSQL                Type = 3
)

// Java JDBC 中的 sql type code
type SqlType int32

const (
	BIT           SqlType = -7
	TINYINT       SqlType = -6
	SMALLINT      SqlType = 5
	INTEGER       SqlType = 4
	BIGINT        SqlType = -5
	FLOAT         SqlType = 6
	REAL          SqlType = 7
	DOUBLE        SqlType = 8
	NUMERIC       SqlType = 2
	DECIMAL       SqlType = 3
	CHAR          SqlType = 1
	VARCHAR       SqlType = 12
	LONGVARCHAR   SqlType = -1
	DATE          SqlType = 91
	TIME          SqlType = 92
	TIMESTAMP     SqlType = 93
	BINARY        SqlType = -2
	VARBINARY     SqlType = -3
	LONGVARBINARY SqlType = -4
	NULL          SqlType = 0
	OTHER         SqlType = 1111
	JAVA_OBJECT   SqlType = 2000
	DISTINCT      SqlType = 2001
	STRUCT        SqlType = 2002
	ARRAY         SqlType = 2003
	BLOB          SqlType = 2004
	CLOB          SqlType = 2005
	REF           SqlType = 2006
	DATALINK      SqlType = 70
	BOOLEAN       SqlType = 16
	/*-------------JDBC 4.0 -----------------*/
	ROWID        SqlType = -8
	NCHAR        SqlType = -15
	NVARCHAR     SqlType = -9
	LONGNVARCHAR SqlType = -16
	NCLOB        SqlType = 2011
	SQLXML       SqlType = 2009
	/*-------------JDBC 4.2 -----------------*/
	REF_CURSOR              SqlType = 2012
	TIME_WITH_TIMEZONE      SqlType = 2013
	TIMESTAMP_WITH_TIMEZONE SqlType = 2014
)
