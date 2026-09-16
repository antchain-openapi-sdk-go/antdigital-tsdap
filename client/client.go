// This file is auto-generated, don't edit it. Thanks.
package client

import (
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	antchainutil "github.com/antchain-openapi-sdk-go/antchain-util/service"
)

// Description:
//
// Model for initing client
type Config struct {
	// accesskey id
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accesskey secret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// security token
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// http protocol
	//
	// example:
	//
	// http
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// read timeout
	//
	// example:
	//
	// 10
	ReadTimeout *int `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	// connect timeout
	//
	// example:
	//
	// 10
	ConnectTimeout *int `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	// http proxy
	//
	// example:
	//
	// http://localhost
	HttpProxy *string `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	// https proxy
	//
	// example:
	//
	// https://localhost
	HttpsProxy *string `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	// endpoint
	//
	// example:
	//
	// cs.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// proxy white list
	//
	// example:
	//
	// http://localhost
	NoProxy *string `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	// max idle conns
	//
	// example:
	//
	// 3
	MaxIdleConns *int `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	// user agent
	//
	// example:
	//
	// Alibabacloud/1
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	// socks5 proxy
	Socks5Proxy *string `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	// socks5 network
	//
	// example:
	//
	// TCP
	Socks5NetWork *string `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	// 长链接最大空闲时长
	MaxIdleTimeMillis *int `json:"maxIdleTimeMillis,omitempty" xml:"maxIdleTimeMillis,omitempty"`
	// 长链接最大连接时长
	KeepAliveDurationMillis *int `json:"keepAliveDurationMillis,omitempty" xml:"keepAliveDurationMillis,omitempty"`
	// 最大连接数（长链接最大总数）
	MaxRequests *int `json:"maxRequests,omitempty" xml:"maxRequests,omitempty"`
	// 每个目标主机的最大连接数（分主机域名的长链接最大总数
	MaxRequestsPerHost *int `json:"maxRequestsPerHost,omitempty" xml:"maxRequestsPerHost,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
}

func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
}

func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
}

func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
}

func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
}

func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
}

func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
}

func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
}

func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
}

func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
}

func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
}

func (s *Config) SetMaxIdleTimeMillis(v int) *Config {
	s.MaxIdleTimeMillis = &v
	return s
}

func (s *Config) SetKeepAliveDurationMillis(v int) *Config {
	s.KeepAliveDurationMillis = &v
	return s
}

func (s *Config) SetMaxRequests(v int) *Config {
	s.MaxRequests = &v
	return s
}

func (s *Config) SetMaxRequestsPerHost(v int) *Config {
	s.MaxRequestsPerHost = &v
	return s
}

// 发票申请额度配置
type AmountSummaryConfig struct {
	// 汇总KEY
	// example:
	//
	// 2334445
	SummaryKey *string `json:"summary_key,omitempty" xml:"summary_key,omitempty" require:"true"`
	// 汇总金额，单位为元
	// example:
	//
	// 230
	SummaryValue *string `json:"summary_value,omitempty" xml:"summary_value,omitempty" require:"true"`
}

func (s AmountSummaryConfig) String() string {
	return tea.Prettify(s)
}

func (s AmountSummaryConfig) GoString() string {
	return s.String()
}

func (s *AmountSummaryConfig) SetSummaryKey(v string) *AmountSummaryConfig {
	s.SummaryKey = &v
	return s
}

func (s *AmountSummaryConfig) SetSummaryValue(v string) *AmountSummaryConfig {
	s.SummaryValue = &v
	return s
}

// 发票行信息
type InvoiceLineVO struct {
	// 发票行金额
	// example:
	//
	// 10.09
	LineAmt *string `json:"line_amt,omitempty" xml:"line_amt,omitempty" require:"true"`
	// 发票行ID
	// example:
	//
	// 20
	LineId *string `json:"line_id,omitempty" xml:"line_id,omitempty" require:"true"`
	// 税额
	// example:
	//
	// 20.09
	TaxAmt *string `json:"tax_amt,omitempty" xml:"tax_amt,omitempty" require:"true"`
	// 税率
	// example:
	//
	// 0.06
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty" require:"true"`
	// 货物或劳务名称，如 软件服务费201612
	// example:
	//
	// 软件服务费201612
	LineProductName *string `json:"line_product_name,omitempty" xml:"line_product_name,omitempty" require:"true"`
	// 不含税金额
	// example:
	//
	// 93.24
	TaxExclusiveAmt *string `json:"tax_exclusive_amt,omitempty" xml:"tax_exclusive_amt,omitempty"`
	// 数量
	// example:
	//
	// 13
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 单位
	// example:
	//
	// 个
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	// example:
	//
	// 2.34
	UnitPrice *string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 规格型号
	// example:
	//
	// 规格型号
	SpecificationModel *string `json:"specification_model,omitempty" xml:"specification_model,omitempty"`
	// 服务大类
	// example:
	//
	// 技术服务费
	TaxClassificationName *string `json:"tax_classification_name,omitempty" xml:"tax_classification_name,omitempty"`
	// 税收分类编号
	// example:
	//
	// 344555
	TaxClassificationCode *string `json:"tax_classification_code,omitempty" xml:"tax_classification_code,omitempty"`
	// 货物或劳务名称后缀
	// example:
	//
	// 202306
	ProductNameSuffix *string `json:"product_name_suffix,omitempty" xml:"product_name_suffix,omitempty"`
}

func (s InvoiceLineVO) String() string {
	return tea.Prettify(s)
}

func (s InvoiceLineVO) GoString() string {
	return s.String()
}

func (s *InvoiceLineVO) SetLineAmt(v string) *InvoiceLineVO {
	s.LineAmt = &v
	return s
}

func (s *InvoiceLineVO) SetLineId(v string) *InvoiceLineVO {
	s.LineId = &v
	return s
}

func (s *InvoiceLineVO) SetTaxAmt(v string) *InvoiceLineVO {
	s.TaxAmt = &v
	return s
}

func (s *InvoiceLineVO) SetTaxRate(v string) *InvoiceLineVO {
	s.TaxRate = &v
	return s
}

func (s *InvoiceLineVO) SetLineProductName(v string) *InvoiceLineVO {
	s.LineProductName = &v
	return s
}

func (s *InvoiceLineVO) SetTaxExclusiveAmt(v string) *InvoiceLineVO {
	s.TaxExclusiveAmt = &v
	return s
}

func (s *InvoiceLineVO) SetQuantity(v int64) *InvoiceLineVO {
	s.Quantity = &v
	return s
}

func (s *InvoiceLineVO) SetUnit(v string) *InvoiceLineVO {
	s.Unit = &v
	return s
}

func (s *InvoiceLineVO) SetUnitPrice(v string) *InvoiceLineVO {
	s.UnitPrice = &v
	return s
}

func (s *InvoiceLineVO) SetSpecificationModel(v string) *InvoiceLineVO {
	s.SpecificationModel = &v
	return s
}

func (s *InvoiceLineVO) SetTaxClassificationName(v string) *InvoiceLineVO {
	s.TaxClassificationName = &v
	return s
}

func (s *InvoiceLineVO) SetTaxClassificationCode(v string) *InvoiceLineVO {
	s.TaxClassificationCode = &v
	return s
}

func (s *InvoiceLineVO) SetProductNameSuffix(v string) *InvoiceLineVO {
	s.ProductNameSuffix = &v
	return s
}

// 销方信息
type ApplyInvoiceSeller struct {
	// 地址
	// example:
	//
	// 上海市黄浦区
	SellerAddress *string `json:"seller_address,omitempty" xml:"seller_address,omitempty"`
	// 银行账号
	// example:
	//
	// 323422244555
	SellerBankAccount *string `json:"seller_bank_account,omitempty" xml:"seller_bank_account,omitempty" require:"true"`
	// 银行名称
	// example:
	//
	// 招商很行
	SellerBankName *string `json:"seller_bank_name,omitempty" xml:"seller_bank_name,omitempty" require:"true"`
	// 公司名称
	// example:
	//
	// 蚂蚁区块链
	SellerCompanyName *string `json:"seller_company_name,omitempty" xml:"seller_company_name,omitempty" require:"true"`
	// 税号
	// example:
	//
	// 33453344556
	SellerTaxNo *string `json:"seller_tax_no,omitempty" xml:"seller_tax_no,omitempty" require:"true"`
	// 电话
	// example:
	//
	// 0571-978655
	SellerTelephone *string `json:"seller_telephone,omitempty" xml:"seller_telephone,omitempty" require:"true"`
	// 销方机构id
	// example:
	//
	// ZL6
	SellerInstId *string `json:"seller_inst_id,omitempty" xml:"seller_inst_id,omitempty" require:"true"`
}

func (s ApplyInvoiceSeller) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoiceSeller) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceSeller) SetSellerAddress(v string) *ApplyInvoiceSeller {
	s.SellerAddress = &v
	return s
}

func (s *ApplyInvoiceSeller) SetSellerBankAccount(v string) *ApplyInvoiceSeller {
	s.SellerBankAccount = &v
	return s
}

func (s *ApplyInvoiceSeller) SetSellerBankName(v string) *ApplyInvoiceSeller {
	s.SellerBankName = &v
	return s
}

func (s *ApplyInvoiceSeller) SetSellerCompanyName(v string) *ApplyInvoiceSeller {
	s.SellerCompanyName = &v
	return s
}

func (s *ApplyInvoiceSeller) SetSellerTaxNo(v string) *ApplyInvoiceSeller {
	s.SellerTaxNo = &v
	return s
}

func (s *ApplyInvoiceSeller) SetSellerTelephone(v string) *ApplyInvoiceSeller {
	s.SellerTelephone = &v
	return s
}

func (s *ApplyInvoiceSeller) SetSellerInstId(v string) *ApplyInvoiceSeller {
	s.SellerInstId = &v
	return s
}

// 寄送信息
type ApplyInvoiceDelivery struct {
	// 寄送类型，如1表示快递，2表示email
	// example:
	//
	// 1
	DeliveryType *string `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
	// 开票人PID
	// example:
	//
	// 2088720671581149
	Pid *string `json:"pid,omitempty" xml:"pid,omitempty" require:"true"`
	// 收件人名称
	// example:
	//
	// 收件人名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 电话
	// example:
	//
	// 13987654321
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 国家
	// example:
	//
	// 中国
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// 省份
	// example:
	//
	// 上海市
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	// example:
	//
	// 上海市
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// 区县
	// example:
	//
	// 黄浦区
	CountyDistrict *string `json:"county_district,omitempty" xml:"county_district,omitempty"`
	// 街道
	// example:
	//
	// 街道
	Street *string `json:"street,omitempty" xml:"street,omitempty"`
	// 详细地址
	// example:
	//
	// 外马路618号
	DetailAddress *string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 收件人邮件
	// example:
	//
	// a@1.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s ApplyInvoiceDelivery) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoiceDelivery) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceDelivery) SetDeliveryType(v string) *ApplyInvoiceDelivery {
	s.DeliveryType = &v
	return s
}

func (s *ApplyInvoiceDelivery) SetPid(v string) *ApplyInvoiceDelivery {
	s.Pid = &v
	return s
}

func (s *ApplyInvoiceDelivery) SetName(v string) *ApplyInvoiceDelivery {
	s.Name = &v
	return s
}

func (s *ApplyInvoiceDelivery) SetTelephone(v string) *ApplyInvoiceDelivery {
	s.Telephone = &v
	return s
}

func (s *ApplyInvoiceDelivery) SetCountry(v string) *ApplyInvoiceDelivery {
	s.Country = &v
	return s
}

func (s *ApplyInvoiceDelivery) SetProvince(v string) *ApplyInvoiceDelivery {
	s.Province = &v
	return s
}

func (s *ApplyInvoiceDelivery) SetCity(v string) *ApplyInvoiceDelivery {
	s.City = &v
	return s
}

func (s *ApplyInvoiceDelivery) SetCountyDistrict(v string) *ApplyInvoiceDelivery {
	s.CountyDistrict = &v
	return s
}

func (s *ApplyInvoiceDelivery) SetStreet(v string) *ApplyInvoiceDelivery {
	s.Street = &v
	return s
}

func (s *ApplyInvoiceDelivery) SetDetailAddress(v string) *ApplyInvoiceDelivery {
	s.DetailAddress = &v
	return s
}

func (s *ApplyInvoiceDelivery) SetEmail(v string) *ApplyInvoiceDelivery {
	s.Email = &v
	return s
}

// 发票信息
type InvoiceInfoVO struct {
	// 发票ID
	// example:
	//
	// 233445656676
	InvoiceId *string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty" require:"true"`
	// 发票编号
	// example:
	//
	// FULL_ELECTRONIC_INVOICE
	InvoiceCode *string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// 发票号码
	// example:
	//
	// 3455666767777
	InvoiceNo *string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 发票金额
	// example:
	//
	// 220.09
	InvoiceAmt *string `json:"invoice_amt,omitempty" xml:"invoice_amt,omitempty" require:"true"`
	// 发票类型,01专票 02普票 03营业税发票 04国际形式发票 05其它发票
	// example:
	//
	// 01
	InvoiceType *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty" require:"true"`
	// 发票的业务状态, 待开票: TO_INV, 开票中: INV_ING, 已开票: INVED, 退票中: INV_RETURN, 换票中: INV_REPLACE, 已失效: INV_DEP
	// example:
	//
	// INVED
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// 发票介质，01：电子，02：纸质
	// example:
	//
	// 01
	InvoiceMaterial *string `json:"invoice_material,omitempty" xml:"invoice_material,omitempty"`
	// 发票行信息
	InvoiceLines []*InvoiceLineVO `json:"invoice_lines,omitempty" xml:"invoice_lines,omitempty" type:"Repeated"`
	// 开票日期
	// example:
	//
	// 2018-10-10T10:10:00Z
	InvoiceDate *string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
}

