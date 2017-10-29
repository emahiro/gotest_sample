package test

import "testing"

func TestPublicMethod(t *testing.T) {
	t.Logf("test public method","")
	PublicMethod()
}

func TestPrivateMethod(t *testing.T){
	t.Logf("test private method","")
	privateMethod()
}