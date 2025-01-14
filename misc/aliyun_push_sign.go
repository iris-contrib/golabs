package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func main() {

	mac := hmac.New(sha1.New, []byte("testsecret&"))
	mac.Write([]byte("GET&%2F&AccessKeyId%3Dtestid%26Action%3DGetDeviceInfos%26AppKey%3D23267207%26Devices%3De2ba19de97604f55b165576736477b74%252C92a1da34bdfd4c9692714917ce22d53d%26Format%3DXML%26RegionId%3Dcn-hangzhou%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3Dc4f5f0de-b3ff-4528-8a89-fa478bda8d80%26SignatureVersion%3D1.0%26Timestamp%3D2016-03-29T03%253A59%253A24Z%26Version%3D2015-08-27"))
	expectedMAC := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	println(expectedMAC)
}
