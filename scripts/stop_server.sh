#!/bin/bash

NUM_servers=$1

# Stop Servers with ssh connection
for p in $(seq 8051 $((8050+NUM_servers))); do
   ssh osgroup16@122.200.68.26 -i ../id_rsa -o StrictHostKeyChecking=no -p $p -t "cd /osdata/osgroup16 && chmod +x terminate_pid.sh && rm -rf /osdata/osgroup16/miner.out && rm -rf /osdata/osgroup16/miner && rm -rf /osdata/osgroup16/utxoset.json && rm -rf /osdata/osgroup16/blockchain.json  && ./terminate_pid.sh" &
done

# Wait for all servers to finish
wait