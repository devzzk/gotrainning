package main

func main() {
	//var (
	//	mchId  						 = "146298098"
	//	mchCertificateSerialNumber 	string = "20507A4AEC45E602A793D9078E5898FB7662BA3D"  // 商户证书序列号
	//	mchAPIv3Key                	string = "Gouni2022Y6mm6ssHangzhouHello88s"
	//)
	//
	//mchPrivateKey, err := utils.LoadPrivateKeyWithPath("./payment/apiclient_key.pem")
	//if err != nil {
	//	log.Fatal("load merchant private key error")
	//	return
	//}
	//
	//ctx := context.Background()
	//// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	//opts := []core.ClientOption{
	//	option.WithWechatPayAutoAuthCipher(mchId, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	//}
	//client, err := core.NewClient(ctx, opts...)
	//if err != nil {
	//	log.Fatalf("new wechat pay client err:%s", err)
	//}
	//
	//// 发送请求，以下载微信支付平台证书为例
	//// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	//svc := certificates.CertificatesApiService{Client: client}
	//resp, result, err := svc.DownloadCertificates(ctx)
	//log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)



}
