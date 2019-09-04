package snow_test

// make sure to load right module
import (
	"server/store/snow"
	"testing"
)

func TestInit(t *testing.T) {
	snst, err := snow.Init("")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if isTest(snst) {
		t.Errorf("Expected snow.ServicenowStore, got different type")
	}
}

func isTest(t interface{}) bool {
	switch t.(type) {
	case snow.ServicenowStore:
		return true
	default:
		return false
	}
}

func TestGetList(t *testing.T) {
	//var err error

	// Success case - GetList should return expected json
	var err error
	testJson := `{
  "Name": "ServiceNowQuery",
  "Report": [
    {
      "number": "INC1234"
    }
  ]
}
`
	snst, err := snow.Init(testJson)

	expectedStr := `{"Name":"ServiceNowQuery","Report":[{"number":"INC1234","assigned_to":"","description":"","state":"","priority":"","severity":""}]}`
	js, err := snst.GetList()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	jsStr := string(*js)

	if jsStr != expectedStr {
		t.Errorf("Expected %s, got %s", expectedStr, jsStr)
	}

	// failure case - no json data
	snst.Content = `{
  "Name": "ServiceNowQuery",
  "Report": [
    {`
	js, err = snst.GetList()

	if err == nil {
		t.Errorf("Expected error, got %v", err)
	}
}
