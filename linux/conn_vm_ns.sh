#!/bin/bash
sudo apt update 
sudo apt upgrade -y
#tools
sudo apt install iproute2
sudo apt install net-tools

sudo ip link add br0 type bridge
sudo ip addr add 192.168.0.1/24 dev br0


#create namespace 

#namespace vm0
sudo ip netns add vm0
sudo ip link add veth0 type veth peer name ceth0
sudo ip link set ceth0 netns vm0
sudo ip link set veth0 up
sudo ip link set veth0 master br0
sudo ip netns exec vm0 ip addr add 192.168.0.10/24 dev ceth0
sudo ip netns exec vm0 ip link set lo up
sudo ip netns exec vm0 ip link set ceth0 up
sudo ip netns exec vm0 ip route add default via 192.168.0.1 dev ceth0

#namespace vm1
sudo ip netns add vm1
sudo ip link add veth1 type veth peer name ceth1
sudo ip link set ceth1 netns vm1
sudo ip link set veth1 up
sudo ip link set veth1 master br0
sudo ip netns exec vm1  ip addr add 192.168.0.11/24 dev ceth1
sudo ip netns exec vm1 ip link set lo up
sudo ip netns exec vm1 ip link set ceth1 up
sudo ip netns exec vm1 ip route add default via 192.168.0.1 dev ceth1


# test coneection
sudo ip netns exec vm0 ping 192.168.0.11
sudo ip netns exec vm1 ping 192.168.0.10


sud sysctl -w net.ipv4.ip_forward=1
sudo iptables -t nat -A POSTROUTING -s 192.168.0.0/24 ! -o br0 -j MASQUERADE