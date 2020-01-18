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
	"fmt"
	util "github.build.ge.com/212359746/wzutil"
	"reflect"
	"gopkg.in/yaml.v2"
	"encoding/json"
	"os"
)

var ()

const ()

type AgentConfig struct{
	Template map[string]interface{}
}
        
func InitAgentConfig(t map[string]interface{}) *AgentConfig {

	return &AgentConfig{
		Template:t,
	}

}

func (a *AgentConfig) LoadFromYAML(fp string) (map[string]interface{},error){

	f,err:=util.ReadFile(fp)
	if err!=nil{
		return nil,err
	}

	t:=make(map[string]interface{})
	
	err = yaml.Unmarshal(f, &t)
        if err != nil {
		return nil,err
        }

	_t:=t[util.HEADER_CONFIG()].(map[interface{}]interface{})
	if _t["conf"]==nil{
		return nil,errors.New("invalid config header."+util.HEADER_CONFIG())
	}
	
	return a.LoadConfigMap(_t["conf"])
}

func (a *AgentConfig) LoadFromJSON(j []byte) (map[string]interface{},error){

	var cfg map[string]interface{}
	if err := json.Unmarshal(j, &cfg); err != nil {
		return nil,err
	}
	
	return a.LoadConfigMap(cfg)
}

func (a *AgentConfig) LoadConfigMap(cfg interface{}) (map[string]interface{},error) {

	//clone template config
	cfm:=make(map[string]interface{})
	for k,v := range a.Template {
		cfm[k] = v
	}

	var op = func(_k,v interface{}, cfm map[string]interface{}) error{

		k:=_k.(string)
		i,ok:=CLI_FLAGS[k]
		if !ok {
			return errors.New("unsupported flag:"+k)
		}
		
		_,ok=cfm[i.([]interface{})[0].(string)]
		if !ok {
			return errors.New("unsupported flag:"+k)
		}

		switch v.(type) {
		case string:
			_v0:=v.(string)
			_v := os.ExpandEnv(_v0)
			cfm[i.([]interface{})[0].(string)]=&_v
		case bool:
			_v:=v.(bool)
			cfm[i.([]interface{})[0].(string)]=&_v
		case int:
			_v:=v.(int)
			cfm[i.([]interface{})[0].(string)]=&_v
		case float64:
			_v:=int(v.(float64))
			cfm[i.([]interface{})[0].(string)]=&_v
		default:
			op:=fmt.Sprintf("v:type: %s",reflect.TypeOf(v))
			return errors.New("flag:"+k+" has unexpected value:"+op)

		}
		
		return nil
	}
	
	switch cfg.(type){
	case map[string]interface{}:
		
		for _k, v := range cfg.(map[string]interface{}) {
			if err:=op(_k,v,cfm);err!=nil{
				return nil,err
			}
			
		}
		
	case map[interface{}]interface{}:
		
		for _k, v := range cfg.(map[interface{}]interface{}) {
			if err:=op(_k,v,cfm);err!=nil{
				return nil,err
			}
		}
		
	default:
		return nil,errors.New("mistyped configuration set.")		
	}
	
	return cfm, nil
}
