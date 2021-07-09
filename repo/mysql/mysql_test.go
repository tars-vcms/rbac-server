package mysql


import (
	"fmt"
	"testing"
)

func TestForDebug(t *testing.T)  {
	fmt.Println(AccessRoleSqlImpl.HaveAccessPointId(1, 2))
}
