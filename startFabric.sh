#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error
set -e

# launch network
cd network
# Raft ordering and couchdb can be disabled by removing corresponding options.
echo y | ./byfn.sh up -a -n -s couchdb -o etcdraft
