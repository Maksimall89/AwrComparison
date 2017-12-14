package main
import (
	"testing"
)

// TODO coverage = 100%

func TestReplaceTag(t *testing.T) {
	t.Parallel()
/*
	type testPair struct {
		original string
		replaced string
	}

	var tests = []testPair{
		{"<span> sdf </span>", " sdf "},
		{"< /p >", "\n"},
		{"q</h3>w", "q</b>\nw"},
		{"< H4>", "<b>"},
		{"sdfs", "sdfs"},
		{`s<df>s`, "ss"},
		{`<a href="http://d1.endata.cx/data/games/58762/34vfmdshvf_b.jpg" border="0"><img src="http://d1.endata.cx/data/games/58762/34vfmdshvf_s.jpg" title="fds"></a>`, `[Спрятанная ссылка: <a href="http://d1.endata.cx/data/games/58762/34vfmdshvf_b.jpg">http://d1.endata.cx/data/games/58762/34vfmdshvf_b.jpg</a> под картинкой: <a href="http://d1.endata.cx/data/games/58762/34vfmdshvf_s.jpg">http://d1.endata.cx/data/games/58762/34vfmdshvf_s.jpg</a>]`},
		{`<a href="http://d1.endata.cx/data/games/58762/34vfmdshvf_b.jpg" ><img src="http://d1.endata.cx/data/games/58762/34vfmdshvf_s.jpg"></a>`, `[Спрятанная ссылка: <a href="http://d1.endata.cx/data/games/58762/34vfmdshvf_b.jpg">http://d1.endata.cx/data/games/58762/34vfmdshvf_b.jpg</a> под картинкой: <a href="http://d1.endata.cx/data/games/58762/34vfmdshvf_s.jpg">http://d1.endata.cx/data/games/58762/34vfmdshvf_s.jpg</a>]`},
	}

	for _, pair := range tests {
		v := replaceTag(pair.original)
		if v != pair.replaced {
			t.Error(
				"For", pair.original,
				"expected", pair.replaced,
				"got", v,
			)
		}
	}
*/
}
