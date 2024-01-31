package notice

import "testing"

func TestStartup(t *testing.T) {
	err := New("'测试'", "'这是内容内容内容'")

	t.Log(err)
}
