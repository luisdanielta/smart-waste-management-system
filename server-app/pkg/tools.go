package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	tp "swms-sa/types"
)

func ValidateRequiredFields(r *http.Request, fields []string) bool {
	for _, field := range fields {
		if r.FormValue(field) == "" {
			return false
		}
	}
	return true
}

func Fetch(url string) (chan interface{}, bool) {
	var err tp.Error_t
	var resp *http.Response
	ch := make(chan interface{})

	go func() {
		defer close(ch)
		resp, err.ERR = http.Get(url)
		err.Check()
		if err.ERR != nil {
			ch <- err.ERR
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			err.ERR = errors.New("status code is not 200")
			ch <- err.ERR
			err.Check()
			return
		}

		var body []byte
		body, err.ERR = io.ReadAll(resp.Body)
		err.Check()

		var data map[string]interface{}
		err.ERR = json.Unmarshal(body, &data)
		err.Check()

		if err.ERR != nil {
			fmt.Println("Error unmarshalling response:", err.ERR)
			fmt.Println("Response content:", string(body))
			ch <- err.ERR
			return
		}

		ch <- data
	}()

	if <-ch == nil {
		return nil, false
	}

	return ch, true

}

/* used example of Fetch function */
/*
func main() {
	resp := Fetch("https://www.google.com")
	response := <-resp

	if response == nil {
		println("response is nil")
		return
	}

	fmt.Println(resp)
}
*/
