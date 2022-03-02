package utils

import (
	"main/config"
	"main/model"
	"strings"
	"time"

	"github.com/signintech/gopdf"
)

func NewInvoiceOrderPdf(order []*model.Order) {

	//获取见积头数据
	var orderInfo model.Order
	for _, item := range order {
		orderInfo = *item
	}

	pdf := gopdf.GoPdf{}
	// pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 850.32, H: 1203.12}}) //595.28, 841.89 = A4
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	pdf.SetFillColor(0, 0, 0)

	err := pdf.AddTTFFont("Shippori Mincho", config.InitConfig().FontPath+"ShipporiMincho-Regular.ttf")
	if err != nil {
		panic(err)
	}
	err = pdf.AddTTFFont("Shippori Mincho B1", config.InitConfig().FontPath+"ShipporiMinchoB1-Bold.ttf")
	if err != nil {
		panic(err)
	}

	// drawGrid(&pdf)
	TitleInvoiceOrder(&pdf, orderInfo)
	CompanyInvoiceOrder(&pdf, orderInfo)
	BodyInvoiceOrder(&pdf, orderInfo)

	if orderInfo.OrderPdfNum != nil {
		pdf.WritePdf(config.InitConfig().FilePath + "/pdf/invoiceorder/" + (*orderInfo.InvoiceOrderPdfNum)[0:4] + "-" + (*orderInfo.InvoiceOrderPdfNum)[4:6] + "-" + (*orderInfo.InvoiceOrderPdfNum)[6:8] + "/" + *orderInfo.InvoiceOrderPdfNum + "_注文請書_" + *orderInfo.CustomerName + "様_" + *orderInfo.ProjectName + ".pdf")
	} else {
		now := time.Now().Format("2006-01-02")
		pdf.WritePdf(config.InitConfig().FilePath + "/pdf/invoiceorder/" + now + "/" + *orderInfo.InvoiceOrderPdfNum + "_注文請書_" + *orderInfo.CustomerName + "様_" + *orderInfo.ProjectName + ".pdf")
	}
}

func NewOrderPdf(order []*model.Order) {
	//获取见积头数据
	var orderInfo model.Order
	for _, item := range order {
		orderInfo = *item
	}

	pdf := gopdf.GoPdf{}
	// pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 850.32, H: 1203.12}}) //595.28, 841.89 = A4
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	pdf.SetFillColor(0, 0, 0)

	err := pdf.AddTTFFont("Shippori Mincho", config.InitConfig().FontPath+"ShipporiMincho-Regular.ttf")
	if err != nil {
		panic(err)
	}
	err = pdf.AddTTFFont("Shippori Mincho B1", config.InitConfig().FontPath+"ShipporiMinchoB1-Bold.ttf")
	if err != nil {
		panic(err)
	}
	TitleOrder(&pdf, orderInfo)
	CompanyOrder(&pdf, orderInfo)
	BodyOrder(&pdf, orderInfo)

	if orderInfo.InvoiceOrderPdfNum != nil {
		pdf.WritePdf(config.InitConfig().FilePath + "/pdf/order/" + (*orderInfo.OrderPdfNum)[0:4] + "-" + (*orderInfo.OrderPdfNum)[4:6] + "-" + (*orderInfo.OrderPdfNum)[6:8] + "/" + *orderInfo.OrderPdfNum + "_注文書_" + *orderInfo.CustomerName + "様_" + *orderInfo.ProjectName + ".pdf")
	} else {
		now := time.Now().Format("2006-01-02")
		pdf.WritePdf(config.InitConfig().FilePath + "/pdf/order/" + now + "/" + *orderInfo.OrderPdfNum + "_注文書_" + *orderInfo.CustomerName + "様_" + *orderInfo.ProjectName + ".pdf")
	}
}

