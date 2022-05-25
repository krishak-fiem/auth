package constants

import "github.com/krishak-fiem/utils/go/env"

var CASSANDRA_HOST = env.Get("AUTH_CASSANDRA_HOST")
var CASSANDRA_PORT = env.Get("AUTH_CASSANDRA_PORT")
var KAFKA_BROKER_URL = env.Get("KAFKA_BROKER_URL")
