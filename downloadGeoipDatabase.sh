#!/bin/bash

set -e

geoip_database_location=http://geolite.maxmind.com/download/geoip/database/GeoLite2-Country.mmdb.gz
destination=geoip_database

wget ${geoip_database_location} -O ${destination}.gz
gzip --decompress --stdout ${destination}.gz > ${destination}.mmdb