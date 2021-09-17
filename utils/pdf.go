package utils

import (
	"fmt"
	"github.com/signintech/gopdf"
)

const FONTPATH = "/Users/yangxianglong/go/Vue_Iris/back-end/static/font/"
const IMGPATH = "/Users/yangxianglong/go/Vue_Iris/back-end/static/img/"
const FILEPATH = "/Users/yangxianglong/go/Vue_Iris/back-end/static/file/"

func NewPdf() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 850.32, H: 1203.12}}) //595.28, 841.89 = A4
	pdf.AddPage()

	Title(&pdf)
	Customer(&pdf)
	drawGrid(&pdf)
	Company(&pdf)
	BodyTitle(&pdf)
	Deliverables(&pdf)
	pdf.WritePdf(FILEPATH + "pdf/見積書.pdf")
}

/*
* ①标题「御見積書」
* ②标题外框
* ③见积日期
 */
type Header struct {
	x1 float64
	y1 float64
	w  float64
	h  float64
}

func Title(pdf *gopdf.GoPdf) {
	//外边框
	var t1 = &Header{
		x1: 310,
		y1: 173,
		w:  228,
		h:  35,
	}
	//内边框
	var t2 = &Header{
		x1: 312.5,
		y1: 175.5,
		w:  223,
		h:  30,
	}
	err := pdf.AddTTFFont("mincho", FONTPATH+"ShipporiAntiqueB1-Regular.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("mincho", "", 26) //フォント、文字サイズ指定
	pdf.SetX(335)                 //x座標指定
	pdf.SetY(178)                 //y座標指定
	pdf.Cell(nil, "御　見　積　書")      //Rect, String
	pdf.SetLineWidth(0.7)

	//描画外边框
	pdf.Line(t1.x1, t1.y1, t1.x1+t1.w, t1.y1)
	pdf.Line(t1.x1+t1.w, t1.y1, t1.x1+t1.w, t1.y1+t1.h)
	pdf.Line(t1.x1, t1.y1+t1.h, t1.x1+t1.w, t1.y1+t1.h)
	pdf.Line(t1.x1, t1.y1, t1.x1, t1.y1+t1.h)

	//描画内边框
	pdf.Line(t2.x1, t2.y1, t2.x1+t2.w, t2.y1)
	pdf.Line(t2.x1+t2.w, t2.y1, t2.x1+t2.w, t2.y1+t2.h)
	pdf.Line(t2.x1, t2.y1+t2.h, t2.x1+t2.w, t2.y1+t2.h)
	pdf.Line(t2.x1, t2.y1, t2.x1, t2.y1+t2.h)

	err = pdf.AddTTFFont("simfang", FONTPATH+"/simfang.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("simfang", "", 14) //フォント、文字サイズ指定
	pdf.SetX(600)                  //x座標指定
	pdf.SetY(210)                  //y座標指定
	pdf.Cell(nil, "見積No.")         //見積No.
	pdf.SetX(655)
	pdf.SetY(210)
	pdf.Cell(nil, "Esh210831145627") //Esh210831145627
	pdf.SetX(655)
	pdf.SetY(225)
	pdf.Cell(nil, "16/09/2021") //16/09/2021
}

/*
* ①得意先名〇〇御中
* ②挨拶
* ③見積No.
 */
func Customer(pdf *gopdf.GoPdf) {
	//①得意先名〇〇
	pdf.SetFont("mincho", "", 18) //フォント、文字サイズ指定
	customerName := "株式会社ヤマダ電機"
	pdf.SetX(115) //x座標指定
	pdf.SetY(256) //y座標指定
	pdf.Cell(nil, customerName)
	//〇〇御中
	pdf.SetX(130 + float64(len(customerName))*6) //x座標指定
	pdf.SetY(256)                                //y座標指定
	pdf.Cell(nil, "御中")
	pdf.SetLineWidth(0.7)
	pdf.Line(115, 275, 115+float64(len(customerName)+9)*6, 275)

	//②挨拶
	//err = pdf.AddTTFFont("simfang", "./font/simfang.ttf")
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("ipaexm", "", 14) //フォント、文字サイズ指定
	pdf.SetX(115)                 //x座標指定
	pdf.SetY(285)                 //y座標指定
	pdf.Cell(nil, "拝啓 貴社御依頼に対し、下記の通り御見積り申し上げます。")

	pdf.SetX(115) //x座標指定
	pdf.SetY(300) //y座標指定
	pdf.Cell(nil, "何卒ご用命のほど、よろしくお願い申し上げます。")
}

func Company(pdf *gopdf.GoPdf) {
	err := pdf.AddTTFFont("simfang", FONTPATH+"simfang.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("simfang", "", 12) //フォント、文字サイズ指定

	pdf.Image(IMGPATH+"bridge_logo.png", 550, 295, nil)

	pdf.SetX(530)
	pdf.SetY(315)
	pdf.Cell(nil, "〒104-0032"+"　　"+"株式会社ブリッジ") //邮编
	pdf.SetX(530)
	pdf.SetY(330)
	pdf.Cell(nil, "　東京都中央区八丁堀4丁目11-10第2SSビル　1F") //地址
	pdf.SetX(530)
	pdf.SetY(345)
	pdf.Cell(nil, "　Tel:03-6222-3222　Fax:03-6222-3228") //联系方式
	pdf.SetX(530)
	pdf.SetY(360)
	pdf.Cell(nil, "※有効期限:30日　　　担当　栗原") //作成者

}

func BodyTitle(pdf *gopdf.GoPdf) {
	err := pdf.AddTTFFont("PingBold", FONTPATH+"PingBold.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("PingBold", "", 12) //フォント、文字サイズ指定

	pdf.SetX(140)
	pdf.SetY(375)
	pdf.Cell(nil, "１. 見積対象案件名")
	pdf.SetX(140)
	pdf.SetY(450)
	pdf.Cell(nil, "２．作業内容")
	pdf.SetX(140)
	pdf.SetY(540)
	pdf.Cell(nil, "３．成果物")
	pdf.SetX(140)
	pdf.SetY(650)
	pdf.Cell(nil, "４．作業場所")
	pdf.SetX(140)
	pdf.SetY(725)
	pdf.Cell(nil, "５．お見積金額")
	pdf.SetX(140)
	pdf.SetY(1030)
	pdf.Cell(nil, "６．支払条件")
}

func Deliverables(pdf *gopdf.GoPdf) {
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
	pdf.SetX(163)
	pdf.SetY(555)
	pdf.Cell(nil, "成果物及び納品予定日は、以下の通りです。")

	lineX1 := 163.0
	lineY1 := 565.0
	lineW := 622.0
	lineH := 12.0

	pdf.SetLineWidth(0.7)

	pdf.Line(lineX1, lineY1, lineX1, lineY1+lineH*4)             //左
	pdf.Line(470, lineY1, 470, lineY1+lineH*4)                   //左2
	pdf.Line(590, lineY1, 590, lineY1+lineH*4)                   //左3
	pdf.Line(650, lineY1, 650, lineY1+lineH*4)                   //右2
	pdf.Line(lineX1+lineW, lineY1, lineX1+lineW, lineY1+lineH*4) //右

	for num := 0.0; num <= 4; num++ {
		pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)
	}

	pdf.SetX(280)
	pdf.SetY(566)
	pdf.Cell(nil, "成　果　物")
	pdf.SetX(510)
	pdf.Cell(nil, "納品媒体")
	pdf.SetX(610)
	pdf.Cell(nil, "部数")
	pdf.SetX(690)
	pdf.Cell(nil, "納品予定日")

}

func drawGrid(pdf *gopdf.GoPdf) {

	err := pdf.AddTTFFont("ping", FONTPATH+"ping.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("ping", "", 8) //フォント、文字サイズ指定

	pdf.SetLineWidth(0.3)

	var x, y float64
	for x = 0; x < 850; x += 50 {
		pdf.Line(x, 0, x, 10)
		pdf.SetX(x)  //x座標指定
		pdf.SetY(10) //y座標指定
		s := fmt.Sprintf("%.0f", x)
		pdf.Cell(nil, s)
	}
	for y = 0; y <= 1200; y += 50 {
		pdf.Line(0, y, 10, y)
		pdf.SetX(10) //x座標指定
		pdf.SetY(y)  //y座標指定
		s := fmt.Sprintf("%.0f", y)
		pdf.Cell(nil, s)
	}
}
