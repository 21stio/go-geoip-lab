SHELL:=sh -x

play:
	test -d vendor || glide install
	test -f geoip_database.mmdb || ./downloadGeoipDatabase.sh ;\
	go run main.go ;\