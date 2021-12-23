package programmers

import "testing"

func TestRunningTruck(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name          string
		bridegeLength int
		weight        int
		truckWeights  []int
		want          int
	}{
		// TODO: Add test cases.
		{"프로그래머스 지나가는 트럭 테스트 해보기", 2, 10, []int{7, 4, 5, 6}, 8},
		{"프로그래머스 지나가는 트럭 테스트 해보기", 100, 100, []int{10}, 101},
		{"프로그래머스 지나가는 트럭 테스트 해보기", 100, 100, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, 110},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunningTruck(tt.bridegeLength, tt.weight, tt.truckWeights); got != tt.want {
				t.Errorf("RunningTruck() = %v, want %v", got, tt.want)
			}
		})
	}
}
