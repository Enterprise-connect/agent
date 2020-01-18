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
	"reflect"
	"testing"
)

func TestInitAgentConfig(t *testing.T) {
	type args struct {
		t map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want *AgentConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitAgentConfig(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitAgentConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentConfig_LoadFromYAML(t *testing.T) {
	type fields struct {
		Template map[string]interface{}
	}
	type args struct {
		fp string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgentConfig{
				Template: tt.fields.Template,
			}
			got, err := a.LoadFromYAML(tt.args.fp)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentConfig.LoadFromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentConfig.LoadFromYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentConfig_LoadFromJSON(t *testing.T) {
	type fields struct {
		Template map[string]interface{}
	}
	type args struct {
		j []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgentConfig{
				Template: tt.fields.Template,
			}
			got, err := a.LoadFromJSON(tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentConfig.LoadFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentConfig.LoadFromJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentConfig_LoadConfigMap(t *testing.T) {
	type fields struct {
		Template map[string]interface{}
	}
	type args struct {
		cfg interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgentConfig{
				Template: tt.fields.Template,
			}
			got, err := a.LoadConfigMap(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("AgentConfig.LoadConfigMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentConfig.LoadConfigMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
