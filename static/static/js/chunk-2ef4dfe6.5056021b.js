(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2ef4dfe6"],{"05e2":function(e,t,r){"use strict";r.d(t,"c",(function(){return o})),r.d(t,"d",(function(){return l})),r.d(t,"a",(function(){return s})),r.d(t,"b",(function(){return i}));var a=r("1bab");function o(e){return Object(a["a"])({url:"/delivery",method:"post",data:e})}function l(e){return Object(a["a"])({url:"/delivery",method:"put",data:e})}function s(e){return Object(a["a"])({url:"/delivery/all/"+e.project_code,method:"get",data:e})}function i(e){return Object(a["a"])({url:"/delivery/one/"+e.delivery_code,method:"get",data:e})}},"35a8":function(e,t,r){"use strict";r.r(t);var a=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("el-form",{ref:"ruleForm",staticClass:"demo-ruleForm",attrs:{model:e.ruleForm,rules:e.rules}},[r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:12}},[r("el-form-item",{attrs:{prop:"customer_name"}},[r("el-input",{attrs:{disabled:!0},model:{value:e.ruleForm.customer_name,callback:function(t){e.$set(e.ruleForm,"customer_name",t)},expression:"ruleForm.customer_name"}},[r("template",{slot:"prepend"},[e._v("得意先名")])],2)],1)],1),r("el-col",{attrs:{span:12}},[r("el-form-item",{attrs:{prop:"project_code"}},[r("el-input",{attrs:{disabled:!0},model:{value:e.ruleForm.project_code,callback:function(t){e.$set(e.ruleForm,"project_code",t)},expression:"ruleForm.project_code"}},[r("template",{slot:"prepend"},[e._v("案件CD")])],2)],1)],1)],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:12}},[r("el-form-item",{attrs:{prop:"estimate_code"}},[r("el-input",{attrs:{disabled:!0},model:{value:e.ruleForm.estimate_code,callback:function(t){e.$set(e.ruleForm,"estimate_code",t)},expression:"ruleForm.estimate_code"}},[r("template",{slot:"prepend"},[e._v("見積CD")])],2)],1)],1),r("el-col",{attrs:{span:12}},[r("el-form-item",{attrs:{prop:"delivery_code"}},[r("el-input",{attrs:{disabled:!0},model:{value:e.ruleForm.delivery_code,callback:function(t){e.$set(e.ruleForm,"delivery_code",t)},expression:"ruleForm.delivery_code"}},[r("template",{slot:"prepend"},[e._v("納品CD")])],2)],1)],1)],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:24}},[r("el-form-item",{attrs:{prop:"project_name"}},[r("el-input",{attrs:{disabled:!0},model:{value:e.ruleForm.project_name,callback:function(t){e.$set(e.ruleForm,"project_name",t)},expression:"ruleForm.project_name"}},[r("template",{slot:"prepend"},[e._v("案件名")])],2)],1)],1)],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:24}},[r("el-card",{staticClass:"box-card"},[r("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[r("el-col",{attrs:{span:10}},[e._v("商品名/摘要")]),r("el-col",{attrs:{span:4}},[e._v("数量")]),r("el-col",{attrs:{span:10}},[e._v("備考")])],1),r("div",[r("el-row",{attrs:{gutter:0}},[r("el-col",{attrs:{span:10}},[r("el-form-item",{attrs:{prop:"deliverables1"}},[r("el-input",{model:{value:e.ruleForm.deliverables1,callback:function(t){e.$set(e.ruleForm,"deliverables1",t)},expression:"ruleForm.deliverables1"}})],1)],1),r("el-col",{attrs:{span:4}},[r("el-form-item",{attrs:{prop:"quantity1"}},[r("el-input",{model:{value:e.ruleForm.quantity1,callback:function(t){e.$set(e.ruleForm,"quantity1",t)},expression:"ruleForm.quantity1"}})],1)],1),r("el-col",{attrs:{span:10}},[r("el-form-item",{attrs:{prop:"memo1"}},[r("el-input",{model:{value:e.ruleForm.memo1,callback:function(t){e.$set(e.ruleForm,"memo1",t)},expression:"ruleForm.memo1"}})],1)],1),r("el-col",{attrs:{span:10}},[r("el-form-item",{attrs:{prop:"deliverables2"}},[r("el-input",{model:{value:e.ruleForm.deliverables2,callback:function(t){e.$set(e.ruleForm,"deliverables2",t)},expression:"ruleForm.deliverables2"}})],1)],1),r("el-col",{attrs:{span:4}},[r("el-form-item",{attrs:{prop:"quantity2"}},[r("el-input",{model:{value:e.ruleForm.quantity2,callback:function(t){e.$set(e.ruleForm,"quantity2",t)},expression:"ruleForm.quantity2"}})],1)],1),r("el-col",{attrs:{span:10}},[r("el-form-item",{attrs:{prop:"memo2"}},[r("el-input",{model:{value:e.ruleForm.memo2,callback:function(t){e.$set(e.ruleForm,"memo2",t)},expression:"ruleForm.memo2"}})],1)],1),r("el-col",{attrs:{span:10}},[r("el-form-item",{attrs:{prop:"deliverables3"}},[r("el-input",{model:{value:e.ruleForm.deliverables3,callback:function(t){e.$set(e.ruleForm,"deliverables3",t)},expression:"ruleForm.deliverables3"}})],1)],1),r("el-col",{attrs:{span:4}},[r("el-form-item",{attrs:{prop:"quantity3"}},[r("el-input",{model:{value:e.ruleForm.quantity3,callback:function(t){e.$set(e.ruleForm,"quantity3",t)},expression:"ruleForm.quantity3"}})],1)],1),r("el-col",{attrs:{span:10}},[r("el-form-item",{attrs:{prop:"memo3"}},[r("el-input",{model:{value:e.ruleForm.memo3,callback:function(t){e.$set(e.ruleForm,"memo3",t)},expression:"ruleForm.memo3"}})],1)],1)],1)],1)])],1)],1),r("br"),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:24}},[r("el-form-item",{attrs:{label:"",prop:"remarks"}},[r("el-input",{model:{value:e.ruleForm.remarks,callback:function(t){e.$set(e.ruleForm,"remarks",t)},expression:"ruleForm.remarks"}},[r("template",{slot:"prepend"},[e._v("※(注)")])],2)],1)],1)],1),r("el-row",{attrs:{type:"flex",gutter:20,justify:"end"}},[r("el-col",{attrs:{span:2}},[r("p",[e._v("以上")])])],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:24}},[r("el-form-item",[r("el-button-group",[r("el-button",{attrs:{type:"warning",icon:"el-icon-document-delete",plain:""},on:{click:function(t){return e.cancelForm()}}},[e._v("閉じる")]),r("el-button",{attrs:{type:"primary",icon:"el-icon-document-checked",plain:""},on:{click:function(t){return e.submitForm("ruleForm")}}},[e._v("登録・更新")])],1)],1)],1)],1)],1)},o=[],l=r("67b1"),s=r("05e2"),i=r("5c96"),m={name:"DeliveryForm",components:{},props:{cDeliveryForm:{type:Object,defaylt:function(){return{project_name:"",project_code:"",customer_name:"",estimate_code:"",delivery_code:"",method:""}}}},data:function(){return{ruleForm:{delivery_code:this.cDeliveryForm.delivery_code,project_name:this.cDeliveryForm.project_name,project_code:this.cDeliveryForm.project_code,estimate_code:this.cDeliveryForm.estimate_code,customer_name:this.cDeliveryForm.customer_name,deliverables1:"",deliverables2:"",deliverables3:"",quantity1:"",quantity2:"",quantity3:"",memo1:"",memo2:"",memo3:"",remarks:""},rules:{customer_name:[{required:!0,message:"得意先名を入力してください",trigger:"change"}],CustomerAddress:[{required:!0,message:"得意先住所を入力してください",trigger:"blur"}],Work:[{required:!0,message:"作業内容を入力してください",trigger:"blur"}],WorkTime:[{required:!0,message:"作業期間を入力してください",trigger:"blur"}],PersonnelI:[{required:!0,message:"主任担当者(甲)を入力してください",trigger:"blur"}],PersonnelII:[{required:!0,message:"主任担当者(乙)を入力してください",trigger:"blur"}],Deliverables:[{required:!0,message:"成果物を入力してください",trigger:"blur"}],DeliverableSpace:[{required:!0,message:"納入場所を入力してください",trigger:"blur"}],Commission:[{required:!0,message:"委託料を入力してください",trigger:"blur"}],PaymentDate:[{required:!0,message:"支払期日を入力してください",trigger:"blur"}],AcceptanceConditions:[{required:!0,message:"検収条件を入力してください",trigger:"blur"}]}}},created:function(){"put"===this.cDeliveryForm.method?this.getDeliveryDataOne():this.getEstimateDataOne(),this.getEstimateDetailsDataAll()},methods:{getEstimateDataOne:function(){var e=this,t={};t["estimate_code"]=this.cDeliveryForm.estimate_code,Object(l["c"])(t).then((function(t){console.log(t),e.ruleForm.deliverables1=t.data.data[0].deliverables1,e.ruleForm.deliverables2=t.data.data[0].deliverables2,e.ruleForm.deliverables3=t.data.data[0].deliverables3,e.ruleForm.quantity1=t.data.data[0].quantity1,e.ruleForm.quantity2=t.data.data[0].quantity2,e.ruleForm.quantity3=t.data.data[0].quantity3})).catch((function(e){console.log(e)}))},getEstimateDetailsDataAll:function(){var e={};e["estimate_code"]=this.cDeliveryForm.estimate_code,Object(l["e"])(e).then((function(e){console.log(e.data)})).catch((function(e){console.log(e)}))},getDeliveryDataOne:function(){var e=this,t={};t["delivery_code"]=this.cDeliveryForm.delivery_code,Object(s["b"])(t).then((function(t){e.ruleForm.delivery_code=t.data.data[0].delivery_code,e.ruleForm.project_name=t.data.data[0].project_name,e.ruleForm.project_code=t.data.data[0].project_code,e.ruleForm.estimate_code=t.data.data[0].estimate_code,e.ruleForm.customer_name=t.data.data[0].customer_name,e.ruleForm.deliverables1=t.data.data[0].deliverables1,e.ruleForm.deliverables2=t.data.data[0].deliverables2,e.ruleForm.deliverables3=t.data.data[0].deliverables3,e.ruleForm.quantity1=t.data.data[0].quantity1,e.ruleForm.quantity2=t.data.data[0].quantity2,e.ruleForm.quantity3=t.data.data[0].quantity3,e.ruleForm.memo1=t.data.data[0].memo1,e.ruleForm.memo2=t.data.data[0].memo2,e.ruleForm.memo3=t.data.data[0].memo3,e.ruleForm.remarks=t.data.data[0].remarks})).catch((function(e){console.log(e)}))},postDeliveryData:function(){var e=this,t={};t["delivery_code"]=this.ruleForm.delivery_code,t["estimate_code"]=this.ruleForm.estimate_code,t["project_code"]=this.ruleForm.project_code,t["project_name"]=this.ruleForm.project_name,t["customer_name"]=this.ruleForm.customer_name,t["deliverables1"]=this.ruleForm.deliverables1,t["deliverables2"]=this.ruleForm.deliverables2,t["deliverables3"]=this.ruleForm.deliverables3,t["quantity1"]=this.ruleForm.quantity1,t["quantity2"]=this.ruleForm.quantity2,t["quantity3"]=this.ruleForm.quantity3,t["memo1"]=this.ruleForm.memo1,t["memo2"]=this.ruleForm.memo2,t["memo3"]=this.ruleForm.memo3,t["remarks"]=this.ruleForm.remarks,t["created_by"]=this.$store.getters.full_name,Object(s["c"])(t).then((function(t){200===t.status?(e.$emit("submit"),Object(i["Message"])({showClose:!0,message:t.data.message,type:"success"})):(e.$emit("cancel"),Object(i["Message"])({showClose:!0,message:t.data.message,type:"error"}))})).catch((function(e){console.log(e)}))},putDeliveryData:function(){var e=this,t={};t["delivery_code"]=this.ruleForm.delivery_code,t["estimate_code"]=this.ruleForm.estimate_code,t["project_code"]=this.ruleForm.project_code,t["project_name"]=this.ruleForm.project_name,t["customer_name"]=this.ruleForm.customer_name,t["deliverables1"]=this.ruleForm.deliverables1,t["deliverables2"]=this.ruleForm.deliverables2,t["deliverables3"]=this.ruleForm.deliverables3,t["quantity1"]=this.ruleForm.quantity1,t["quantity2"]=this.ruleForm.quantity2,t["quantity3"]=this.ruleForm.quantity3,t["memo1"]=this.ruleForm.memo1,t["memo2"]=this.ruleForm.memo2,t["memo3"]=this.ruleForm.memo3,t["remarks"]=this.ruleForm.remarks,t["modified_by"]=this.$store.getters.full_name,Object(s["d"])(t).then((function(t){200===t.status?(e.$emit("submit"),Object(i["Message"])({showClose:!0,message:t.data.message,type:"success"})):(e.$emit("cancel"),Object(i["Message"])({showClose:!0,message:t.data.message,type:"error"}))})).catch((function(e){console.log(e)}))},submitForm:function(e){var t=this;this.$refs[e].validate((function(e){if(!e)return console.log("error submit!!"),!1;"put"===t.cDeliveryForm.method?t.putDeliveryData():t.postDeliveryData()}))},cancelForm:function(){this.$emit("cancel")}}},u=m,n=r("2877"),c=Object(n["a"])(u,a,o,!1,null,"0646015a",null);t["default"]=c.exports},"67b1":function(e,t,r){"use strict";r.d(t,"g",(function(){return o})),r.d(t,"i",(function(){return l})),r.d(t,"d",(function(){return s})),r.d(t,"b",(function(){return i})),r.d(t,"c",(function(){return m})),r.d(t,"k",(function(){return u})),r.d(t,"h",(function(){return n})),r.d(t,"j",(function(){return c})),r.d(t,"e",(function(){return d})),r.d(t,"a",(function(){return p})),r.d(t,"f",(function(){return b}));var a=r("1bab");function o(e){return Object(a["a"])({url:"/estimate",method:"post",data:e})}function l(e){return Object(a["a"])({url:"/estimate",method:"put",data:e})}function s(e){return Object(a["a"])({url:"/estimate/",method:"get",data:e})}function i(e){return Object(a["a"])({url:"/estimate/all/"+e.project_code,method:"get",data:e})}function m(e){return Object(a["a"])({url:"/estimate/one/"+e.estimate_code,method:"get",data:e})}function u(e){return Object(a["a"])({url:"/estimate/search",method:"get",params:e})}function n(e){return Object(a["a"])({url:"/estimate/detail",method:"post",data:e})}function c(e){return Object(a["a"])({url:"/estimate/detail",method:"put",data:e})}function d(e){return Object(a["a"])({url:"/estimate/detail/all/"+e.estimate_code,method:"get",data:e})}function p(e){return Object(a["a"])({url:"/estimate/detail/"+e.estimate_details_code,method:"delete",data:e})}function b(e){return Object(a["a"])({url:"/estimate/pdf/"+e.estimate_code,method:"get",data:e})}}}]);
//# sourceMappingURL=chunk-2ef4dfe6.5056021b.js.map