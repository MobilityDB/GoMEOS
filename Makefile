# Define the directory where the MobilityDB repository will be cloned
MOBILITYDB_DIR = MobilityDB

# Check if the repository is already cloned
$(MOBILITYDB_DIR):
	git clone https://github.com/MobilityDB/MobilityDB.git $(MOBILITYDB_DIR)

# Define the build steps
install: $(MOBILITYDB_DIR)
	cd $(MOBILITYDB_DIR) && mkdir -p build && cd build && \
	cmake -DMEOS=on .. && \
	make && \
	sudo -A make install