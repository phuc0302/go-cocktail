package mongo

import (
	"testing"

	"gopkg.in/mgo.v2"
)

func TestCreateSession(t *testing.T) {
	CreateSession(Config{Host: "", Port: "", Database: "", Username: "", Password: ""})
	CreateSession(Config{Host: "localhost", Port: "27017", Database: "concept"})
	CreateSession(Config{Host: "localhost", Port: "27017", Database: "concept", Username: "", Password: ""})
}

func TestGetEventualSession(t *testing.T) {
	session, database := GetEventualSession()

	if session.Mode() != mgo.Eventual {
		t.Error("Not Eventual Session.")
	}

	if database.Name != "concept" {
		t.Error("Invalid database name.")
	}
}
