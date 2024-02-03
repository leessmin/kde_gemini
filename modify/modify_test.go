package modify

import "testing"

func TestJudgeTime(t *testing.T) {
	s := judgeTime("07:00", "19:00")
	t.Log(s)
}

func TestModifyTheme(t *testing.T) {
	ModifyTheme()
}
