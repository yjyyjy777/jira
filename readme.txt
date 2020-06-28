vender
把密钥导出
openssl pkcs12 -in east_MDM.p12 -nocerts -out key.pem
openssl rsa -in key.pem -out vender.key
准备mdm.cer
密码：要输入密码

生成客户证书请求文件 


openssl genrsa -des3 -out customerPrivateKey.pem 2048
openssl req -new -key customerPrivateKey.pem -out customer.csr
如果需要去掉私钥的key，参考第四行命令


签名
要求具有python 2 环境
python mdm_vendor_sign.py  --csr customer.csr --key vender.key --mdm mdm.cer 



参考1：客户推送证书转p12
openssl pkcs12 -export -in MDM_Certificate.pem -out MDM_Certificate.p12 -inkey customerPrivateKey.pem
参考2：测试证书是否有效
openssl s_client -connect gateway.push.apple.com:2195 -cert customer_apns_sys.pem -key customerPrivateKey.pem -debug -showcerts -status
如果提示了错误，或连接直接closed则证书有问题。如果一直一直处于等待输入状态，输入任意、退出则证书是有效的。


参考2：判断证书和私钥是否匹配
openssl pkey -in customerPrivateKey.pem -pubout -outform pem | sha256sum
openssl x509 -in customer_apns_sys.pem -pubkey -noout -outform pem | sha256sum
openssl req -in customer.csr -pubkey -noout -outform pem | sha256sum


