package tpl

var TypeMssqlDict = map[string]string{
	"smallint":       "int16",
	"bigint":         "int64",
	"nvarchar":       "string",
	"datetimeoffset": " *time.Time",
}
