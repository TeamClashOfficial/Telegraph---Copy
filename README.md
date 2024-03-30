# TSS Multi Party Server

## How to build new binary and send to servers
1. `cd multiPartyServer`
2. `cd scripts`
3. `bash servers.sh` Note: This command will build a new binary and send to all of the ec2 servers. 
4. If you get this error `scp: /home/ubuntu/tss/multiPartyServer: Text file busy`, then stop the running multiPartyServer binary on all servers and run the command `bash servers.sh` again.
## How to start binary on ec2 servers
1. `cd multiPartyServer/scripts`
2. Start terminator on your terminal and open four new tabs. And execute these commands in each of the tab. 
```
    ssh -i tss-key.pem ubuntu@18.229.137.38
    ssh -i tss-key.pem ubuntu@18.228.192.203
    ssh -i tss-key.pem ubuntu@54.233.170.242
    ssh -i tss-key.pem ubuntu@18.231.112.8
```
3. `cd tss` in each of the servers.
4. Make sure there is a "config.json" file in tss directory and the configuration are valid for all servers.
5. Run `./multiPartServer` in each of the servers.
6. If you get this error `-bash: ./multiPartyServer: Text file busy` then delete multiPartServer and send it again using `bash servers.sh` from your development computer

## Send keygen request to ec2 servers
1. `cd multiPartyServer/scripts`
2. `./keygen.sh` Note: This command will start keygen process on ec2 servers

## Send keysign request to ec2 servers
1. `cd multiPartyServer/scripts`
2. `./keysign.sh` Note: This command will start sign process on ec2 servers

## Updating log level:
    The log level can be updated in the config.json file as following:
        ErrorLevel  -   2 
        WarnLevel   -   3
        InfoLevel   -   4
        DebugLevel  -   5

## Generate go code from abi

### Install packages:
    sudo npm i -g solc@0.4.26
    sudo add-apt-repository -y ppa:ethereum/ethereum
    sudo apt-get update
    sudo apt-get install ethereum -y 
### Generate go code
	/usr/local/bin/solcjs --abi Store.sol
	/usr/local/bin/solcjs --bin Store.sol
	abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go