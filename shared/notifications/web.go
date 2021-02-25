package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"code.jtg.tools/ayush.singhal/notifications-microservice/configuration"
	"github.com/pkg/errors"
)

// Web struct implements Notifications interface
type Web struct {
	To    string
	Title string
	Body  string
}

// SendNotification function sends a web notification to the specified deviceToken given the server key and title, body of the notification
func (web *Web) SendNotification() error {
	url := configuration.GetResp().WebNotification.URL

	var jsonStr = []byte(fmt.Sprintf(`{"notification": {
		"title": "%s", 
		"body": "%s"
		},
		"to" : "%s"}`, web.Title, web.Body, web.To))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", fmt.Sprintf("key=%s", configuration.GetResp().WebNotification.ServerKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Declared an empty map interface
	var result map[string]interface{}

	if resp.Status != "200 OK" {
		return errors.New("Non 200 status received")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "Error Reading Response")
	}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(body, &result)

	if result["success"] != 1.0 {
		return errors.New("Notification Sending failed")
	}
	return nil
}