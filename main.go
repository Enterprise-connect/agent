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
	"flag"
	"os"
	"fmt"
	util "github.build.ge.com/212359746/wzutil"
	//remove the pprof dependency to simplify agent binary
	//_ "net/http/pprof"
)

const (
	EC_LOGO = `
           ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄                                            
          ▐░░░░░░░░░░░▌▐░░░░░░░░░░░
          ▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀   
          ▐░▌          ▐░▌            
          ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌            
          ▐░░░░░░░░░░░▌▐░▌            
          ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌            
          ▐░▌          ▐░▌            
          ▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄   
          ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌  
           ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  @Digital Connect 
`
	COPY_RIGHT = "Digital Connect,  @GE Corporate"
	ISSUE_TRACKER = "https://github.com/Enterprise-connect/ec-x-sdk/issues"
	TC_HEADER = "X-Thread-Connect"
	XCALR_URL = "https://x-thread-connect.run.pcs.aws-usw02-dev.ice.predix.io"
)

var (
	CFM = make(map[string]interface{})
	CERT_BASE64 string = ""

	//bootstrap the cli
	CLI_FLAGS = map[string]interface{}{
		"fup":[]interface{}{"_fu","",`Specify a file to upload to the server agent.`},
		"fdw":[]interface{}{"_fd","",`Specify a file to download from the client agent.`},
		"cfg":[]interface{}{"_cf","",`Specify the config file to launch the agent.`},
		"mod":[]interface{}{"_md","agent",`Specify the EC Agent Mode in "client", "server", or "gateway".`},
		"ver":[]interface{}{"_v", false, "Show EC Agent's version."},
		"cer":[]interface{}{"_rc", false, "Show EC Agent Cert."},
		"lpt":[]interface{}{"_lp","7990",`Specify the default EC port#.`},
		"gpt":[]interface{}{"_gt","8990",`Specify the gateway port# in fuse-mode. (gw:server|gw:client)`},
		//deprecated. pprof imported in main pkg
		//"911":[]interface{}{"_91",false,`Internal system profiling`},
		"apt":[]interface{}{"_ap","17990",`Specify the EC http endpoint port# if the agent when -api is set.`},
		"cps":[]interface{}{"_cp",0,`Specify the Websocket compression-ratio for agent's inter-communication purpose. E.g. [0-9]. "0" is no compression whereas "9" is the best. "-1" indicates system default. "-2" is HuffmanOnly. See RFC 1951 for more detail.`},
		"tid":[]interface{}{"_ts","",`Specify the Target EC Server Id if the "client" mode is set`},
		"hst":[]interface{}{"_gh","","Specify the EC Gateway URI. E.g. wss://<somedomain>:8989"},
		"sst":[]interface{}{"_sh","","Specify the EC Service URI. E.g. https://<service.of.predix.io>"},
		"pth":[]interface{}{"_ph","","Specify the directory to the certificate/key."},
		"dat":[]interface{}{"_da","","Specify the string to be encrypted."},
		"psp":[]interface{}{"_pw","","Specify the passphrase of the private key, combined with the flags (-dec -pth <dir-pvtkey> -psp <passphrase> -dat <encypted-data>) for one-step decryption."},
		"sgn":[]interface{}{"_sg",false,"Start a CA Cert-Signing process."},
		"enc":[]interface{}{"_ep",false,"Encrypt the string for validation."},
		"dec":[]interface{}{"_dp",false,"Decrypt the string for validation. can combine with the flags (-dec -pth <dir-pvtkey> -psp <passphrase> -dat <encypted-data>) for one-step decryption."},
		"rht":[]interface{}{"_rh","",`Specify the Resource Host if the "server" mode is set. E.g. <someip>, <somedomain>. value will be discard when TLS is specified.`},
		"rpt":[]interface{}{"_rp","0",`Specify the Resource Port# if the "server" mode is set. E.g. 8989, 38989`},
		"aid":[]interface{}{"_id","","Specify the agent Id assigned by the EC Service. You may find it in the Cloud Foundry VCAP_SERVICE"},
		"tkn":[]interface{}{"_tk","","Specify the OAuth Token. The token may expire depending on your OAuth provisioner. This flag is ignored if OAuth2 Auto-Refresh were set."},
		"pxy":[]interface{}{"_px","","Specify a local Proxy service. E.g. http://hello.world.com:8080"},
		"cid":[]interface{}{"_ci","","Specify the client Id to auto-refresh the OAuth2 token."},
		"csc":[]interface{}{"_cs","","Specify the client secret to auto-refresh the OAuth2 token."},
		"oa2":[]interface{}{"_oa","","Specify URL of the OAuth2 provisioner. E.g. https://<somedomain>/oauth/token"},
		"dur":[]interface{}{"_du",0,"Specify the duration for the next token refresh in seconds. (default 100 years)"},
		"crt":[]interface{}{"_ct","","Specify the relative path of a digital certificate to operate the EC agent. (.pfx, .cer, .p7s, .der, .pem, .crt)"},
		
		"wtl":[]interface{}{"_wl","0.0.0.0/0,::/0","Specify the ip(s) whitelist in the cidr net format. Concatenate ips by comma. E.g. 89.24.9.0/24, 7.6.0.0/16"},
		"bkl":[]interface{}{"_bl","","Specify the ip(s) blocklist in the IPv4/IPv6 format. Concatenate ips by comma. E.g. 10.20.30.5, 2002:4559:1FE2::4559:1FE2"},
		"plg":[]interface{}{"_pg",false,`Enable EC plugin list. This requires the plugins.yml file presented in the agent path.`},
		"inf":[]interface{}{"_if",false,"The Product Information."},
		"dbg":[]interface{}{"_dg",false,"Turn on debug mode. This will introduce more error information. E.g. connection error."},
	
		"vfy":[]interface{}{"_vf",false,"Verify the legitimacy of a digital certificate."},
		"rnw":[]interface{}{"_rn",false,"Renew a previous-issued x509 certificate."},
		"zon":[]interface{}{"_zn","",`Specify the Zone/Service Inst. Id. required in the "gateway" mode.`},
		"pks":[]interface{}{"_ps","","Specify the private key in base64 encoded format to decrypt the token from the agent. E.g. -pkg LS0tLS1CRUdJTiBSU0EgU.."},
		"pky":[]interface{}{"_pk","","Specify the relative path to a TLS key when operate as the gateway as desired. E.g. ./path/to/key.pem."},
		"pct":[]interface{}{"_pc","","Specify the relative path to a TLS cert when operate as the gateway as desired. E.g. ./path/to/cert.pem."},
		"hca":[]interface{}{"_hc","",`Specify a port# to turn on the Healthcheck API. This flag is always on when in the "gateway mode" with the provisioned local port. Upon provisioned, the api is available at <agent_uri>/health.`},
		"gen":[]interface{}{"_gc",false,"Generate a certificate request for the usage validation purpose."},
		"shc":[]interface{}{"_sc",false,"Health API requires basic authentication for Health APIs."},
		"vln":[]interface{}{"_vn",false,"Enable support for EC VLAN Network."},
		"grp":[]interface{}{"_gp","","GroupID needed for Agent Client/Server."},
		"gsg":[]interface{}{"_gs",false,"Generate the signature based on the given private key string(-pks), passphrase (-psp), and message (-dat) to be signed."},

		"vsg":[]interface{}{"_vs",false,"Verify the signature based on the given public key string (-pbk), passphrase, original message (-osg), and the signature (-dat), all are base64 encoded."},
		"pbk":[]interface{}{"_pb","","Base64 encoded certificate string."},
		"osg":[]interface{}{"_od","","Signature string in base64 encoded format."},
		"tse":[]interface{}{"_et",false,"Create a EC-compatible Token, with publickey (-pbk) and an optional 32-digits uuid (-dat). "},
		"tsd":[]interface{}{"_dt",false,"Check the timestamp of the EC token (-dat) "},
		"api":[]interface{}{"_ht",false,"Operate agent in HTTP mode."},
		"smp":[]interface{}{"_sm",false,"Simplifying the output for integration purpose."},
	}
)
		
