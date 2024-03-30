#!/bin/bash

mkdir -p /etc/apt/keyrings 2>/dev/null

wget -qO- https://download.01.org/intel-sgx/sgx_repo/ubuntu/intel-sgx-deb.key | tee /etc/apt/keyrings/intel-sgx-keyring.asc >/dev/null

echo "deb [signed-by=/etc/apt/keyrings/intel-sgx-keyring.asc arch=amd64] https://download.01.org/intel-sgx/sgx_repo/ubuntu $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/intel-sgx.list

apt update

EGO_DEB=ego_1.4.0_amd64_ubuntu-$(lsb_release -rs).deb

wget https://github.com/edgelesssys/ego/releases/download/v1.4.0/"${EGO_DEB}"

apt install ./"${EGO_DEB}"
