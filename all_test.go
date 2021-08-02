package ja

import (
	"testing"

	. "github.com/otiai10/mint"
)

const str = `沖縄地方は、台風第６号の影響で曇りや雨の天気となっています。
２１日から２２日にかけて台風の影響で、先島諸島は大荒れ、沖縄本島地方と大東島地方は荒れた天気となり、所により雷を伴い、激しい雨や非常に激しい雨が降るでしょう。
宮古島地方では２２日にかけて、暴風に警戒してください。八重山地方では、北の風が強く吹いており、２２日未明からは暴風となる所がある見込みです。
沖縄地方では２２日にかけて、発達した積乱雲の下での落雷や竜巻などの激しい突風に十分注意してください。
沖縄本島地方と先島諸島では、２１日夜のはじめ頃にかけて、台風の接近と大潮の時期が重なるため潮位が高くなる見込みです。海岸や河口付近の低地での高潮による浸水や冠水に注意してください。
沿岸の海域はうねりを伴い、沖縄本島地方と先島諸島では大しけとなっています。沿岸の海域ではうねりを伴う高波に警戒してください。大東島地方ではうねりを伴い波が高くなっています。高波に注意してください。`

func TestParagraph(t *testing.T) {
	blocks := Paragraph(str, Within(60))
	Expect(t, len(blocks)).ToBe(6)
	Expect(t, len(blocks[0])).ToBe(1)
	Expect(t, len(blocks[1])).ToBe(2)
	Expect(t, len(blocks[2])).ToBe(2)
	Expect(t, len(blocks[3])).ToBe(1)
}

func TestCut(t *testing.T) {
	s := "予報期間　８月３日から８月９日まで\n　向こう一週間は、期間の前半は高気圧に覆われて晴れる日もありますが、気圧の谷や湿った空気の影響で雲が広がりやすいでしょう。\n　最高気温と最低気温はともに、平年並か平年より高い見込みです。熱中症など健康管理に注意してください。\n　降水量は、平年並か平年より少ないでしょう。"
	x1 := Cut(s)
	Expect(t, x1[1]).ToBe("　向こう一週間は、")
	x2 := Cut(s, true)
	Expect(t, x2[1]).ToBe("向こう一週間は、")
}
