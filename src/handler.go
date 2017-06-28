package main

import (
	"dao"
	"fmt"
	"net/http"
	"strconv"
	"unsafe"
	"util"
	"util/json"
	"util/log"

	"github.com/julienschmidt/httprouter"
)

var tokenMap map[string]int

func init() {
	tokenMap = make(map[string]int)
}

// ------------------------------------------------------------------------------
func token(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.H("/token  -------------------------------------------------")
	r.ParseForm()

	appkey := r.Form.Get("appkey")
	if appkey == "" {
		log.E("appkey missed")
		writeObj(w, http.StatusBadRequest, "")
		return
	}
	log.I("Receive appKey: ", appkey)

	username := r.Form.Get("username")
	if username == "" {
		log.E("username missed")
		writeObj(w, http.StatusBadRequest, "")
		return
	}
	log.I("Receive username: ", username)

	password := r.Form.Get("password")
	if password == "" {
		log.E("password missed")
		writeObj(w, http.StatusBadRequest, "")
		return
	}

	ak, e := strconv.Atoi(appkey)
	if e != nil {
		log.E("appKey is not number: ", appkey)
		writeObj(w, http.StatusBadRequest, "")
		return
	}

	if !dao.ValidateUsr(ak, username, password) {
		log.E("Validate failed for user: ", username)
		writeObj(w, http.StatusUnauthorized, "")
		return
	}
	token := util.GenToken()
	log.I("Generate token: ", token)
	tokenMap[token] = ak
	writeObj(w, http.StatusOK, token)
}

// ------------------------------------------------------------------------------
func validate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.H("/validate  -------------------------------------------------")
	appkey := r.Header.Get("appkey")
	if appkey == "" {
		log.E("appkey missed")
		writeObj(w, http.StatusBadRequest, false)
		return
	}
	log.I("Receive appkey: ", appkey)

	token := r.Header.Get("token")
	if token == "" {
		log.E("token missed")
		writeObj(w, http.StatusBadRequest, false)
		return
	}
	log.I("Receive token: ", token)

	appname := r.Header.Get("appname")
	if appname == "" {
		log.E("appname missed")
		writeObj(w, http.StatusBadRequest, false)
		return
	}
	log.I("Receive appname: ", appname)

	key, e := strconv.Atoi(appkey)
	if e != nil {
		log.E("appkey is not number: ", appkey)
		writeObj(w, http.StatusBadRequest, "")
	}
	result := checkToken(key, token, appname)
	log.I("Check result: ", result)
	writeObj(w, http.StatusOK, result)
}

// ------------------------------------------------------------------------------
func writeObj(w http.ResponseWriter, status int, obj interface{}) {
	log.I(status, " ", http.StatusText(status))
	w.WriteHeader(status)
	bytes := json.ToJSON(obj)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, token")
	w.Header().Set("content-type", "application/json")

	fmt.Fprintf(w, *(*string)(unsafe.Pointer(&bytes)))
}

func checkToken(appkey int, token, appname string) bool {
	key := tokenMap[token]
	if key == 0 {
		return false
	}
	// TODO
	// doa.Getappname(appkey)
	return true
}
