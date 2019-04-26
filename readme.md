# OIPWG/docker

Docker Compose configurations for basic FLO/OIP related daemons.


##Prerequisites
- Docker Compose - https://docs.docker.com/compose/install/


## Currently supported daemons
- Go-Flo Mainnet
- Go-Flo Testnet
- IPFS


### Go-Flo Mainnet
Modify `go-flo/mainnet.conf`

- Build image - necessary if config file is modified `docker-compose build goFloMainnet`
- Run daemon in background `docker-compose up -d goFloMainnet`
- Tail daemon logs `docker-compose logs -f goFloMainnet`



### Go-Flo Testnet
Modify `go-flo/testnet.conf`

- Build image - necessary if config file is modified `docker-compose build goFloTestnet`
- Run daemon in background `docker-compose up -d goFloTestnet`
- Tail daemon logs `docker-compose logs -f goFloTestnet`


### IPFS
A local directory at `./ipfsStaging` is available and mounted as `/export` within the IPFS container

- Build image - necessary if config file is modified `docker-compose build ipfs`
- Run daemon in background `docker-compose up -d ipfs`
- Tail daemon logs `docker-compose logs -f ipfs`

