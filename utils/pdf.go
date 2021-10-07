package utils

import (
	"fmt"
	"io/ioutil"
	"main/model"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/signintech/gopdf"
)

const FONTPATH = "/Users/yangxianglong/go/Vue_Iris/back-end/static/font/"
const IMGPATH = "/Users/yangxianglong/go/Vue_Iris/back-end/static/img/"
const FILEPATH = "/Users/yangxianglong/go/Vue_Iris/back-end/static/file/pdf/estimate/"

func NewPdf(estimate []*model.Estimate, estimateDetail []*model.EstimateDetail) {
	//获取见积头数据
	var estimateInfo model.Estimate
	for _, item := range estimate {
		estimateInfo = *item
	}

	//获取见积详细数据
	var estimateDetailInfoInitial, estimateDetailInfoRunning []model.EstimateDetail
	for _, item := range estimateDetail {
		if !item.MainFlag {
			estimateDetailInfoInitial = append(estimateDetailInfoInitial, *item)
		} else {
			estimateDetailInfoRunning = append(estimateDetailInfoRunning, *item)
		}

	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 850.32, H: 1203.12}}) //595.28, 841.89 = A4
	pdf.AddPage()
	pdf.SetFillColor(0, 0, 0)
	Title(&pdf, estimateInfo)
	Customer(&pdf, estimateInfo)
	// drawGrid(&pdf)
	Company(&pdf, estimateInfo)
	EstimateName(&pdf, estimateInfo)
	BodyTitle(&pdf)
	Work(&pdf, estimateInfo)
	Deliverables(&pdf, estimateInfo)
	WorkSpace(&pdf, estimateInfo)
	EstimateDetail(&pdf, estimateInfo, estimateDetailInfoInitial, estimateDetailInfoRunning)

	now := time.Now().Format("2006-01-02")
	_, err := os.Stat(FILEPATH + now)
	if err != nil {
		os.Mkdir(FILEPATH+now, os.ModePerm)
	}

	fileInfo, _ := ioutil.ReadDir(FILEPATH + now)

	var files []string
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}

	iris.New().Logger().Info(len(files))

	// PaymentConditions(&pdf, estimateInfo)
	pdf.WritePdf(FILEPATH + now + "/" + estimateInfo.EstimateCode + ".pdf")
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

func Title(pdf *gopdf.GoPdf, info model.Estimate) {
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
	pdf.Cell(nil, info.EstimateCode) //Esh210831145627
	pdf.SetX(655)
	pdf.SetY(225)
	pdf.Cell(nil, info.CreatedAt.Format("02/01/2006")) //16/09/2021
}

/*
* ①得意先名〇〇御中
* ②挨拶
* ③見積No.
 */
func Customer(pdf *gopdf.GoPdf, info model.Estimate) {
	//①得意先名〇〇
	pdf.SetFont("mincho", "", 18) //フォント、文字サイズ指定
	customerName := info.CustomerName
	pdf.SetX(95)  //x座標指定
	pdf.SetY(256) //y座標指定
	pdf.Cell(nil, customerName)
	//〇〇御中
	pdf.SetX(110 + float64(len(customerName))*6) //x座標指定
	pdf.SetY(256)                                //y座標指定
	pdf.Cell(nil, "御中")
	pdf.SetLineWidth(0.7)
	pdf.Line(95, 275, 95+float64(len(customerName)+9)*6, 275)

	//②挨拶
	//err = pdf.AddTTFFont("simfang", "./font/simfang.ttf")
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("ipaexm", "", 14) //フォント、文字サイズ指定
	pdf.SetX(95)                  //x座標指定
	pdf.SetY(285)                 //y座標指定
	pdf.Cell(nil, "拝啓 貴社御依頼に対し、下記の通り御見積り申し上げます。")

	pdf.SetX(95)  //x座標指定
	pdf.SetY(300) //y座標指定
	pdf.Cell(nil, "何卒ご用命のほど、よろしくお願い申し上げます。")
}

