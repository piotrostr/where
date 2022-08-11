# this is going to have to be for the small images like alpine
# so some of the build flags will need to be included
build:
	go build .

run:
	go build -o where . && ./where -port 8080

install:
	go install