func main() {	
	defer func(){
		if r:=recover();r!=nil{
			//util.PanicRecovery(r)
			fmt.Println(" [EC Agent] main pacakge exception:",r)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	util.InfoLog("loading application parameters..")
	//dynamically assign flags
	for k, v := range CLI_FLAGS {
		_v:=v.([]interface{})
		switch _v[1].(type) {
		case string:
			CFM[_v[0].(string)]=flag.String(k,_v[1].(string),_v[2].(string))
		case bool:
			CFM[_v[0].(string)]=flag.Bool(k,_v[1].(bool),_v[2].(string))
		case int:
			CFM[_v[0].(string)]=flag.Int(k,_v[1].(int),_v[2].(string))
		default:
			panic("flags "+k+" is not implemented.")

		}
	}
	
	flag.Parse()

	util.Branding("/.ec","ec-plugin","ec-config",TC_HEADER,"EC",EC_LOGO,COPY_RIGHT,XCALR_URL,ISSUE_TRACKER)

	if *CFM["_ht"].(*bool) {
		util.Init(*CFM["_md"].(*string),*CFM["_dg"].(*bool))
		InitAgentAPI(CFM,*CFM["_pb"].(*string))
		return
	}

	if *CFM["_cf"].(*string)!="" {
		ac:=InitAgentConfig(CFM)
		yc,err:=ac.LoadFromYAML(*CFM["_cf"].(*string))
		if err!=nil {
			panic(err)
		}
		CFM = yc
	}
	
	util.Init(*CFM["_md"].(*string),*CFM["_dg"].(*bool))
	agtOps,err:=NewAgentOps(CFM)
	if err!=nil{
		panic(err)
	}
	agtOps.Start()
}
