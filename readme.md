# OIPWG/docker

Docker Compose configurations for basic FLO/OIP related daemons.


## Prerequisites
- Docker Compose - https://docs.docker.com/compose/install/


## Currently supported daemons
- Go-Flo Mainnet
- Go-Flo Testnet
- OIP Web Wallet
- OIP Mainnet
- OIP Testnet
- Elasticsearch
- Kibana
- Caddy
- IPFS

# Common


### Caddy
Caddy provides a reverse proxy to the potential running services herein  
Configuration/Customization may take place by modifying `caddy/Caddyfile` and adding plugins to `caddy/with_plugins.go`

By default Caddy will provide http:// on port 80 listening on all interfaces

```
http://<ip>/                    Web wallet
http://<ip>/mainnet/oip         OIP daemon apis running on mainnet
http://<ip>/mainnet/kibana      Kibana instance running on mainnet
                                user:pass = kibana:mainnet
http://<ip>/testnet/oip         OIP daemon apis running on testnet
http://<ip>/testnet/kibana      Kibana instance running on testnet
                                user:pass = kibana:testnet
http://<ip>/ipfs/               IPFS Gateway
```

- Build image - necessary if Caddyfile or with_plugins.go are modified `docker-compose build caddy`
- Run daemon in background `docker-compose up -d caddy`
- Tail daemon logs `docker-compose logs -f caddy`



### IPFS
A local directory at `./ipfsStaging` is available and mounted as `/export` within the IPFS container

- Build image - necessary if config file is modified `docker-compose build ipfs`
- Run daemon in background `docker-compose up -d ipfs`
- Tail daemon logs `docker-compose logs -f ipfs`


### Web Wallet
Web wallet provides a multi currency web accessible wallet

- Run daemon in background `docker-compose up -d webwallet`
- Tail daemon logs `docker-compose logs -f webwallet`



# Mainnet


### Go-Flo
Modify `go-flo/mainnet.conf`

- Build image - necessary if config file is modified `docker-compose build goflomainnet`
- Run daemon in background `docker-compose up -d goflomainnet`
- Tail daemon logs `docker-compose logs -f goflomainnet`


### OIP
Modify `oip/config.mainnet.yml`

- Build image - necessary if config file is modified `docker-compose build oipmainnet`
- Run daemon in background `docker-compose up -d oipmainnet`
- Tail daemon logs `docker-compose logs -f oipmainnet`


### Elasticsearch
Elasticsearch mainnet defaults to 4GB Heap size, adjust `ES_JAVA_OPTS` as appropriate

- Run daemon in background `docker-compose up -d esmainnet`
- Tail daemon logs `docker-compose logs -f esmainnet`


### Kibana

- Run daemon in background `docker-compose up -d kibanamainnet`
- Tail daemon logs `docker-compose logs -f kibanamainnet`



# Testnet


### Go-Flo
Modify `go-flo/testnet.conf`

- Build image - necessary if config file is modified `docker-compose build goflotestnet`
- Run daemon in background `docker-compose up -d goflotestnet`
- Tail daemon logs `docker-compose logs -f goflotestnet`


### OIP
Modify `oip/config.testnet.yml`

- Build image - necessary if config file is modified `docker-compose build oiptestnet`
- Run daemon in background `docker-compose up -d oiptestnet`
- Tail daemon logs `docker-compose logs -f oiptestnet`


### Elasticsearch
Elasticsearch testnet defaults to 2GB Heap size, adjust `ES_JAVA_OPTS` as appropriate

- Run daemon in background `docker-compose up -d estestnet`
- Tail daemon logs `docker-compose logs -f estestnet`


### Kibana

- Run daemon in background `docker-compose up -d kibanatestnet`
- Tail daemon logs `docker-compose logs -f kibanatestnet`

