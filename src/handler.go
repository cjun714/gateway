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

	appKey := r.Form.Get("appKey")
	if appKey == "" {
		log.E("appKey missed")
		writeObj(w, http.StatusBadRequest, "")
		return
	}
	log.I("Receive appKey: ", appKey)

	usrname := r.Form.Get("usrname")
	if usrname == "" {
		log.E("usrname missed")
		writeObj(w, http.StatusBadRequest, "")
		return
	}
	log.I("Receive usrname: ", usrname)

	passwd := r.Form.Get("passwd")
	if passwd == "" {
		log.E("passwd missed")
		writeObj(w, http.StatusBadRequest, "")
		return
	}

	ak, e := strconv.Atoi(appKey)
	if e != nil {
		log.E("appKey is not number: ", appKey)
		writeObj(w, http.StatusBadRequest, "")
		return
	}

	if !dao.ValidateUsr(ak, usrname, passwd) {
		log.E("Validate failed for user: ", usrname)
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
	appKey := r.Header.Get("appKey")
	if appKey == "" {
		log.E("appKey missed")
		writeObj(w, http.StatusBadRequest, false)
		return
	}
	log.I("Receive appKey: ", appKey)

	token := r.Header.Get("token")
	if token == "" {
		log.E("token missed")
		writeObj(w, http.StatusBadRequest, false)
		return
	}
	log.I("Receive token: ", token)

	appName := r.Header.Get("appName")
	if appName == "" {
		log.E("appName missed")
		writeObj(w, http.StatusBadRequest, false)
		return
	}
	log.I("Receive appName: ", appName)

	key, e := strconv.Atoi(appKey)
	if e != nil {
		log.E("appKey is not number: ", appKey)
		writeObj(w, http.StatusBadRequest, "")
	}
	result := checkToken(key, token, appName)
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

func checkToken(appKey int, token, appName string) bool {
	key := tokenMap[token]
	if key == 0 {
		return false
	}
	// TODO
	// doa.GetAppName(appKey)
	return true
}