func (s InvoiceInfoVO) String() string {
	return tea.Prettify(s)
}

func (s InvoiceInfoVO) GoString() string {
	return s.String()
}

func (s *InvoiceInfoVO) SetInvoiceId(v string) *InvoiceInfoVO {
	s.InvoiceId = &v
	return s
}

func (s *InvoiceInfoVO) SetInvoiceCode(v string) *InvoiceInfoVO {
	s.InvoiceCode = &v
	return s
}

func (s *InvoiceInfoVO) SetInvoiceNo(v string) *InvoiceInfoVO {
	s.InvoiceNo = &v
	return s
}

func (s *InvoiceInfoVO) SetInvoiceAmt(v string) *InvoiceInfoVO {
	s.InvoiceAmt = &v
	return s
}

func (s *InvoiceInfoVO) SetInvoiceType(v string) *InvoiceInfoVO {
	s.InvoiceType = &v
	return s
}

func (s *InvoiceInfoVO) SetStatus(v string) *InvoiceInfoVO {
	s.Status = &v
	return s
}

func (s *InvoiceInfoVO) SetInvoiceMaterial(v string) *InvoiceInfoVO {
	s.InvoiceMaterial = &v
	return s
}

func (s *InvoiceInfoVO) SetInvoiceLines(v []*InvoiceLineVO) *InvoiceInfoVO {
	s.InvoiceLines = v
	return s
}

func (s *InvoiceInfoVO) SetInvoiceDate(v string) *InvoiceInfoVO {
	s.InvoiceDate = &v
	return s
}

// 发票申请金额信息
type ApplyInvoiceAmountAllocation struct {
	// 额度汇总信息，标准可开票单据下，KEY为需要占用的单据号，VALUE为需要占用的额度
	// example:
	//
	// xxx
	AmountSummaryConfig []*AmountSummaryConfig `json:"amount_summary_config,omitempty" xml:"amount_summary_config,omitempty" type:"Repeated"`
	// 额度来源，默认INVOICE_RCPT
	// example:
	//
	// INVOICE_RCPT
	AmountSource *string `json:"amount_source,omitempty" xml:"amount_source,omitempty"`
	// 币种，默认156
	// example:
	//
	// 156
	Ccy *string `json:"ccy,omitempty" xml:"ccy,omitempty" require:"true"`
}

func (s ApplyInvoiceAmountAllocation) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoiceAmountAllocation) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceAmountAllocation) SetAmountSummaryConfig(v []*AmountSummaryConfig) *ApplyInvoiceAmountAllocation {
	s.AmountSummaryConfig = v
	return s
}

func (s *ApplyInvoiceAmountAllocation) SetAmountSource(v string) *ApplyInvoiceAmountAllocation {
	s.AmountSource = &v
	return s
}

func (s *ApplyInvoiceAmountAllocation) SetCcy(v string) *ApplyInvoiceAmountAllocation {
	s.Ccy = &v
	return s
}

// 发票行信息
type PreviewInvoiceLine struct {
	// 货物名称(商品名称)
	// example:
	//
	// 技术服务费
	LineProductName *string `json:"line_product_name,omitempty" xml:"line_product_name,omitempty" require:"true"`
	// 税率
	// example:
	//
	// 0.13
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty" require:"true"`
	// 发票行含税金额，单位:元
	// example:
	//
	// 102.98
	LineAmt *string `json:"line_amt,omitempty" xml:"line_amt,omitempty" require:"true"`
	// 币种，默认156
	// example:
	//
	// 156
	Ccy *string `json:"ccy,omitempty" xml:"ccy,omitempty"`
	// 单位
	// example:
	//
	// 个
	MeasurementUnit *string `json:"measurement_unit,omitempty" xml:"measurement_unit,omitempty"`
	// 关联的L5商品
	// example:
	//
	// SQBRFSZL601262688
	RelateCommodityCode *string `json:"relate_commodity_code,omitempty" xml:"relate_commodity_code,omitempty" require:"true"`
	// 服务类型 如：AFTECH_SERVICE
	// example:
	//
	// AFTECH_SERVICE
	ServiceMode *string `json:"service_mode,omitempty" xml:"service_mode,omitempty"`
	// 规格型号
	// example:
	//
	// 规格型号
	ProductSpecification *string `json:"product_specification,omitempty" xml:"product_specification,omitempty"`
}

func (s PreviewInvoiceLine) String() string {
	return tea.Prettify(s)
}

func (s PreviewInvoiceLine) GoString() string {
	return s.String()
}

func (s *PreviewInvoiceLine) SetLineProductName(v string) *PreviewInvoiceLine {
	s.LineProductName = &v
	return s
}

func (s *PreviewInvoiceLine) SetTaxRate(v string) *PreviewInvoiceLine {
	s.TaxRate = &v
	return s
}

func (s *PreviewInvoiceLine) SetLineAmt(v string) *PreviewInvoiceLine {
	s.LineAmt = &v
	return s
}

func (s *PreviewInvoiceLine) SetCcy(v string) *PreviewInvoiceLine {
	s.Ccy = &v
	return s
}

func (s *PreviewInvoiceLine) SetMeasurementUnit(v string) *PreviewInvoiceLine {
	s.MeasurementUnit = &v
	return s
}

func (s *PreviewInvoiceLine) SetRelateCommodityCode(v string) *PreviewInvoiceLine {
	s.RelateCommodityCode = &v
	return s
}

func (s *PreviewInvoiceLine) SetServiceMode(v string) *PreviewInvoiceLine {
	s.ServiceMode = &v
	return s
}

func (s *PreviewInvoiceLine) SetProductSpecification(v string) *PreviewInvoiceLine {
	s.ProductSpecification = &v
	return s
}

// 发票购方信息
type ApplyInvoiceBuyer struct {
	// 购方地址
	// example:
	//
	// 上海市黄浦区
	BuyerAddress *string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// 银行账号
	// example:
	//
	// 23355434
	BuyerBankAccount *string `json:"buyer_bank_account,omitempty" xml:"buyer_bank_account,omitempty"`
	// 银行名称
	// example:
	//
	// 招商银行
	BuyerBankName *string `json:"buyer_bank_name,omitempty" xml:"buyer_bank_name,omitempty"`
	// 抬头
	// example:
	//
	// xxxx公司
	BuyerInvoiceTitle *string `json:"buyer_invoice_title,omitempty" xml:"buyer_invoice_title,omitempty" require:"true"`
	// 税号
	// example:
	//
	// 3344555
	BuyerTaxNo *string `json:"buyer_tax_no,omitempty" xml:"buyer_tax_no,omitempty"`
	// 电话
	// example:
	//
	// 0571-94848
	BuyerTelephone *string `json:"buyer_telephone,omitempty" xml:"buyer_telephone,omitempty"`
	// 纳税人资格类型
	// example:
	//
	// 01
	TaxPayerQualification *string `json:"tax_payer_qualification,omitempty" xml:"tax_payer_qualification,omitempty" require:"true"`
	// 客户开票配置ID
	// example:
	//
	// 2344
	BuyerConfigId *string `json:"buyer_config_id,omitempty" xml:"buyer_config_id,omitempty"`
}

func (s ApplyInvoiceBuyer) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoiceBuyer) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceBuyer) SetBuyerAddress(v string) *ApplyInvoiceBuyer {
	s.BuyerAddress = &v
	return s
}

func (s *ApplyInvoiceBuyer) SetBuyerBankAccount(v string) *ApplyInvoiceBuyer {
	s.BuyerBankAccount = &v
	return s
}

func (s *ApplyInvoiceBuyer) SetBuyerBankName(v string) *ApplyInvoiceBuyer {
	s.BuyerBankName = &v
	return s
}

func (s *ApplyInvoiceBuyer) SetBuyerInvoiceTitle(v string) *ApplyInvoiceBuyer {
	s.BuyerInvoiceTitle = &v
	return s
}

func (s *ApplyInvoiceBuyer) SetBuyerTaxNo(v string) *ApplyInvoiceBuyer {
	s.BuyerTaxNo = &v
	return s
}

func (s *ApplyInvoiceBuyer) SetBuyerTelephone(v string) *ApplyInvoiceBuyer {
	s.BuyerTelephone = &v
	return s
}

func (s *ApplyInvoiceBuyer) SetTaxPayerQualification(v string) *ApplyInvoiceBuyer {
	s.TaxPayerQualification = &v
	return s
}

func (s *ApplyInvoiceBuyer) SetBuyerConfigId(v string) *ApplyInvoiceBuyer {
	s.BuyerConfigId = &v
	return s
}

// 发票行信息
type IntlInvoiceInfoItem struct {
	// 发票号
	// example:
	//
	// 34555
	InvoiceNo *string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty" require:"true"`
	// 发票金额
	// example:
	//
	// 34.23
	InvoiceAmt *string `json:"invoice_amt,omitempty" xml:"invoice_amt,omitempty" require:"true"`
	// 不含税金
	// example:
	//
	// 23.12
	ExcludingTaxInvoiceAmt *string `json:"excluding_tax_invoice_amt,omitempty" xml:"excluding_tax_invoice_amt,omitempty" require:"true"`
	// 税额
	// example:
	//
	// 2.23
	TaxAmt *string `json:"tax_amt,omitempty" xml:"tax_amt,omitempty" require:"true"`
	// 税率
	// example:
	//
	// 0.09
	Tax *string `json:"tax,omitempty" xml:"tax,omitempty" require:"true"`
	// 状态
	// example:
	//
	// INVED
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// 发票ID
	// example:
	//
	// 20260716107315002131450000476631
	InvoiceId *string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty" require:"true"`
}

func (s IntlInvoiceInfoItem) String() string {
	return tea.Prettify(s)
}

func (s IntlInvoiceInfoItem) GoString() string {
	return s.String()
}

func (s *IntlInvoiceInfoItem) SetInvoiceNo(v string) *IntlInvoiceInfoItem {
	s.InvoiceNo = &v
	return s
}

func (s *IntlInvoiceInfoItem) SetInvoiceAmt(v string) *IntlInvoiceInfoItem {
	s.InvoiceAmt = &v
	return s
}

func (s *IntlInvoiceInfoItem) SetExcludingTaxInvoiceAmt(v string) *IntlInvoiceInfoItem {
	s.ExcludingTaxInvoiceAmt = &v
	return s
}

func (s *IntlInvoiceInfoItem) SetTaxAmt(v string) *IntlInvoiceInfoItem {
	s.TaxAmt = &v
	return s
}

func (s *IntlInvoiceInfoItem) SetTax(v string) *IntlInvoiceInfoItem {
	s.Tax = &v
	return s
}

func (s *IntlInvoiceInfoItem) SetStatus(v string) *IntlInvoiceInfoItem {
	s.Status = &v
	return s
}

func (s *IntlInvoiceInfoItem) SetInvoiceId(v string) *IntlInvoiceInfoItem {
	s.InvoiceId = &v
	return s
}

// 发票申请场景下发票行信息
type ApplyInvoiceLine struct {
	// 含税金额
	// example:
	//
	// 133880
	Amt *string `json:"amt,omitempty" xml:"amt,omitempty" require:"true"`
	// 税额
	// example:
	//
	// 45
	TaxAmt *string `json:"tax_amt,omitempty" xml:"tax_amt,omitempty" require:"true"`
	// 税率
	// example:
	//
	// 0.13
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty" require:"true"`
	// 不含税金额
	// example:
	//
	// 24556
	TaxExclusiveAmt *string `json:"tax_exclusive_amt,omitempty" xml:"tax_exclusive_amt,omitempty" require:"true"`
	// 含税单价
	// example:
	//
	// 234
	UnitAmt *string `json:"unit_amt,omitempty" xml:"unit_amt,omitempty" require:"true"`
	// 服务大类编号
	// example:
	//
	// 2345455
	TaxClassificationCode *string `json:"tax_classification_code,omitempty" xml:"tax_classification_code,omitempty" require:"true"`
	// 货物或劳务名称
	// example:
	//
	// 技术服务费
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	// 产品CODE
	// example:
	//
	// SPU33445
	ProductCode *string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 规格型号
	// example:
	//
	// 333
	ProductSpecification *string `json:"product_specification,omitempty" xml:"product_specification,omitempty"`
	// 计量单位
	// example:
	//
	// 个
	MeasurementNnit *string `json:"measurement_nnit,omitempty" xml:"measurement_nnit,omitempty"`
	// 数量，默认为1
	// example:
	//
	// 1
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 发票行ID
	// example:
	//
	// 33455666
	InvoiceLineId *string `json:"invoice_line_id,omitempty" xml:"invoice_line_id,omitempty" require:"true"`
	// 劳务与货物名称的后缀，主要有账期（202309）、PID（2088XXXX）等
	// example:
	//
	// 202309
	ProductNameSuffix *string `json:"product_name_suffix,omitempty" xml:"product_name_suffix,omitempty"`
}

func (s ApplyInvoiceLine) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoiceLine) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceLine) SetAmt(v string) *ApplyInvoiceLine {
	s.Amt = &v
	return s
}

