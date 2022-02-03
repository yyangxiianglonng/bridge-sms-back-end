package utils

import (
	"main/config"
	"main/model"
	"strings"
	"time"

	"github.com/signintech/gopdf"
)

func NewDeliveryPdf(delivery []*model.Delivery) {

	//获取见积头数据
	var deliveryInfo model.Delivery
	for _, item := range delivery {
		deliveryInfo = *item
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
	CustomerDelivery(&pdf, deliveryInfo)
	CompanyDelivery(&pdf, deliveryInfo)
	TitleDelivery(&pdf, deliveryInfo)
	BodyDelivery(&pdf, deliveryInfo)

	if deliveryInfo.DeliveryPdfNum != nil {
		pdf.WritePdf(config.InitConfig().FilePath + "/pdf/delivery/" + (*deliveryInfo.DeliveryPdfNum)[0:4] + "-" + (*deliveryInfo.DeliveryPdfNum)[4:6] + "-" + (*deliveryInfo.DeliveryPdfNum)[6:8] + "/" + *deliveryInfo.DeliveryPdfNum + "_納品書_" + *deliveryInfo.CustomerName + "様_" + *deliveryInfo.ProjectName + ".pdf")
	} else {
		now := time.Now().Format("2006-01-02")
		pdf.WritePdf(config.InitConfig().FilePath + "/pdf/delivery/" + now + "/" + *deliveryInfo.DeliveryPdfNum + "_納品書_" + *deliveryInfo.CustomerName + "様_" + *deliveryInfo.ProjectName + ".pdf")
	}
}

/*
* ①得意先名〇〇御中
* ②挨拶
* ③納品書No.
 */
func CustomerDelivery(pdf *gopdf.GoPdf, info model.Delivery) {
	cu := Header{
		x1: 70,
		y1: 120,
	}
	//①得意先名〇〇
	pdf.SetFont("Shippori Mincho", "", 14) //フォント、文字サイズ指定
	customerName := info.CustomerName
	pdf.SetX(cu.x1) //x座標指定
	pdf.SetY(cu.y1) //y座標指定
	pdf.Cell(nil, *customerName)
	//〇〇御中
	pdf.SetX(cu.x1 + 10 + float64(len(*customerName))*5) //x座標指定
	pdf.SetY(cu.y1)                                      //y座標指定
	pdf.Cell(nil, "御中")
}

func CompanyDelivery(pdf *gopdf.GoPdf, info model.Delivery) {

	co := Header{
		x1: 380,
		y1: 165,
	}

	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定

	pdf.SetX(co.x1 + 40)    //x座標指定
	pdf.SetY(co.y1 - 60)    //y座標指定
	pdf.Cell(nil, "納品書No.") //見積No.
	pdf.SetX(co.x1 + 100)
	pdf.SetY(co.y1 - 60)
	pdf.Cell(nil, *info.DeliveryPdfNum)
	pdf.SetX(co.x1 + 78)
	pdf.SetY(co.y1 - 45)
	pdf.Cell(nil, info.CreatedAt.Format("2006年01月02日")) //16/09/2021

	// pdf.Image(config.InitConfig().ImgPath+"bridge_logo.png", co.x1+10, co.y1-20, nil)
	pdf.Image(config.InitConfig().ImgPath+"stamp-mini.png", co.x1+120, co.y1, nil)

	pdf.SetX(co.x1 + 75)
	pdf.SetY(co.y1)
	pdf.Cell(nil, "株式会社ブリッジ") //邮编
	pdf.SetX(co.x1 + 5)
	pdf.SetY(co.y1 + 15)
	pdf.Cell(nil, "東京都中央区八丁堀4丁目11-10") //地址
	pdf.SetX(co.x1 + 95)
	pdf.SetY(co.y1 + 30)
	pdf.Cell(nil, "第2SSビル 1F") //地址
	pdf.SetX(co.x1 + 80)
	pdf.SetY(co.y1 + 45)
	pdf.Cell(nil, "Tel:03-6222-3222") //联系方式
}

func TitleDelivery(pdf *gopdf.GoPdf, info model.Delivery) {
	td := Header{
		x1: 260,
		y1: 250,
		w:  480,
	}
	pdf.SetFont("Shippori Mincho", "", 16) //フォント、文字サイズ指定
	pdf.SetX(td.x1)                        //x座標指定
	pdf.SetY(td.y1 + 5)                    //y座標指定
	pdf.Cell(nil, "納　品　書")                 //Rect, String

	pdf.SetLineWidth(1.0)
	pdf.Line(td.x1-190, td.y1, td.x1-190+td.w, td.y1)
	pdf.Line(td.x1-190, td.y1+30, td.x1-190+td.w, td.y1+30)

	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定
	pdf.SetX(td.x1 - 190)                  //x座標指定
	pdf.SetY(td.y1 + 50)                   //y座標指定
	pdf.Cell(nil, "以下内容の通り納品させて頂きました。")

	pdf.SetX(td.x1 - 190) //x座標指定
	pdf.SetY(td.y1 + 70)  //y座標指定
	pdf.Cell(nil, "ご利用頂ける状態である為、ご確認の程よろしくお願い致します。")

	pdf.SetX(td.x1 + 30)  //x座標指定
	pdf.SetY(td.y1 + 100) //y座標指定
	pdf.Cell(nil, "記")
}
func BodyDelivery(pdf *gopdf.GoPdf, info model.Delivery) {
	bd := Header{
		x1: 70,
		y1: 440,
		w:  480,
		h:  20,
	}
	//1.件名
	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定
	pdf.SetX(bd.x1)                        //x座標指定
	pdf.SetY(bd.y1 - 50)                   //y座標指定
	pdf.Cell(nil, "1.件名")
	pdf.SetX(bd.x1 + 60)
	pdf.Cell(nil, *info.ProjectName)
	//2.成果物
	pdf.SetX(bd.x1)      //x座標指定
	pdf.SetY(bd.y1 - 20) //y座標指定
	pdf.Cell(nil, "2.成果物")

	//追加背景颜色
	pdf.SetFillColor(203, 203, 203)
	pdf.RectFromUpperLeftWithStyle(bd.x1, bd.y1, bd.w, bd.h, "FD")
	pdf.SetFillColor(203, 203, 203)

	pdf.SetLineWidth(1.0)
	pdf.Line(bd.x1, bd.y1, bd.x1, bd.y1+bd.h*4)           //左
	pdf.Line(285, bd.y1, 285, bd.y1+bd.h*4)               //左2
	pdf.Line(340, bd.y1, 340, bd.y1+bd.h*4)               //右2
	pdf.Line(bd.x1+bd.w, bd.y1, bd.x1+bd.w, bd.y1+bd.h*4) //右

	for num := 0.0; num <= 4; num++ {
		pdf.Line(bd.x1, bd.y1+bd.h*num, bd.x1+bd.w, bd.y1+bd.h*num)
	}
	pdf.SetFont("Shippori Mincho", "", 10)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(140)
	pdf.SetY(bd.y1 + 4)
	pdf.Cell(nil, "商品名称/摘要")
	pdf.SetX(300)
	pdf.Cell(nil, "部数")
	pdf.SetX(430)
	pdf.Cell(nil, "備考")

	pdf.SetX(bd.x1 + 1)
	pdf.SetY(bd.y1 + bd.h + 6)
	pdf.Cell(nil, *info.Deliverables1)
	pdf.SetX(310)
	pdf.Cell(nil, *info.Quantity1)
	pdf.SetX(341)
	pdf.Cell(nil, *info.Memo1)

	pdf.SetX(bd.x1 + 1)
	pdf.SetY(bd.y1 + bd.h*2 + 6)
	pdf.Cell(nil, *info.Deliverables2)
	pdf.SetX(310)
	pdf.Cell(nil, *info.Quantity2)
	pdf.SetX(341)
	pdf.Cell(nil, *info.Memo2)

	pdf.SetX(bd.x1 + 1)
	pdf.SetY(bd.y1 + bd.h*3 + 6)
	pdf.Cell(nil, *info.Deliverables3)
	pdf.SetX(310)
	pdf.Cell(nil, *info.Quantity3)
	pdf.SetX(341)
	pdf.Cell(nil, *info.Memo3)

	//3.備考
	pdf.SetFont("Shippori Mincho", "", 12) //フォント、文字サイズ指定
	pdf.SetX(bd.x1)                        //x座標指定
	pdf.SetY(bd.y1 + 100)                  //y座標指定
	pdf.Cell(nil, "3.備考")

	arr_str := strings.Split(*info.Remarks, "\n")
	for index, str := range arr_str {
		pdf.SetFont("Shippori Mincho", "", 10) //フォント、文字サイズ指定
		pdf.SetX(bd.x1)
		pdf.SetY(bd.y1 + 130 + float64(index*25))
		pdf.Cell(nil, str)
	}

	var lenRow int
	if len(arr_str) < 3 {
		lenRow = 3
	} else {
		lenRow = len(arr_str)
	}

	pdf.SetLineWidth(1.0)
	for num := 0.0; num < float64(lenRow); num++ {
		pdf.Line(bd.x1, bd.y1+140+(bd.h+5)*num, bd.x1+bd.w, bd.y1+140+(bd.h+5)*num)
	}
}
