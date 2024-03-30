ips="
18.229.137.38
18.228.192.203
54.233.170.242
18.231.112.8
"
echo "Building"
cd ../
go build
cd scripts

for ip in $ips; do
    echo $ip
    scp -i "tss-key.pem" ../telegraph ubuntu@$ip:/home/ubuntu/tss &
    pids[${i}]=$!
done

# wait for all pids
for pid in ${pids[*]}; do
    wait $pid
done

sleep 100


# ssh -i "tss-key.pem" ubuntu@18.229.137.38
# ssh -i "tss-key.pem" ubuntu@18.228.192.203
# ssh -i "tss-key.pem" ubuntu@54.233.170.242
# ssh -i "tss-key.pem" ubuntu@18.231.112.8