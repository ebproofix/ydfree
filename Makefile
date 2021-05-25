run: createDir 
	go run main.go -y ./testing/yd_dir -b ./testing/rc_dir -m del

test: createDir
	go run main.go -y ./testing/yd_dir -b ./testing/rc_dir -m print

createDir:
	./scripts/makeDirs.sh