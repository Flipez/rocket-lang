package object_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/flipez/rocket-lang/object"
)

func TestHTTPObjectMethods(t *testing.T) {
	tests := []inputTestCase{
		{`HTTP.nope()`, "undefined method `.nope()` for HTTP"},
		{`HTTP.handle(1, "test")`, "wrong argument type on position 0: got=INTEGER, want=STRING"},
		{`HTTP.handle("/", "test")`, "wrong argument type on position 1: got=STRING, want=FUNCTION"},
		{`HTTP.listen(3000)`, "Invalid handler. Call only supported on instance."},
		{`def test(){};HTTP.handle("/", test)`, "Invalid handler. Call only supported on instance."},
		{`a = HTTP.new(); a.listen(-1)`, "listening on port -1: listen tcp: address -1: invalid port"},
		{`a = HTTP.new(); a.listen(80)`, "listening on port 80: listen tcp :80: bind: permission denied"},
		{"HTTP.new().to_json()", "HTTP is not serializable"},
	}

	testInput(t, tests)
}
func TestHTTPType(t *testing.T) {
	tests := []inputTestCase{
		{"HTTP", "HTTP"},
	}

	for _, tt := range tests {
		def := testEval(tt.input).(*object.HTTP)
		defInspect := def.Inspect()

		if defInspect != tt.expected {
			t.Errorf("wrong string. expected=%#v, got=%#v", tt.expected, defInspect)
		}
	}
}

func TestHTTPServerMethods(t *testing.T) {
	httpServer := `def test(){response["body"] = "test"};
def test_a(){response = "test"};
a = HTTP.new();
a.handle("/", test);
a.handle("/test2", test_a);
a.listen(3000)`

	go testEval(httpServer)
	time.Sleep(100 * time.Millisecond) // workaround to give testIntput time to evaluate the input and start the http handle

	tests := []inputTestCase{
		{"/", "test"},
		{"/test2", ""},
	}

	client := &http.Client{}

	for _, tt := range tests {
		var data = strings.NewReader("servus")
		req, err := http.NewRequest("POST", "http://127.0.0.1:3000"+tt.input, data)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if string(body) != tt.expected {
			t.Errorf("wrong string. expected=%#v, got=%#v", tt.expected, string(body))
		}
	}
}
