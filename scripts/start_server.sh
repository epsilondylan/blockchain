
#!/bin/bash
NUM_servers=$1

# Start Servers with ssh connection
for p in $(seq 8051 $((8050+NUM_servers))); do
   ssh osgroup16@122.200.68.26 -i ../id_rsa -o StrictHostKeyChecking=no -p $p -t "nohup bash -c 'cd /osdata/osgroup16 && chmod +x miner && ./miner > miner.out 2>&1 &'"&
done

# Wait for all servers to finish
wait