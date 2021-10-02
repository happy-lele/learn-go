package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	m1 := map[string]int{}
	m1["one"] = 1
	m2 := make(map[string]int, 10)

	t.Log(m, m1, m2)
}

func TestAccessNotExistingKey(t *testing.T)  {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
    // 返回两个值，v是值，ok为布尔类型，true代表存在
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	// key value
	for k, v := range m1 {
		t.Log(k, v)
	}
}