func (s *ApplyInvoiceLine) SetTaxAmt(v string) *ApplyInvoiceLine {
	s.TaxAmt = &v
	return s
}

func (s *ApplyInvoiceLine) SetTaxRate(v string) *ApplyInvoiceLine {
	s.TaxRate = &v
	return s
}

func (s *ApplyInvoiceLine) SetTaxExclusiveAmt(v string) *ApplyInvoiceLine {
	s.TaxExclusiveAmt = &v
	return s
}

func (s *ApplyInvoiceLine) SetUnitAmt(v string) *ApplyInvoiceLine {
	s.UnitAmt = &v
	return s
}

func (s *ApplyInvoiceLine) SetTaxClassificationCode(v string) *ApplyInvoiceLine {
	s.TaxClassificationCode = &v
	return s
}

func (s *ApplyInvoiceLine) SetProductName(v string) *ApplyInvoiceLine {
	s.ProductName = &v
	return s
}

func (s *ApplyInvoiceLine) SetProductCode(v string) *ApplyInvoiceLine {
	s.ProductCode = &v
	return s
}

func (s *ApplyInvoiceLine) SetProductSpecification(v string) *ApplyInvoiceLine {
	s.ProductSpecification = &v
	return s
}

func (s *ApplyInvoiceLine) SetMeasurementNnit(v string) *ApplyInvoiceLine {
	s.MeasurementNnit = &v
	return s
}

func (s *ApplyInvoiceLine) SetQuantity(v string) *ApplyInvoiceLine {
	s.Quantity = &v
	return s
}

func (s *ApplyInvoiceLine) SetInvoiceLineId(v string) *ApplyInvoiceLine {
	s.InvoiceLineId = &v
	return s
}

func (s *ApplyInvoiceLine) SetProductNameSuffix(v string) *ApplyInvoiceLine {
	s.ProductNameSuffix = &v
	return s
}

// 申请记录
type IntlInvoiceApplyInfoItem struct {
	// 租户ID
	// example:
	//
	// 20882838383
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty" require:"true"`
	// 发票申请ID
	// example:
	//
	// 29394
	InvoiceApplyId *string `json:"invoice_apply_id,omitempty" xml:"invoice_apply_id,omitempty" require:"true"`
	// 发票申请金额
	// example:
	//
	// 13.44
	InvoiceAmt *string `json:"invoice_amt,omitempty" xml:"invoice_amt,omitempty" require:"true"`
	// 发票币种
	// example:
	//
	// 840
	InvoiceCcy *string `json:"invoice_ccy,omitempty" xml:"invoice_ccy,omitempty" require:"true"`
	// 发票申请时间
	// example:
	//
	// 2025-04-93 12:03:33
	InvoiceDate *string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 申请人名称
	// example:
	//
	// 客户名称
	OperatorName *string `json:"operator_name,omitempty" xml:"operator_name,omitempty" require:"true"`
	// 申请人ID
	// example:
	//
	// 20882838383
	OperatorId *string `json:"operator_id,omitempty" xml:"operator_id,omitempty" require:"true"`
	// 申请状态
	// example:
	//
	// 03
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// 申请业务号
	// example:
	//
	// 335455
	BsnNo *string `json:"bsn_no,omitempty" xml:"bsn_no,omitempty" require:"true"`
	// 发票列表
	// example:
	//
	// undefined
	RelateInvoices []*IntlInvoiceInfoItem `json:"relate_invoices,omitempty" xml:"relate_invoices,omitempty" type:"Repeated"`
	// 发票类型
	// example:
	//
	// 42
	InvoiceType *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty" require:"true"`
	// 申请类型
	// example:
	//
	// new
	ApplyType *string `json:"apply_type,omitempty" xml:"apply_type,omitempty" require:"true"`
	// 形式发票文件映射
	// example:
	//
	// xx
	ElcFileMap *string `json:"elc_file_map,omitempty" xml:"elc_file_map,omitempty" require:"true"`
}

func (s IntlInvoiceApplyInfoItem) String() string {
	return tea.Prettify(s)
}

func (s IntlInvoiceApplyInfoItem) GoString() string {
	return s.String()
}

