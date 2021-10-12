package utils

import (
	"fmt"
	"io/ioutil"
	"main/model"
	"os"
	"strconv"
	"strings"
	"time"

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
	// pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 850.32, H: 1203.12}}) //595.28, 841.89 = A4
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
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
	t1 := Header{
		x1: 192,
		y1: 48,
		w:  218,
		h:  35,
	}
	//内边框
	t2 := Header{
		x1: t1.x1 + 2.5,
		y1: t1.y1 + 2.5,
		w:  t1.w - 5,
		h:  t1.h - 5,
	}
	err := pdf.AddTTFFont("mincho", FONTPATH+"ShipporiAntiqueB1-Regular.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("mincho", "", 26) //フォント、文字サイズ指定
	pdf.SetX(212)                 //x座標指定
	pdf.SetY(t1.y1 + 5)           //y座標指定
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
	cu := Header{
		x1: 70,
		y1: 120,
	}
	//①得意先名〇〇
	pdf.SetFont("mincho", "", 14) //フォント、文字サイズ指定
	customerName := info.CustomerName
	pdf.SetX(cu.x1) //x座標指定
	pdf.SetY(cu.y1) //y座標指定
	pdf.Cell(nil, customerName)
	//〇〇御中
	pdf.SetX(cu.x1 + 10 + float64(len(customerName))*5) //x座標指定
	pdf.SetY(cu.y1)                                     //y座標指定
	pdf.Cell(nil, "御中")
	pdf.SetLineWidth(0.7)
	pdf.Line(cu.x1, cu.y1+15, cu.x1+float64(len(customerName)+6)*5+10, cu.y1+15)

	//②挨拶
	//err = pdf.AddTTFFont("simfang", "./font/simfang.ttf")
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
	pdf.SetX(cu.x1)               //x座標指定
	pdf.SetY(cu.y1 + 29)          //y座標指定
	pdf.Cell(nil, "拝啓 貴社御依頼に対し、下記の通り御見積り申し上げます。")

	pdf.SetX(cu.x1)      //x座標指定
	pdf.SetY(cu.y1 + 41) //y座標指定
	pdf.Cell(nil, "何卒ご用命のほど、よろしくお願い申し上げます。")
}

func Company(pdf *gopdf.GoPdf, info model.Estimate) {

	co := Header{
		x1: 360,
		y1: 149,
	}

	err := pdf.AddTTFFont("simfang", FONTPATH+"simfang.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("simfang", "", 9) //フォント、文字サイズ指定

	pdf.Image(IMGPATH+"bridge_logo.png", co.x1+10, co.y1-20, nil)

	pdf.SetX(co.x1)
	pdf.SetY(co.y1)
	pdf.Cell(nil, "〒104-0032"+"　　"+"株式会社ブリッジ") //邮编
	pdf.SetX(co.x1)
	pdf.SetY(co.y1 + 15)
	pdf.Cell(nil, "　東京都中央区八丁堀4丁目11-10第2SSビル 1F") //地址
	pdf.SetX(co.x1)
	pdf.SetY(co.y1 + 30)
	pdf.Cell(nil, "　Tel:03-6222-3222　Fax:03-6222-3228") //联系方式
	pdf.SetX(co.x1)
	pdf.SetY(co.y1 + 45)
	pdf.Cell(nil, "※有効期限:30日"+"　　　　　"+info.CreatedBy) //作成者

}

func BodyTitle(pdf *gopdf.GoPdf) {
	bo := Header{
		x1: 85,
		y1: 220,
	}
	err := pdf.AddTTFFont("PingBold", FONTPATH+"PingBold.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("PingBold", "", 11) //フォント、文字サイズ指定

	pdf.SetX(bo.x1)
	pdf.SetY(bo.y1)
	pdf.Cell(nil, "１. 見積対象案件名")
	pdf.SetX(bo.x1)
	pdf.SetY(bo.y1 + 45) //285
	pdf.Cell(nil, "２．作業内容")
	pdf.SetX(bo.x1)
	pdf.SetY(bo.y1 + 110) //350
	pdf.Cell(nil, "３．成果物")
	pdf.SetX(bo.x1)
	pdf.SetY(bo.y1 + 200) //440
	pdf.Cell(nil, "４．作業場所")
	pdf.SetX(bo.x1)
	pdf.SetY(bo.y1 + 245) //485
	pdf.Cell(nil, "５．お見積金額")
}

//见积对象案件名
func EstimateName(pdf *gopdf.GoPdf, info model.Estimate) {
	es := Header{
		x1: 100,
		y1: 235,
	}
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}

	pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
	pdf.SetX(es.x1)
	pdf.SetY(es.y1)
	pdf.Cell(nil, info.EstimateName)
}

//作业内容
func Work(pdf *gopdf.GoPdf, info model.Estimate) {
	wo := Header{
		x1: 100,
		y1: 280,
	}
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}

	pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定

	if len(info.Work1) != 0 && len(info.Work2) != 0 && len(info.Work3) != 0 {
		pdf.SetX(wo.x1)
		pdf.SetY(wo.y1)
		pdf.Cell(nil, "・"+info.Work1)
		pdf.SetX(wo.x1)
		pdf.SetY(wo.y1 + 13)
		pdf.Cell(nil, "・"+info.Work2)
		pdf.SetX(wo.x1)
		pdf.SetY(wo.y1 + 23)
		pdf.Cell(nil, "・"+info.Work3)
	} else if len(info.Work1) != 0 && len(info.Work2) != 0 {
		pdf.SetX(wo.x1)
		pdf.SetY(wo.y1)
		pdf.Cell(nil, "・"+info.Work1)
		pdf.SetX(wo.x1)
		pdf.SetY(wo.y1 + 13)
		pdf.Cell(nil, "・"+info.Work2)
	} else {
		pdf.SetX(wo.x1)
		pdf.SetY(wo.y1)
		pdf.Cell(nil, "・"+info.Work1)
	}

}

//成果物
func Deliverables(pdf *gopdf.GoPdf, info model.Estimate) {
	de := Header{
		x1: 100,
		y1: 355,
		w:  445,
		h:  11,
	}

	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
	pdf.SetX(de.x1)
	pdf.SetY(de.y1 - 10) //365
	pdf.Cell(nil, "成果物及び納品予定日は、以下の通りです。")

	//追加背景颜色
	pdf.SetFillColor(255, 255, 153)
	pdf.RectFromUpperLeftWithStyle(de.x1, de.y1, de.w, de.h, "FD")
	pdf.SetFillColor(255, 255, 153)

	pdf.SetLineWidth(0.7)
	pdf.Line(de.x1, de.y1, de.x1, de.y1+de.h*4)           //左
	pdf.Line(340, de.y1, 340, de.y1+de.h*4)               //左2
	pdf.Line(400, de.y1, 400, de.y1+de.h*4)               //左3
	pdf.Line(440, de.y1, 440, de.y1+de.h*4)               //右2
	pdf.Line(de.x1+de.w, de.y1, de.x1+de.w, de.y1+de.h*4) //右

	for num := 0.0; num <= 4; num++ {
		pdf.Line(de.x1, de.y1+de.h*num, de.x1+de.w, de.y1+de.h*num)
	}
	pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(200)
	pdf.SetY(de.y1 + 1)
	pdf.Cell(nil, "成　果　物")
	pdf.SetX(356)
	pdf.Cell(nil, "納品媒体")
	pdf.SetX(410)
	pdf.Cell(nil, "部数")
	pdf.SetX(470)
	pdf.Cell(nil, "納品予定日")

	pdf.SetX(de.x1 + 1)
	pdf.SetY(de.y1 + 13)
	pdf.Cell(nil, info.Deliverables1)
	pdf.SetX(341)
	pdf.Cell(nil, info.Media1)
	pdf.SetX(420)
	pdf.Cell(nil, info.Quantity1)
	pdf.SetX(441)
	pdf.Cell(nil, info.DeliveryDate1)

	pdf.SetX(de.x1 + 1)
	pdf.SetY(de.y1 + 25)
	pdf.Cell(nil, info.Deliverables2)
	pdf.SetX(341)
	pdf.Cell(nil, info.Media2)
	pdf.SetX(420)
	pdf.Cell(nil, info.Quantity2)
	pdf.SetX(441)
	pdf.Cell(nil, info.DeliveryDate2)

	pdf.SetX(de.x1 + 1)
	pdf.SetY(de.y1 + 36)
	pdf.Cell(nil, info.Deliverables3)
	pdf.SetX(341)
	pdf.Cell(nil, info.Media3)
	pdf.SetX(420)
	pdf.Cell(nil, info.Quantity3)
	pdf.SetX(441)
	pdf.Cell(nil, info.DeliveryDate3)
}

//作业场所
func WorkSpace(pdf *gopdf.GoPdf, info model.Estimate) {
	ws := Header{
		x1: 100,
		y1: 435,
	}
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}
	pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
	pdf.SetX(ws.x1)
	pdf.SetY(ws.y1)
	pdf.Cell(nil, info.WorkSpace)
}