func Company(pdf *gopdf.GoPdf, info model.Estimate) {
	err := pdf.AddTTFFont("simfang", FONTPATH+"simfang.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("simfang", "", 12) //フォント、文字サイズ指定

	pdf.Image(IMGPATH+"bridge_logo.png", 530, 295, nil)

	pdf.SetX(510)
	pdf.SetY(315)
	pdf.Cell(nil, "〒104-0032"+"　　"+"株式会社ブリッジ") //邮编
	pdf.SetX(510)
	pdf.SetY(330)
	pdf.Cell(nil, "　東京都中央区八丁堀4丁目11-10第2SSビル　1F") //地址
	pdf.SetX(510)
	pdf.SetY(345)
	pdf.Cell(nil, "　Tel:03-6222-3222　Fax:03-6222-3228") //联系方式
	pdf.SetX(510)
	pdf.SetY(360)
	pdf.Cell(nil, "※有効期限:30日"+"　　　　　"+info.CreatedBy) //作成者

}

func BodyTitle(pdf *gopdf.GoPdf) {
	err := pdf.AddTTFFont("PingBold", FONTPATH+"PingBold.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("PingBold", "", 12) //フォント、文字サイズ指定

	pdf.SetX(120)
	pdf.SetY(375)
	pdf.Cell(nil, "１. 見積対象案件名")
	pdf.SetX(120)
	pdf.SetY(430)
	pdf.Cell(nil, "２．作業内容")
	pdf.SetX(120)
	pdf.SetY(500)
	pdf.Cell(nil, "３．成果物")
	pdf.SetX(120)
	pdf.SetY(585)
	pdf.Cell(nil, "４．作業場所")
	pdf.SetX(120)
	pdf.SetY(620)
	pdf.Cell(nil, "５．お見積金額")
}

//见积对象案件名
func EstimateName(pdf *gopdf.GoPdf, info model.Estimate) {
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}

	pdf.SetFont("ipaexm", "", 11) //フォント、文字サイズ指定
	pdf.SetX(153)
	pdf.SetY(390)
	pdf.Cell(nil, info.EstimateName)
}

//作业内容
func Work(pdf *gopdf.GoPdf, info model.Estimate) {
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}

	pdf.SetFont("ipaexm", "", 11) //フォント、文字サイズ指定

	if len(info.Work1) != 0 && len(info.Work2) != 0 && len(info.Work3) != 0 {
		pdf.SetX(153)
		pdf.SetY(445)
		pdf.Cell(nil, "・"+info.Work1)
		pdf.SetX(153)
		pdf.SetY(458)
		pdf.Cell(nil, "・"+info.Work2)
		pdf.SetX(153)
		pdf.SetY(471)
		pdf.Cell(nil, "・"+info.Work3)
	} else if len(info.Work1) != 0 && len(info.Work2) != 0 {
		pdf.SetX(153)
		pdf.SetY(445)
		pdf.Cell(nil, "・"+info.Work1)
		pdf.SetX(153)
		pdf.SetY(458)
		pdf.Cell(nil, "・"+info.Work2)
	} else {
		pdf.SetX(153)
		pdf.SetY(445)
		pdf.Cell(nil, "・"+info.Work1)
	}

}

