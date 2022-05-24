package simplemath

import "testing"

/*
测试方法
 */
func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1,2), failed. Got %d, expected 3.", r)
	}
}