//见积详细
func EstimateDetail(pdf *gopdf.GoPdf, info model.Estimate, infoInitial []model.EstimateDetail, infoRunning []model.EstimateDetail) {
	ed := Header{
		x1: 100,
		y1: 490,
		w:  445,
		h:  11,
	}
	so := 5.6
	err := pdf.AddTTFFont("ipaexm", FONTPATH+"ipaexm.ttf")
	if err != nil {
		panic(err)
	}
	err = pdf.AddTTFFont("simfang", FONTPATH+"simfang.ttf")
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
		pdf.RectFromUpperLeftWithStyle(ed.x1, ed.y1, ed.w, ed.h, "FD")
		pdf.SetFillColor(255, 255, 153)

		pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
		pdf.SetX(ed.x1)
		pdf.SetY(ed.y1 - 10)
		pdf.Cell(nil, "5.1　お見積金額（イニシャル費用）")

		pdf.SetLineWidth(1.0)
		pdf.Line(ed.x1, ed.y1, ed.x1, ed.y1+ed.h*float64(lenInfoInitial))           //左
		pdf.Line(120, ed.y1, 120, ed.y1+ed.h*float64(lenInfoInitial))               //左2
		pdf.Line(280, ed.y1, 280, ed.y1+ed.h*float64(lenInfoInitial))               //左3
		pdf.Line(340, ed.y1, 340, ed.y1+ed.h*float64(lenInfoInitial))               //左4
		pdf.Line(380, ed.y1, 380, ed.y1+ed.h*float64(lenInfoInitial))               //右3
		pdf.Line(440, ed.y1, 440, ed.y1+ed.h*float64(lenInfoInitial))               //右2
		pdf.Line(ed.x1+ed.w, ed.y1, ed.x1+ed.w, ed.y1+ed.h*float64(lenInfoInitial)) //右

		for num := 0.0; num <= float64(lenInfoInitial); num++ {
			if num == 0 || num == float64(lenInfoInitial) {
				pdf.SetLineWidth(1.0)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
				pdf.SetTextColor(0, 0, 0)
				pdf.SetX(ed.x1 + 1)
				pdf.SetY(ed.y1 + 1)
				pdf.Cell(nil, "No.")
				pdf.SetX(170)
				pdf.Cell(nil, "項　目")
				pdf.SetX(300)
				pdf.Cell(nil, "単　価")
				pdf.SetX(350)
				pdf.Cell(nil, "数　量")
				pdf.SetX(400)
				pdf.Cell(nil, "金　額")
				pdf.SetX(480)
				pdf.Cell(nil, "備　考")
			} else if num == 1 {
				pdf.SetLineWidth(1.0)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
				pdf.SetX(ed.x1 + 1)
				pdf.SetY(ed.y1 + ed.h*num + 1)
				// pdf.Cell(nil, strconv.FormatFloat((float64(num)-1), 'g', -1, 32))
				pdf.Cell(nil, "1")
				pdf.SetX(121)
				pdf.Cell(nil, infoInitial[0].ProductName)
				pdf.SetX(340 - float64(len(infoInitial[0].Price))*so)
				pdf.Cell(nil, convertStr(infoInitial[0].Price))
				pdf.SetX(360)
				pdf.Cell(nil, convertStr(infoInitial[0].Quantity))
				pdf.SetX(440 - float64(len(infoInitial[0].SubTotal))*so)
				pdf.Cell(nil, convertStr(infoInitial[0].SubTotal))
				pdf.SetX(441)
				pdf.Cell(nil, " ")
			} else if num == float64(lenInfoInitial)-1 {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定

				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "合　計")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumInitial*110/100)))*so)
				pdf.Cell(nil, convert(sumInitial*110/100))
			} else if num == float64(lenInfoInitial)-2 {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定
				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "消費税")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumInitial/10)))*so)
				pdf.Cell(nil, convert(sumInitial/10))
				pdf.SetX(441)
				pdf.Cell(nil, "10%")
			} else if num == float64(lenInfoInitial)-3 {
				pdf.SetLineWidth(1.5)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定
				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "小　計")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumInitial)))*so)
				pdf.Cell(nil, convert(sumInitial))
			} else {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				//处理值引颜色问题,以及缩小金额的缩紧 7>6.5
				price, _ := strconv.Atoi(infoInitial[int(num)-1].Price)
				if price < 0 {
					pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
					pdf.SetTextColor(255, 0, 0)
					pdf.SetX(ed.x1 + 1)
					pdf.SetY(ed.y1 + ed.h*num + 1)
					pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
					pdf.SetX(121)
					pdf.Cell(nil, infoInitial[int(num)-1].ProductName)
					// pdf.SetX(510 - float64(len(infoInitial[int(num)-1].Price))*7)
					// pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Price))
					// pdf.SetX(516)
					// pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Quantity))
					pdf.SetX(440 - float64(len(infoInitial[int(num)-1].SubTotal))*(so-0.5))
					pdf.Cell(nil, convertStr(infoInitial[int(num)-1].SubTotal))
					pdf.SetX(441)
					pdf.Cell(nil, " ")
					pdf.SetTextColor(0, 0, 0)
				} else {
					pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
					pdf.SetX(ed.x1 + 1)
					pdf.SetY(ed.y1 + ed.h*num + 1)
					pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
					pdf.SetX(121)
					pdf.Cell(nil, infoInitial[int(num)-1].ProductName)
					pdf.SetX(340 - float64(len(infoInitial[int(num)-1].Price))*so)
					pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Price))
					pdf.SetX(360)
					pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Quantity))
					pdf.SetX(440 - float64(len(infoInitial[int(num)-1].SubTotal))*so)
					pdf.Cell(nil, convertStr(infoInitial[int(num)-1].SubTotal))
					pdf.SetX(441)
					pdf.Cell(nil, " ")
				}

			}
		}

		lenInfoRunning := len(infoRunning) + 4

		ed.y1 = ed.y1 + ed.h*float64(lenInfoInitial) + 15

		//追加背景颜色
		pdf.SetFillColor(255, 255, 153)
		pdf.RectFromUpperLeftWithStyle(ed.x1, ed.y1, ed.w, ed.h, "FD")
		pdf.SetFillColor(255, 255, 153)

		pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
		pdf.SetX(ed.x1)
		pdf.SetY(ed.y1 - 10)
		pdf.Cell(nil, "5.2　お見積金額（ランニング費用）")

		pdf.SetLineWidth(1.0)
		pdf.Line(ed.x1, ed.y1, ed.x1, ed.y1+ed.h*float64(lenInfoRunning))           //左
		pdf.Line(120, ed.y1, 120, ed.y1+ed.h*float64(lenInfoRunning))               //左2
		pdf.Line(280, ed.y1, 280, ed.y1+ed.h*float64(lenInfoRunning))               //左3
		pdf.Line(340, ed.y1, 340, ed.y1+ed.h*float64(lenInfoRunning))               //左4
		pdf.Line(380, ed.y1, 380, ed.y1+ed.h*float64(lenInfoRunning))               //右3
		pdf.Line(440, ed.y1, 440, ed.y1+ed.h*float64(lenInfoRunning))               //右2
		pdf.Line(ed.x1+ed.w, ed.y1, ed.x1+ed.w, ed.y1+ed.h*float64(lenInfoRunning)) //右

		for num := 0.0; num <= float64(lenInfoRunning); num++ {
			if num == 0 || num == float64(lenInfoRunning) {
				pdf.SetLineWidth(1.0)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
				pdf.SetTextColor(0, 0, 0)
				pdf.SetX(ed.x1 + 1)
				pdf.SetY(ed.y1 + 1)
				pdf.Cell(nil, "No.")
				pdf.SetX(170)
				pdf.Cell(nil, "項　目")
				pdf.SetX(300)
				pdf.Cell(nil, "単価")
				pdf.SetX(350)
				pdf.Cell(nil, "数量")
				pdf.SetX(400)
				pdf.Cell(nil, "金額")
				pdf.SetX(480)
				pdf.Cell(nil, "備　考")
			} else if num == 1 {
				pdf.SetLineWidth(1.0)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
				pdf.SetX(ed.x1 + 1)
				pdf.SetY(ed.y1 + ed.h*num + 1)
				// pdf.Cell(nil, strconv.FormatFloat((float64(num)-1), 'g', -1, 32))
				pdf.Cell(nil, "1")
				pdf.SetX(121)
				pdf.Cell(nil, infoRunning[0].ProductName)
				pdf.SetX(340 - float64(len(infoRunning[0].Price))*so)
				pdf.Cell(nil, convertStr(infoRunning[0].Price))
				pdf.SetX(360)
				pdf.Cell(nil, convertStr(infoRunning[0].Quantity))
				pdf.SetX(440 - float64(len(infoRunning[0].SubTotal))*so)
				pdf.Cell(nil, convertStr(infoRunning[0].SubTotal))
				pdf.SetX(441)
				pdf.Cell(nil, " ")
			} else if num == float64(lenInfoRunning)-1 {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定

				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "合　計")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumRunning*110/100)))*so)
				pdf.Cell(nil, convert(sumRunning*110/100))
			} else if num == float64(lenInfoRunning)-2 {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定
				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "消費税")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumRunning/10)))*so)
				pdf.Cell(nil, convert(sumRunning/10))
				pdf.SetX(441)
				pdf.Cell(nil, "10%")
			} else if num == float64(lenInfoRunning)-3 {
				pdf.SetLineWidth(1.5)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定
				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "小　計")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumRunning)))*so)
				pdf.Cell(nil, convert(sumRunning))
			} else {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				price, _ := strconv.Atoi(infoRunning[int(num)-1].Price)
				if price < 0 {
					pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
					pdf.SetTextColor(255, 0, 0)
					pdf.SetX(ed.x1 + 1)
					pdf.SetY(ed.y1 + ed.h*num + 1)
					pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
					pdf.SetX(121)
					pdf.Cell(nil, infoRunning[int(num)-1].ProductName)
					// pdf.SetX(510 - float64(len(infoRunning[int(num)-1].Price))*7)
					// pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Price))
					// pdf.SetX(516)
					// pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Quantity))
					pdf.SetX(440 - float64(len(infoRunning[int(num)-1].SubTotal))*(so-0.5))
					pdf.Cell(nil, convertStr(infoRunning[int(num)-1].SubTotal))
					pdf.SetX(441)
					pdf.Cell(nil, " ")
					pdf.SetTextColor(0, 0, 0)
				} else {
					pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
					pdf.SetX(ed.x1 + 1)
					pdf.SetY(ed.y1 + ed.h*num + 1)
					pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
					pdf.SetX(121)
					pdf.Cell(nil, infoRunning[int(num)-1].ProductName)
					pdf.SetX(340 - float64(len(infoRunning[int(num)-1].Price))*so)
					pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Price))
					pdf.SetX(360)
					pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Quantity))
					pdf.SetX(440 - float64(len(infoRunning[int(num)-1].SubTotal))*so)
					pdf.Cell(nil, convertStr(infoRunning[int(num)-1].SubTotal))
					pdf.SetX(441)
					pdf.Cell(nil, " ")
				}

			}
		}

		//补足部分
		ed.y1 += ed.h*float64(lenInfoRunning) + 1
		arr_str := strings.Split(info.Supplement, "\n")
		if len(info.Supplement) != 0 {
			pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
			pdf.SetX(ed.x1)
			pdf.SetY(ed.y1)
			pdf.Cell(nil, "【補　足】")
			for index, str := range arr_str {
				pdf.SetFont("simfang", "", 10) //フォント、文字サイズ指定
				pdf.SetX(ed.x1 + 30)
				pdf.SetY(ed.y1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}

		//5.3见积书部分

		if len(info.Other) != 0 {
			ed.y1 += float64((len(arr_str)+1)*10) + 10
			pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
			pdf.SetX(ed.x1)
			pdf.SetY(ed.y1)
			pdf.Cell(nil, "5.3　その他費用")
			arr_str = strings.Split(info.Other, "\n")
			for index, str := range arr_str {
				pdf.SetFont("simfang", "", 10) //フォント、文字サイズ指定
				pdf.SetX(ed.x1 + 30)
				pdf.SetY(ed.y1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}

		ed.y1 += float64((len(arr_str)+1)*10) + 15
		PaymentConditions(pdf, ed.y1, info)
	} else if len(infoInitial) != 0 {
		//通过数据个数计算表的列数 数据个数 + 表头一行 + 末尾三行
		lenInfoInitial := len(infoInitial) + 4

		//追加背景颜色
		pdf.SetFillColor(255, 255, 153)
		pdf.RectFromUpperLeftWithStyle(ed.x1, ed.y1, ed.w, ed.h, "FD")
		pdf.SetFillColor(255, 255, 153)

		pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
		pdf.SetX(ed.x1)
		pdf.SetY(ed.y1 - 10)
		pdf.Cell(nil, "5.1　お見積金額（イニシャル費用）")

		pdf.SetLineWidth(1.0)
		pdf.Line(ed.x1, ed.y1, ed.x1, ed.y1+ed.h*float64(lenInfoInitial))           //左
		pdf.Line(120, ed.y1, 120, ed.y1+ed.h*float64(lenInfoInitial))               //左2
		pdf.Line(280, ed.y1, 280, ed.y1+ed.h*float64(lenInfoInitial))               //左3
		pdf.Line(340, ed.y1, 340, ed.y1+ed.h*float64(lenInfoInitial))               //左4
		pdf.Line(380, ed.y1, 380, ed.y1+ed.h*float64(lenInfoInitial))               //右3
		pdf.Line(440, ed.y1, 440, ed.y1+ed.h*float64(lenInfoInitial))               //右2
		pdf.Line(ed.x1+ed.w, ed.y1, ed.x1+ed.w, ed.y1+ed.h*float64(lenInfoInitial)) //右

		for num := 0.0; num <= float64(lenInfoInitial); num++ {
			if num == 0 || num == float64(lenInfoInitial) {
				pdf.SetLineWidth(1.0)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
				pdf.SetTextColor(0, 0, 0)
				pdf.SetX(ed.x1 + 1)
				pdf.SetY(ed.y1 + 1)
				pdf.Cell(nil, "No.")
				pdf.SetX(170)
				pdf.Cell(nil, "項　目")
				pdf.SetX(300)
				pdf.Cell(nil, "単　価")
				pdf.SetX(350)
				pdf.Cell(nil, "数　量")
				pdf.SetX(400)
				pdf.Cell(nil, "金　額")
				pdf.SetX(480)
				pdf.Cell(nil, "備　考")
			} else if num == 1 {
				pdf.SetLineWidth(1.0)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
				pdf.SetX(ed.x1 + 1)
				pdf.SetY(ed.y1 + ed.h*num + 1)
				// pdf.Cell(nil, strconv.FormatFloat((float64(num)-1), 'g', -1, 32))
				pdf.Cell(nil, "1")
				pdf.SetX(121)
				pdf.Cell(nil, infoInitial[0].ProductName)
				pdf.SetX(340 - float64(len(infoInitial[0].Price))*so)
				pdf.Cell(nil, convertStr(infoInitial[0].Price))
				pdf.SetX(360)
				pdf.Cell(nil, convertStr(infoInitial[0].Quantity))
				pdf.SetX(440 - float64(len(infoInitial[0].SubTotal))*so)
				pdf.Cell(nil, convertStr(infoInitial[0].SubTotal))
				pdf.SetX(441)
				pdf.Cell(nil, " ")
			} else if num == float64(lenInfoInitial)-1 {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定

				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "合　計")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumInitial*110/100)))*so)
				pdf.Cell(nil, convert(sumInitial*110/100))
			} else if num == float64(lenInfoInitial)-2 {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定
				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "消費税")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumInitial/10)))*so)
				pdf.Cell(nil, convert(sumInitial/10))
				pdf.SetX(441)
				pdf.Cell(nil, "10%")
			} else if num == float64(lenInfoInitial)-3 {
				pdf.SetLineWidth(1.5)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定
				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "小　計")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumInitial)))*so)
				pdf.Cell(nil, convert(sumInitial))
			} else {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				//处理值引颜色问题,以及缩小金额的缩紧 7>6.5
				price, _ := strconv.Atoi(infoInitial[int(num)-1].Price)
				if price < 0 {
					pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
					pdf.SetTextColor(255, 0, 0)
					pdf.SetX(ed.x1 + 1)
					pdf.SetY(ed.y1 + ed.h*num + 1)
					pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
					pdf.SetX(121)
					pdf.Cell(nil, infoInitial[int(num)-1].ProductName)
					// pdf.SetX(510 - float64(len(infoInitial[int(num)-1].Price))*7)
					// pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Price))
					// pdf.SetX(516)
					// pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Quantity))
					pdf.SetX(440 - float64(len(infoInitial[int(num)-1].SubTotal))*(so-0.5))
					pdf.Cell(nil, convertStr(infoInitial[int(num)-1].SubTotal))
					pdf.SetX(441)
					pdf.Cell(nil, " ")
					pdf.SetTextColor(0, 0, 0)
				} else {
					pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
					pdf.SetX(ed.x1 + 1)
					pdf.SetY(ed.y1 + ed.h*num + 1)
					pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
					pdf.SetX(121)
					pdf.Cell(nil, infoInitial[int(num)-1].ProductName)
					pdf.SetX(340 - float64(len(infoInitial[int(num)-1].Price))*so)
					pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Price))
					pdf.SetX(360)
					pdf.Cell(nil, convertStr(infoInitial[int(num)-1].Quantity))
					pdf.SetX(440 - float64(len(infoInitial[int(num)-1].SubTotal))*so)
					pdf.Cell(nil, convertStr(infoInitial[int(num)-1].SubTotal))
					pdf.SetX(441)
					pdf.Cell(nil, " ")
				}

			}
		}
		//补足部分
		ed.y1 += ed.h*float64(lenInfoInitial) + 1
		arr_str := strings.Split(info.Supplement, "\n")

		if len(info.Supplement) != 0 {
			pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
			pdf.SetX(ed.x1)
			pdf.SetY(ed.y1)
			pdf.Cell(nil, "【補　足】")
			for index, str := range arr_str {
				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(ed.x1 + 30)
				pdf.SetY(ed.y1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}

		//5.3见积书部分
		if len(info.Other) != 0 {
			ed.y1 += float64((len(arr_str)+1)*10) + 10
			pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
			pdf.SetX(ed.x1)
			pdf.SetY(ed.y1)
			pdf.Cell(nil, "5.2　その他費用")
			arr_str = strings.Split(info.Other, "\n")
			for index, str := range arr_str {
				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(ed.x1 + 30)
				pdf.SetY(ed.y1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}
		ed.y1 += float64((len(arr_str)+1)*10) + 15
		PaymentConditions(pdf, ed.y1, info)

	} else {
		//通过数据个数计算表的列数 数据个数 + 表头一行 + 末尾三行

		lenInfoRunning := len(infoRunning) + 4

		//追加背景颜色
		pdf.SetFillColor(255, 255, 153)
		pdf.RectFromUpperLeftWithStyle(ed.x1, ed.y1, ed.w, ed.h, "FD")
		pdf.SetFillColor(255, 255, 153)

		pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
		pdf.SetX(ed.x1)
		pdf.SetY(ed.y1 - 10)
		pdf.Cell(nil, "5.2　お見積金額（ランニング費用）")

		pdf.SetLineWidth(1.0)
		pdf.Line(ed.x1, ed.y1, ed.x1, ed.y1+ed.h*float64(lenInfoRunning))           //左
		pdf.Line(120, ed.y1, 120, ed.y1+ed.h*float64(lenInfoRunning))               //左2
		pdf.Line(280, ed.y1, 280, ed.y1+ed.h*float64(lenInfoRunning))               //左3
		pdf.Line(340, ed.y1, 340, ed.y1+ed.h*float64(lenInfoRunning))               //左4
		pdf.Line(380, ed.y1, 380, ed.y1+ed.h*float64(lenInfoRunning))               //右3
		pdf.Line(440, ed.y1, 440, ed.y1+ed.h*float64(lenInfoRunning))               //右2
		pdf.Line(ed.x1+ed.w, ed.y1, ed.x1+ed.w, ed.y1+ed.h*float64(lenInfoRunning)) //右

		for num := 0.0; num <= float64(lenInfoRunning); num++ {
			if num == 0 || num == float64(lenInfoRunning) {
				pdf.SetLineWidth(1.0)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
				pdf.SetTextColor(0, 0, 0)
				pdf.SetX(ed.x1 + 1)
				pdf.SetY(ed.y1 + 1)
				pdf.Cell(nil, "No.")
				pdf.SetX(170)
				pdf.Cell(nil, "項　目")
				pdf.SetX(300)
				pdf.Cell(nil, "単価")
				pdf.SetX(350)
				pdf.Cell(nil, "数量")
				pdf.SetX(400)
				pdf.Cell(nil, "金額")
				pdf.SetX(480)
				pdf.Cell(nil, "備　考")
			} else if num == 1 {
				pdf.SetLineWidth(1.0)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
				pdf.SetX(ed.x1 + 1)
				pdf.SetY(ed.y1 + ed.h*num + 1)
				// pdf.Cell(nil, strconv.FormatFloat((float64(num)-1), 'g', -1, 32))
				pdf.Cell(nil, "1")
				pdf.SetX(121)
				pdf.Cell(nil, infoRunning[0].ProductName)
				pdf.SetX(340 - float64(len(infoRunning[0].Price))*so)
				pdf.Cell(nil, convertStr(infoRunning[0].Price))
				pdf.SetX(360)
				pdf.Cell(nil, convertStr(infoRunning[0].Quantity))
				pdf.SetX(440 - float64(len(infoRunning[0].SubTotal))*so)
				pdf.Cell(nil, convertStr(infoRunning[0].SubTotal))
				pdf.SetX(441)
				pdf.Cell(nil, " ")
			} else if num == float64(lenInfoRunning)-1 {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定

				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "合　計")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumRunning*110/100)))*so)
				pdf.Cell(nil, convert(sumRunning*110/100))
			} else if num == float64(lenInfoRunning)-2 {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定
				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "消費税")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumRunning/10)))*so)
				pdf.Cell(nil, convert(sumRunning/10))
				pdf.SetX(441)
				pdf.Cell(nil, "10%")
			} else if num == float64(lenInfoRunning)-3 {
				pdf.SetLineWidth(1.5)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)

				pdf.SetFont("ipaexm", "b", 8) //フォント、文字サイズ指定
				pdf.SetY(ed.y1 + ed.h*num + 1)
				pdf.SetX(255)
				pdf.Cell(nil, "小　計")
				pdf.SetX(440 - float64(len(strconv.Itoa(sumRunning)))*so)
				pdf.Cell(nil, convert(sumRunning))
			} else {
				pdf.SetLineWidth(0.7)
				pdf.Line(ed.x1, ed.y1+ed.h*num, ed.x1+ed.w, ed.y1+ed.h*num)
				price, _ := strconv.Atoi(infoRunning[int(num)-1].Price)
				if price < 0 {
					pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
					pdf.SetTextColor(255, 0, 0)
					pdf.SetX(ed.x1 + 1)
					pdf.SetY(ed.y1 + ed.h*num + 1)
					pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
					pdf.SetX(121)
					pdf.Cell(nil, infoRunning[int(num)-1].ProductName)
					// pdf.SetX(510 - float64(len(infoRunning[int(num)-1].Price))*7)
					// pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Price))
					// pdf.SetX(516)
					// pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Quantity))
					pdf.SetX(440 - float64(len(infoRunning[int(num)-1].SubTotal))*(so-0.5))
					pdf.Cell(nil, convertStr(infoRunning[int(num)-1].SubTotal))
					pdf.SetX(441)
					pdf.Cell(nil, " ")
					pdf.SetTextColor(0, 0, 0)
				} else {
					pdf.SetFont("ipaexm", "", 8) //フォント、文字サイズ指定
					pdf.SetX(ed.x1 + 1)
					pdf.SetY(ed.y1 + ed.h*num + 1)
					pdf.Cell(nil, strconv.FormatFloat((float64(num)), 'g', -1, 32))
					pdf.SetX(121)
					pdf.Cell(nil, infoRunning[int(num)-1].ProductName)
					pdf.SetX(340 - float64(len(infoRunning[int(num)-1].Price))*so)
					pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Price))
					pdf.SetX(360)
					pdf.Cell(nil, convertStr(infoRunning[int(num)-1].Quantity))
					pdf.SetX(440 - float64(len(infoRunning[int(num)-1].SubTotal))*so)
					pdf.Cell(nil, convertStr(infoRunning[int(num)-1].SubTotal))
					pdf.SetX(441)
					pdf.Cell(nil, " ")
				}

			}
		}

		//补足部分
		ed.y1 += ed.h*float64(lenInfoRunning) + 1
		arr_str := strings.Split(info.Supplement, "\n")
		if len(info.Supplement) != 0 {
			pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
			pdf.SetX(ed.x1)
			pdf.SetY(ed.y1)
			pdf.Cell(nil, "【補　足】")
			for index, str := range arr_str {
				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(ed.x1 + 30)
				pdf.SetY(ed.y1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}

		//5.3见积书部分
		if len(info.Other) != 0 {
			ed.y1 += float64((len(arr_str)+1)*10) + 10
			pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
			pdf.SetX(ed.x1)
			pdf.SetY(ed.y1)
			pdf.Cell(nil, "5.2　その他費用")
			arr_str = strings.Split(info.Other, "\n")
			for index, str := range arr_str {
				pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
				pdf.SetX(ed.x1 + 30)
				pdf.SetY(ed.y1 + 10 + float64(index*10))
				pdf.Cell(nil, str)
			}
		}

		ed.y1 += float64((len(arr_str)+1)*10) + 15
		PaymentConditions(pdf, ed.y1, info)

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
	pdf.SetFont("PingBold", "", 11) //フォント、文字サイズ指定

	pdf.SetX(85)
	pdf.SetY(lineY1)
	pdf.Cell(nil, "６．支払条件")

	arr_str := strings.Split(info.PaymentConditions, "\n")
	for index, str := range arr_str {
		pdf.SetFont("ipaexm", "", 10) //フォント、文字サイズ指定
		pdf.SetX(100)
		pdf.SetY(lineY1 + 15 + float64(index*10))
		pdf.Cell(nil, str)
	}

}

func convert(integer int) string {
	res := ""
	if integer >= 0 {
		arr := strings.Split(fmt.Sprintf("%d", integer), "")
		cnt := len(arr) - 1
		i2 := 0
		for i := cnt; i >= 0; i-- {
			if i2 > 2 && i2%3 == 0 {
				res = fmt.Sprintf(",%s", res)
			}
			res = fmt.Sprintf("%s%s", arr[i], res)
			i2++
		}
	} else {
		integer = -integer
		arr := strings.Split(fmt.Sprintf("%d", integer), "")
		cnt := len(arr) - 1
		i2 := 0
		for i := cnt; i >= 0; i-- {
			if i2 > 2 && i2%3 == 0 {
				res = fmt.Sprintf(",%s", res)
			}
			res = fmt.Sprintf("%s%s", arr[i], res)
			i2++
		}
		res = "-" + res
	}
	return res
}

func convertStr(str string) string {
	res := ""
	if str[0:1] != "-" {
		arr := strings.Split(str, "")
		cnt := len(arr) - 1
		i2 := 0
		for i := cnt; i >= 0; i-- {
			if i2 > 2 && i2%3 == 0 {
				res = fmt.Sprintf(",%s", res)
			}
			res = fmt.Sprintf("%s%s", arr[i], res)
			i2++
		}
	} else {
		str = str[1:]
		arr := strings.Split(str, "")
		cnt := len(arr) - 1
		i2 := 0
		for i := cnt; i >= 0; i-- {
			if i2 > 2 && i2%3 == 0 {
				res = fmt.Sprintf(",%s", res)
			}
			res = fmt.Sprintf("%s%s", arr[i], res)
			i2++
		}
		res = "-" + res
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
	for x = 0; x < 850; x += 20 {
		pdf.Line(x, 0, x, 30)
		pdf.SetX(x)  //x座標指定
		pdf.SetY(10) //y座標指定
		s := fmt.Sprintf("%.0f", x)
		pdf.Cell(nil, s)
	}
	for y = 0; y <= 1200; y += 20 {
		pdf.Line(0, y, 30, y)
		pdf.SetX(10) //x座標指定
		pdf.SetY(y)  //y座標指定
		s := fmt.Sprintf("%.0f", y)
		pdf.Cell(nil, s)
	}
}
