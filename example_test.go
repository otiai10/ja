package ja

import "fmt"

func ExampleLinebreak() {
	s := `沖縄本島地方と先島諸島では、２１日夜のはじめ頃にかけて、台風の接近と大潮の時期が重なるため潮位が高くなる見込みです。海岸や河口付近の低地での高潮による浸水や冠水に注意してください。`

	lines := Linebreak(s, MoreThan(20))
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// 沖縄本島地方と先島諸島では、２１日夜のはじめ頃にかけて、
	// 台風の接近と大潮の時期が重なるため潮位が高くなる見込みです。
	// 海岸や河口付近の低地での高潮による浸水や冠水に注意してください。
}
