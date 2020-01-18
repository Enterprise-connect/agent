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
	"errors"
	api "github.build.ge.com/212359746/wzapi"
	config "github.build.ge.com/212359746/wzconf"
	util "github.build.ge.com/212359746/wzutil"
	core "github.build.ge.com/212359746/wzcore"
)

type AgentOps struct {
	//cli-compatible vars
	//Config  map[string]interface{}
	Agent  core.AgentIntr
}

func NewAgentOps(c map[string]interface{}) (*AgentOps,error) {

	conf,err:=config.Init(c)
	if err!=nil{
		return nil,err
	}

	switch conf.ClientType {
	case "server":	
		if err:=conf.InitServerConfig(c);err!=nil{
			return nil,err
		}
		
	case "client":

		if err:=conf.InitClientConfig(c);err!=nil{
			return nil,err
		}
		
	case "gateway":
		if err:=conf.InitGatewayConfig(c);err!=nil{
			return nil,err
		}
			
	case "gw:client":
		if err:=conf.InitGWClientConfig(c);err!=nil{
			return nil,err
		}

	case "gw:server":
		if err:=conf.InitGWServerConfig(c);err!=nil{
			return nil,err
		}
	default:
		return nil,errors.New("Unknown agent mode.")
	}

	var _agt core.AgentIntr
	switch conf.ClientType {
	case "server":
		_agt=core.NewServer(conf,api.HealthCheckSrv)
	case "client":
		_agt=core.NewClient(conf,api.HealthCheckClt)
	case "gateway":
		_agt=core.NewGateway(conf,api.HealthCheckGwy)
	case "gw:client":
		//health api schem stays the same to the gateway
		_agt=core.NewGWClient(conf,api.HealthCheckGwy)
	case "gw:server":
		_agt=core.NewGWServer(conf,api.HealthCheckGWServer)
	default:
		return nil,errors.New("no agent mode specified.")
	}

	return &AgentOps{
		Agent: _agt,
	},nil
}

func (a *AgentOps) Start() {
	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		}
	}()
	
	a.Agent.Hire()
	
	core.Operation(a.Agent)
}

func (a *AgentOps) Stop() {
	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		}
	}()

	a.Agent.Fire()
}

func (a *AgentOps) Config()(*config.Config){
	return a.Agent.Config()
}
