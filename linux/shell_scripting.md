Setting up Virtual Network between Namespaces
This guide outlines the steps to create two namespaces named blue-namespace and lemon-namespace, and establish a virtual Ethernet network between them using veth interfaces. The goal is to enable communication between the namespaces and allow them to ping each other.

Prerequisites
Linux operating system

Root or sudo access

Packages

sudo apt update
sudo apt upgrade -y
sudo apt install iproute2
sudo apt install net-tools

Steps
1. Enable IP forwarding in the Linux kernel:

sudo sysctl -w net.ipv4.ip_forward=1

2. Create namespaces:
sudo ip netns add blue-namespace
sudo ip netns add lemon-namespace

3. Create the virtual Ethernet link pair:
sudo ip link add veth-blue type veth peer name veth-lemon
In order to verify, run sudo ip link list

4. Set the cable as NIC
sudo ip link set veth-blue netns blue-namespace
sudo ip link set veth-lemon netns lemon-namespace
To verify run sudo ip netns exec blue-namespace ip link and sudo ip netns exec lemon-namespace ip link

But as we see, interface has been created but it's DOWN and has no ip. Now assign a ip address and turn it UP.

5. Assign IP Addresses to the Interfaces
sudo ip netns exec blue-namespace ip addr add 192.168.0.1/24 dev veth-blue
sudo ip netns exec lemon-namespace ip addr add 192.168.0.2/24 dev veth-lemon
To verify run sudo ip netns exec blue-namespace ip addr and sudo ip netns exec lemon-namespace ip addr

6. Set the Interfaces Up
sudo ip netns exec blue-namespace ip link set veth-blue up
sudo ip netns exec lemon-namespace ip link set veth-lemon up
Now run again sudo ip netns exec blue-namespace ip link and sudo ip netns exec lemon-namespace ip link to verify

7. Set Default Routes
sudo ip netns exec blue-namespace ip route add default via 192.168.0.1 dev veth-blue
sudo ip netns exec lemon-namespace ip route add default via 192.168.0.2 dev veth-lemon

8. Test Connectivity
sudo ip netns exec blue-namespace ping 192.168.0.2
sudo ip netns exec lemon-namespace ping 192.168.0.1
Use these commands to test the connectivity between the namespaces by pinging each other's IP address.

9. Clean Up (optional)
sudo ip netns del blue-namespace
sudo ip netns del lemon-namespace
If you want to remove the namespaces run these commands to clean up the setup.