func TitleInvoiceOrder(pdf *gopdf.GoPdf, info model.Order) {
	pdf.SetFont("Shippori Mincho", "", 26) //フォント、文字サイズ指定
	pdf.SetX(212)                          //x座標指定
	pdf.SetY(53)                           //y座標指定
	pdf.Cell(nil, "注　文　請　書")               //Rect, String
	pdf.SetLineWidth(0.7)

	pdf.SetFont("Shippori Mincho", "", 10) //フォント、文字サイズ指定
	pdf.SetX(405)                          //x座標指定
	pdf.SetY(90)                           //y座標指定
	pdf.Cell(nil, "注文請書No.")               //注文請書No.
	pdf.SetX(469)
	pdf.SetY(90)
	pdf.Cell(nil, *info.InvoiceOrderPdfNum)
	pdf.SetX(447)
	pdf.SetY(105)
	pdf.Cell(nil, info.InvoiceOrderDate.Format("2006年01月02日")) //16/09/2021

	cu := Header{
		x1: 70,
		y1: 120,
	}
	//①得意先名〇〇
	pdf.SetFont("Shippori Mincho", "", 14) //フォント、文字サイズ指定
	pdf.SetX(cu.x1)                        //x座標指定
	pdf.SetY(cu.y1)                        //y座標指定
	pdf.Cell(nil, *info.CustomerName)
	//〇〇御中
	pdf.SetX(cu.x1 + 10 + float64(len(*info.CustomerName))*5) //x座標指定
	pdf.SetY(cu.y1)                                           //y座標指定
	pdf.Cell(nil, "御中")
}

func TitleOrder(pdf *gopdf.GoPdf, info model.Order) {
	pdf.SetFont("Shippori Mincho", "", 26) //フォント、文字サイズ指定
	pdf.SetX(212)                          //x座標指定
	pdf.SetY(53)                           //y座標指定
	pdf.Cell(nil, "注　文　書")                 //Rect, String
	pdf.SetLineWidth(0.7)

	pdf.SetFont("Shippori Mincho", "", 10) //フォント、文字サイズ指定
	pdf.SetX(420)                          //x座標指定
	pdf.SetY(90)                           //y座標指定
	pdf.Cell(nil, "注文書No.")                //注文請書No.
	pdf.SetX(469)
	pdf.SetY(90)
	pdf.Cell(nil, *info.OrderPdfNum)
	pdf.SetX(470)
	pdf.SetY(105)
	pdf.Cell(nil, "年　月　日") // 年　　　月　　　日

	cu := Header{
		x1: 70,
		y1: 120,
	}
	//①得意先名〇〇
	pdf.SetFont("Shippori Mincho", "", 14) //フォント、文字サイズ指定
	pdf.SetX(cu.x1)                        //x座標指定
	pdf.SetY(cu.y1)                        //y座標指定
	pdf.Cell(nil, "株式会社ブリッジ")
	//〇〇御中
	pdf.SetX(cu.x1 + 10 + float64(len("株式会社ブリッジ"))*5) //x座標指定
	pdf.SetY(cu.y1)                                   //y座標指定
	pdf.Cell(nil, "御中")
}

func CompanyInvoiceOrder(pdf *gopdf.GoPdf, info model.Order) {

	co := Header{
		x1: 300,
		y1: 140,
		w:  445,
	}

	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定

	pdf.Image(config.InitConfig().ImgPath+"stamp.png", co.x1+140, co.y1, nil)

	pdf.SetX(co.x1)
	pdf.SetY(co.y1)
	pdf.Cell(nil, "（住所）東京都中央区八丁堀4丁目11-10") //地址
	pdf.SetX(co.x1 + 140)
	pdf.SetY(co.y1 + 15)
	pdf.Cell(nil, "第2SSビル 1F") //地址

	pdf.SetX(co.x1)
	pdf.SetY(co.y1 + 50)
	pdf.Cell(nil, "（社名）株式会社ブリッジ") //邮编

	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定

	if len(*info.EstimateName) > 75 {
		cutSize := CutStringAsSize(*info.EstimateName, 72)

		pdf.SetX(co.x1 - 220)
		pdf.SetY(co.y1 + 85)
		pdf.Cell(nil, "件名（業務名）："+(*info.EstimateName)[0:cutSize])
		pdf.SetX(co.x1 - 125)
		pdf.SetY(co.y1 + 99)
		pdf.Cell(nil, (*info.EstimateName)[cutSize:])
		pdf.SetLineWidth(1.5)
		pdf.Line(co.x1-220, co.y1+113, co.x1-220+co.w, co.y1+113)
	} else {
		pdf.SetX(co.x1 - 220)
		pdf.SetY(co.y1 + 85)
		pdf.Cell(nil, "件名（業務名）："+*info.EstimateName)
		pdf.SetLineWidth(1.5)
		pdf.Line(co.x1-220, co.y1+98, co.x1-220+co.w, co.y1+98)
	}

	pdf.SetX(co.x1 - 220)
	pdf.SetY(co.y1 + 120)
	pdf.Cell(nil, "標題の件につき"+"　"+*info.EstimateOfOrder)
}

