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
	"net/http"
	"testing"
)

func TestInitAgentAPI(t *testing.T) {
	type args struct {
		t   map[string]interface{}
		crt string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitAgentAPI(tt.args.t, tt.args.crt)
		})
	}
}

func TestAgentAPI_Token(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		token string
		opsId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
			}
			i.Token(tt.args.w, tt.args.r, tt.args.token, tt.args.opsId)
		})
	}
}

func TestAgentAPI_POSTConfig(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		opsId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
			}
			i.POSTConfig(tt.args.w, tt.args.r, tt.args.opsId)
		})
	}
}

func TestAgentAPI_GETConfig(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		opsId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
			}
			i.GETConfig(tt.args.w, tt.args.r, tt.args.opsId)
		})
	}
}

func TestAgentAPI_PUTConfig(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		opsId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
			}
			i.PUTConfig(tt.args.w, tt.args.r, tt.args.opsId)
		})
	}
}

func TestAgentAPI_Hire(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		opsId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
			}
			i.Hire(tt.args.w, tt.args.r, tt.args.opsId)
		})
	}
}

func TestAgentAPI_Resume(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		opsId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
			}
			i.Resume(tt.args.w, tt.args.r, tt.args.opsId)
		})
	}
}

func TestAgentAPI_Suspend(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		opsId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
			}
			i.Suspend(tt.args.w, tt.args.r, tt.args.opsId)
		})
	}
}

func TestAgentAPI_Status(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		opsId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
			}
			i.Status(tt.args.w, tt.args.r, tt.args.opsId)
		})
	}
}

func TestAgentAPI_Fire(t *testing.T) {
	type fields struct {
		AgentOpsList *ConfLocker
	}
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		opsId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AgentAPI{
				AgentOpsList: tt.fields.AgentOpsList,
			}
			i.Fire(tt.args.w, tt.args.r, tt.args.opsId)
		})
	}
}