func (s *IntlInvoiceApplyInfoItem) SetTenantId(v string) *IntlInvoiceApplyInfoItem {
	s.TenantId = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetInvoiceApplyId(v string) *IntlInvoiceApplyInfoItem {
	s.InvoiceApplyId = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetInvoiceAmt(v string) *IntlInvoiceApplyInfoItem {
	s.InvoiceAmt = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetInvoiceCcy(v string) *IntlInvoiceApplyInfoItem {
	s.InvoiceCcy = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetInvoiceDate(v string) *IntlInvoiceApplyInfoItem {
	s.InvoiceDate = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetOperatorName(v string) *IntlInvoiceApplyInfoItem {
	s.OperatorName = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetOperatorId(v string) *IntlInvoiceApplyInfoItem {
	s.OperatorId = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetStatus(v string) *IntlInvoiceApplyInfoItem {
	s.Status = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetBsnNo(v string) *IntlInvoiceApplyInfoItem {
	s.BsnNo = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetRelateInvoices(v []*IntlInvoiceInfoItem) *IntlInvoiceApplyInfoItem {
	s.RelateInvoices = v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetInvoiceType(v string) *IntlInvoiceApplyInfoItem {
	s.InvoiceType = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetApplyType(v string) *IntlInvoiceApplyInfoItem {
	s.ApplyType = &v
	return s
}

func (s *IntlInvoiceApplyInfoItem) SetElcFileMap(v string) *IntlInvoiceApplyInfoItem {
	s.ElcFileMap = &v
	return s
}

// 发票预览请求
type PreviewInvoiceRequest struct {
	// 开票场景，默认基于订单开票
	// example:
	//
	// LTC_RCPT_BILL
	InvoiceBizScene *string `json:"invoice_biz_scene,omitempty" xml:"invoice_biz_scene,omitempty"`
	// 租户ID
	// example:
	//
	// 2088720671581149
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty" require:"true"`
	// 开票操作，默认预览开票PREVIEW_INVOICING
	// example:
	//
	// PREVIEW_INVOICING
	InvoiceBizAction *string `json:"invoice_biz_action,omitempty" xml:"invoice_biz_action,omitempty"`
	// 币种
	// example:
	//
	// 156
	Ccy *string `json:"ccy,omitempty" xml:"ccy,omitempty" require:"true"`
	// ou
	// example:
	//
	// ZL6
	Ou *string `json:"ou,omitempty" xml:"ou,omitempty" require:"true"`
	// 合同号
	// example:
	//
	// 2088720671581149-ZNHYFM01222234
	ArNo *string `json:"ar_no,omitempty" xml:"ar_no,omitempty" require:"true"`
	// 调用来源
	// example:
	//
	// IOT
	Source *string `json:"source,omitempty" xml:"source,omitempty" require:"true"`
	// 开票人ID
	// example:
	//
	// 2088720671581149
	OperatorNo *string `json:"operator_no,omitempty" xml:"operator_no,omitempty" require:"true"`
	// 操作人名称
	// example:
	//
	// XXXD
	OperatorName *string `json:"operator_name,omitempty" xml:"operator_name,omitempty" require:"true"`
	// 外部申请单据号，长度不超过32位
	// example:
	//
	// 3455444
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 发票类型， 01 专票 02 普票
	// example:
	//
	// 01
	InvoiceType *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty" require:"true"`
	// 发票介质 01 电子发票 02 纸质发票
	// example:
	//
	// 01
	InvoiceMaterial *string `json:"invoice_material,omitempty" xml:"invoice_material,omitempty" require:"true"`
	// 票面备注 该内容会原样展示到发票上
	// example:
	//
	// 备注
	InvoiceNote *string `json:"invoice_note,omitempty" xml:"invoice_note,omitempty"`
	// 申请说明
	// example:
	//
	// 申请说明
	ApplyReason *string `json:"apply_reason,omitempty" xml:"apply_reason,omitempty"`
	// 购方信息
	InvoiceBuyer *ApplyInvoiceBuyer `json:"invoice_buyer,omitempty" xml:"invoice_buyer,omitempty" require:"true"`
	// 寄送信息
	ApplyInvoiceDelivery *ApplyInvoiceDelivery `json:"apply_invoice_delivery,omitempty" xml:"apply_invoice_delivery,omitempty" require:"true"`
	// 发票需要占用的额度信息
	ApplyInvoiceQuota *ApplyInvoiceAmountAllocation `json:"apply_invoice_quota,omitempty" xml:"apply_invoice_quota,omitempty" require:"true"`
	// 发票行列表
	PreviewInvoiceLines []*PreviewInvoiceLine `json:"preview_invoice_lines,omitempty" xml:"preview_invoice_lines,omitempty" type:"Repeated"`
}

func (s PreviewInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviewInvoiceRequest) GoString() string {
	return s.String()
}

func (s *PreviewInvoiceRequest) SetInvoiceBizScene(v string) *PreviewInvoiceRequest {
	s.InvoiceBizScene = &v
	return s
}

func (s *PreviewInvoiceRequest) SetTenantId(v string) *PreviewInvoiceRequest {
	s.TenantId = &v
	return s
}

func (s *PreviewInvoiceRequest) SetInvoiceBizAction(v string) *PreviewInvoiceRequest {
	s.InvoiceBizAction = &v
	return s
}

func (s *PreviewInvoiceRequest) SetCcy(v string) *PreviewInvoiceRequest {
	s.Ccy = &v
	return s
}

func (s *PreviewInvoiceRequest) SetOu(v string) *PreviewInvoiceRequest {
	s.Ou = &v
	return s
}

func (s *PreviewInvoiceRequest) SetArNo(v string) *PreviewInvoiceRequest {
	s.ArNo = &v
	return s
}

func (s *PreviewInvoiceRequest) SetSource(v string) *PreviewInvoiceRequest {
	s.Source = &v
	return s
}

func (s *PreviewInvoiceRequest) SetOperatorNo(v string) *PreviewInvoiceRequest {
	s.OperatorNo = &v
	return s
}

func (s *PreviewInvoiceRequest) SetOperatorName(v string) *PreviewInvoiceRequest {
	s.OperatorName = &v
	return s
}

func (s *PreviewInvoiceRequest) SetOutBizNo(v string) *PreviewInvoiceRequest {
	s.OutBizNo = &v
	return s
}

func (s *PreviewInvoiceRequest) SetInvoiceType(v string) *PreviewInvoiceRequest {
	s.InvoiceType = &v
	return s
}

func (s *PreviewInvoiceRequest) SetInvoiceMaterial(v string) *PreviewInvoiceRequest {
	s.InvoiceMaterial = &v
	return s
}

func (s *PreviewInvoiceRequest) SetInvoiceNote(v string) *PreviewInvoiceRequest {
	s.InvoiceNote = &v
	return s
}

func (s *PreviewInvoiceRequest) SetApplyReason(v string) *PreviewInvoiceRequest {
	s.ApplyReason = &v
	return s
}

func (s *PreviewInvoiceRequest) SetInvoiceBuyer(v *ApplyInvoiceBuyer) *PreviewInvoiceRequest {
	s.InvoiceBuyer = v
	return s
}

func (s *PreviewInvoiceRequest) SetApplyInvoiceDelivery(v *ApplyInvoiceDelivery) *PreviewInvoiceRequest {
	s.ApplyInvoiceDelivery = v
	return s
}

func (s *PreviewInvoiceRequest) SetApplyInvoiceQuota(v *ApplyInvoiceAmountAllocation) *PreviewInvoiceRequest {
	s.ApplyInvoiceQuota = v
	return s
}

func (s *PreviewInvoiceRequest) SetPreviewInvoiceLines(v []*PreviewInvoiceLine) *PreviewInvoiceRequest {
	s.PreviewInvoiceLines = v
	return s
}

// 退换票信息
type ReturnInvoiceInfo struct {
	// 发票ID
	// example:
	//
	// 33445566556
	InvoiceId *string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty" require:"true"`
	// 退换票类型,01:退票 02:换票
	// example:
	//
	// 01
	ReturnOrderType *string `json:"return_order_type,omitempty" xml:"return_order_type,omitempty" require:"true"`
	// 退换票原因类型, 01：发票介质修改  02：发票类型修改  03：发票信息修改 05：其他 06：不需要发票
	// example:
	//
	// 06
	ReturnReasonType *string `json:"return_reason_type,omitempty" xml:"return_reason_type,omitempty" require:"true"`
	// 备注
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 快递单号
	// example:
	//
	// 3444
	TrackingNo *string `json:"tracking_no,omitempty" xml:"tracking_no,omitempty"`
	// 快递公司名称
	// example:
	//
	// XX快递
	ExpressCompanyName *string `json:"express_company_name,omitempty" xml:"express_company_name,omitempty"`
	// 是否认证，1：已认证抵扣 0：未认证抵扣
	// example:
	//
	// 0
	Auth *string `json:"auth,omitempty" xml:"auth,omitempty"`
}

func (s ReturnInvoiceInfo) String() string {
	return tea.Prettify(s)
}

func (s ReturnInvoiceInfo) GoString() string {
	return s.String()
}

func (s *ReturnInvoiceInfo) SetInvoiceId(v string) *ReturnInvoiceInfo {
	s.InvoiceId = &v
	return s
}

func (s *ReturnInvoiceInfo) SetReturnOrderType(v string) *ReturnInvoiceInfo {
	s.ReturnOrderType = &v
	return s
}

func (s *ReturnInvoiceInfo) SetReturnReasonType(v string) *ReturnInvoiceInfo {
	s.ReturnReasonType = &v
	return s
}

func (s *ReturnInvoiceInfo) SetMemo(v string) *ReturnInvoiceInfo {
	s.Memo = &v
	return s
}

func (s *ReturnInvoiceInfo) SetTrackingNo(v string) *ReturnInvoiceInfo {
	s.TrackingNo = &v
	return s
}

func (s *ReturnInvoiceInfo) SetExpressCompanyName(v string) *ReturnInvoiceInfo {
	s.ExpressCompanyName = &v
	return s
}

func (s *ReturnInvoiceInfo) SetAuth(v string) *ReturnInvoiceInfo {
	s.Auth = &v
	return s
}

// 发票介质
type InvoiceMaterialVO struct {
	// 发票介质，01: 电子发票; 02: 纸质发票
	// example:
	//
	// 01
	Material *string `json:"material,omitempty" xml:"material,omitempty" require:"true"`
	// 名称
	// example:
	//
	// 电子发票
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s InvoiceMaterialVO) String() string {
	return tea.Prettify(s)
}

func (s InvoiceMaterialVO) GoString() string {
	return s.String()
}

func (s *InvoiceMaterialVO) SetMaterial(v string) *InvoiceMaterialVO {
	s.Material = &v
	return s
}

func (s *InvoiceMaterialVO) SetName(v string) *InvoiceMaterialVO {
	s.Name = &v
	return s
}

// 国家信息
type CountryCnEnItem struct {
	// 国家字母编号
	// example:
	//
	// CN
	CountryCode *string `json:"country_code,omitempty" xml:"country_code,omitempty" require:"true"`
	// 国家中文名称
	// example:
	//
	// 中国
	CountryCn *string `json:"country_cn,omitempty" xml:"country_cn,omitempty" require:"true"`
	// 国家英文名称
	// example:
	//
	// China
	CountryEn *string `json:"country_en,omitempty" xml:"country_en,omitempty" require:"true"`
}

func (s CountryCnEnItem) String() string {
	return tea.Prettify(s)
}

func (s CountryCnEnItem) GoString() string {
	return s.String()
}

func (s *CountryCnEnItem) SetCountryCode(v string) *CountryCnEnItem {
	s.CountryCode = &v
	return s
}

func (s *CountryCnEnItem) SetCountryCn(v string) *CountryCnEnItem {
	s.CountryCn = &v
	return s
}

func (s *CountryCnEnItem) SetCountryEn(v string) *CountryCnEnItem {
	s.CountryEn = &v
	return s
}

// 发票预览信息
type InvoicePreviewVO struct {
	// 发票号码
	// example:
	//
	// 3344
	InvoiceNo *string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 发票代码
	// example:
	//
	// 323422244555
	InvoiceCode *string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// 开票时间
	// example:
	//
	// 2023-09-08
	InvoiceDate *string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
	// 发票类型
	// example:
	//
	// 01
	InvoiceType *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty" require:"true"`
	// 开票金额
	// example:
	//
	// 34.98
	InvoiceAmt *string `json:"invoice_amt,omitempty" xml:"invoice_amt,omitempty" require:"true"`
	// 税额
	// example:
	//
	// 34
	TaxAmt *string `json:"tax_amt,omitempty" xml:"tax_amt,omitempty" require:"true"`
	// 销方信息
	InvoiceSeller *ApplyInvoiceSeller `json:"invoice_seller,omitempty" xml:"invoice_seller,omitempty" require:"true"`
	// 购方信息
	InvoiceBuyer *ApplyInvoiceBuyer `json:"invoice_buyer,omitempty" xml:"invoice_buyer,omitempty" require:"true"`
	// 发票票面备注
	// example:
	//
	// 备注
	InvoiceNote *string `json:"invoice_note,omitempty" xml:"invoice_note,omitempty"`
	// 租户id
	// example:
	//
	// 2088720671581149
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty" require:"true"`
	// 发票ID
	// example:
	//
	// 355566677676
	InvoiceId *string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 发票介质
	// example:
	//
	// 01
	InvoiceMaterial *string `json:"invoice_material,omitempty" xml:"invoice_material,omitempty" require:"true"`
	// 申请原因
	// example:
	//
	// 申请原因
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 发票行信息
	InvoiceLines []*InvoiceLineVO `json:"invoice_lines,omitempty" xml:"invoice_lines,omitempty" require:"true" type:"Repeated"`
	// 发票预览记录号
	// example:
	//
	// 3344
	InvoicePreviewLogNo *string `json:"invoice_preview_log_no,omitempty" xml:"invoice_preview_log_no,omitempty" require:"true"`
}

func (s InvoicePreviewVO) String() string {
	return tea.Prettify(s)
}

func (s InvoicePreviewVO) GoString() string {
	return s.String()
}

func (s *InvoicePreviewVO) SetInvoiceNo(v string) *InvoicePreviewVO {
	s.InvoiceNo = &v
	return s
}

func (s *InvoicePreviewVO) SetInvoiceCode(v string) *InvoicePreviewVO {
	s.InvoiceCode = &v
	return s
}

func (s *InvoicePreviewVO) SetInvoiceDate(v string) *InvoicePreviewVO {
	s.InvoiceDate = &v
	return s
}

func (s *InvoicePreviewVO) SetInvoiceType(v string) *InvoicePreviewVO {
	s.InvoiceType = &v
	return s
}

func (s *InvoicePreviewVO) SetInvoiceAmt(v string) *InvoicePreviewVO {
	s.InvoiceAmt = &v
	return s
}

func (s *InvoicePreviewVO) SetTaxAmt(v string) *InvoicePreviewVO {
	s.TaxAmt = &v
	return s
}

func (s *InvoicePreviewVO) SetInvoiceSeller(v *ApplyInvoiceSeller) *InvoicePreviewVO {
	s.InvoiceSeller = v
	return s
}

func (s *InvoicePreviewVO) SetInvoiceBuyer(v *ApplyInvoiceBuyer) *InvoicePreviewVO {
	s.InvoiceBuyer = v
	return s
}

func (s *InvoicePreviewVO) SetInvoiceNote(v string) *InvoicePreviewVO {
	s.InvoiceNote = &v
	return s
}

func (s *InvoicePreviewVO) SetTenantId(v string) *InvoicePreviewVO {
	s.TenantId = &v
	return s
}

func (s *InvoicePreviewVO) SetInvoiceId(v string) *InvoicePreviewVO {
	s.InvoiceId = &v
	return s
}

func (s *InvoicePreviewVO) SetInvoiceMaterial(v string) *InvoicePreviewVO {
	s.InvoiceMaterial = &v
	return s
}

func (s *InvoicePreviewVO) SetMemo(v string) *InvoicePreviewVO {
	s.Memo = &v
	return s
}

func (s *InvoicePreviewVO) SetInvoiceLines(v []*InvoiceLineVO) *InvoicePreviewVO {
	s.InvoiceLines = v
	return s
}

func (s *InvoicePreviewVO) SetInvoicePreviewLogNo(v string) *InvoicePreviewVO {
	s.InvoicePreviewLogNo = &v
	return s
}

// 开票申请项
type InvoiceApplyItem struct {
	// 商品名称
	// example:
	//
	// xxx
	ItemName *string `json:"item_name,omitempty" xml:"item_name,omitempty" require:"true"`
	// 不含税金额
	// example:
	//
	// 5520
	ExcludingTaxOfAmount *string `json:"excluding_tax_of_amount,omitempty" xml:"excluding_tax_of_amount,omitempty" require:"true"`
	// 商品CODE
	// example:
	//
	// 001
	ItemCode *string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 含税金额
	// example:
	//
	// 3452
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 关联单据号
	// example:
	//
	// 123123
	InvoiceRcptNo *string `json:"invoice_rcpt_no,omitempty" xml:"invoice_rcpt_no,omitempty"`
}

func (s InvoiceApplyItem) String() string {
	return tea.Prettify(s)
}

func (s InvoiceApplyItem) GoString() string {
	return s.String()
}

func (s *InvoiceApplyItem) SetItemName(v string) *InvoiceApplyItem {
	s.ItemName = &v
	return s
}

func (s *InvoiceApplyItem) SetExcludingTaxOfAmount(v string) *InvoiceApplyItem {
	s.ExcludingTaxOfAmount = &v
	return s
}

func (s *InvoiceApplyItem) SetItemCode(v string) *InvoiceApplyItem {
	s.ItemCode = &v
	return s
}

func (s *InvoiceApplyItem) SetAmount(v string) *InvoiceApplyItem {
	s.Amount = &v
	return s
}

func (s *InvoiceApplyItem) SetInvoiceRcptNo(v string) *InvoiceApplyItem {
	s.InvoiceRcptNo = &v
	return s
}

// 国际可开票单据
type IntlRcptDetailItem struct {
	// 可开票单据号
	// example:
	//
	// 33345454
	ReceiptNo *string `json:"receipt_no,omitempty" xml:"receipt_no,omitempty" require:"true"`
	// 租户ID
	// example:
	//
	// 20882838383
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty" require:"true"`
	// 合同号或订单号
	// example:
	//
	// 3333
	ArNo *string `json:"ar_no,omitempty" xml:"ar_no,omitempty" require:"true"`
	// 计费类型
	// example:
	//
	// PREPAY/AFTER_PAY
	ChargeType *string `json:"charge_type,omitempty" xml:"charge_type,omitempty" require:"true"`
	// 合同商品总金额
	// example:
	//
	// 34.99
	RcptContractAmt *string `json:"rcpt_contract_amt,omitempty" xml:"rcpt_contract_amt,omitempty" require:"true"`
	// 合同商品已开票金额
	// example:
	//
	// 22.12
	ReptInvedAmt *string `json:"rept_inved_amt,omitempty" xml:"rept_inved_amt,omitempty" require:"true"`
	// 合同商品剩余可开票金额
	// example:
	//
	// 23.21
	RcptContractRemainAmt *string `json:"rcpt_contract_remain_amt,omitempty" xml:"rcpt_contract_remain_amt,omitempty" require:"true"`
	// 商品CODE
	// example:
	//
	// abd
	CommodityCode *string `json:"commodity_code,omitempty" xml:"commodity_code,omitempty" require:"true"`
	// 商品名称
	// example:
	//
	// 测试商品
	CommodityName *string `json:"commodity_name,omitempty" xml:"commodity_name,omitempty" require:"true"`
	// ou
	// example:
	//
	// Z37
	Ou *string `json:"ou,omitempty" xml:"ou,omitempty" require:"true"`
	// 币种
	// example:
	//
	// 840
	Ccy *string `json:"ccy,omitempty" xml:"ccy,omitempty" require:"true"`
	// 税率
	// example:
	//
	// 0.09
	Rate *string `json:"rate,omitempty" xml:"rate,omitempty" require:"true"`
	// 合同金额（不含税）
	// example:
	//
	// 22
	RcptContractAmtExclTax *string `json:"rcpt_contract_amt_excl_tax,omitempty" xml:"rcpt_contract_amt_excl_tax,omitempty" require:"true"`
	// 已开票金额（不含税）
	// example:
	//
	// 23
	ReptInvedAmtExclTax *string `json:"rept_inved_amt_excl_tax,omitempty" xml:"rept_inved_amt_excl_tax,omitempty" require:"true"`
	// 剩余可开票金额（不含税）
	// example:
	//
	// 23
	RcptContractRemainAmtExclTax *string `json:"rcpt_contract_remain_amt_excl_tax,omitempty" xml:"rcpt_contract_remain_amt_excl_tax,omitempty" require:"true"`
}

func (s IntlRcptDetailItem) String() string {
	return tea.Prettify(s)
}

func (s IntlRcptDetailItem) GoString() string {
	return s.String()
}

func (s *IntlRcptDetailItem) SetReceiptNo(v string) *IntlRcptDetailItem {
	s.ReceiptNo = &v
	return s
}

func (s *IntlRcptDetailItem) SetTenantId(v string) *IntlRcptDetailItem {
	s.TenantId = &v
	return s
}

func (s *IntlRcptDetailItem) SetArNo(v string) *IntlRcptDetailItem {
	s.ArNo = &v
	return s
}

func (s *IntlRcptDetailItem) SetChargeType(v string) *IntlRcptDetailItem {
	s.ChargeType = &v
	return s
}

func (s *IntlRcptDetailItem) SetRcptContractAmt(v string) *IntlRcptDetailItem {
	s.RcptContractAmt = &v
	return s
}

func (s *IntlRcptDetailItem) SetReptInvedAmt(v string) *IntlRcptDetailItem {
	s.ReptInvedAmt = &v
	return s
}

func (s *IntlRcptDetailItem) SetRcptContractRemainAmt(v string) *IntlRcptDetailItem {
	s.RcptContractRemainAmt = &v
	return s
}

func (s *IntlRcptDetailItem) SetCommodityCode(v string) *IntlRcptDetailItem {
	s.CommodityCode = &v
	return s
}

func (s *IntlRcptDetailItem) SetCommodityName(v string) *IntlRcptDetailItem {
	s.CommodityName = &v
	return s
}

func (s *IntlRcptDetailItem) SetOu(v string) *IntlRcptDetailItem {
	s.Ou = &v
	return s
}

func (s *IntlRcptDetailItem) SetCcy(v string) *IntlRcptDetailItem {
	s.Ccy = &v
	return s
}

func (s *IntlRcptDetailItem) SetRate(v string) *IntlRcptDetailItem {
	s.Rate = &v
	return s
}

func (s *IntlRcptDetailItem) SetRcptContractAmtExclTax(v string) *IntlRcptDetailItem {
	s.RcptContractAmtExclTax = &v
	return s
}

func (s *IntlRcptDetailItem) SetReptInvedAmtExclTax(v string) *IntlRcptDetailItem {
	s.ReptInvedAmtExclTax = &v
	return s
}

func (s *IntlRcptDetailItem) SetRcptContractRemainAmtExclTax(v string) *IntlRcptDetailItem {
	s.RcptContractRemainAmtExclTax = &v
	return s
}

// 金额类
type MultiCurrencyMoneyOpenApi struct {
	// 最小币种单位
	// example:
	//
	// 1233
	Cent *string `json:"cent,omitempty" xml:"cent,omitempty" require:"true"`
	// 币种
	// example:
	//
	// 156
	CurrencyValue *string `json:"currency_value,omitempty" xml:"currency_value,omitempty" require:"true"`
}

func (s MultiCurrencyMoneyOpenApi) String() string {
	return tea.Prettify(s)
}

func (s MultiCurrencyMoneyOpenApi) GoString() string {
	return s.String()
}

func (s *MultiCurrencyMoneyOpenApi) SetCent(v string) *MultiCurrencyMoneyOpenApi {
	s.Cent = &v
	return s
}

func (s *MultiCurrencyMoneyOpenApi) SetCurrencyValue(v string) *MultiCurrencyMoneyOpenApi {
	s.CurrencyValue = &v
	return s
}

// 开票单据明细
type RcptDetailVO struct {
	// 合同号
	// example:
	//
	// 2088720671581149-ZNHYFM01222234
	ArNo *string `json:"ar_no,omitempty" xml:"ar_no,omitempty" require:"true"`
	// OU
	// example:
	//
	// ZL6
	Ou *string `json:"ou,omitempty" xml:"ou,omitempty" require:"true"`
	// 租户ID
	// example:
	//
	// 2088720671581149
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty" require:"true"`
	// 客户名称
	// example:
	//
	// 测试客户
	TenantName *string `json:"tenant_name,omitempty" xml:"tenant_name,omitempty" require:"true"`
	// 单据唯一号
	// example:
	//
	// 20230928107305000028710015937380
	ReceiptNo *string `json:"receipt_no,omitempty" xml:"receipt_no,omitempty" require:"true"`
	// 币种
	// example:
	//
	// 156
	Ccy *string `json:"ccy,omitempty" xml:"ccy,omitempty" require:"true"`
	// 税率
	// example:
	//
	// 0.06
	Tax *string `json:"tax,omitempty" xml:"tax,omitempty"`
	// 商品CODE
	// example:
	//
	// SQBRFSZL601262688
	CommodityCode *string `json:"commodity_code,omitempty" xml:"commodity_code,omitempty" require:"true"`
	// 开票模式，01：先款后票、02：先票后款
	// example:
	//
	// 01
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty" require:"true"`
	// 单据可开票总金额
	// example:
	//
	// 13.45
	TotalAmt *string `json:"total_amt,omitempty" xml:"total_amt,omitempty" require:"true"`
	// 已开票金额
	// example:
	//
	// 2.23
	InvedAmt *string `json:"inved_amt,omitempty" xml:"inved_amt,omitempty" require:"true"`
	// 剩余可开票金额
	// example:
	//
	// 234.99
	RemainAmt *string `json:"remain_amt,omitempty" xml:"remain_amt,omitempty" require:"true"`
	// 商品名称
	// example:
	//
	// 商品名称
	CommodityName *string `json:"commodity_name,omitempty" xml:"commodity_name,omitempty" require:"true"`
}

func (s RcptDetailVO) String() string {
	return tea.Prettify(s)
}

func (s RcptDetailVO) GoString() string {
	return s.String()
}

func (s *RcptDetailVO) SetArNo(v string) *RcptDetailVO {
	s.ArNo = &v
	return s
}

func (s *RcptDetailVO) SetOu(v string) *RcptDetailVO {
	s.Ou = &v
	return s
}

func (s *RcptDetailVO) SetTenantId(v string) *RcptDetailVO {
	s.TenantId = &v
	return s
}

func (s *RcptDetailVO) SetTenantName(v string) *RcptDetailVO {
	s.TenantName = &v
	return s
}

func (s *RcptDetailVO) SetReceiptNo(v string) *RcptDetailVO {
	s.ReceiptNo = &v
	return s
}

func (s *RcptDetailVO) SetCcy(v string) *RcptDetailVO {
	s.Ccy = &v
	return s
}

func (s *RcptDetailVO) SetTax(v string) *RcptDetailVO {
	s.Tax = &v
	return s
}

func (s *RcptDetailVO) SetCommodityCode(v string) *RcptDetailVO {
	s.CommodityCode = &v
	return s
}

func (s *RcptDetailVO) SetMode(v string) *RcptDetailVO {
	s.Mode = &v
	return s
}

func (s *RcptDetailVO) SetTotalAmt(v string) *RcptDetailVO {
	s.TotalAmt = &v
	return s
}

func (s *RcptDetailVO) SetInvedAmt(v string) *RcptDetailVO {
	s.InvedAmt = &v
	return s
}

func (s *RcptDetailVO) SetRemainAmt(v string) *RcptDetailVO {
	s.RemainAmt = &v
	return s
}

func (s *RcptDetailVO) SetCommodityName(v string) *RcptDetailVO {
	s.CommodityName = &v
	return s
}

// 发票申请
type ApplyInvoiceRequest struct {
	// 租户ID
	// example:
	//
	// 2088720671581149
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty" require:"true"`
	// OU
	// example:
	//
	// ZL6
	Ou *string `json:"ou,omitempty" xml:"ou,omitempty" require:"true"`
	// 合同号
	// example:
	//
	// 2088720671581149-ZNHYFM01222234
	ArNo *string `json:"ar_no,omitempty" xml:"ar_no,omitempty" require:"true"`
	// 开票业务场景
	// example:
	//
	// LTC_RCPT_BILL
	BizScene *string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty" require:"true"`
	// 业务操作
	// example:
	//
	// PREVIEW_INVOICING
	BizAction *string `json:"biz_action,omitempty" xml:"biz_action,omitempty" require:"true"`
	// 系统来源
	// example:
	//
	// IOT
	Source *string `json:"source,omitempty" xml:"source,omitempty" require:"true"`
	// 操作人员工号
	// example:
	//
	// 2334
	OperatorNo *string `json:"operator_no,omitempty" xml:"operator_no,omitempty" require:"true"`
	// 操作人员名称
	// example:
	//
	// 测试
	OperatorName *string `json:"operator_name,omitempty" xml:"operator_name,omitempty" require:"true"`
	// 外部业务号
	// example:
	//
	// 34445
	InvoiceApplyBizNo *string `json:"invoice_apply_biz_no,omitempty" xml:"invoice_apply_biz_no,omitempty" require:"true"`
	// 发票介质
	// example:
	//
	// 01
	InvoiceMaterial *string `json:"invoice_material,omitempty" xml:"invoice_material,omitempty" require:"true"`
	// 申请原因
	// example:
	//
	// 原因
	ApplyReason *string `json:"apply_reason,omitempty" xml:"apply_reason,omitempty"`
	// 寄送信息
	ApplyInvoiceDelivery *ApplyInvoiceDelivery `json:"apply_invoice_delivery,omitempty" xml:"apply_invoice_delivery,omitempty" require:"true"`
	// 发票需要占用的额度信息
	ApplyInvoiceQuota *ApplyInvoiceAmountAllocation `json:"apply_invoice_quota,omitempty" xml:"apply_invoice_quota,omitempty" require:"true"`
	// 发票信息列表，一次申请可能会拆分出多张票
	ApplyInvoices []*ApplyInvoiceLine `json:"apply_invoices,omitempty" xml:"apply_invoices,omitempty" require:"true" type:"Repeated"`
	// 预览记录号
	// example:
	//
	// 20234566767
	PreviewLogNo *string `json:"preview_log_no,omitempty" xml:"preview_log_no,omitempty" require:"true"`
}

func (s ApplyInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoiceRequest) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceRequest) SetTenantId(v string) *ApplyInvoiceRequest {
	s.TenantId = &v
	return s
}

func (s *ApplyInvoiceRequest) SetOu(v string) *ApplyInvoiceRequest {
	s.Ou = &v
	return s
}

func (s *ApplyInvoiceRequest) SetArNo(v string) *ApplyInvoiceRequest {
	s.ArNo = &v
	return s
}

func (s *ApplyInvoiceRequest) SetBizScene(v string) *ApplyInvoiceRequest {
	s.BizScene = &v
	return s
}

func (s *ApplyInvoiceRequest) SetBizAction(v string) *ApplyInvoiceRequest {
	s.BizAction = &v
	return s
}

func (s *ApplyInvoiceRequest) SetSource(v string) *ApplyInvoiceRequest {
	s.Source = &v
	return s
}

func (s *ApplyInvoiceRequest) SetOperatorNo(v string) *ApplyInvoiceRequest {
	s.OperatorNo = &v
	return s
}

func (s *ApplyInvoiceRequest) SetOperatorName(v string) *ApplyInvoiceRequest {
	s.OperatorName = &v
	return s
}

func (s *ApplyInvoiceRequest) SetInvoiceApplyBizNo(v string) *ApplyInvoiceRequest {
	s.InvoiceApplyBizNo = &v
	return s
}

func (s *ApplyInvoiceRequest) SetInvoiceMaterial(v string) *ApplyInvoiceRequest {
	s.InvoiceMaterial = &v
	return s
}

func (s *ApplyInvoiceRequest) SetApplyReason(v string) *ApplyInvoiceRequest {
	s.ApplyReason = &v
	return s
}

func (s *ApplyInvoiceRequest) SetApplyInvoiceDelivery(v *ApplyInvoiceDelivery) *ApplyInvoiceRequest {
	s.ApplyInvoiceDelivery = v
	return s
}

func (s *ApplyInvoiceRequest) SetApplyInvoiceQuota(v *ApplyInvoiceAmountAllocation) *ApplyInvoiceRequest {
	s.ApplyInvoiceQuota = v
	return s
}

func (s *ApplyInvoiceRequest) SetApplyInvoices(v []*ApplyInvoiceLine) *ApplyInvoiceRequest {
	s.ApplyInvoices = v
	return s
}

func (s *ApplyInvoiceRequest) SetPreviewLogNo(v string) *ApplyInvoiceRequest {
	s.PreviewLogNo = &v
	return s
}

// 用户开票信息
type UserInvoiceInfo struct {
	// 公司标题（发票抬头）
	// example:
	//
	// xxxx Group Limited
	Title *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	// 纳税人类型
	// example:
	//
	// 03
	TaxPayerQualification *string `json:"tax_payer_qualification,omitempty" xml:"tax_payer_qualification,omitempty" require:"true"`
	// 注册国家编号
	// example:
	//
	// HK
	RegisterCountry *string `json:"register_country,omitempty" xml:"register_country,omitempty" require:"true"`
	// 公司注册地址
	// example:
	//
	// xxxxxx
	Address *string `json:"address,omitempty" xml:"address,omitempty" require:"true"`
	// 纳税人识别号
	// example:
	//
	// 123
	TaxNo *string `json:"tax_no,omitempty" xml:"tax_no,omitempty"`
	// 公司注册电话
	// example:
	//
	// 17797768855
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 开户行
	// example:
	//
	// 中国人民银行
	BankName *string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// 银行账号
	// example:
	//
	// 12312312
	BankAccount *string `json:"bank_account,omitempty" xml:"bank_account,omitempty"`
}

func (s UserInvoiceInfo) String() string {
	return tea.Prettify(s)
}

func (s UserInvoiceInfo) GoString() string {
	return s.String()
}

func (s *UserInvoiceInfo) SetTitle(v string) *UserInvoiceInfo {
	s.Title = &v
	return s
}

func (s *UserInvoiceInfo) SetTaxPayerQualification(v string) *UserInvoiceInfo {
	s.TaxPayerQualification = &v
	return s
}

func (s *UserInvoiceInfo) SetRegisterCountry(v string) *UserInvoiceInfo {
	s.RegisterCountry = &v
	return s
}

func (s *UserInvoiceInfo) SetAddress(v string) *UserInvoiceInfo {
	s.Address = &v
	return s
}

func (s *UserInvoiceInfo) SetTaxNo(v string) *UserInvoiceInfo {
	s.TaxNo = &v
	return s
}

func (s *UserInvoiceInfo) SetTelephone(v string) *UserInvoiceInfo {
	s.Telephone = &v
	return s
}

func (s *UserInvoiceInfo) SetBankName(v string) *UserInvoiceInfo {
	s.BankName = &v
	return s
}

func (s *UserInvoiceInfo) SetBankAccount(v string) *UserInvoiceInfo {
	s.BankAccount = &v
	return s
}

// 发票申请信息
type InvoiceApplyInfoVO struct {
	// 租户ID
	// example:
	//
	// 2088720671581149
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty" require:"true"`
	// 发票申请ID
	// example:
	//
	// 2234345667
	InvoiceApplyId *string `json:"invoice_apply_id,omitempty" xml:"invoice_apply_id,omitempty"`
	// 申请金额
	// example:
	//
	// 34.98
	InvoiceAmt *string `json:"invoice_amt,omitempty" xml:"invoice_amt,omitempty" require:"true"`
	// 发票申请日期
	// example:
	//
	// 2018-10-10T10:10:00Z
	InvoiceApplyDate *string `json:"invoice_apply_date,omitempty" xml:"invoice_apply_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 发票类型
	// example:
	//
	// 01
	InvoiceType *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty" require:"true"`
	// 申请类型
	// example:
	//
	// new
	ApplyType *string `json:"apply_type,omitempty" xml:"apply_type,omitempty"`
	// 申请人名称
	// example:
	//
	// 测试
	OperatorName *string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 申请人ID
	// example:
	//
	// 2088720671581149
	OperatorId *string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 申请状态，01：处理中，03：已完成，04：审批中，05：已驳回，06：已撤回，10：申请终止
	// example:
	//
	// 03
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// 关联的发票列表
	RelateInvoices []*InvoiceInfoVO `json:"relate_invoices,omitempty" xml:"relate_invoices,omitempty" type:"Repeated"`
	// 业务号
	// example:
	//
	// 323422244555
	BsnNo *string `json:"bsn_no,omitempty" xml:"bsn_no,omitempty" require:"true"`
}

func (s InvoiceApplyInfoVO) String() string {
	return tea.Prettify(s)
}

func (s InvoiceApplyInfoVO) GoString() string {
	return s.String()
}

func (s *InvoiceApplyInfoVO) SetTenantId(v string) *InvoiceApplyInfoVO {
	s.TenantId = &v
	return s
}

func (s *InvoiceApplyInfoVO) SetInvoiceApplyId(v string) *InvoiceApplyInfoVO {
	s.InvoiceApplyId = &v
	return s
}

func (s *InvoiceApplyInfoVO) SetInvoiceAmt(v string) *InvoiceApplyInfoVO {
	s.InvoiceAmt = &v
	return s
}

func (s *InvoiceApplyInfoVO) SetInvoiceApplyDate(v string) *InvoiceApplyInfoVO {
	s.InvoiceApplyDate = &v
	return s
}

func (s *InvoiceApplyInfoVO) SetInvoiceType(v string) *InvoiceApplyInfoVO {
	s.InvoiceType = &v
	return s
}

func (s *InvoiceApplyInfoVO) SetApplyType(v string) *InvoiceApplyInfoVO {
	s.ApplyType = &v
	return s
}

func (s *InvoiceApplyInfoVO) SetOperatorName(v string) *InvoiceApplyInfoVO {
	s.OperatorName = &v
	return s
}

func (s *InvoiceApplyInfoVO) SetOperatorId(v string) *InvoiceApplyInfoVO {
	s.OperatorId = &v
	return s
}

func (s *InvoiceApplyInfoVO) SetStatus(v string) *InvoiceApplyInfoVO {
	s.Status = &v
	return s
}

func (s *InvoiceApplyInfoVO) SetRelateInvoices(v []*InvoiceInfoVO) *InvoiceApplyInfoVO {
	s.RelateInvoices = v
	return s
}

func (s *InvoiceApplyInfoVO) SetBsnNo(v string) *InvoiceApplyInfoVO {
	s.BsnNo = &v
	return s
}

// 申请的发票信息
type ApplyInvoice struct {
	// 发票类型，01,增值税专用发票; * 02,增值税普通发票; * 04,国际形式发票; * 05,其它发票
	// example:
	//
	// 01
	InvoiceType *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty" require:"true"`
	// 发票金额
	// example:
	//
	// 3244.98
	InvoiceAmt *string `json:"invoice_amt,omitempty" xml:"invoice_amt,omitempty" require:"true"`
	// 发票行信息
	ApplyInvoiceLines []*ApplyInvoiceLine `json:"apply_invoice_lines,omitempty" xml:"apply_invoice_lines,omitempty" require:"true" type:"Repeated"`
	// 销方信息
	ApplyInvoiceSeller *ApplyInvoiceSeller `json:"apply_invoice_seller,omitempty" xml:"apply_invoice_seller,omitempty" require:"true"`
	// 购方信息
	ApplyInvoiceBuyer *ApplyInvoiceBuyer `json:"apply_invoice_buyer,omitempty" xml:"apply_invoice_buyer,omitempty" require:"true"`
	// 发票备注
	// example:
	//
	// 备注
	InvoiceNote *string `json:"invoice_note,omitempty" xml:"invoice_note,omitempty"`
	// 币种
	// example:
	//
	// 156
	Ccy *string `json:"ccy,omitempty" xml:"ccy,omitempty"`
	// 发票ID
	// example:
	//
	// 334556666
	ApplyInvoiceId *string `json:"apply_invoice_id,omitempty" xml:"apply_invoice_id,omitempty" require:"true"`
}

func (s ApplyInvoice) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoice) GoString() string {
	return s.String()
}

func (s *ApplyInvoice) SetInvoiceType(v string) *ApplyInvoice {
	s.InvoiceType = &v
	return s
}

func (s *ApplyInvoice) SetInvoiceAmt(v string) *ApplyInvoice {
	s.InvoiceAmt = &v
	return s
}

func (s *ApplyInvoice) SetApplyInvoiceLines(v []*ApplyInvoiceLine) *ApplyInvoice {
	s.ApplyInvoiceLines = v
	return s
}

func (s *ApplyInvoice) SetApplyInvoiceSeller(v *ApplyInvoiceSeller) *ApplyInvoice {
	s.ApplyInvoiceSeller = v
	return s
}

func (s *ApplyInvoice) SetApplyInvoiceBuyer(v *ApplyInvoiceBuyer) *ApplyInvoice {
	s.ApplyInvoiceBuyer = v
	return s
}

func (s *ApplyInvoice) SetInvoiceNote(v string) *ApplyInvoice {
	s.InvoiceNote = &v
	return s
}

func (s *ApplyInvoice) SetCcy(v string) *ApplyInvoice {
	s.Ccy = &v
	return s
}

func (s *ApplyInvoice) SetApplyInvoiceId(v string) *ApplyInvoice {
	s.ApplyInvoiceId = &v
	return s
}

// 客户的开票配置信息
type UserInvoiceConfigVO struct {
	// 租户ID
	// example:
	//
	// 2088720671581149
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty" require:"true"`
	// 企业名称
	// example:
	//
	// XX公司
	CompanyName *string `json:"company_name,omitempty" xml:"company_name,omitempty" require:"true"`
	// 企业电话
	// example:
	//
	// 0571-877776
	CompanyPhoneNo *string `json:"company_phone_no,omitempty" xml:"company_phone_no,omitempty"`
	// 公司地址(详细地址)
	// example:
	//
	// 上海市黄浦区外马路
	CompanyAddress *string `json:"company_address,omitempty" xml:"company_address,omitempty"`
	// 银行名称
	// example:
	//
	// 招商银行
	BankName *string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// 银行账号
	// example:
	//
	// 34355565
	BankAccount *string `json:"bank_account,omitempty" xml:"bank_account,omitempty"`
	// 税号
	// example:
	//
	// 35556T5
	TaxNo *string `json:"tax_no,omitempty" xml:"tax_no,omitempty"`
	// 是否是一般纳税人
	// example:
	//
	// true, false
	GeneralTaxpayer *bool `json:"general_taxpayer,omitempty" xml:"general_taxpayer,omitempty" require:"true"`
}

func (s UserInvoiceConfigVO) String() string {
	return tea.Prettify(s)
}

func (s UserInvoiceConfigVO) GoString() string {
	return s.String()
}

func (s *UserInvoiceConfigVO) SetTenantId(v string) *UserInvoiceConfigVO {
	s.TenantId = &v
	return s
}

func (s *UserInvoiceConfigVO) SetCompanyName(v string) *UserInvoiceConfigVO {
	s.CompanyName = &v
	return s
}

func (s *UserInvoiceConfigVO) SetCompanyPhoneNo(v string) *UserInvoiceConfigVO {
	s.CompanyPhoneNo = &v
	return s
}

func (s *UserInvoiceConfigVO) SetCompanyAddress(v string) *UserInvoiceConfigVO {
	s.CompanyAddress = &v
	return s
}

func (s *UserInvoiceConfigVO) SetBankName(v string) *UserInvoiceConfigVO {
	s.BankName = &v
	return s
}

func (s *UserInvoiceConfigVO) SetBankAccount(v string) *UserInvoiceConfigVO {
	s.BankAccount = &v
	return s
}

func (s *UserInvoiceConfigVO) SetTaxNo(v string) *UserInvoiceConfigVO {
	s.TaxNo = &v
	return s
}

func (s *UserInvoiceConfigVO) SetGeneralTaxpayer(v bool) *UserInvoiceConfigVO {
	s.GeneralTaxpayer = &v
	return s
}

// 发票邮寄信息
type InvoiceMailInfo struct {
	// PID，指定发票邮寄地址归属的PID
	// example:
	//
	// 123
	Pid *string `json:"pid,omitempty" xml:"pid,omitempty" require:"true"`
	// 联系人名字
	// example:
	//
	// 张三
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	// 国家
	// example:
	//
	// 中国
	Country *string `json:"country,omitempty" xml:"country,omitempty" require:"true"`
	// 收件人邮箱
	// example:
	//
	// 715680094@qq.com
	Email *string `json:"email,omitempty" xml:"email,omitempty" require:"true"`
	// 联系人电话
	// example:
	//
	// 17797768851
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 省份
	// example:
	//
	// 河南
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
	// 城市
	// example:
	//
	// 郑州
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// 区/县
	// example:
	//
	// 金水区
	CountyDistrict *string `json:"county_district,omitempty" xml:"county_district,omitempty"`
	// 街道
	// example:
	//
	// 良秀路街道
	Street *string `json:"street,omitempty" xml:"street,omitempty"`
	// 详细地址
	// example:
	//
	// 良秀路180号
	DetailAddress *string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// BD邮箱
	// example:
	//
	// 715680093@qq.com
	BdEmail *string `json:"bd_email,omitempty" xml:"bd_email,omitempty"`
}

func (s InvoiceMailInfo) String() string {
	return tea.Prettify(s)
}

func (s InvoiceMailInfo) GoString() string {
	return s.String()
}

func (s *InvoiceMailInfo) SetPid(v string) *InvoiceMailInfo {
	s.Pid = &v
	return s
}

func (s *InvoiceMailInfo) SetName(v string) *InvoiceMailInfo {
	s.Name = &v
	return s
}

func (s *InvoiceMailInfo) SetCountry(v string) *InvoiceMailInfo {
	s.Country = &v
	return s
}

func (s *InvoiceMailInfo) SetEmail(v string) *InvoiceMailInfo {
	s.Email = &v
	return s
}

func (s *InvoiceMailInfo) SetTelephone(v string) *InvoiceMailInfo {
	s.Telephone = &v
	return s
}

func (s *InvoiceMailInfo) SetProvince(v string) *InvoiceMailInfo {
	s.Province = &v
	return s
}

func (s *InvoiceMailInfo) SetCity(v string) *InvoiceMailInfo {
	s.City = &v
	return s
}

func (s *InvoiceMailInfo) SetCountyDistrict(v string) *InvoiceMailInfo {
	s.CountyDistrict = &v
	return s
}

func (s *InvoiceMailInfo) SetStreet(v string) *InvoiceMailInfo {
	s.Street = &v
	return s
}

func (s *InvoiceMailInfo) SetDetailAddress(v string) *InvoiceMailInfo {
	s.DetailAddress = &v
	return s
}

func (s *InvoiceMailInfo) SetBdEmail(v string) *InvoiceMailInfo {
	s.BdEmail = &v
	return s
}

// 出站swift报文元素
type OutboundSwiftMessage struct {
	// DAP给报文消息分配的唯一ID，使用无横线 UUID，固定 32 位
	// example:
	//
	// 550e8400e29b41d4a716446655440000
	MessageId *string `json:"message_id,omitempty" xml:"message_id,omitempty" require:"true"`
	// 报文格式：MT 或 MX
	// example:
	//
	// MT
	Format *string `json:"format,omitempty" xml:"format,omitempty" require:"true"`
	// 报文类型，例如 MT542、pacs.009
	// example:
	//
	// MT542
	MessageType *string `json:"message_type,omitempty" xml:"message_type,omitempty" require:"true"`
	// 报文原文 SHA-256 摘要
	// example:
	//
	// 52a44d7c6d3a2b6c8f1a9d0e7c6b5a4f3e2d1c0b9a887766554433221100abcd
	MessageHash *string `json:"message_hash,omitempty" xml:"message_hash,omitempty" require:"true"`
	// 报文预期接收方BIC
	// example:
	//
	// DAP000xxxx
	ReceiverBic *string `json:"receiver_bic,omitempty" xml:"receiver_bic,omitempty"`
	// 报文预期接收方DN
	// example:
	//
	// DAP000xxxx
	ReceiverDn *string `json:"receiver_dn,omitempty" xml:"receiver_dn,omitempty"`
	// 本次领取时间，ISO-8601 UTC 字符串
	// example:
	//
	// 2026-08-03T07:25:30.123Z
	ClaimedAt *string `json:"claimed_at,omitempty" xml:"claimed_at,omitempty" require:"true"`
	// 建议在该时间前回传 ACK/NACK
	// example:
	//
	// 2026-08-03T07:30:30.123Z
	AckDeadlineAt *string `json:"ack_deadline_at,omitempty" xml:"ack_deadline_at,omitempty" require:"true"`
	// swift报文消息原文
	// example:
	//
	// {1:F01CMUOHKHHAXXX0000000000}{2:I542BANKHKHHXXXXN}{4:
	//
	// :16R:GENL
	//
	// :20C::SEME//DAP-SEME-20260629-000089
	//
	// :23G:NEWM
	//
	// :16S:GENL
	//
	// -}
	RawMessage *string `json:"raw_message,omitempty" xml:"raw_message,omitempty" require:"true"`
}

func (s OutboundSwiftMessage) String() string {
	return tea.Prettify(s)
}

func (s OutboundSwiftMessage) GoString() string {
	return s.String()
}

func (s *OutboundSwiftMessage) SetMessageId(v string) *OutboundSwiftMessage {
	s.MessageId = &v
	return s
}

func (s *OutboundSwiftMessage) SetFormat(v string) *OutboundSwiftMessage {
	s.Format = &v
	return s
}

func (s *OutboundSwiftMessage) SetMessageType(v string) *OutboundSwiftMessage {
	s.MessageType = &v
	return s
}

func (s *OutboundSwiftMessage) SetMessageHash(v string) *OutboundSwiftMessage {
	s.MessageHash = &v
	return s
}

func (s *OutboundSwiftMessage) SetReceiverBic(v string) *OutboundSwiftMessage {
	s.ReceiverBic = &v
	return s
}

func (s *OutboundSwiftMessage) SetReceiverDn(v string) *OutboundSwiftMessage {
	s.ReceiverDn = &v
	return s
}

func (s *OutboundSwiftMessage) SetClaimedAt(v string) *OutboundSwiftMessage {
	s.ClaimedAt = &v
	return s
}

func (s *OutboundSwiftMessage) SetAckDeadlineAt(v string) *OutboundSwiftMessage {
	s.AckDeadlineAt = &v
	return s
}

func (s *OutboundSwiftMessage) SetRawMessage(v string) *OutboundSwiftMessage {
	s.RawMessage = &v
	return s
}

type PagequeryAntcloudInvoiceRcptDetailRequest struct {
	// OAuth模式下的授权token
	AuthToken *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	// 合同号
	ArNo *string `json:"ar_no,omitempty" xml:"ar_no,omitempty"`
	// 租户ID
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 当前页码，不传默认为1
	CurrentPage *int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 分页条数，不传默认20条
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 系统来源
	Source *string `json:"source,omitempty" xml:"source,omitempty" require:"true"`
}

func (s PagequeryAntcloudInvoiceRcptDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s PagequeryAntcloudInvoiceRcptDetailRequest) GoString() string {
	return s.String()
}

func (s *PagequeryAntcloudInvoiceRcptDetailRequest) SetAuthToken(v string) *PagequeryAntcloudInvoiceRcptDetailRequest {
	s.AuthToken = &v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailRequest) SetArNo(v string) *PagequeryAntcloudInvoiceRcptDetailRequest {
	s.ArNo = &v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailRequest) SetTenantId(v string) *PagequeryAntcloudInvoiceRcptDetailRequest {
	s.TenantId = &v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailRequest) SetCurrentPage(v int64) *PagequeryAntcloudInvoiceRcptDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailRequest) SetPageSize(v int64) *PagequeryAntcloudInvoiceRcptDetailRequest {
	s.PageSize = &v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailRequest) SetSource(v string) *PagequeryAntcloudInvoiceRcptDetailRequest {
	s.Source = &v
	return s
}

type PagequeryAntcloudInvoiceRcptDetailResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 单据明细列表
	DetailList []*RcptDetailVO `json:"detail_list,omitempty" xml:"detail_list,omitempty" type:"Repeated"`
	// 当前页
	CurrentPage *int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 分页大小
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 查询结果总数
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s PagequeryAntcloudInvoiceRcptDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s PagequeryAntcloudInvoiceRcptDetailResponse) GoString() string {
	return s.String()
}

func (s *PagequeryAntcloudInvoiceRcptDetailResponse) SetReqMsgId(v string) *PagequeryAntcloudInvoiceRcptDetailResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailResponse) SetResultCode(v string) *PagequeryAntcloudInvoiceRcptDetailResponse {
	s.ResultCode = &v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailResponse) SetResultMsg(v string) *PagequeryAntcloudInvoiceRcptDetailResponse {
	s.ResultMsg = &v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailResponse) SetDetailList(v []*RcptDetailVO) *PagequeryAntcloudInvoiceRcptDetailResponse {
	s.DetailList = v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailResponse) SetCurrentPage(v int64) *PagequeryAntcloudInvoiceRcptDetailResponse {
	s.CurrentPage = &v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailResponse) SetPageSize(v int64) *PagequeryAntcloudInvoiceRcptDetailResponse {
	s.PageSize = &v
	return s
}

func (s *PagequeryAntcloudInvoiceRcptDetailResponse) SetTotalCount(v int64) *PagequeryAntcloudInvoiceRcptDetailResponse {
	s.TotalCount = &v
	return s
}

type PushSwiftInboundRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 发送方生成的唯一请求标识，32位长度字符串
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
	// 报文格式：MT 或 MX
	Format *string `json:"format,omitempty" xml:"format,omitempty" require:"true"`
	// 原始 MT/MX 报文，必须保留原始换行、空白和字段顺序
	RawMessage *string `json:"raw_message,omitempty" xml:"raw_message,omitempty" require:"true"`
	// 发送方BIC
	SenderBic *string `json:"sender_bic,omitempty" xml:"sender_bic,omitempty"`
	// 发送方DN
	SenderDn *string `json:"sender_dn,omitempty" xml:"sender_dn,omitempty"`
	// 接收方BIC
	ReceiverBic *string `json:"receiver_bic,omitempty" xml:"receiver_bic,omitempty"`
	// 接收方DN
	ReceiverDn *string `json:"receiver_dn,omitempty" xml:"receiver_dn,omitempty"`
}

func (s PushSwiftInboundRequest) String() string {
	return tea.Prettify(s)
}

func (s PushSwiftInboundRequest) GoString() string {
	return s.String()
}

func (s *PushSwiftInboundRequest) SetAuthToken(v string) *PushSwiftInboundRequest {
	s.AuthToken = &v
	return s
}

func (s *PushSwiftInboundRequest) SetProductInstanceId(v string) *PushSwiftInboundRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *PushSwiftInboundRequest) SetRequestId(v string) *PushSwiftInboundRequest {
	s.RequestId = &v
	return s
}

func (s *PushSwiftInboundRequest) SetFormat(v string) *PushSwiftInboundRequest {
	s.Format = &v
	return s
}

func (s *PushSwiftInboundRequest) SetRawMessage(v string) *PushSwiftInboundRequest {
	s.RawMessage = &v
	return s
}

func (s *PushSwiftInboundRequest) SetSenderBic(v string) *PushSwiftInboundRequest {
	s.SenderBic = &v
	return s
}

func (s *PushSwiftInboundRequest) SetSenderDn(v string) *PushSwiftInboundRequest {
	s.SenderDn = &v
	return s
}

func (s *PushSwiftInboundRequest) SetReceiverBic(v string) *PushSwiftInboundRequest {
	s.ReceiverBic = &v
	return s
}

func (s *PushSwiftInboundRequest) SetReceiverDn(v string) *PushSwiftInboundRequest {
	s.ReceiverDn = &v
	return s
}

type PushSwiftInboundResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// DAP 生成的唯一消息ID，32位长度字符串
	MessageId *string `json:"message_id,omitempty" xml:"message_id,omitempty"`
	// 是否已接收并持久化
	Accepted *bool `json:"accepted,omitempty" xml:"accepted,omitempty"`
	// 是否为重复提交
	Duplicate *bool `json:"duplicate,omitempty" xml:"duplicate,omitempty"`
	// 消息接收时间，ISO-8601 UTC 格式
	ReceivedAt *string `json:"received_at,omitempty" xml:"received_at,omitempty"`
	// 报文原文 SHA-256 摘要
	MessageHash *string `json:"message_hash,omitempty" xml:"message_hash,omitempty"`
}

func (s PushSwiftInboundResponse) String() string {
	return tea.Prettify(s)
}

func (s PushSwiftInboundResponse) GoString() string {
	return s.String()
}

func (s *PushSwiftInboundResponse) SetReqMsgId(v string) *PushSwiftInboundResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PushSwiftInboundResponse) SetResultCode(v string) *PushSwiftInboundResponse {
	s.ResultCode = &v
	return s
}

func (s *PushSwiftInboundResponse) SetResultMsg(v string) *PushSwiftInboundResponse {
	s.ResultMsg = &v
	return s
}

func (s *PushSwiftInboundResponse) SetMessageId(v string) *PushSwiftInboundResponse {
	s.MessageId = &v
	return s
}

func (s *PushSwiftInboundResponse) SetAccepted(v bool) *PushSwiftInboundResponse {
	s.Accepted = &v
	return s
}

func (s *PushSwiftInboundResponse) SetDuplicate(v bool) *PushSwiftInboundResponse {
	s.Duplicate = &v
	return s
}

func (s *PushSwiftInboundResponse) SetReceivedAt(v string) *PushSwiftInboundResponse {
	s.ReceivedAt = &v
	return s
}

func (s *PushSwiftInboundResponse) SetMessageHash(v string) *PushSwiftInboundResponse {
	s.MessageHash = &v
	return s
}

type ClaimSwiftOutboundRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 单次领取数量，默认 50，最大 50
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 支持筛选报文格式，可以传入MT或者MX，为空或不传时不过滤
	SupportedFormats []*string `json:"supported_formats,omitempty" xml:"supported_formats,omitempty" type:"Repeated"`
}

func (s ClaimSwiftOutboundRequest) String() string {
	return tea.Prettify(s)
}

func (s ClaimSwiftOutboundRequest) GoString() string {
	return s.String()
}

func (s *ClaimSwiftOutboundRequest) SetAuthToken(v string) *ClaimSwiftOutboundRequest {
	s.AuthToken = &v
	return s
}

func (s *ClaimSwiftOutboundRequest) SetProductInstanceId(v string) *ClaimSwiftOutboundRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ClaimSwiftOutboundRequest) SetLimit(v int64) *ClaimSwiftOutboundRequest {
	s.Limit = &v
	return s
}

