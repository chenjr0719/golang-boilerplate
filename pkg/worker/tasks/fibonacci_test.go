package tasks

import "testing"

func TestFibonacci(t *testing.T) {
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
		{"fibonacci: 5", args{"1", `{"target":5}`}, "5", false},
		{"fibonacci: 10", args{"1", `{"target":10}`}, "55", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fibonacci(tt.args.jobID, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fibonacci() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
