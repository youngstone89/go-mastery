package singleton_test

import "testing"

// super type
type Singleton interface {
	AddOne() int
}

// sub type
type singleton struct {
	count int
}

// duck typing to Singleton interface
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

var instance *singleton

// get instance function that returns Singleton interface
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		//Test of acceptance criteria 1 failed
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}
	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}
	counter2 := GetInstance()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Expected same instance in counter2 but it got a different instance")
	}
}