func (s *ClaimSwiftOutboundRequest) SetSupportedFormats(v []*string) *ClaimSwiftOutboundRequest {
	s.SupportedFormats = v
	return s
}

type ClaimSwiftOutboundResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 本次查询到的报文数量
	MessageCount *int64 `json:"message_count,omitempty" xml:"message_count,omitempty"`
	// 已领取报文列表
	Messages []*OutboundSwiftMessage `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// 当前调用方和过滤条件下仍可领取的剩余报文数量
	RemainingPreparedCount *int64 `json:"remaining_prepared_count,omitempty" xml:"remaining_prepared_count,omitempty"`
}

func (s ClaimSwiftOutboundResponse) String() string {
	return tea.Prettify(s)
}

func (s ClaimSwiftOutboundResponse) GoString() string {
	return s.String()
}

func (s *ClaimSwiftOutboundResponse) SetReqMsgId(v string) *ClaimSwiftOutboundResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ClaimSwiftOutboundResponse) SetResultCode(v string) *ClaimSwiftOutboundResponse {
	s.ResultCode = &v
	return s
}

func (s *ClaimSwiftOutboundResponse) SetResultMsg(v string) *ClaimSwiftOutboundResponse {
	s.ResultMsg = &v
	return s
}

func (s *ClaimSwiftOutboundResponse) SetMessageCount(v int64) *ClaimSwiftOutboundResponse {
	s.MessageCount = &v
	return s
}

func (s *ClaimSwiftOutboundResponse) SetMessages(v []*OutboundSwiftMessage) *ClaimSwiftOutboundResponse {
	s.Messages = v
	return s
}

func (s *ClaimSwiftOutboundResponse) SetRemainingPreparedCount(v int64) *ClaimSwiftOutboundResponse {
	s.RemainingPreparedCount = &v
	return s
}

type AckSwiftOutboundRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 在claim的时候获得的DAP给报文消息分配的唯一id，32位
	MessageId *string `json:"message_id,omitempty" xml:"message_id,omitempty" require:"true"`
	// 本次 ACK/NACK 请求 ID，32位，调用方设置，用于审计和排查
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
	// 获得确定发送结果的时间，ISO-8601 UTC 字符串
	AckedAt *string `json:"acked_at,omitempty" xml:"acked_at,omitempty" require:"true"`
	// 原始 ACK/NACK的 MT/MX 报文，必须保留原始换行、空白和字段顺序
	RawMessage *string `json:"raw_message,omitempty" xml:"raw_message,omitempty" require:"true"`
	// 填写 SAA_XML_V2 或者 MQ_MT
	ReceiptFormat *string `json:"receipt_format,omitempty" xml:"receipt_format,omitempty" require:"true"`
}

func (s AckSwiftOutboundRequest) String() string {
	return tea.Prettify(s)
}

func (s AckSwiftOutboundRequest) GoString() string {
	return s.String()
}

func (s *AckSwiftOutboundRequest) SetAuthToken(v string) *AckSwiftOutboundRequest {
	s.AuthToken = &v
	return s
}

func (s *AckSwiftOutboundRequest) SetProductInstanceId(v string) *AckSwiftOutboundRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *AckSwiftOutboundRequest) SetMessageId(v string) *AckSwiftOutboundRequest {
	s.MessageId = &v
	return s
}

func (s *AckSwiftOutboundRequest) SetRequestId(v string) *AckSwiftOutboundRequest {
	s.RequestId = &v
	return s
}

func (s *AckSwiftOutboundRequest) SetAckedAt(v string) *AckSwiftOutboundRequest {
	s.AckedAt = &v
	return s
}

func (s *AckSwiftOutboundRequest) SetRawMessage(v string) *AckSwiftOutboundRequest {
	s.RawMessage = &v
	return s
}

func (s *AckSwiftOutboundRequest) SetReceiptFormat(v string) *AckSwiftOutboundRequest {
	s.ReceiptFormat = &v
	return s
}

type AckSwiftOutboundResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// DAP平台分配的报文消息唯一id，32位
	MessageId *string `json:"message_id,omitempty" xml:"message_id,omitempty"`
	// 是否重复发送消息
	Duplicate *bool `json:"duplicate,omitempty" xml:"duplicate,omitempty"`
	// DAP平台接收并处理 ACK/NACK 的时间，ISO-8601 UTC 字符串
	AckedAt *string `json:"acked_at,omitempty" xml:"acked_at,omitempty"`
	// RECEIVED：回执已保存，包括解析或关联校验失败；CONFLICT：回执已保存，但投递已终态，不覆盖原结果。两种情况均无需重发。
	ReceiptStatus *string `json:"receipt_status,omitempty" xml:"receipt_status,omitempty"`
}

func (s AckSwiftOutboundResponse) String() string {
	return tea.Prettify(s)
}

func (s AckSwiftOutboundResponse) GoString() string {
	return s.String()
}

func (s *AckSwiftOutboundResponse) SetReqMsgId(v string) *AckSwiftOutboundResponse {
	s.ReqMsgId = &v
	return s
}

func (s *AckSwiftOutboundResponse) SetResultCode(v string) *AckSwiftOutboundResponse {
	s.ResultCode = &v
	return s
}

func (s *AckSwiftOutboundResponse) SetResultMsg(v string) *AckSwiftOutboundResponse {
	s.ResultMsg = &v
	return s
}

func (s *AckSwiftOutboundResponse) SetMessageId(v string) *AckSwiftOutboundResponse {
	s.MessageId = &v
	return s
}

func (s *AckSwiftOutboundResponse) SetDuplicate(v bool) *AckSwiftOutboundResponse {
	s.Duplicate = &v
	return s
}

func (s *AckSwiftOutboundResponse) SetAckedAt(v string) *AckSwiftOutboundResponse {
	s.AckedAt = &v
	return s
}

func (s *AckSwiftOutboundResponse) SetReceiptStatus(v string) *AckSwiftOutboundResponse {
	s.ReceiptStatus = &v
	return s
}

type QueryclaimedSwiftOutboundRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 单次查询数量，默认 50，最大 50
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 支持的报文格式；为空或不传时不过滤
	SupportedFormats []*string `json:"supported_formats,omitempty" xml:"supported_formats,omitempty" type:"Repeated"`
	// 只查询 ACK/NACK 建议时间早于该时间点的报文，ISO-8601 UTC 字符串
	AckDeadlineBefore *string `json:"ack_deadline_before,omitempty" xml:"ack_deadline_before,omitempty"`
}

func (s QueryclaimedSwiftOutboundRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryclaimedSwiftOutboundRequest) GoString() string {
	return s.String()
}

func (s *QueryclaimedSwiftOutboundRequest) SetAuthToken(v string) *QueryclaimedSwiftOutboundRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryclaimedSwiftOutboundRequest) SetProductInstanceId(v string) *QueryclaimedSwiftOutboundRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryclaimedSwiftOutboundRequest) SetLimit(v int64) *QueryclaimedSwiftOutboundRequest {
	s.Limit = &v
	return s
}

func (s *QueryclaimedSwiftOutboundRequest) SetSupportedFormats(v []*string) *QueryclaimedSwiftOutboundRequest {
	s.SupportedFormats = v
	return s
}

func (s *QueryclaimedSwiftOutboundRequest) SetAckDeadlineBefore(v string) *QueryclaimedSwiftOutboundRequest {
	s.AckDeadlineBefore = &v
	return s
}

type QueryclaimedSwiftOutboundResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 本次查询到的报文数量
	MessageCount *int64 `json:"message_count,omitempty" xml:"message_count,omitempty"`
	// 出站swift报文元素
	Messages []*OutboundSwiftMessage `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s QueryclaimedSwiftOutboundResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryclaimedSwiftOutboundResponse) GoString() string {
	return s.String()
}

