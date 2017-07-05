#!/usr/bin/python

import sys
import json

query = json.loads(sys.stdin.read())
ret   = dict()

netmodifier =  101 if query['vis'] == 'public' else 1
num_subnets = int(query['subnets'])

vpc_net = query['vpc']
netaddr, netmask = vpc_net.split("/")
subnetmask = int(netmask) + 8

for i in xrange(num_subnets):
    key = "%d" % i
    netlist = netaddr.split('.')
    netlist[2] = str(netmodifier + i)
    ret[key] = '%(ip)s/%(subnetmask)d' % \
        { "ip": '.'.join(netlist), "subnetmask": subnetmask }

sys.stdout.write(str(json.dumps(ret)))
