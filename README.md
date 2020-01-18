# build.Agent

GoDoc | SDK | Benchmark
--- | --- | ---
[Production](https://ec-sdk-doc-hokkaido.run.aws-usw02-pr.ice.predix.io) | [v1.hokkaido.203](https://github.com/Enterprise-connect/ec-x-sdk/releases/tag/v1.hokkaido.203) |  N/A
[Beta](https://ec-sdk-doc-fukuoka.run.aws-usw02-dev.ice.predix.io) | [v1beta.fukuoka.1667](https://github.com/Enterprise-connect/ec-x-sdk/releases/tag/v1beta.fukuoka.1667) | N/A

## Design
The agent inherits both the Four-way|Two-way bi-directional connectivity pattern from [the original design](https://github.com/Enterprise-connect/ec-sdk/wiki/EC-Agent#usage-scenario-one--four-way-bi-directional-virtual-connectivity). Additionally, the following flows documented several key mechanism adopted in the agent connectivity-
* [Service and Agents Superconn connectivity flow](https://github.build.ge.com/Enterprise-Connect/ec-service#ec-service-and-agents-superconn-connectivity-flow)
## Build
### Build up an agent with the agent-buildpack docker image

This might be the easiest and quickiest way to build an agent binary without any dependencies.
```shell
#clone this repo
git clone https://github.build.ge.com/Enterprise-Connect/ec-agent.git

#go to the repo path
cd ./ec-agent

#download the sdk and get the latest library. Get latest <agent-release-lib> E.g. v1beta.fukuoka.1668.lib, v1.hokkaido.204.lib.
curl https://codeload.github.com/Enterprise-connect/ec-x-sdk/zip/<agent-release-lib> -lOk

#unzip the donwloaded sdk 
unzip <path/to/release/zip>

#replace the /pkg folder with the go packages, which can be found now from the sdk
cp -r </path/to/sdk/lib/go/pkg> ./pkg

#pull the beta agent builder
docker pull dtr.predix.io/dig-digiconnect/ec-agent-builder:v1beta 

#for docker proxy environment please refer to https://docs.docker.com/network/proxy/
#to build an agent binary. This step will generate an agent artifact based on the Makefile, which can be found in the repo
docker run --network host -v $(pwd):/build -e HTTPS_PROXY=${HTTPS_PROXY} -e NO_PROXY=${NO_PROXY} -e DIND_PATH=./ -i --name <container_name> dtr.predix.io/dig-digiconnect/ec-agent-builder:v1beta

docker run -v $(pwd):/build -i --name <container> <ec-agent-builder>:v1beta

#check the agent binary revisioin
./agent -ver
```

By default, the Makefile does not specify the output OS/Environment. Please refer to [go offical build flags](https://golang.org/cmd/go/). For the convenience, developers may refer to the [ec-sdk-buildpack](https://github.build.ge.com/Enterprise-Connect/ec-sdk-buildpack) for the build detail.

### Access the Digital-Foundry Jenkins

To conveniently access these pre-configured Jenkins jobs as a build Admin, please refer to [these steps to get a Digital-Foundry Jenkins instance](https://github.build.ge.com/Enterprise-Connect/ec-sdk-buildpack/blob/beta/README.md#if-you-are-granted-access-to-the-digi-digiconnect-org-dtrpredixio-you-may-use-the-digital-foundry-jenkins-to-access-all-builds-) 


### Build an agent binary from your environment
 - clone this repo.
 - copy both folders src/pkg in the repo to your ${GOPATH}/
 - build the artifact-
```shellscript
$ go version # this to ensure your go environment/version is compatible with linker from the EC library in the sdk.
$ cd /path/to/the/repo
$ go build
```

## Available Artifact Types (In EC-SDK)
`ecagent_<os>_<dns-resolving>`
For instance, if my machine ran a Linux distribution with a configurable dns resolver, I should use the artifact `ecagent_linux_var`

## Usage (outdated. please verify in the main.go in this repo. standard man file in progress)
```shellscript
$ ./ecagent -h
Usage of ./ecagent:
...
``` 

## Configuration (Example only, users may need to modify the usage accordingly)
### Launch agent for Connectivity via API endpoints
```script
./ecagent_linux_var -api -pbk <certificate generated by the agent ```-gen```>
```
 Valid endpoints are as follows-
 * /<agent_revision: v1beta|v1>/token (GET, receive a EC Agent-specific cipher key)
 * /<agent_revision: v1beta|v1>/<operation#>/config (POST, configure the agent)
 * /<agent_revision: v1beta|v1>/<operation#>/status (POST)
 * /<agent_revision: v1beta|v1>/<operation#>/hire (POST)
 * /<agent_revision: v1beta|v1>/<operation#>/fire (POST)
 * /<agent_revision: v1beta|v1>/<operation#>/suspend (POST)
 * /<agent_revision: v1beta|v1>/<operation#>/resume (POST)
 
 The ```-gen``` flag will generate a CSR (certificate-signing request) for your API operations. You may choose to generate your self-signed cert via agent, however the self-signed certificate will not be valid for the agent 
 API operations, unless your cipher key-pair is an authorised CA issuer previously issued by EC team. Please contact EC/TC team for the digital-signing detail. Or you may submit your CSR direcrlty via [xcaler](https://x-thread-connect.run.pcs.aws-usw02-pr.ice.predix.io/v1/swagger-ui.html). 
### Launch the agent via the yaml file. This will help orginise a complex agent detail. Here shows some example below.
```yaml
ec-config:
  conf:
    mod: gateway
    gpt: "17990"
    zon: <Cloud-foundry Service/Zone Id>
    grp: 012019-int
    sst: https://e27fc834-28be-4851-9d6a-b7033d568270.run.aws-usw02-dev.ice.predix.io
    dbg: true
    tkn: <Admin token from the Cloud Foundry VCAP>
    hst: http://localhost:17990
  cred:
    cert: <Cert-Issued-By-TC/EC team>
    developerId: 6a6154f2-d9d9-4096-bad6-69832d34a3242d
    privateKey: <Private key generated by the agent -gen>
    passPhrase: mypass
```
  You may request the developer's credential (Optional, listed as the ```cred``` field) with EC/TC team. Or again simply submit/create the CSR via [xcaler](https://x-thread-connect.run.pcs.aws-usw02-pr.ice.predix.io/v1/swagger-ui.html). Credential will send to you once it's proxied signed by the team's admin.

### Launch the agent in command-line
#### Server
```script
./ecagent_linux_var -mod server -aid <server-id> -hst <gateway-url>/agent /
-rht <amazon-postgres-url, schema excluded> -rpt <amazon-postgres-port> /
-cid <uaa-client-id> -csc <uaa-client-secret> -oa2 <uaa-url>/oauth/token /
-dur 300 -dbg -hca ${PORT} -zon <ec-zone-id> -sst <ec-cf-service-url>
```
#### Gateway
```script
./ecagent_linux_var -mod gateway -lpt ${PORT} -zon <ec-zone-id> /
-sst <ec-cf-service-url> -dbg -tkn <ec-admin-token-from-vcap>
```
#### Client
```script
./ecagent_linux_var -mod client -aid <client-id> -hst <gateway-url>/agent /
-lpt 7990 -tid <server-id> -oa2 <uaa-url>/oauth/token -cid <uaa-client-id> /
-csc <uaa-client-secret> -dur 300 -dbg -pxy <local-proxy, needed from Jenkins to CF>
```

#### Gateway-Client FuseMode
The connectivity between ```<server> <gw-client>``` depends on a open-access fusemode agent. In this case, the agent which declares in gw-client mode will need to expose the ```-gpt``` to public. Here is the usage example
```shell
agent \
-mod gw:client \
-lpt 8991 <port available to client TCP access>\
-gpt <the always-gateway-occupied port, should be always made available to public for the agent communication> \
-zon <predix zone id in vcap> \
-grp <EC group id. The predix zone id (-zon) added by default> \
-sst <EC service instance URL available in VCAP> \
-dbg \
-tkn <EC admin token in predix VCAP> \
-aid <agent id> \
-tid <target agent id> \
-pxy <optional: http://PITC-Zscaler-AmericasZ.proxy.corporate.ge.com:80>
```

#### Gateway-Server FuseMode
Like gw-client mode, the connectivity between ```<client> <gw-server>``` depends on a open-access fusemode agent. In the case of gw-server mode, the agent setting ```-gpt``` will expose to public. Here is the usage example
```shell
agent \
-mod gw:server \
-gpt <the always-gateway-occupied port, should be always made available to public for the agent communication> \
-zon <predix zone id in vcap> \
-grp <EC group id. The predix zone id (-zon) added by default> \
-sst <EC service instance URL available in VCAP> \
-dbg \
-tkn <EC admin token in predix VCAP> \
-pxy <optional> \
-rpt <TCP resource host port# to where the traffic will be forwarded> \
-rht <TCP resource host ip (or resolvable host name> to where the traffic will be forwarded> \
-plg <force to load EC plugin(s) via the plugin.yml>
```
### Cryptography API
Although EC agent has always been developer-centric, yet it does not compromise at security. The agent distibution is now embedded with several cipher/decipher tools. With the agent's security helper utility, developers are empowered to operate/own a CA and become an certificate issuer by will. The agent basic cryptography command as follows (```-smp``` for simplifying the output for machine-readable )

```shell
#generate pkcs15 signature hash
./ecagent -gsg -pks <base64_private_key> -psp <passphrase_of_the_private_key> -dat <base64_string>

#validate the signature hash
./ecagent -vsg -pbk <base64_cert> -osg <base64_original_string> -dat <base64_signature_hash>

#pkcs15 encrypt
./ecagent -enc -pbk <base64_cert> -dat <plaintext_string>

#pkcs15 decrypt
./ecagent -dec -pks <base64_private_key> -psp <passphrase_of_the_private_key> -dat <base64_cipher>

#Generate an agent-specific time-sensative token
./ecagent -tse -pbk <base64_cert> <-dat optional for further verification>

#Decipher the agent-specific time-sensative token from -tse
./ecagent -tsd -dat <base64 decoded cipher_generated_from -tse>

```
[Thread-Connect Subscription API-series](https://github.build.ge.com/Thread-Connect/tc-subscription) has implemented the EC's Cryptography API for its purpose of the authentication. See [in swagger](https://x-thread-connect.run.pcs.aws-usw02-pr.ice.predix.io/v1/swagger-ui.html)

### Enterprise-Connect CA Issuer's API
An step-through instruction and example will display when execute the following commands-
```
#generate a x509 CSR and a private key pair. In this usage, agent will offer an option as a final step to submit the CSR via xcalr on behalf of the developer
./ecagent -gen

#create a x509 certificate. This function requries the issuer's certificate and private key pair
./ecagent -sgn

#validate the authenticity of the inquiry certificate
./ecagent -vfy

#renew a x509 certificate. To renew, the developer must provision a certificate previously issued by the agent api/xcalr
./ecagent -rnw
```
You may aquire a CSR and/or a certificate|credential via the [xcaler. See its swagger for detail](https://x-thread-connect.run.pcs.aws-usw02-pr.ice.predix.io/v1/swagger-ui.html#/)

Certificates issued via xcalr is preset to 180 days. To extend the valid usage, you may request a renewal, which is currently available via the agent only.

### File Transfer between Agents
Beginning v1.hokkaido.203 | v1beta.fukuoka.1665, agent is capable of transferring files amongst other compatible agents. The usage is simple and straight forward
```
#upload a file
./ecagent -mod <client|gw:client> <some flags...> -fup \
<path/to/source/file>:<path/to/target/file>

#download a file
./ecagent -mod <client|gw:client> <some flags...> -fdw \
<path/to/source/file/in/server/agent>:<path/to/target/file/in/localhost>
```
When launching a client|gw:client agent, the agent will attempt to complete the file transfers prior to ready for further TCP requests.
