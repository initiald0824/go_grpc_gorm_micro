syntax = "proto3";

package {{.TableName}};
{{- $hasDateTime := true }}{{range .Fields}}{{if $hasDateTime}}{{if eq .DataType "timestamp"}}
import public "google/protobuf/timestamp.proto";
{{- $hasDateTime = false }}{{end}}{{end}}{{end}}
option go_package = "proto;proto";

message {{.ModelName}} {
{{range $i, $v := .Fields}}
{{if eq $v.EXTRA "auto_increment"}}
    // @inject_tag: json:"{{case2CamelAndLcfirst $v.ColumnName}}" gorm:"primary_key;AUTO_INCREMENT;default:{{$v.ColumnComment}};comment:{{$v.ColumnComment}};{{- if $v.DataTypeLong -}}size:{{$v.DataTypeLong}};{{- end -}}"{{else}}
    // @inject_tag: json:"{{case2CamelAndLcfirst $v.ColumnName}}" gorm:"default:{{$v.ColumnComment}};comment:{{$v.ColumnComment}};{{- if $v.DataTypeLong -}}size:{{$v.DataTypeLong}};{{- end -}}"{{end}}{{if eq $v.DataType "bigint"}}
   int64 {{case2CamelAndLcfirst $v.ColumnName}} = {{add $i}};{{else if eq $v.DataType "tinyint"}}
   int32 {{case2CamelAndLcfirst $v.ColumnName}} = {{add $i}};{{else if eq $v.DataType "int"}}
   int32 {{case2CamelAndLcfirst $v.ColumnName}} = {{add $i}};{{else if eq $v.DataType "timestamp"}}
   google.protobuf.Timestamp {{case2CamelAndLcfirst $v.ColumnName}} = {{add $i}};{{else}}
   string {{case2CamelAndLcfirst $v.ColumnName}} = {{add $i}};{{end}}
{{end}}

}