func (s *QueryclaimedSwiftOutboundResponse) SetReqMsgId(v string) *QueryclaimedSwiftOutboundResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryclaimedSwiftOutboundResponse) SetResultCode(v string) *QueryclaimedSwiftOutboundResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryclaimedSwiftOutboundResponse) SetResultMsg(v string) *QueryclaimedSwiftOutboundResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryclaimedSwiftOutboundResponse) SetMessageCount(v int64) *QueryclaimedSwiftOutboundResponse {
	s.MessageCount = &v
	return s
}

func (s *QueryclaimedSwiftOutboundResponse) SetMessages(v []*OutboundSwiftMessage) *QueryclaimedSwiftOutboundResponse {
	s.Messages = v
	return s
}

type Client struct {
	Endpoint                *string
	RegionId                *string
	AccessKeyId             *string
	AccessKeySecret         *string
	Protocol                *string
	UserAgent               *string
	ReadTimeout             *int
	ConnectTimeout          *int
	HttpProxy               *string
	HttpsProxy              *string
	Socks5Proxy             *string
	Socks5NetWork           *string
	NoProxy                 *string
	MaxIdleConns            *int
	SecurityToken           *string
	MaxIdleTimeMillis       *int
	KeepAliveDurationMillis *int
	MaxRequests             *int
	MaxRequestsPerHost      *int
}

