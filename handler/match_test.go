package handler

import (
	"testing"
)

func TestMatch(t *testing.T) {

	assertTrue(t, Match("/", "/"))
	assertTrue(t, Match("/account/1", "account/{id}"))
	assertTrue(t, Match("/account/1", "/account/{id}/"))
	assertTrue(t, Match("/user/143/account/333/address/333", "/user/{uid}/account/{aid}/address/{adid}"))

	assertFalse(t, Match("/user/1", "/user/2"))
	assertFalse(t, Match("/user/1", "/user/1/account"))
	assertFalse(t, Match("/user/1/account/3", "/user/{uid}/account/{aid}/address"))
}

func assertTrue(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func assertFalse(t *testing.T, b bool) {
	if b {
		t.Fail()
	}
}
