package tasks

import "testing"

func TestSleep(t *testing.T) {
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
		{"sleep", args{"1", `{"seconds":5}`}, `{"nums":[0,1,2,3,4]}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sleep(tt.args.jobID, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sleep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sleep() = %v, want %v", got, tt.want)
			}
		})
	}
}
