/*
 * Copyright (c) 2016 General Electric Company. All rights reserved.
 *
 * The copyright to the computer software herein is the property of
 * General Electric Company. The software may be used and/or copied only
 * with the written permission of General Electric Company or in accordance
 * with the terms and conditions stipulated in the agreement/contract
 * under which the software has been supplied.
 *
 * author: apolo.yasuda@ge.com
 */

package main

import (
	//"fmt"
	//util "github.build.ge.com/212359746/wzutil"
	"sync"
	"errors"
	"io/ioutil"
	api "github.build.ge.com/212359746/wzapi"
	"net/http"
	config "github.build.ge.com/212359746/wzconf"
	util "github.build.ge.com/212359746/wzutil"
	"encoding/json"
)

var (
	tmplt map[string]interface{}
)
const (
	
)

func InitAgentAPI(t map[string]interface{}, crt string) {
	tmplt = t
	
	aapi:= &AgentAPI{
		AgentOpsList: &ConfLocker{
			m: make(map[string]interface{}),
		},
	}
	
	config.StartAgentAPI(aapi,*t["_ap"].(*string), crt)
}

type ConfLocker struct{
	sync.RWMutex
	m map[string]interface{}
}

type AgentAPI struct{
	AgentOpsList *ConfLocker
}

func (i *AgentAPI) Token(w http.ResponseWriter, r *http.Request, token, opsId string){
	
	//api.ErrResponse(w, 500, errors.New("internal error."), "internal error.")

	//create a placeholder for this operation forward
	i.AgentOpsList.Lock()
	i.AgentOpsList.m[opsId]="*"
	i.AgentOpsList.Unlock()
	
	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)
	
	w.Write([]byte(token))
	return
}

func (i *AgentAPI) POSTConfig(w http.ResponseWriter, r *http.Request, opsId string){

	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		}
	}()

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsId]
	i.AgentOpsList.RUnlock()
	if !ok{
		
		api.ErrResponse(w, 500, errors.New("operation unauthorised."), "operation unauthorised.")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ErrResponse(w, 500, errors.New("internal error"), err.Error())
		return
	}

	ac:=InitAgentConfig(tmplt)
	cfm,err:=ac.LoadFromJSON(body)
	if err!=nil{
		api.ErrResponse(w, 500,err, err.Error())		
		return
	}
	
	s:="verified"
	cfm["status"]=&s
	cfm["opsId"]=&opsId

	//cfg needs to be in flags format
	agtOps,err:=NewAgentOps(cfm)
	if err!=nil{
		api.ErrResponse(w, 500, err, err.Error())
		return
	}
	i.AgentOpsList.Lock()
	i.AgentOpsList.m[opsId]=agtOps
	i.AgentOpsList.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)	
	_cfg, _ := json.Marshal(agtOps.Config())
	w.Write(_cfg)
	return
}
		
func (i *AgentAPI) GETConfig(w http.ResponseWriter, r *http.Request, opsId string){
	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsId]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised."), "operation unauthorised.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	var agtOps *AgentOps
	i.AgentOpsList.RLock()
	agtOps,k:=i.AgentOpsList.m[opsId].(*AgentOps)
	if !k {
		api.ErrResponse(w, 500, errors.New("agent operation not found"), "operation not found")
		return
	}
	i.AgentOpsList.RUnlock()

	_cfg, _ := json.Marshal(agtOps.Config())
	w.Write(_cfg)
	return
}

func (i *AgentAPI) PUTConfig(w http.ResponseWriter, r *http.Request, opsId string){
	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsId]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised."), "operation unauthorised.")
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"updated"}`))
	return
}

func (i *AgentAPI) Hire(w http.ResponseWriter, r *http.Request, opsId string){

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsId]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised."), "operation unauthorised.")
		return
	}
	
	var agtOps *AgentOps
	i.AgentOpsList.RLock()
	agtOps,k:=i.AgentOpsList.m[opsId].(*AgentOps)
	if !k {
		api.ErrResponse(w, 500, errors.New("agent operation not found"), "operation not found")
		return
	}
	i.AgentOpsList.RUnlock()

	go agtOps.Start()
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"running"}`))
	return
}

func (i *AgentAPI) Resume(w http.ResponseWriter, r *http.Request, opsId string){

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsId]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised."), "operation unauthorised.")
		return
	}
	
/*	var agtOps *AgentOps
	i.RLock()
	agtOps,ok:=i.m[opsId].(*AgentOps)
	if !ok {
		api.ErrResponse(w, 500, err, err.Error())
		return
	}
	i.RUnlock()

*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"running"}`))
	return

}

func (i *AgentAPI) Suspend(w http.ResponseWriter, r *http.Request, opsId string){

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsId]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised."), "operation unauthorised.")
		return
	}
	
/*	var agtOps *AgentOps
	i.RLock()
	agtOps,ok:=i.m[opsId].(*AgentOps)
	if !ok {
		api.ErrResponse(w, 500, err, err.Error())
		return
	}
	i.RUnlock()
*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"stop"}`))
	return

}

func (i *AgentAPI) Status(w http.ResponseWriter, r *http.Request, opsId string){

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsId]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised."), "operation unauthorised.")
		return
	}
	
/*	var agtOps *AgentOps
	i.RLock()
	agtOps,ok:=i.m[opsId].(*AgentOps)
	if !ok {
		api.ErrResponse(w, 500, err, err.Error())
		return
	}
	i.RUnlock()
*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
	return

}

func (i *AgentAPI) Fire(w http.ResponseWriter, r *http.Request, opsId string){

	var ok bool
	i.AgentOpsList.RLock()
	_,ok = i.AgentOpsList.m[opsId]
	i.AgentOpsList.RUnlock()
	if !ok{
		api.ErrResponse(w, 500, errors.New("operation unauthorised."), "operation unauthorised.")
		return
	}

	var agtOps *AgentOps

	i.AgentOpsList.RLock()
	agtOps,k:=i.AgentOpsList.m[opsId].(*AgentOps)
	i.AgentOpsList.RUnlock()
	if !k {
		api.ErrResponse(w, 500, errors.New("agent is invalid."), "invalid")
		return
	}
	go agtOps.Stop()
	
	i.AgentOpsList.Lock()
	delete(i.AgentOpsList.m,opsId)
	i.AgentOpsList.Unlock()
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"removed"}`))
	return
}
