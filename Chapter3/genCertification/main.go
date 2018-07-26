package genCertification

import (
	"crypto/x509"
	"crypto/rand"
	"math/big"
	"crypto/x509/pkix"
	"time"
	"net"
	"crypto/rsa"
	"log"
	"os"
	"encoding/pem"
)


func main(){

	max :=new(big.Int).Lsh(big.NewInt(1),128)     //1<<128
	serialnumber,_ :=rand.Int(rand.Reader,max)

	subject :=pkix.Name{
		Organization:[]string{"Vode Co."},
		OrganizationalUnit:[]string{"Github.com repo"},
		CommonName:"Go Web programming",
	}

	certificateTemplate :=x509.Certificate{
		SerialNumber:serialnumber,            //证书的唯一的序列号
		Subject:subject,                      //证书拥有者的distinguished name
		NotBefore:time.Now(),                 //生效日期
		NotAfter:time.Now().Add(time.Hour*24*365), //有效日期一年
		KeyUsage:x509.KeyUsageDigitalSignature|x509.KeyUsageKeyEncipherment,  // KeyUsage represents the set of actions that are valid for a given key.
                                                                              // 数字签名&加密

		ExtKeyUsage:[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},           // ExtKeyUsage represents an extended set of actions that are valid for a given key.
		                                                                      // Each of the ExtKeyUsage* constants define a unique action.
		                                                                      //证书用于服务器的身份验证
		IPAddresses:[]net.IP{net.ParseIP("127.0.0.1")},                    //证书只在本地使用
	}
	pk, err := rsa.GenerateKey(rand.Reader, 2048)                        // GenerateKey generates an RSA keypair of the given bit size using the
														                      // random source random (for example, crypto/rand.Reader).
														                      //2048bit 的秘钥对 随机数源是rand.Reader
    if err!=nil{
    	log.Fatalln(err)
	}

	// The certificate is signed by parent. If parent is equal to template then the
	// certificate is self-signed. The parameter pub is the public key of the
	// signee and priv is the private key of the signer.
	// pub和priv是产生的服务器的ssl的公钥 私钥,公钥自签名产生证书
	// All keys types that are implemented via crypto.Signer are supported (This
	// includes *rsa.PublicKey and *ecdsa.PublicKey.)
	derCertBytes,err :=x509.CreateCertificate(rand.Reader,&certificateTemplate,&certificateTemplate,pk.Public(),pk)
	if err!=nil{
		log.Fatalln(err)
	}
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derCertBytes})
	certOut.Close()

	keyOut,_ :=os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()


	//http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)

}

/*
Ref:
1.TonyBai的博客: https://tonybai.com/2015/04/30/go-and-https/
数字证书的组成:
a. C：证书相关信息（对象名称+过期时间+证书发布者+证书签名算法….）
b. S：证书的数字签名
其中的数字签名是通过公式S = F(Digest(C))得到的。
Digest为摘要函数，也就是 md5、sha-1或sha256等单向散列算法，用于将无限输入值转换为一个有限长度的“浓缩”输出值。比如我们常用md5值来验证下载的大文件是否完 整。大文件的内容就是一个无限输入。大文件被放在网站上用于下载时，网站会对大文件做一次md5计算，得出一个128bit的值作为大文件的 摘要一同放在网站上。用户在下载文件后，对下载后的文件再进行一次本地的md5计算，用得出的值与网站上的md5值进行比较，如果一致，则大 文件下载完好，否则下载过程大文件内容有损坏或源文件被篡改。
F为签名函数。CA自己的私钥是唯一标识CA签名的，因此CA用于生成数字证书的签名函数一定要以自己的私钥作为一个输入参数。在RSA加密 系统中，发送端的解密函数就是一个以私钥作 为参数的函数，因此常常被用作签名函数使用。签名算法是与证书一并发送给接收 端的，比如apple的一个服务的证书中关于签名算法的描述是“带 RSA 加密的 SHA-256 ( 1.2.840.113549.1.1.11 )”。
接收端会运用下面算法对数字证书的签名进行校验：F'(S) ?= Digest(C)  这里的F'(S)是签名的逆过程,也就是解密的过程,用的是CA的公钥作为输入参数(CA),这里面用到CA自身的数字证书,(包含CA自己的公钥)。而且为了保证CA证书的真实性，浏览器是在出厂时就内置了 这些CA证书的，而不是后期通过通信的方式获取的。CA证书就是用来校验由该CA颁发的数字证书的。

*/