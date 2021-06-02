package asciicast

import (
	"encoding/json"
	"fmt"

	// json "github.com/pquerna/ffjson/ffjson"

)

type Frame struct {
	Delay float64
	InOut string
	Data  []byte
}

func (f *Frame) String() string {
	s, _ := json.Marshal(string(f.Data))
	return fmt.Sprintf(`[%.6f, "%v", %s]`, f.Delay, f.InOut, s)
}

func (f *Frame) MarshalJSON() ([]byte, error) {
	s, _ := json.Marshal(string(f.Data))
	jsonData := fmt.Sprintf(`[%.6f, "%v", %s]`, f.Delay, f.InOut, s)
	// fmt.Println("jsonData: ", jsonData)
	return []byte(jsonData), nil
}

func (f *Frame) UnmarshalJSON(data []byte) error {
	var x interface{}

	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}

	f.Delay = x.([]interface{})[0].(float64)
	f.InOut = x.([]interface{})[1].(string)
	s := []byte(x.([]interface{})[2].(string))
	b := make([]byte, len(s))
	copy(b, s)
	f.Data = b

	return nil
}
