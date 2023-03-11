package main

import (
	"github.com/pnguyen215/gobase-voip-core/pkg/ami"
)

func NewClient() (*ami.AMI, error) {
	client, err := ami.NewAmi("127.0.0.1", 5038, "u_username", "u_password")
	return client, err
}

func main() {

}
