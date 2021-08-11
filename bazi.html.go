package bazi

import (
	"fmt"

	htmlgo "github.com/yangtizi/htmlgo"
)

// ToHTML 转换成网页形式
func (m *TBazi) ToHTML() string {
	html := htmlgo.NewHTML()
	html.SetTitle("三清宫八字")

	row := htmlgo.NewRow()
	row.AddTo(html.GetBody())

	ganzhi := func(strText1, strText2 string, pZhu *TZhu) *htmlgo.TDiv {
		column := htmlgo.NewColumn().SetFlex(1)

		// column.SetBackground(fmt.Sprintf("rgb(%d,%d,%d)", rand.Intn(255), rand.Intn(255), rand.Intn(255)))

		div := htmlgo.NewDiv().AddTo(column)
		htmlgo.NewFont().P().SetText(strText1).SetSize(2).SetColor("gray").AddTo(div)
		htmlgo.NewFont().P().SetText(strText2).SetSize(2).SetColor("gray").AddTo(div)

		{
			// 天干
			row := htmlgo.NewRow().SetPadding("10px").AddTo(column)

			htmlgo.NewFont().SetText(pZhu.Gan().String()).SetSize(8).SetColor(pZhu.Gan().ToWuXing().Color()).AddTo(row)
			htmlgo.NewFont().SetText(pZhu.ShiShen().String()).SetSize(3).SetColor(pZhu.Gan().ToWuXing().Color()).AddTo(row)

		}

		{
			// 地支
			row := htmlgo.NewRow().SetPadding("10px").AddTo(column)
			htmlgo.NewFont().SetText(pZhu.Zhi().String()).SetSize(8).SetColor(pZhu.Zhi().ToWuXing().Color()).AddTo(row)
			column := htmlgo.NewColumn().AddTo(row)
			// 藏干
			for j := 0; j < pZhu.CangGan().Size(); j++ {
				htmlgo.NewFont().SetText(pZhu.CangGan().ShiShen(j).String()).SetSize(3).SetColor(pZhu.CangGan().Gan(j).ToWuXing().Color()).AddTo(column)
			}

		}

		return column
	}

	// 四柱
	ganzhi(fmt.Sprintf("%d年", m.Date().Year()), m.LunarDate().Year(), m.SiZhu().YearZhu()).AddTo(row)
	ganzhi(fmt.Sprintf("%02d月", m.Date().Month()), m.LunarDate().Month(), m.SiZhu().MonthZhu()).AddTo(row)
	ganzhi(fmt.Sprintf("%02d日", m.Date().Day()), m.LunarDate().Day(), m.SiZhu().DayZhu()).AddTo(row)
	ganzhi(fmt.Sprintf("%02d时", m.Date().Hour()), m.LunarDate().Hour(), m.SiZhu().HourZhu()).AddTo(row)

	// 分隔符
	htmlgo.NewDiv().SetBackground("rgb(238,238,238)").SetMargin("10px 0px").SetHeight("5px").AddTo(html.GetBody())

	// htmlgo.NewDiv().SetMargin("10px").AddChild(
	// 	htmlgo.NewFont().SetText("大运"),
	// )

	for j := 0; j < 2; j++ {
		row = htmlgo.NewRow().SetBackground("lightgray").AddTo(html.GetBody())

		for k := 0; k < 5; k++ {
			i := j*5 + k
			column := htmlgo.NewColumn().SetBackground("white").SetMargin("1px").SetFlex(1).AddTo(row)

			htmlgo.NewDiv().SetMargin("10px").AddTo(column).AddChild(
				htmlgo.NewFont().SetColor("gray").SetSize(2).SetText(fmt.Sprintf("%d", m.DaYun().Age(i)+m.Date().Year())))

			htmlgo.NewDiv().SetMargin("3px").AddTo(column).AddChild(
				htmlgo.NewFont().SetColor(m.DaYun().Zhu(i).Gan().ToWuXing().Color()).SetSize(5).SetText(m.DaYun().Zhu(i).Gan().String()))

			htmlgo.NewDiv().SetMargin("3px").AddTo(column).AddChild(
				htmlgo.NewFont().SetColor(m.DaYun().Zhu(i).Zhi().ToWuXing().Color()).SetSize(5).SetText(m.DaYun().Zhu(i).Zhi().String()))

			htmlgo.NewDiv().SetMargin("10px").AddTo(column).AddChild(
				htmlgo.NewFont().SetColor("gray").SetSize(2).SetText(fmt.Sprintf("%d", m.DaYun().Age(i))))
		}
	}

	return html.String()
}
