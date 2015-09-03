package main

import (
	// "encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dtoebe/3dcartrestapi/receiver"
	"io/ioutil"
	"os"
	// "strings"
)

/* Will work on Mac/Linux/BSD Not sure about Windows
* Put your 3dcartapiconf.json in your ~/Documents
* Uses $HOME env var to find the json file
**** EXAMPLE JSON: 3dcartapiconf.json ****
* {
*	"private_key": "Your apps Private Key",
*	"token": "Your Apps token"
* }
 */
//TODO: Add Documentation
var Config struct {
	PrivateKey string `json:"private_key"`
	Token      string `json:"token"`
}

var Data struct {
	Item []struct {
		data map[string]string
	}
}

func main() {

	service := flag.String("s", "Products", "The Type of results.  ie: Products, or Categories")
	filter := flag.String("c", "none", "If you want to filter a specific CatId")
	addData := flag.String("f", "none", "If you want to filter a spacific type of data.  ie: AdvancedOptions")
	// output := flag.String("o", "OUtput.csv", "The CSV output file. MUST HAVE .csv extention")
	storeUrl := flag.String("u", "https://yourstore.com", "The store's Secure URL. Has to start with https://")

	flag.Parse()
	confDir := os.Getenv("HOME") + "/Documents"

	confFile, err := ioutil.ReadFile(confDir + "/3dcartapiconf.json")
	if err != nil {
		fmt.Println("Err opening The JSON Config File: ", err)
		return
	}

	if err = json.Unmarshal([]byte(confFile), &Config); err != nil {
		fmt.Println("Malformed JSON file:", err)
	}

	res := receiver.NewConf(Config.PrivateKey, Config.Token, 1)

	var data []byte

	if *storeUrl == "https://yourstore.com" {
		fmt.Println("You need to provide your own https:// URL")
		return
	}

	if *filter == "none" {
		_, _, data = res.GetData(*service)
	} else if *filter != "none" && *addData == "none" {
		_, _, data = res.GetData(*storeUrl, *service, *filter)
	} else {
		_, _, data = res.GetData(*storeUrl, *service, *filter, *addData)
	}

	// csvOut, err := os.Create(os.Getenv("HOME") + "/Desktop/" + *output)
	// if err != nil {
	// 	fmt.Println("Err creating CSV File:", err)
	// 	return
	// }
	// defer csvOut.Close()
	//
	// filterData := strings.Replace(string(data), "[", "", -1)
	// filterData = strings.Replace(filterData, "]", "", -1)
	// filterData = strings.Replace(filterData, "{", "", -1)
	// filterData = strings.Replace(filterData, "}", "", -1)
	// filterData = strings.Replace(filterData, "\"SKUInfo\":", "", -1)
	// splitData := strings.Split(filterData, ",")
	//
	// w := csv.NewWriter(csvOut)
	//
	// for _, sd := range splitData {
	// 	var record []string
	// 	record = strings.Split(sd, ":")
	// 	w.Write(record)
	// }

	fmt.Println(string(data))
}
