BUILD_DIR := my-acr
all: mongo_list my_acr.deb

mongo_list:
	go build -o $(BUILD_DIR)/var/lib/my/bin/mongo_list mongo_list.go

my_acr.deb: $(BUILD_DIR)/var/lib/my/bin/mongo_list my-acr/DEBIAN/control
	dpkg-deb --build my-acr
