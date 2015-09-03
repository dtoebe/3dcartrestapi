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
	PrivateKey string
	AppToken   string
	APIVersion int
}

//APIURL const: Base url to connect ot 3dCart's Restful API
const APIURL = "https://apirest.3dcart.com"

//NewConf Initializes the Conf struct returns that struct
func NewConf(key string, token string, ver int) *Conf {
	return &Conf{key, token, ver}
}

//GetData Takes the initialized Conf struct and which service you to call
//	Sevice: Example "Products" to receive all products from your store
//The output:
//	Response Status (string)
//	Response Header (http.Header)
//	Response Body	([]byte)
func (conf Conf) GetData(storeUrl string, services ...string) (string, http.Header, []byte) {

	url := APIURL + "/3dCartWebAPI/v" + strconv.Itoa(conf.APIVersion)
	for _, service := range services {
		url += "/" + service
	}

	return resData(url, storeUrl, conf.PrivateKey, conf.AppToken)
}

//GetSpecificData was made with getting /Products/{CatID} in mind.
//	Right now it takes the service (example: Products), and a filter (example CatId) then outputs
//		Status
//		Header
//		body ([]byte)

func resData(url string, secureURL, privKey, token string) (string, http.Header, []byte) {
	req, err := http.NewRequest("GET", url, nil)
	errOutF("Error with GET:", err)

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("SecureUrl", secureURL)
	req.Header.Set("PrivateKey", privKey)
	req.Header.Set("Token", token)

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
