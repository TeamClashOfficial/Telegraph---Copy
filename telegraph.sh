#!/bin/bash

# Set the name of your binary
binary=telegraph

# Set the URL where the binary is hosted
URL=https://telegraph-binary.s3.us-east-2.amazonaws.com/telegraph

# Install and verify Golang
if ! which go &> /dev/null; then
    echo "Installing go..."
    
    # Remove previous versions of go
    sudo rm -rf /usr/local/go

    # Install the required go version using snap
    sudo snap install go --channel=1.19/stable --classic

    # Add go binary to the system PATH
    echo 'export PATH=$PATH:/snap/bin' >> ~/.bashrc

    # Make Go available for the current script
    export PATH=$PATH:/snap/bin
fi

# Verify go installation
go_version=$(go version | awk -F\  '{print $3}' | cut -d 'o' -f2)
if [ -z "$go_version" ]; then
    echo "Go installation failed or not installed."
    exit 1
else
    echo "Go installation successful. Version: $go_version"
fi

# Install GCC
echo "Installing GCC..."
sudo apt update
sudo apt install build-essential -y

# Check if the GCC installation was successful
if [ $? -eq 0 ]; then
    echo "GCC installation completed."
else
    echo "GCC installation failed."
    exit 1
fi

# Check if MongoDB is installed
if ! which mongod &> /dev/null; then
    echo "Installing MongoDB..."

    # Update the apt packages list
    sudo apt update
    sudo apt upgrade

    # Install gpg2 for validating mongodb packages
    sudo apt install gnupg2 

    # Download the mongodb GPG key
    wget -nc https://www.mongodb.org/static/pgp/server-6.0.asc 

    # Add the key to your apt keyring
    cat server-6.0.asc | gpg --dearmor | sudo tee /etc/apt/keyrings/mongodb.gpg >/dev/null 

    # Add the mongodb repo to your apt sources list
    sudo sh -c 'echo "deb [ arch=amd64,arm64 signed-by=/etc/apt/keyrings/mongodb.gpg] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/6.0 multiverse" >> /etc/apt/sources.list.d/mongo.list'

    # Update the apt packages list again
    sudo apt update 

    # Install mongodb
    sudo apt install mongodb-org 

    # Start the mongodb service
    sudo systemctl start mongod
fi

# Verify MongoDB installation
if ! pgrep mongod > /dev/null; then
    echo "Error: MongoDB is not running!"
    exit 1
else
    echo "MongoDB installation successful."
fi

if ! which telegraph; then
	# Download the binary file from the URL
	echo "Downloading binary .."
	wget -q $URL -O $binary

	# Make the binary file executable
	chmod +x $binary

	# Create a directory for your command line tool
	sudo mkdir -p /usr/local/$binary

	# Copy the binary file to the directory you just created
	sudo cp $binary /usr/local/$binary/

	# Create a symlink to the binary file in /usr/local/bin
	sudo ln -sf /usr/local/$binary/$binary /usr/local/bin/$binary

	env_file=$HOME/.env
	[[ ! -f $env_file ]] &&
		cat <<- EOF > $env_file
			CHAIN_ID=1
			CHAIN_TYPE="EVM"
			CONTRACT_ADDRESS="0x27E66e0ea5Abf34d412e21FDADB83B32080A255e"
			DB_URL="mongodb://localhost:27017/telegraph"
			EVM_HTTP_URL=""
			EVM_WSS_URL=""
			FINISHED_SETUP="true"
			GENESIS_IP=""
			HSM_CKA_ID=""
			HSM_CKA_LABEL=""
			HSM_CONFIG_PATH="/etc/softhsm/softhsm2.conf"
			HSM_PATH="/usr/lib/softhsm/libsofthsm2.so"
			HSM_PIN=123456
			HSM_TOKEN_LABEL="telegraph"
			HSM_USE_GCM_IV="false"
			ID="ID"
			IP="0.0.0.0"
			IS_GENESIS="true"
			KEY="314159265358979323846264338327950288419716939937510582097494459"
			LOG_LEVEL=5
			MAIN_PORT_ADDRESS=""
			MAX_CONN=0
			MONIKER="Moniker String"
			NAME=""
			PARTY_PASSWORD="password string"
			PORT=7044
			PUBLIC_KEY=""
			THRESHOLD=0
			TLS_CERT=""
			TLS_KEY=""
		EOF

	service_file=/etc/systemd/system/$binary.service
	[[ ! -f $service_file ]] &&
		cat <<- EOF | sudo tee $service_file &> /dev/null
			[Unit]
			Description=Telegraph Daemon

			[Service]
			Type=simple
			WorkingDirectory=$HOME
			ExecStart=/usr/local/bin/telegraph
			ExecReload=/bin/kill -HUP $MAINPID
			Restart=always
			StandardOutput=syslog
			StandardError=syslog
			SyslogIdentifier=Telegraph
			User=root
			Group=root

			[Install]
			WantedBy=multi-user.target
		EOF

	service_running=$(systemctl status telegraph.service |
		awk '/Active:/ { print $3 ~ "running" }')

	sudo systemctl enable telegraph.service
	((service_running)) || sudo systemctl start telegraph.service
fi

[ -f "$binary" ] && rm $binary