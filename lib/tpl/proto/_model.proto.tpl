syntax = "proto3";

package {{.TableName}};
{{- $hasDateTime := true }}{{range .Fields}}{{if $hasDateTime}}{{if eq .DataType "datetime"}}
import public "google/protobuf/timestamp.proto";
{{- $hasDateTime = false }}{{end}}{{end}}{{end}}
option go_package = "proto;proto";

message {{.ModelName}} {
{{range $i, $v := .Fields}}
{{if eq $v.EXTRA "auto_increment"}}
    // @inject_tag: json:"{{$v.ColumnName}}" gorm:"primary_key;AUTO_INCREMENT;default:{{$v.ColumnComment}};comment:{{$v.ColumnComment}};{{- if $v.DataTypeLong -}}size:{{$v.DataTypeLong}};{{- end -}}"{{else}}
    // @inject_tag: json:"{{$v.ColumnName}}" gorm:"default:{{$v.ColumnComment}};comment:{{$v.ColumnComment}};{{- if $v.DataTypeLong -}}size:{{$v.DataTypeLong}};{{- end -}}"{{end}}{{if eq $v.DataType "bigint"}}
   int64 {{$v.ColumnName}} = {{add $i}};{{else if eq $v.DataType "datetime"}}
   google.protobuf.Timestamp {{$v.ColumnName}} = {{add $i}};{{else}}
   string {{$v.ColumnName}} = {{add $i}};{{end}}
{{end}}

}
