#!/bin/bash

# Directory to clone the MobilityDB repository
MOBILITYDB_DIR="MobilityDB"

# Check if the repository is already cloned
if [ ! -d "$MOBILITYDB_DIR" ]; then
    git clone https://github.com/MobilityDB/MobilityDB.git $MOBILITYDB_DIR
fi

# Build and install the MEOS library
cd $MOBILITYDB_DIR
mkdir -p build
cd build
cmake -DMEOS=on ..
make -j
sudo make install