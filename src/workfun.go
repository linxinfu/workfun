package src

import (
	"encoding/json"
	"errors"
	"github.com/gen2brain/dlgs"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
)

type SweetMessage struct {
	Code      int      `json:"code"`
	Message   string   `json:"message"`
	ReturnObj []string `json:"returnObj"`
}

func StartMakeFun() {
	express()
	sweetTask()
}

func sweetTask() {
	c := cron.New()
	spec := "0 */10 * * * ?"
	err := c.AddFunc(spec, express)
	if err != nil {
		log.Println(err)
		return
	}
	c.Start()
	select {}
}

func express() {
	msg, err := getMessage()
	if err != nil {
		sweetDialog(err.Error())
	} else {
		sweetDialog(msg)
	}
}

func getMessage() (string, error) {
	var message SweetMessage
	
	url := "https://api.lovelive.tools/api/SweetNothings/Serialization/Json"
	method := "GET"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &message)
	if err != nil {
		return "", err
	}
	if len(message.ReturnObj) > 0 {
		return message.ReturnObj[0], nil
	} else {
		return "", errors.New("can not get the message")
	}
}

func sweetDialog(msg string) {
	_, _ = dlgs.Info("SWEET TIME", msg)
}
