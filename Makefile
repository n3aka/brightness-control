build:
	@echo "building.."
	go build .
	@echo "moving and perm"
	chmod +x brightness-control
	sudo mv brightness-control /usr/bin
	# sudo cp -n brightness-control.sh /usr/bin
	sudo cp bc.service /etc/systemd/system/
	sudo systemctl daemon-reload
	sudo systemctl restart bc