func CompanyOrder(pdf *gopdf.GoPdf, info model.Order) {

	co := Header{
		x1: 250,
		y1: 140,
		w:  445,
	}

	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定

	// pdf.Image(config.InitConfig().ImgPath+"stamp.png", co.x1+140, co.y1, nil)

	pdf.SetX(co.x1)
	pdf.SetY(co.y1)
	if len(*info.CustomerAddress) > 120 {
		cutSize1 := CutStringAsSize(*info.CustomerAddress, 60)
		pdf.Cell(nil, "（住所）"+(*info.CustomerAddress)[:cutSize1]) //地址
		pdf.SetX(co.x1 + 48)
		pdf.SetY(co.y1 + 15)

		cutSize2 := CutStringAsSize((*info.CustomerAddress)[cutSize1:], 60)
		pdf.Cell(nil, (*info.CustomerAddress)[cutSize1:cutSize1+cutSize2]) //地址
		pdf.SetX(co.x1 + 48)
		pdf.SetY(co.y1 + 30)
		pdf.Cell(nil, (*info.CustomerAddress)[cutSize1+cutSize2:]) //地址
	} else if len(*info.CustomerAddress) > 60 {
		cutSize := CutStringAsSize(*info.CustomerAddress, 60)
		pdf.Cell(nil, "（住所）"+(*info.CustomerAddress)[:cutSize]) //地址
		pdf.SetX(co.x1 + 48)
		pdf.SetY(co.y1 + 15)
		pdf.Cell(nil, (*info.CustomerAddress)[cutSize:]) //地址
	} else {
		pdf.Cell(nil, "（住所）"+*info.CustomerAddress) //地址
	}

	pdf.SetX(co.x1)
	pdf.SetY(co.y1 + 50)
	if len(*info.CustomerName) > 60 {
		pdf.Cell(nil, "（社名）"+(*info.CustomerName)[:60])
		pdf.SetX(co.x1 + 48)
		pdf.SetY(co.y1 + 65)
		pdf.Cell(nil, (*info.CustomerName)[60:])
	} else {
		pdf.Cell(nil, "（社名）"+*info.CustomerName)
	}

	pdf.SetX(co.x1 + 250)
	pdf.SetY(co.y1 + 70)
	pdf.Cell(nil, "印")

	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定

	if len(*info.EstimateName) > 75 {
		cutSize := CutStringAsSize(*info.EstimateName, 72)

		pdf.SetX(co.x1 - 170)
		pdf.SetY(co.y1 + 85)
		pdf.Cell(nil, "件名（業務名）："+(*info.EstimateName)[0:cutSize])
		pdf.SetX(co.x1 - 75)
		pdf.SetY(co.y1 + 99)
		pdf.Cell(nil, (*info.EstimateName)[cutSize:])
		pdf.SetLineWidth(1.5)
		pdf.Line(co.x1-170, co.y1+113, co.x1-170+co.w, co.y1+113)
	} else {
		pdf.SetX(co.x1 - 170)
		pdf.SetY(co.y1 + 85)
		pdf.Cell(nil, "件名（業務名）："+*info.EstimateName)
		pdf.SetLineWidth(1.5)
		pdf.Line(co.x1-170, co.y1+98, co.x1-170+co.w, co.y1+98)
	}

	pdf.SetX(co.x1 - 170)
	pdf.SetY(co.y1 + 120)
	pdf.Cell(nil, "標題の件につき"+"　"+*info.EstimateOfOrder)
}