// Description:
//
// # Init client with Config
//
// @param config - config contains the necessary information to create a client
func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if tea.BoolValue(util.IsUnset(config)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "ParameterMissing",
			"message": "'config' can not be unset",
		})
		return _err
	}

	client.AccessKeyId = config.AccessKeyId
	client.AccessKeySecret = config.AccessKeySecret
	client.SecurityToken = config.SecurityToken
	client.Endpoint = config.Endpoint
	client.Protocol = config.Protocol
	client.UserAgent = config.UserAgent
	client.ReadTimeout = util.DefaultNumber(config.ReadTimeout, tea.Int(20000))
	client.ConnectTimeout = util.DefaultNumber(config.ConnectTimeout, tea.Int(20000))
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = util.DefaultNumber(config.MaxIdleConns, tea.Int(60000))
	client.MaxIdleTimeMillis = util.DefaultNumber(config.MaxIdleTimeMillis, tea.Int(5))
	client.KeepAliveDurationMillis = util.DefaultNumber(config.KeepAliveDurationMillis, tea.Int(5000))
	client.MaxRequests = util.DefaultNumber(config.MaxRequests, tea.Int(100))
	client.MaxRequestsPerHost = util.DefaultNumber(config.MaxRequestsPerHost, tea.Int(100))
	return nil
}

