package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

func sign_bytes(originalData string, keyPath string) string {
	privateKey, _ := ioutil.ReadFile(keyPath)
	block, _ := pem.Decode(privateKey)
	prvKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	hash := sha1.New()
	hash.Write([]byte(originalData))
	sign, _ := rsa.SignPKCS1v15(rand.Reader, prvKey, crypto.SHA1, hash.Sum(nil))
	encodeString := base64.StdEncoding.EncodeToString(sign)
	return encodeString
}

func main() {
	dataMap := make(map[string]string)

	dataMap["merchant_no"] = "1018"
	dataMap["bank_code"] = "900000001"
	dataMap["tied_card_type"] = "1000"
	dataMap["amount"] = "10000"
	dataMap["payer_id_card"] = "310109198210181535"
	dataMap["bank_mobile"] = "18117321018"
	dataMap["sign_order_no"] = "AP2I111018212312121259000"
	dataMap["payer_bank_card_no"] = "6217560800023362683"
	dataMap["payer_name"] = "陈怡海"
	dataMap["version"] = "1.0"
	fmt.Println(dataMap)
	var keys []string
	for k := range dataMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sortedDataMap string
	for index, k := range keys {
		if index == 0 {
			sortedDataMap += "{"
		}
		if index+1 != len(keys) {
			sortedDataMap += "\"" + k + "\""
			sortedDataMap += ":"
			sortedDataMap += "\"" + dataMap[k] + "\""
			sortedDataMap += ","
		} else {
			sortedDataMap += "\"" + k + "\""
			sortedDataMap += ":"
			sortedDataMap += "\"" + dataMap[k] + "\""
			sortedDataMap += "}"
		}
	}
	fmt.Println(sortedDataMap)
	keyPath := "private_key.pem"

	signData := sign_bytes(sortedDataMap, keyPath)
	fmt.Println(signData)
	dataMap["sign"] = signData
	jsonDataMap, _ := json.Marshal(dataMap)

	requestObj, _ := http.NewRequest("POST", "", bytes.NewBuffer(jsonDataMap))

	requestObj.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, _ := client.Do(requestObj)

	fmt.Println("status", resp.Status)
	fmt.Println("response:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
