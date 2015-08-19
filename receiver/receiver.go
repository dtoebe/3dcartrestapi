package receiver

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//Conf struct:	This sets the config variables;
//	SecureURL:	Your store's HTTPS URL (Find: Settings -> General -> Store Settings);
//	PrivateKey:	Your apps private key (Find: https://devportal.3dcart.com/);
//	AppToken:	Token given when your store has authorized your app (Find: https://devportal.3dcart.com/);
//	APIVersion: 	The version of the Rest API (Currentaly: 1);
type Conf struct {
	SecureURL  string
	PrivateKey string
	AppToken   string
	APIVersion int
}

//APIURL const: Base url to connect ot 3dCart's Restful API
const APIURL = "https://apirest.3dcart.com"

//NewConf Initializes the Conf struct returns that struct
func NewConf(secure string, key string, token string, ver int) *Conf {
	return &Conf{secure, key, token, ver}
}

//GetData Takes the initialized Conf struct and which service you to call
//	Sevice: Example "Products" to receive all products from your store
//The output:
//	Response Status (string)
//	Response Header (http.Header)
//	Response Body	([]byte)
func (conf Conf) GetData(service string) (string, http.Header, []byte) {
	url := APIURL + "/3dCartWebAPI/v" + strconv.Itoa(conf.APIVersion) + "/" + service

	req, err := http.NewRequest("GET", url, nil)
	errOutF("Error with GET:", err)

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("SecureUrl", conf.SecureURL)
	req.Header.Set("PrivateKey", conf.PrivateKey)
	req.Header.Set("Token", conf.AppToken)

	client := &http.Client{}
	res, err := client.Do(req)
	errOutF("Error Sending Request:", err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	errOutF("Error reading response:", err)

	return res.Status, res.Header, body
}

func errOutF(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
