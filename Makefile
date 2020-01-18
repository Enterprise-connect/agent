#
#  Copyright (c) 2016 General Electric Company. All rights reserved.
#
#  The copyright to the computer software herein is the property of
#  General Electric Company. The software may be used and/or copied only
#  with the written permission of General Electric Company or in accordance
#  with the terms and conditions stipulated in the agreement/contract
#  under which the software has been supplied.
#
#  author: apolo.yasuda@ge.com
#

ecagent=agent

.DEFAULT_GOAL: $(ecagent)

$(ecagent): agent-build

pre-install:
	@ls -la
agent-build:
	@echo test..
	@go test -vet=off
	@echo Creating artifact..
	@go build -o ./agent .

.PHONY: install
install:
	ls -al
