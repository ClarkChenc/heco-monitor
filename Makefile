BINARY="heco-monitor"
DOCKER_IMAGE="clarkchenc/heco-monitor:v0.0.3"

abi:
	abigen --abi=abi/erc20/erc20.abi --pkg=erc20 --out=abi/erc20/erc20.go

build: abi	
	go build -o ${BINARY}

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

image: build
	docker build -t ${DOCKER_IMAGE} .

.PHONY: contracts build