func BodyInvoiceOrder(pdf *gopdf.GoPdf, info model.Order) {
	bo := Header{
		x1: 70,
		y1: 300,
		w:  465,
		h:  80,
	}

	pdf.SetFont("Shippori Mincho", "", 10) //フォント、文字サイズ指定
	pdf.SetX(bo.x1 + 230)
	pdf.SetY(bo.y1 - 11) //365
	pdf.Cell(nil, "記")

	//追加背景颜色
	// pdf.SetFillColor(255, 255, 153)
	// pdf.RectFromUpperLeftWithStyle(bo.x1, bo.y1, bo.w, bo.h, "FD")
	// pdf.SetFillColor(255, 255, 153)

	pdf.SetLineWidth(1.5)
	pdf.Line(bo.x1, bo.y1, bo.x1, bo.y1+bo.h*5)           //左
	pdf.Line(178, bo.y1, 178, bo.y1+bo.h*5)               //左2
	pdf.Line(180, bo.y1, 180, bo.y1+bo.h*5)               //左2
	pdf.Line(bo.x1+bo.w, bo.y1, bo.x1+bo.w, bo.y1+bo.h*5) //右

	pdf.Line(bo.x1, bo.y1+100, bo.x1+bo.w, bo.y1+100)
	pdf.Line(bo.x1, bo.y1+220, bo.x1+bo.w, bo.y1+220)
	pdf.Line(bo.x1, bo.y1+300, bo.x1+bo.w, bo.y1+300)
	pdf.Line(bo.x1, bo.y1+340, bo.x1+bo.w, bo.y1+340)

	for num := 0.0; num <= 5; num++ {
		pdf.Line(bo.x1, bo.y1+bo.h*num, bo.x1+bo.w, bo.y1+bo.h*num)
	}

	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定
	pdf.SetX(bo.x1 + 10)
	pdf.SetY(bo.y1 + 1)
	pdf.Cell(nil, "作　業　内　容")

	work_str := strings.Split(*info.Work, "\n")
	for index, str := range work_str {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 1 + float64(index*12))
		pdf.Cell(nil, str)
	}

	pdf.SetX(bo.x1 + 4)
	pdf.SetY(bo.y1 + 81)
	pdf.Cell(nil, "作業期間")
	pdf.SetFont("Shippori Mincho", "", 8) //フォント、文字サイズ指定
	pdf.SetX(bo.x1 + 54)
	pdf.SetY(bo.y1 + 85)
	pdf.Cell(nil, "または")
	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定
	pdf.SetX(bo.x1 + 79)
	pdf.SetY(bo.y1 + 81)
	pdf.Cell(nil, "納期")

	pdf.SetX(bo.x1 + 111)
	pdf.SetY(bo.y1 + 81)
	pdf.Cell(nil, *info.WorkTime)

	pdf.SetX(bo.x1 + 21)
	pdf.SetY(bo.y1 + 101)
	pdf.Cell(nil, "主任担当者")

	if len(*info.Personnel2) != 0 {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 101)
		pdf.Cell(nil, "（御社）"+*info.Personnel1)
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 131)
		pdf.Cell(nil, "（弊社）"+*info.Personnel2)
	} else {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 101)
		pdf.Cell(nil, "（御社）"+*info.Personnel1)
	}

	// personnel_str := strings.Split(*info.Personnel1, "\n")
	// for index, str := range personnel_str {
	// 	pdf.SetX(bo.x1 + 111)
	// 	pdf.SetY(bo.y1 + 101 + float64(index*30))
	// 	pdf.Cell(nil, str)
	// }

	pdf.SetX(bo.x1 + 21)
	pdf.SetY(bo.y1 + 161)
	pdf.Cell(nil, "納　入　物")

	deliverables_str := strings.Split(*info.Deliverables, "\n")
	for index, str := range deliverables_str {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 161 + float64(index*12))
		pdf.Cell(nil, str)
	}

	pdf.SetX(bo.x1 + 2)
	pdf.SetY(bo.y1 + 221)
	pdf.Cell(nil, "提供場所(納入場所)")

	pdf.SetX(bo.x1 + 111)
	pdf.SetY(bo.y1 + 221)
	pdf.Cell(nil, "御社"+*info.DeliverableSpace)

	pdf.SetX(bo.x1 + 21)
	pdf.SetY(bo.y1 + 241)
	pdf.Cell(nil, "委　託　料")

	commission_str := strings.Split(*info.Commission, "\n")
	for index, str := range commission_str {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 241 + float64(index*12))
		pdf.Cell(nil, str)
	}

	pdf.SetX(bo.x1 + 10)
	pdf.SetY(bo.y1 + 301)
	pdf.Cell(nil, "支　払　期　日")

	pdf.SetX(bo.x1 + 111)
	pdf.SetY(bo.y1 + 301)
	pdf.Cell(nil, *info.PaymentDate)

	pdf.SetX(bo.x1 + 10)
	pdf.SetY(bo.y1 + 321)
	pdf.Cell(nil, "検　収　条　件")

	pdf.SetX(bo.x1 + 111)
	pdf.SetY(bo.y1 + 321)
	pdf.Cell(nil, *info.AcceptanceConditions)

	pdf.SetX(bo.x1 + 21)
	pdf.SetY(bo.y1 + 341)
	pdf.Cell(nil, "そ　の　他")

	other_str := strings.Split(*info.Other, "\n")
	for index, str := range other_str {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 341 + float64(index*13))
		pdf.Cell(nil, str)
	}

	pdf.SetX(bo.x1 + 10)
	pdf.SetY(bo.y1 + 420)
	pdf.Cell(nil, "※注：")

	note_str := strings.Split(*info.Note, "\n")
	for index, str := range note_str {
		pdf.SetX(bo.x1 + 45)
		pdf.SetY(bo.y1 + 420 + float64(index*13))
		pdf.Cell(nil, str)
	}

	pdf.SetX(bo.x1 + 380)
	pdf.SetY(bo.y1 + 420 + float64(len(note_str)*13))
	pdf.Cell(nil, "以上")

}
func BodyOrder(pdf *gopdf.GoPdf, info model.Order) {
	bo := Header{
		x1: 70,
		y1: 300,
		w:  465,
		h:  80,
	}

	pdf.SetFont("Shippori Mincho", "", 10) //フォント、文字サイズ指定
	pdf.SetX(bo.x1 + 230)
	pdf.SetY(bo.y1 - 11) //365
	pdf.Cell(nil, "記")

	//追加背景颜色
	// pdf.SetFillColor(255, 255, 153)
	// pdf.RectFromUpperLeftWithStyle(bo.x1, bo.y1, bo.w, bo.h, "FD")
	// pdf.SetFillColor(255, 255, 153)

	pdf.SetLineWidth(1.5)
	pdf.Line(bo.x1, bo.y1, bo.x1, bo.y1+bo.h*5)           //左
	pdf.Line(178, bo.y1, 178, bo.y1+bo.h*5)               //左2
	pdf.Line(180, bo.y1, 180, bo.y1+bo.h*5)               //左2
	pdf.Line(bo.x1+bo.w, bo.y1, bo.x1+bo.w, bo.y1+bo.h*5) //右

	pdf.Line(bo.x1, bo.y1+100, bo.x1+bo.w, bo.y1+100)
	pdf.Line(bo.x1, bo.y1+220, bo.x1+bo.w, bo.y1+220)
	pdf.Line(bo.x1, bo.y1+300, bo.x1+bo.w, bo.y1+300)
	pdf.Line(bo.x1, bo.y1+340, bo.x1+bo.w, bo.y1+340)

	for num := 0.0; num <= 5; num++ {
		pdf.Line(bo.x1, bo.y1+bo.h*num, bo.x1+bo.w, bo.y1+bo.h*num)
	}

	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定
	pdf.SetX(bo.x1 + 10)
	pdf.SetY(bo.y1 + 1)
	pdf.Cell(nil, "作　業　内　容")

	work_str := strings.Split(*info.Work, "\n")
	for index, str := range work_str {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 1 + float64(index*12))
		pdf.Cell(nil, str)
	}

	pdf.SetX(bo.x1 + 4)
	pdf.SetY(bo.y1 + 81)
	pdf.Cell(nil, "作業期間")
	pdf.SetFont("Shippori Mincho", "", 8) //フォント、文字サイズ指定
	pdf.SetX(bo.x1 + 54)
	pdf.SetY(bo.y1 + 85)
	pdf.Cell(nil, "または")
	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定
	pdf.SetX(bo.x1 + 79)
	pdf.SetY(bo.y1 + 81)
	pdf.Cell(nil, "納期")

	pdf.SetX(bo.x1 + 111)
	pdf.SetY(bo.y1 + 81)
	pdf.Cell(nil, *info.WorkTime)

	pdf.SetX(bo.x1 + 21)
	pdf.SetY(bo.y1 + 101)
	pdf.Cell(nil, "主任担当者")

	if len(*info.Personnel2) != 0 {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 101)
		pdf.Cell(nil, "（御社）"+*info.Personnel2)
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 131)
		pdf.Cell(nil, "（弊社）"+*info.Personnel1)
	} else {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 101)
		pdf.Cell(nil, "（弊社）"+*info.Personnel1)
	}

	// personnel_str := strings.Split(*info.Personnel1, "\n")
	// for index, str := range personnel_str {
	// 	pdf.SetX(bo.x1 + 111)
	// 	pdf.SetY(bo.y1 + 101 + float64(index*30))
	// 	pdf.Cell(nil, str)
	// }

	pdf.SetX(bo.x1 + 21)
	pdf.SetY(bo.y1 + 161)
	pdf.Cell(nil, "納　入　物")

	deliverables_str := strings.Split(*info.Deliverables, "\n")
	for index, str := range deliverables_str {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 161 + float64(index*12))
		pdf.Cell(nil, str)
	}

	pdf.SetX(bo.x1 + 2)
	pdf.SetY(bo.y1 + 221)
	pdf.Cell(nil, "提供場所(納入場所)")

	pdf.SetX(bo.x1 + 111)
	pdf.SetY(bo.y1 + 221)
	pdf.Cell(nil, "弊社"+*info.DeliverableSpace)

	pdf.SetX(bo.x1 + 21)
	pdf.SetY(bo.y1 + 241)
	pdf.Cell(nil, "委　託　料")

	commission_str := strings.Split(*info.Commission, "\n")
	for index, str := range commission_str {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 241 + float64(index*12))
		pdf.Cell(nil, str)
	}

	pdf.SetX(bo.x1 + 10)
	pdf.SetY(bo.y1 + 301)
	pdf.Cell(nil, "支　払　期　日")

	pdf.SetX(bo.x1 + 111)
	pdf.SetY(bo.y1 + 301)
	pdf.Cell(nil, *info.PaymentDate)

	pdf.SetX(bo.x1 + 10)
	pdf.SetY(bo.y1 + 321)
	pdf.Cell(nil, "検　収　条　件")

	pdf.SetX(bo.x1 + 111)
	pdf.SetY(bo.y1 + 321)
	pdf.Cell(nil, *info.AcceptanceConditions)

	pdf.SetX(bo.x1 + 21)
	pdf.SetY(bo.y1 + 341)
	pdf.Cell(nil, "そ　の　他")

	other_str := strings.Split(*info.Other, "\n")
	for index, str := range other_str {
		pdf.SetX(bo.x1 + 111)
		pdf.SetY(bo.y1 + 341 + float64(index*13))
		pdf.Cell(nil, str)
	}

	pdf.SetX(bo.x1 + 10)
	pdf.SetY(bo.y1 + 420)
	pdf.Cell(nil, "※注：")

	note_str := strings.Split(*info.Note, "\n")
	for index, str := range note_str {
		pdf.SetX(bo.x1 + 45)
		pdf.SetY(bo.y1 + 420 + float64(index*13))
		pdf.Cell(nil, str)
	}

	pdf.SetX(bo.x1 + 380)
	pdf.SetY(bo.y1 + 420 + float64(len(note_str)*13))
	pdf.Cell(nil, "以上")

}