// Description:
//
// # Encapsulate the request and invoke the network
//
// @param action - api name
//
// @param protocol - http or https
//
// @param method - e.g. GET
//
// @param pathname - pathname of every api
//
// @param request - which contains request params
//
// @param runtime - which controls some details of call api, such as retry times
//
// @return the response
func (client *Client) DoRequest(version *string, action *string, protocol *string, method *string, pathname *string, request map[string]interface{}, headers map[string]*string, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":          "retry",
		"readTimeout":        tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout":     tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":          tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":         tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":            tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"maxIdleConns":       tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"maxIdleTimeMillis":  tea.IntValue(client.MaxIdleTimeMillis),
		"keepAliveDuration":  tea.IntValue(client.KeepAliveDurationMillis),
		"maxRequests":        tea.IntValue(client.MaxRequests),
		"maxRequestsPerHost": tea.IntValue(client.MaxRequestsPerHost),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = pathname
			request_.Query = map[string]*string{
				"method":           action,
				"version":          version,
				"sign_type":        tea.String("HmacSHA1"),
				"req_time":         antchainutil.GetTimestamp(),
				"req_msg_id":       antchainutil.GetNonce(),
				"access_key":       client.AccessKeyId,
				"base_sdk_version": tea.String("TeaSDK-2.0"),
				"sdk_version":      tea.String("1.0.4"),
				"_prod_code":       tea.String("TSDAP"),
				"_prod_channel":    tea.String("default"),
			}
			if !tea.BoolValue(util.Empty(client.SecurityToken)) {
				request_.Query["security_token"] = client.SecurityToken
			}

			request_.Headers = tea.Merge(map[string]*string{
				"host":       util.DefaultString(client.Endpoint, tea.String("openapi.antchain.antgroup.com")),
				"user-agent": util.GetUserAgent(client.UserAgent),
			}, headers)
			tmp := util.AnyifyMapValue(rpcutil.Query(request))
			request_.Body = tea.ToReader(util.ToFormString(tmp))
			request_.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			signedParam := tea.Merge(request_.Query,
				rpcutil.Query(request))
			request_.Query["sign"] = antchainutil.GetSignature(signedParam, client.AccessKeySecret)
			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			raw, _err := util.ReadAsString(response_.Body)
			if _err != nil {
				return _result, _err
			}

			obj := util.ParseJSON(raw)
			res, _err := util.AssertAsMap(obj)
			if _err != nil {
				return _result, _err
			}

			resp, _err := util.AssertAsMap(res["response"])
			if _err != nil {
				return _result, _err
			}

			if tea.BoolValue(antchainutil.HasError(raw, client.AccessKeySecret)) {
				_err = tea.NewSDKError(map[string]interface{}{
					"message": resp["result_msg"],
					"data":    resp,
					"code":    resp["result_code"],
				})
				return _result, _err
			}

			_result = resp
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

// Description:
//
// Description: 支持租户或合同号分页查询开票单据详细信息
//
// Summary: 支持租户或合同号分页查询开票单据详细信息
func (client *Client) PagequeryAntcloudInvoiceRcptDetail(request *PagequeryAntcloudInvoiceRcptDetailRequest) (_result *PagequeryAntcloudInvoiceRcptDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PagequeryAntcloudInvoiceRcptDetailResponse{}
	_body, _err := client.PagequeryAntcloudInvoiceRcptDetailEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Description: 支持租户或合同号分页查询开票单据详细信息
//
// Summary: 支持租户或合同号分页查询开票单据详细信息
func (client *Client) PagequeryAntcloudInvoiceRcptDetailEx(request *PagequeryAntcloudInvoiceRcptDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PagequeryAntcloudInvoiceRcptDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PagequeryAntcloudInvoiceRcptDetailResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antcloud.invoice.rcpt.detail.pagequery"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// Description: 用于提交swift入站报文
//
// Summary: 用于提交swift入站报文
func (client *Client) PushSwiftInbound(request *PushSwiftInboundRequest) (_result *PushSwiftInboundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushSwiftInboundResponse{}
	_body, _err := client.PushSwiftInboundEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Description: 用于提交swift入站报文
//
// Summary: 用于提交swift入站报文
func (client *Client) PushSwiftInboundEx(request *PushSwiftInboundRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushSwiftInboundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PushSwiftInboundResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antdigital.tsdap.swift.inbound.push"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// Description: SWIFT 出站报文领取
//
// Summary: SWIFT 出站报文领取
func (client *Client) ClaimSwiftOutbound(request *ClaimSwiftOutboundRequest) (_result *ClaimSwiftOutboundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ClaimSwiftOutboundResponse{}
	_body, _err := client.ClaimSwiftOutboundEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Description: SWIFT 出站报文领取
//
// Summary: SWIFT 出站报文领取
func (client *Client) ClaimSwiftOutboundEx(request *ClaimSwiftOutboundRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ClaimSwiftOutboundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ClaimSwiftOutboundResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antdigital.tsdap.swift.outbound.claim"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// Description: SWIFT 出站报文 ACK/NACK 回传
//
// Summary: SWIFT 出站报文 ACK/NACK 回传
func (client *Client) AckSwiftOutbound(request *AckSwiftOutboundRequest) (_result *AckSwiftOutboundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AckSwiftOutboundResponse{}
	_body, _err := client.AckSwiftOutboundEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Description: SWIFT 出站报文 ACK/NACK 回传
//
// Summary: SWIFT 出站报文 ACK/NACK 回传
func (client *Client) AckSwiftOutboundEx(request *AckSwiftOutboundRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AckSwiftOutboundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AckSwiftOutboundResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antdigital.tsdap.swift.outbound.ack"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// Description: SWIFT 已领取超时报文查询
//
// Summary: SWIFT 已领取超时报文查询
func (client *Client) QueryclaimedSwiftOutbound(request *QueryclaimedSwiftOutboundRequest) (_result *QueryclaimedSwiftOutboundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryclaimedSwiftOutboundResponse{}
	_body, _err := client.QueryclaimedSwiftOutboundEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Description: SWIFT 已领取超时报文查询
//
// Summary: SWIFT 已领取超时报文查询
func (client *Client) QueryclaimedSwiftOutboundEx(request *QueryclaimedSwiftOutboundRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryclaimedSwiftOutboundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryclaimedSwiftOutboundResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antdigital.tsdap.swift.outbound.queryclaimed"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
