package main

import "fmt"

const PassStatus = "pass"
const FailStatus = "fail"

type HealthCheck struct {
	ServiceID string
	status    string
}

func GenerateCheck() []HealthCheck {
	Slice := make([]HealthCheck, 5, 5)
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			Slice = append(Slice, HealthCheck{
				ServiceID: string(rune(i)),
				status:    PassStatus,
			})
		}
		Slice = append(Slice, HealthCheck{
			ServiceID: string(rune(i)),
			status:    FailStatus,
		})
	}
	return Slice
}

func main() {
	result := GenerateCheck()
	for _, slice := range result {
		if slice.status == PassStatus {
			fmt.Println(slice.status)
		}
	}
}