//成果物
func Deliverables(pdf *gopdf.GoPdf, info model.Estimate) {
	lineX1 := 153.0
	lineY1 := 525.0
	lineW := 622.0
	lineH := 12.0

	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
	pdf.SetX(lineX1)
	pdf.SetY(515)
	pdf.Cell(nil, "成果物及び納品予定日は、以下の通りです。")

	//追加背景颜色
	pdf.SetFillColor(255, 255, 153)
	pdf.RectFromUpperLeftWithStyle(lineX1, lineY1, lineW, lineH, "FD")
	pdf.SetFillColor(255, 255, 153)

	pdf.SetLineWidth(0.7)
	pdf.Line(lineX1, lineY1, lineX1, lineY1+lineH*4)             //左
	pdf.Line(460, lineY1, 460, lineY1+lineH*4)                   //左2
	pdf.Line(580, lineY1, 580, lineY1+lineH*4)                   //左3
	pdf.Line(640, lineY1, 640, lineY1+lineH*4)                   //右2
	pdf.Line(lineX1+lineW, lineY1, lineX1+lineW, lineY1+lineH*4) //右

	for num := 0.0; num <= 4; num++ {
		pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)
	}

	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(270)
	pdf.SetY(lineY1 + 1)
	pdf.Cell(nil, "成　果　物")
	pdf.SetX(500)
	pdf.Cell(nil, "納品媒体")
	pdf.SetX(600)
	pdf.Cell(nil, "部数")
	pdf.SetX(680)
	pdf.Cell(nil, "納品予定日")

	pdf.SetX(155)
	pdf.SetY(lineY1 + 13)
	pdf.Cell(nil, info.Deliverables1)
	pdf.SetX(462)
	pdf.Cell(nil, info.Media1)
	pdf.SetX(582)
	pdf.Cell(nil, info.Quantity1)
	pdf.SetX(642)
	pdf.Cell(nil, info.DeliveryDate1)

	pdf.SetX(155)
	pdf.SetY(lineY1 + 25)
	pdf.Cell(nil, info.Deliverables2)
	pdf.SetX(462)
	pdf.Cell(nil, info.Media2)
	pdf.SetX(582)
	pdf.Cell(nil, info.Quantity2)
	pdf.SetX(642)
	pdf.Cell(nil, info.DeliveryDate2)

	pdf.SetX(155)
	pdf.SetY(lineY1 + 37)
	pdf.Cell(nil, info.Deliverables3)
	pdf.SetX(462)
	pdf.Cell(nil, info.Media3)
	pdf.SetX(582)
	pdf.Cell(nil, info.Quantity3)
	pdf.SetX(642)
	pdf.Cell(nil, info.DeliveryDate3)
}

//作业场所
func WorkSpace(pdf *gopdf.GoPdf, info model.Estimate) {
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("ipaexm", "", 11) //フォント、文字サイズ指定
	pdf.SetX(153)
	pdf.SetY(600)
	pdf.Cell(nil, info.WorkSpace)
}

