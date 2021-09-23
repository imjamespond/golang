package test

import (
	"log"
	"testing"

	"test-casbin/utils"

	"github.com/casbin/casbin/v2"
)

func Test1(t *testing.T) {
	e, err := casbin.NewEnforcer("./test1.model.conf", "./test1.policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	utils.Check(e, "dajun", "data1", "read")
	utils.Check(e, "lizi", "data2", "write")
	utils.Check(e, "dajun", "data1", "write")
	utils.Check(e, "dajun", "data2", "read")

	utils.Check(e, "root", "data1", "read")
	utils.Check(e, "root", "data2", "write")
	utils.Check(e, "root", "data1", "execute")
	utils.Check(e, "root", "data3", "rwx")

}
