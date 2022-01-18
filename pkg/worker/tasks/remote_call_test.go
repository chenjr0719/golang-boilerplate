package tasks

import (
	"encoding/json"
	"testing"
)

func TestRemoteHTTPCall(t *testing.T) {
	type args struct {
		jobID string
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"remoteCall: postman-echo", args{"1", `{"url": "https://httpbin.org/post", "body": {"message": "Hello World"}}`}, `{"args":{},"data":"{\"message\":\"Hello World\"}","files":{},"form":{},"json":{"message":"Hello World"},"url":"https://httpbin.org/post"}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoteHTTPCall(tt.args.jobID, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoteHTTPCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var responseBody map[string]interface{}
			json.Unmarshal([]byte(got), &responseBody)
			delete(responseBody, "headers")
			delete(responseBody, "origin")
			resp, err := json.Marshal(responseBody)
			if err != nil {
				t.Errorf("RemoteHTTPCall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			result := string(resp)
			if result != tt.want {
				t.Errorf("RemoteHTTPCall() = %v, want %v", result, tt.want)
			}
		})
	}
}