//见积详细
func EstimateDetail(pdf *gopdf.GoPdf, info model.Estimate, infoInitial []model.EstimateDetail, infoRunning []model.EstimateDetail) {
	lineX1 := 153.0
	lineY1 := 645.0
	lineW := 622.0
	lineH := 12.0
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}

	var sumInitial = 0
	for _, item := range infoInitial {
		subTotalInitial, _ := strconv.Atoi(item.SubTotal)
		sumInitial += subTotalInitial
	}
	var sumRunning = 0
	for _, item := range infoRunning {
		subTotalRunning, _ := strconv.Atoi(item.SubTotal)
		sumRunning += subTotalRunning
	}

	if len(infoInitial) != 0 && len(infoRunning) != 0 {
		//通过数据个数计算表的列数 数据个数 + 表头一行 + 末尾三行
		lenInfoInitial := len(infoInitial) + 4

		//追加背景颜色
		pdf.SetFillColor(255, 255, 153)
		pdf.RectFromUpperLeftWithStyle(lineX1, lineY1, lineW, lineH, "FD")
		pdf.SetFillColor(255, 255, 153)

		pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
		pdf.SetX(lineX1)
		pdf.SetY(635)
		pdf.Cell(nil, "5.1　お見積金額（イニシャル費用）")

		pdf.SetLineWidth(1.0)
		pdf.Line(lineX1, lineY1, lineX1, lineY1+lineH*float64(lenInfoInitial))             //左
		pdf.Line(171, lineY1, 171, lineY1+lineH*float64(lenInfoInitial))                   //左2
		pdf.Line(420, lineY1, 420, lineY1+lineH*float64(lenInfoInitial))                   //左3
		pdf.Line(510, lineY1, 510, lineY1+lineH*float64(lenInfoInitial))                   //左4
		pdf.Line(550, lineY1, 550, lineY1+lineH*float64(lenInfoInitial))                   //右3
		pdf.Line(640, lineY1, 640, lineY1+lineH*float64(lenInfoInitial))                   //右2
		pdf.Line(lineX1+lineW, lineY1, lineX1+lineW, lineY1+lineH*float64(lenInfoInitial)) //右

		for num := 0.0; num <= float64(lenInfoInitial); num++ {
			if num == 0 || num == float64(lenInfoInitial) {
				pdf.SetLineWidth(1.0)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetX(lineX1 + 1)
				pdf.SetY(lineY1 + 1)
				pdf.Cell(nil, "No.")
				pdf.SetX(270)
				pdf.Cell(nil, "項　目")
				pdf.SetX(450)
				pdf.Cell(nil, "単　価")
				pdf.SetX(515)
				pdf.Cell(nil, "数　量")
				pdf.SetX(580)
				pdf.Cell(nil, "金　額")
				pdf.SetX(690)
				pdf.Cell(nil, "備　考")
			} else if num == 1 {
				pdf.SetLineWidth(1.0)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(156)
				pdf.SetY(lineY1 + lineH*num + 1)
				// pdf.Cell(nil, strconv.FormatFloat((float64(num)-1), 'g', -1, 32))
				pdf.Cell(nil, "1")
				pdf.SetX(172)
				pdf.Cell(nil, infoInitial[0].ProductName)
				pdf.SetX(510 - float64(len(infoInitial[0].Price))*7)
				pdf.Cell(nil, convertStr(infoInitial[0].Price))
				pdf.SetX(516)
				pdf.Cell(nil, convertStr(infoInitial[0].Quantity))
				pdf.SetX(640 - float64(len(infoInitial[0].SubTotal))*7)
				pdf.Cell(nil, convertStr(infoInitial[0].SubTotal))
				pdf.SetX(641)
				pdf.Cell(nil, " ")
			} else if num == float64(lenInfoInitial)-1 {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)
				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定

				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "合　計")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumInitial*110/100)))*7)
				pdf.Cell(nil, convert(sumInitial*110/100))
			} else if num == float64(lenInfoInitial)-2 {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "消費税")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumInitial/10)))*7)
				pdf.Cell(nil, convert(sumInitial/10))
				pdf.SetX(641)
				pdf.Cell(nil, "10%")
			} else if num == float64(lenInfoInitial)-3 {
				pdf.SetLineWidth(1.5)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "小　計")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumInitial)))*7)
				pdf.Cell(nil, convert(sumInitial))
			} else {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(156)
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
				pdf.SetX(172)
				pdf.Cell(nil, infoInitial[int(num)-1].ProductName)
				pdf.SetX(510 - float64(len(infoInitial[int(num)-1].Price))*7)
				pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Price))
				pdf.SetX(516)
				pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Quantity))
				pdf.SetX(640 - float64(len(infoInitial[int(num)-1].SubTotal))*7)
				pdf.Cell(nil, convertStr(infoInitial[int(num)-1].SubTotal))
				pdf.SetX(641)
				pdf.Cell(nil, " ")
			}
		}

		lenInfoRunning := len(infoRunning) + 4

		lineY1 = lineY1 + lineH*float64(lenInfoInitial) + 30

		//追加背景颜色
		pdf.SetFillColor(255, 255, 153)
		pdf.RectFromUpperLeftWithStyle(lineX1, lineY1, lineW, lineH, "FD")
		pdf.SetFillColor(255, 255, 153)

		pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
		pdf.SetX(lineX1)
		pdf.SetY(lineY1 - 10)
		pdf.Cell(nil, "5.2　お見積金額（ランニング費用）")

		pdf.SetLineWidth(1.0)
		pdf.Line(lineX1, lineY1, lineX1, lineY1+lineH*float64(lenInfoRunning))             //左
		pdf.Line(171, lineY1, 171, lineY1+lineH*float64(lenInfoRunning))                   //左2
		pdf.Line(420, lineY1, 420, lineY1+lineH*float64(lenInfoRunning))                   //左3
		pdf.Line(510, lineY1, 510, lineY1+lineH*float64(lenInfoRunning))                   //左4
		pdf.Line(550, lineY1, 550, lineY1+lineH*float64(lenInfoRunning))                   //右3
		pdf.Line(640, lineY1, 640, lineY1+lineH*float64(lenInfoRunning))                   //右2
		pdf.Line(lineX1+lineW, lineY1, lineX1+lineW, lineY1+lineH*float64(lenInfoRunning)) //右

		for num := 0.0; num <= float64(lenInfoRunning); num++ {
			if num == 0 || num == float64(lenInfoRunning) {
				pdf.SetLineWidth(1.0)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetX(lineX1 + 1)
				pdf.SetY(lineY1 + 1)
				pdf.Cell(nil, "No.")
				pdf.SetX(270)
				pdf.Cell(nil, "項　目")
				pdf.SetX(450)
				pdf.Cell(nil, "単　価")
				pdf.SetX(515)
				pdf.Cell(nil, "数　量")
				pdf.SetX(580)
				pdf.Cell(nil, "金　額")
				pdf.SetX(690)
				pdf.Cell(nil, "備　考")
			} else if num == 1 {
				pdf.SetLineWidth(1.0)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(156)
				pdf.SetY(lineY1 + lineH*num + 1)
				// pdf.Cell(nil, strconv.FormatFloat((float64(num)-1), 'g', -1, 32))
				pdf.Cell(nil, "1")
				pdf.SetX(172)
				pdf.Cell(nil, infoRunning[0].ProductName)
				pdf.SetX(510 - float64(len(infoRunning[0].Price))*7)
				pdf.Cell(nil, convertStr(infoRunning[0].Price))
				pdf.SetX(516)
				pdf.Cell(nil, convertStr(infoRunning[0].Quantity))
				pdf.SetX(640 - float64(len(infoRunning[0].SubTotal))*7)
				pdf.Cell(nil, convertStr(infoRunning[0].SubTotal))
				pdf.SetX(641)
				pdf.Cell(nil, " ")
			} else if num == float64(lenInfoRunning)-1 {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)
				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定

				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "合　計")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumRunning*110/100)))*7)
				pdf.Cell(nil, convert(sumRunning*110/100))
			} else if num == float64(lenInfoRunning)-2 {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "消費税")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumRunning/10)))*7)
				pdf.Cell(nil, convert(sumRunning/10))
				pdf.SetX(641)
				pdf.Cell(nil, "10%")
			} else if num == float64(lenInfoRunning)-3 {
				pdf.SetLineWidth(1.5)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "小　計")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumRunning)))*7)
				pdf.Cell(nil, convert(sumRunning))
			} else {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(156)
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
				pdf.SetX(172)
				pdf.Cell(nil, infoRunning[int(num)-1].ProductName)
				pdf.SetX(510 - float64(len(infoRunning[int(num)-1].Price))*7)
				pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Price))
				pdf.SetX(516)
				pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Quantity))
				pdf.SetX(640 - float64(len(infoRunning[int(num)-1].SubTotal))*7)
				pdf.Cell(nil, convertStr(infoRunning[int(num)-1].SubTotal))
				pdf.SetX(641)
				pdf.Cell(nil, " ")
			}
		}

		//补足部分
		lineY1 += lineH*float64(lenInfoRunning) + 1
		arr_str := strings.Split(info.Supplement, "\n")
		if len(info.Supplement) != 0 {
			pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
			pdf.SetX(lineX1)
			pdf.SetY(lineY1)
			pdf.Cell(nil, "【補　足】")
			for index, str := range arr_str {
				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(183)
				pdf.SetY(lineY1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}

		//5.3见积书部分

		if len(info.Other) != 0 {
			lineY1 += float64((len(arr_str)+1)*10) + 10
			pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
			pdf.SetX(lineX1)
			pdf.SetY(lineY1)
			pdf.Cell(nil, "5.3　その他費用")
			arr_str = strings.Split(info.Other, "\n")
			for index, str := range arr_str {
				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(183)
				pdf.SetY(lineY1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}

		lineY1 += float64((len(arr_str)+1)*10) + 15
		PaymentConditions(pdf, lineY1, info)
	} else if len(infoInitial) != 0 {
		//通过数据个数计算表的列数 数据个数 + 表头一行 + 末尾三行
		lenInfoInitial := len(infoInitial) + 4

		//追加背景颜色
		pdf.SetFillColor(255, 255, 153)
		pdf.RectFromUpperLeftWithStyle(lineX1, lineY1, lineW, lineH, "FD")
		pdf.SetFillColor(255, 255, 153)

		pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
		pdf.SetX(lineX1)
		pdf.SetY(635)
		pdf.Cell(nil, "5.1　お見積金額（イニシャル費用）")

		pdf.SetLineWidth(1.0)
		pdf.Line(lineX1, lineY1, lineX1, lineY1+lineH*float64(lenInfoInitial))             //左
		pdf.Line(171, lineY1, 171, lineY1+lineH*float64(lenInfoInitial))                   //左2
		pdf.Line(420, lineY1, 420, lineY1+lineH*float64(lenInfoInitial))                   //左3
		pdf.Line(510, lineY1, 510, lineY1+lineH*float64(lenInfoInitial))                   //左4
		pdf.Line(550, lineY1, 550, lineY1+lineH*float64(lenInfoInitial))                   //右3
		pdf.Line(640, lineY1, 640, lineY1+lineH*float64(lenInfoInitial))                   //右2
		pdf.Line(lineX1+lineW, lineY1, lineX1+lineW, lineY1+lineH*float64(lenInfoInitial)) //右

		for num := 0.0; num <= float64(lenInfoInitial); num++ {
			if num == 0 || num == float64(lenInfoInitial) {
				pdf.SetLineWidth(1.0)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetX(lineX1 + 1)
				pdf.SetY(lineY1 + 1)
				pdf.Cell(nil, "No.")
				pdf.SetX(270)
				pdf.Cell(nil, "項　目")
				pdf.SetX(450)
				pdf.Cell(nil, "単　価")
				pdf.SetX(515)
				pdf.Cell(nil, "数　量")
				pdf.SetX(580)
				pdf.Cell(nil, "金　額")
				pdf.SetX(690)
				pdf.Cell(nil, "備　考")
			} else if num == 1 {
				pdf.SetLineWidth(1.0)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(156)
				pdf.SetY(lineY1 + lineH*num + 1)
				// pdf.Cell(nil, strconv.FormatFloat((float64(num)-1), 'g', -1, 32))
				pdf.Cell(nil, "1")
				pdf.SetX(172)
				pdf.Cell(nil, infoInitial[0].ProductName)
				pdf.SetX(510 - float64(len(infoInitial[0].Price))*7)
				pdf.Cell(nil, convertStr(infoInitial[0].Price))
				pdf.SetX(516)
				pdf.Cell(nil, convertStr(infoInitial[0].Quantity))
				pdf.SetX(640 - float64(len(infoInitial[0].SubTotal))*7)
				pdf.Cell(nil, convertStr(infoInitial[0].SubTotal))
				pdf.SetX(641)
				pdf.Cell(nil, " ")
			} else if num == float64(lenInfoInitial)-1 {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)
				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定

				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "合　計")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumInitial*110/100)))*7)
				pdf.Cell(nil, convert(sumInitial*110/100))
			} else if num == float64(lenInfoInitial)-2 {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "消費税")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumInitial/10)))*7)
				pdf.Cell(nil, convert(sumInitial/10))
				pdf.SetX(641)
				pdf.Cell(nil, "10%")
			} else if num == float64(lenInfoInitial)-3 {
				pdf.SetLineWidth(1.5)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "小　計")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumInitial)))*7)
				pdf.Cell(nil, convert(sumInitial))
			} else {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(156)
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
				pdf.SetX(172)
				pdf.Cell(nil, infoInitial[int(num)-1].ProductName)
				pdf.SetX(510 - float64(len(infoInitial[int(num)-1].Price))*7)
				pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Price))
				pdf.SetX(516)
				pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Quantity))
				pdf.SetX(640 - float64(len(infoInitial[int(num)-1].SubTotal))*7)
				pdf.Cell(nil, convertStr(infoInitial[int(num)-1].SubTotal))
				pdf.SetX(641)
				pdf.Cell(nil, " ")
			}
		}

		//补足部分
		lineY1 += lineH*float64(lenInfoInitial) + 1
		arr_str := strings.Split(info.Supplement, "\n")

		iris.New().Logger().Info(len(info.Supplement))
		iris.New().Logger().Info(info.Supplement)
		if len(info.Supplement) != 0 {
			pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
			pdf.SetX(lineX1)
			pdf.SetY(lineY1)
			pdf.Cell(nil, "【補　足】")
			for index, str := range arr_str {
				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(183)
				pdf.SetY(lineY1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}

		//5.3见积书部分
		if len(info.Other) != 0 {
			lineY1 += float64((len(arr_str)+1)*10) + 10
			pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
			pdf.SetX(lineX1)
			pdf.SetY(lineY1)
			pdf.Cell(nil, "5.2　その他費用")
			arr_str = strings.Split(info.Other, "\n")
			for index, str := range arr_str {
				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(183)
				pdf.SetY(lineY1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}
		lineY1 += float64((len(arr_str)+1)*10) + 15
		PaymentConditions(pdf, lineY1, info)

	} else {
		//通过数据个数计算表的列数 数据个数 + 表头一行 + 末尾三行

		lenInfoRunning := len(infoRunning) + 4

		//追加背景颜色
		pdf.SetFillColor(255, 255, 153)
		pdf.RectFromUpperLeftWithStyle(lineX1, lineY1, lineW, lineH, "FD")
		pdf.SetFillColor(255, 255, 153)

		pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
		pdf.SetX(lineX1)
		pdf.SetY(lineY1 - 10)
		pdf.Cell(nil, "5.1　お見積金額（ランニング費用）")

		pdf.SetLineWidth(1.0)
		pdf.Line(lineX1, lineY1, lineX1, lineY1+lineH*float64(lenInfoRunning))             //左
		pdf.Line(171, lineY1, 171, lineY1+lineH*float64(lenInfoRunning))                   //左2
		pdf.Line(420, lineY1, 420, lineY1+lineH*float64(lenInfoRunning))                   //左3
		pdf.Line(510, lineY1, 510, lineY1+lineH*float64(lenInfoRunning))                   //左4
		pdf.Line(550, lineY1, 550, lineY1+lineH*float64(lenInfoRunning))                   //右3
		pdf.Line(640, lineY1, 640, lineY1+lineH*float64(lenInfoRunning))                   //右2
		pdf.Line(lineX1+lineW, lineY1, lineX1+lineW, lineY1+lineH*float64(lenInfoRunning)) //右

		for num := 0.0; num <= float64(lenInfoRunning); num++ {
			if num == 0 || num == float64(lenInfoRunning) {
				pdf.SetLineWidth(1.0)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)
				pdf.SetTextColor(0, 0, 0)
				pdf.SetX(lineX1 + 1)
				pdf.SetY(lineY1 + 1)
				pdf.Cell(nil, "No.")
				pdf.SetX(270)
				pdf.Cell(nil, "項　目")
				pdf.SetX(450)
				pdf.Cell(nil, "単　価")
				pdf.SetX(515)
				pdf.Cell(nil, "数　量")
				pdf.SetX(580)
				pdf.Cell(nil, "金　額")
				pdf.SetX(690)
				pdf.Cell(nil, "備　考")
			} else if num == 1 {
				pdf.SetLineWidth(1.0)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(156)
				pdf.SetY(lineY1 + lineH*num + 1)
				// pdf.Cell(nil, strconv.FormatFloat((float64(num)-1), 'g', -1, 32))
				pdf.Cell(nil, "1")
				pdf.SetX(172)
				pdf.Cell(nil, infoRunning[0].ProductName)
				pdf.SetX(510 - float64(len(infoRunning[0].Price))*7)
				pdf.Cell(nil, convertStr(infoRunning[0].Price))
				pdf.SetX(516)
				pdf.Cell(nil, convertStr(infoRunning[0].Quantity))
				pdf.SetX(640 - float64(len(infoRunning[0].SubTotal))*7)
				pdf.Cell(nil, convertStr(infoRunning[0].SubTotal))
				pdf.SetX(641)
				pdf.Cell(nil, " ")
			} else if num == float64(lenInfoRunning)-1 {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)
				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定

				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "合　計")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumRunning*110/100)))*7)
				pdf.Cell(nil, convert(sumRunning*110/100))
			} else if num == float64(lenInfoRunning)-2 {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "消費税")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumRunning/10)))*7)
				pdf.Cell(nil, convert(sumRunning/10))
				pdf.SetX(641)
				pdf.Cell(nil, "10%")
			} else if num == float64(lenInfoRunning)-3 {
				pdf.SetLineWidth(1.5)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.SetX(388)
				pdf.Cell(nil, "小　計")
				pdf.SetX(640 - float64(len(strconv.Itoa(sumRunning)))*7)
				pdf.Cell(nil, convert(sumRunning))
			} else {
				pdf.SetLineWidth(0.7)
				pdf.Line(lineX1, lineY1+lineH*num, lineX1+lineW, lineY1+lineH*num)

				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(156)
				pdf.SetY(lineY1 + lineH*num + 1)
				pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
				pdf.SetX(172)
				pdf.Cell(nil, infoRunning[int(num)-1].ProductName)
				pdf.SetX(510 - float64(len(infoRunning[int(num)-1].Price))*7)
				pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Price))
				pdf.SetX(516)
				pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Quantity))
				pdf.SetX(640 - float64(len(infoRunning[int(num)-1].SubTotal))*7)
				pdf.Cell(nil, convertStr(infoRunning[int(num)-1].SubTotal))
				pdf.SetX(641)
				pdf.Cell(nil, " ")
			}
		}

		//补足部分
		lineY1 += lineH*float64(lenInfoRunning) + 1
		arr_str := strings.Split(info.Supplement, "\n")
		if len(info.Supplement) != 0 {
			pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
			pdf.SetX(lineX1)
			pdf.SetY(lineY1)
			pdf.Cell(nil, "【補　足】")
			for index, str := range arr_str {
				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(183)
				pdf.SetY(lineY1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}

		//5.3见积书部分
		if len(info.Other) != 0 {
			lineY1 += float64((len(arr_str)+1)*10) + 10
			pdf.SetFont("ipaexm", "b", 10) //フォント、文字サイズ指定
			pdf.SetX(lineX1)
			pdf.SetY(lineY1)
			pdf.Cell(nil, "5.2　その他費用")
			arr_str = strings.Split(info.Other, "\n")
			for index, str := range arr_str {
				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(183)
				pdf.SetY(lineY1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}

		lineY1 += float64((len(arr_str)+1)*10) + 15
		PaymentConditions(pdf, lineY1, info)

	}

}

//支付条件
func PaymentConditions(pdf *gopdf.GoPdf, lineY1 float64, info model.Estimate) {
	err := pdf.AddTTFFont("PingBold", FONTPATH+"PingBold.ttf")
	if err != nil {
		panic(err)
	}
	err = pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("PingBold", "", 12) //フォント、文字サイズ指定

	pdf.SetX(130)
	pdf.SetY(lineY1)
	pdf.Cell(nil, "６．支払条件")

	arr_str := strings.Split(info.PaymentConditions, "\n")
	for index, str := range arr_str {
		pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
		pdf.SetX(153)
		pdf.SetY(lineY1 + 15 + float64(index*10))
		pdf.Cell(nil, str)
	}

}

func convert(integer int) string {
	arr := strings.Split(fmt.Sprintf("%d", integer), "")
	cnt := len(arr) - 1
	res := ""
	i2 := 0
	for i := cnt; i >= 0; i-- {
		if i2 > 2 && i2%3 == 0 {
			res = fmt.Sprintf(",%s", res)
		}
		res = fmt.Sprintf("%s%s", arr[i], res)
		i2++
	}
	return res
}

func convertStr(str string) string {
	arr := strings.Split(str, "")
	cnt := len(arr) - 1
	res := ""
	i2 := 0
	for i := cnt; i >= 0; i-- {
		if i2 > 2 && i2%3 == 0 {
			res = fmt.Sprintf(",%s", res)
		}
		res = fmt.Sprintf("%s%s", arr[i], res)
		i2++
	}
	return res
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
