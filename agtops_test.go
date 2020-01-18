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

	config "github.build.ge.com/212359746/wzconf"
	core "github.build.ge.com/212359746/wzcore"
)

func TestNewAgentOps(t *testing.T) {
	type args struct {
		c map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *AgentOps
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAgentOps(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAgentOps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAgentOps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAgentOps_Start(t *testing.T) {
	type fields struct {
		Agent core.AgentIntr
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgentOps{
				Agent: tt.fields.Agent,
			}
			a.Start()
		})
	}
}

func TestAgentOps_Stop(t *testing.T) {
	type fields struct {
		Agent core.AgentIntr
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgentOps{
				Agent: tt.fields.Agent,
			}
			a.Stop()
		})
	}
}

func TestAgentOps_Config(t *testing.T) {
	type fields struct {
		Agent core.AgentIntr
	}
	tests := []struct {
		name   string
		fields fields
		want   *config.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AgentOps{
				Agent: tt.fields.Agent,
			}
			if got := a.Config(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AgentOps.Config() = %v, want %v", got, tt.want)
			}
		})
	}
}
