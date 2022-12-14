package mapset_test

import "testing"

func TestMapToFactoryModel(t *testing.T) {
	//  使用map实现工厂模式
	m := map[string]func(num int) int{}

	m["getSelf"] = func(num int) int { return num }
	m["sqrtNum"] = func(num int) int { return num * num }

	t.Log(m["getSelf"](10))
	t.Log(m["sqrtNum"](10))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]string{1: "true", 2: "true", 0: "false"}

	n := 5
	if v, ok := mySet[n]; ok {
		t.Logf("%s is existing", v)
	} else {
		// 添加k-v
		t.Log("k-v no existing...")
		t.Logf("append myset[%d]=%s...", n, "hello")
		mySet[n] = "hello"
		t.Log("append finish...")
	}

	t.Log(mySet)
}
