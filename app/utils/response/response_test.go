package response

import (
	"encoding/json"
	"testing"
)

type Te struct {
	Tee string `json:"tee"`
	Pee string `json:"pee"`
}

func TestSuccess(t *testing.T) {
	t.Log("123")
	var ts []*Te
	ts = append(ts, &Te{"222", "333"})
	ts = append(ts, &Te{"444", "555"})
	resp := Resp{
		Code: "00",
		Msg:  "成功",
		Data: ts,
	}
	b, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}
