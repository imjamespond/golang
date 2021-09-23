package test

import (
	"log"
	"testing"

	"test-casbin/utils"

	"github.com/casbin/casbin/v2"
)

func Test2(t *testing.T) {
	e, err := casbin.NewEnforcer("./test2.model.conf", "./test2.policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	utils.Check(e, "dajun", "data", "read")
	utils.Check(e, "dajun", "data", "write")
	utils.Check(e, "lizi", "data", "read")
	utils.Check(e, "lizi", "data", "write")

	utils.Check(e, "dajun", "prod.data", "read")
	utils.Check(e, "dajun", "prod.data", "write")
	utils.Check(e, "lizi", "dev.data", "read")
	utils.Check(e, "lizi", "dev.data", "write")
	utils.Check(e, "lizi", "prod.data", "write")

}
