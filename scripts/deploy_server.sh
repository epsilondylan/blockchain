#!/bin/bash go run
NUM_servers=$1

# Start Servers with ssh connection
for p in $(seq 8051 $((8050+NUM_servers))); do
   #ssh osgroup16@122.200.68.26 -i id_rsa -o StrictHostKeyChecking=no -p $p -t "mkdir -p /osdata/osgroup16" #在多台服务器上建立文件夹
   scp  -P $p -i ../id_rsa /home/fkd21/blockchain/miner/miner  osgroup16@122.200.68.26:/osdata/osgroup16/miner #传输文件到多台服务器上，只用执行一次
   scp  -P $p -i ../id_rsa /home/fkd21/blockchain/miner/peerlist.json  osgroup16@122.200.68.26:/osdata/osgroup16/peerlist.json
   scp  -P $p -i ../id_rsa /home/fkd21/blockchain/scripts/terminate_pid.sh  osgroup16@122.200.68.26:/osdata/osgroup16/terminate_pid.sh
done


# Wait for all servers to finish
wait


