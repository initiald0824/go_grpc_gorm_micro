syntax = "proto3";

package {{.TableName}};

import "{{.TableName}}_model.proto";
import "common.proto";
import "google/api/annotations.proto";

option go_package = "proto;proto";

service {{.ModelName}}Service {
	rpc Create({{.ModelName}}) returns (common.Response) {
		option (google.api.http) = {
			post:"/v1/{{.RouterName}}/create"
			body:"*"
		};
	}
	
	rpc Delete({{.ModelName}}) returns (common.Response) {
		option (google.api.http) = {
			post:"/v1/{{.RouterName}}/delete"
			body:"*"
		};
	}

	rpc DeleteById({{.ModelName}}) returns (common.Response) {
		option (google.api.http) = {
			post:"/v1/{{.RouterName}}/deleteById"
			body:"*"
		};
	}

	rpc Update({{.ModelName}}) returns (common.Response) {
		option (google.api.http) = {
			post:"/v1/{{.RouterName}}/update"
			body:"*"
		};
	}

	rpc Find({{.ModelName}}) returns (common.Response) {
		option (google.api.http) = {
			post:"/v1/{{.RouterName}}/find"
			body:"*"
		};
	}

	rpc Lists(common.Request) returns (common.Responses) {
		option (google.api.http) = {
			post:"/v1/{{.RouterName}}/lists"
			body:"*"
		};
	}
}
