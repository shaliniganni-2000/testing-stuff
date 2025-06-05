package agent

import "testing"

func TestFetchConstant(t *testing.T) {
	c := &Client{}
	res, err := c.FetchConstant("test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Name != "GeneratedConst" || res.Value != "generated" {
		t.Fatalf("unexpected constant: %#v", res)
	}
}
