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

func Test_MatchTemplatesOk(t *testing.T) {
	path := []string{"api", "v1", "users", "2344"}
	template := []string{"api", "{version}", "users", "{userid}"}

	assertTrue(t, MatchTemplate(path, template))
}

func Test_MatchTemplatesNotOk(t *testing.T) {
	path := []string{"api", "v1", "users", "2344"}
	template := []string{"api", "{version}", "list", "{userid}"}

	assertFalse(t, MatchTemplate(path, template))
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
