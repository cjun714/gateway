package dao

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"util/json"
	"util/log"
)

// ValidateUsr is used to check user is a leagal user or not
func ValidateUsr(appkey int, username, password string) bool {
	url := "http://localhost:8090/check" + "?appkey=" + strconv.Itoa(appkey) + "&username=" + username + "&password=" + password
	log.H(url)
	client := http.Client{}
	request, e := http.NewRequest("GET", url, nil)
	if e != nil {
		log.E(e)
		return false
	}
	resp, e := client.Do(request)
	if e != nil {
		log.E(e)
		return false
	}
	defer resp.Body.Close()

	log.I("resp: ", resp.Status)

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.E(e)
		return false
	}

	var result bool
	json.ToObj(body, &result)
	log.H(result)

	return result
}

// QueryApp is used to get application name
func QueryApp(appkey int) string {
	url := "http://localhost:8090/query-app/" + strconv.Itoa(appkey)
	log.H(url)
	client := http.Client{}
	request, e := http.NewRequest("GET", url, nil)
	if e != nil {
		log.E(e)
		return ""
	}
	resp, e := client.Do(request)
	if e != nil {
		log.E(e)
		return ""
	}
	defer resp.Body.Close()

	log.I("resp: ", resp.Status)

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.E(e)
		return ""
	}

	result := string(body)
	log.H(result)

	return result
}
