package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/pem"
	"errors"
	"log"
	"strings"

	"github.com/macduyhai/loadBalancingGrpcApi/config"
	"github.com/macduyhai/loadBalancingGrpcApi/daos"
)

type EncryptService interface {
	CheckKey(id int64, username string) error
	DecodeToken(token_str string) (int64, error)
	Base64Enc(b1 []byte) string
	Base64Dec(s1 string) ([]byte, error)
	RsaDecrypt(ciphertext []byte, key []byte) ([]byte, error)
	RsaEncrypt(origData []byte, key []byte) ([]byte, error)
}

type ecryptServiceImpl struct {
	config  *config.Config
	userDao daos.UserDao
}

//=================== MA HOA ==========================
func (service *ecryptServiceImpl) CheckKey(id int64, token_str string) error {
	// user, err := service.userDao.CheckUserID(id)
	// if err != nil {
	// 	return err
	// }

	// log.Printf("UserID: %d/t", id)
	// log.Println("Token: " + token_str)
	tokenDe, err := Base64Dec(token_str)
	if err != nil {
		log.Println(err)
		err = errors.New("Key invalid")
		return err

	}
	tokenID, err := RsaDecrypt(tokenDe, []byte(service.config.PrivateKey))
	if err != nil {
		log.Println(err)
		err = errors.New("Key invalid")
		return err
	}
	if id != int64(binary.LittleEndian.Uint64(tokenID)) { // convert []byte to int64
		err := errors.New("Key invalid")
		log.Println(err)
		return err
	} else {
		//log.Println("Key is Valid")
		return nil
	}
}
func (service *userServiceImpl) DecodeToken(token_str string) (int64, error) {
	tokenDe, err := Base64Dec(token_str)
	if err != nil {
		log.Println(err)
		err = errors.New("Key invalid")
		return 0, err

	}
	tokenID, err := RsaDecrypt(tokenDe, []byte(service.config.PrivateKey))
	if err != nil {
		log.Println(err)
		err = errors.New("Key invalid")
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(tokenID)), nil

}
func Base64Enc(b1 []byte) string {
	s1 := base64.StdEncoding.EncodeToString(b1)
	s2 := ""
	var LEN int = 76
	for len(s1) > 76 {
		s2 = s2 + s1[:LEN] + "\n"
		s1 = s1[LEN:]
	}
	s2 = s2 + s1
	return s2
}

func Base64Dec(s1 string) ([]byte, error) {
	s1 = strings.Replace(s1, "\n", "", -1)
	s1 = strings.Replace(s1, "\r", "", -1)
	s1 = strings.Replace(s1, " ", "", -1)
	return base64.StdEncoding.DecodeString(s1)
}

func RsaDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func RsaEncrypt(origData []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}
