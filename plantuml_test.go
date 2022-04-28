package plantuml

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const txt_service = "https://www.plantuml.com/plantuml/txt/"
const raw = `@startuml
Bob -> Alice : hello
@enduml`

const ascii = `
     ┌───┐          ┌─────┐
     │Bob│          │Alice│
     └─┬─┘          └──┬──┘
       │    hello      │   
       │──────────────>│   
     ┌─┴─┐          ┌──┴──┐
     │Bob│          │Alice│
     └───┘          └─────┘`

func TestEncode(t *testing.T) {
	encoded := Encode(raw)
	resp, err := http.Get(txt_service + encoded)
	if err != nil {
		t.Fatal("unexpected http error : ", err)
	} else {
		if resp.StatusCode != http.StatusOK {
			t.Fatal("unexpected status : ", resp.Status)
		}
		recv, _ := ioutil.ReadAll(resp.Body)
		recv_ascii := string(recv)
		t.Logf("\n @@@ RECV ASCII @@@ \n%s", recv_ascii)
		expt_ascii := ascii
		// remove all the ws/lf/lb
		recv_ascii = strings.Replace(recv_ascii, " ", "", -1)
		recv_ascii = strings.Replace(recv_ascii, "\n", "", -1)
		recv_ascii = strings.Replace(recv_ascii, "\r", "", -1)
		recv_ascii = strings.Replace(recv_ascii, "\t", "", -1)
		// remove all the ws/lf/lb
		expt_ascii = strings.Replace(expt_ascii, " ", "", -1)
		expt_ascii = strings.Replace(expt_ascii, "\n", "", -1)
		expt_ascii = strings.Replace(expt_ascii, "\r", "", -1)
		expt_ascii = strings.Replace(expt_ascii, "\t", "", -1)
		if recv_ascii != expt_ascii {
			t.Log("recv.Trim(\" \", \"\\t\", \"\\r\", \"\\n\") : ", recv_ascii)
			t.Log("expt.Trim(\" \", \"\\t\", \"\\r\", \"\\n\") : ", expt_ascii)
			t.Fatal("ASCII mismatch : ", len(recv_ascii), len(expt_ascii))
		}
	}
}

func TestDecode(t *testing.T) {
	// only works when
	encoded := Encode(raw)
	decoded, err := Decode(encoded)
	if err != nil {
		t.Error("unexpected decode error : ", err)
	} else {
		if decoded != raw {
			t.Logf("\n @@@ EXPECT @@@ \n%s", raw)
			t.Logf("\n @@@ DECODE @@@ \n%s", decoded)
			t.Error("mismatch with raw")
		}
	}
}
