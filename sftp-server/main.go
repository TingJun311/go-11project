package main

import (
	r "crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	//"math"
	"math/big"
	"math/rand"
	"pkg/config"
	"time"

	//"io"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type sftpDetails struct {
	UserName string
	Password string
	Address string
	Connection *ssh.Client
	Client *sftp.Client
}

func main() {
	objectSftp := sftpDetails{
		UserName: config.SFTP_USERNAME,
		Password: config.SFTP_PASSWORD,
		Address: config.SFTP_ADDRESS,
	}
	objectSftp.Connection = objectSftp.SFTPConnect()
	objectSftp.Client = objectSftp.NewSFTPClient()
	defer func() {
		objectSftp.Connection.Close()
		objectSftp.Client.Close()
	}()
	
	num, err := GenerateId(10)
	if err != nil {
		fmt.Println(err)
	}

	num2, _ := generateRandomBigInt(4)
	if l := len(num2); l < 10 {
		fmt.Println("err")
	}
	fmt.Println(num, num2, RandomNum(10))
//	const filePath = "./SCRC/REDEEM_TRANSACTION_ONLINE/"
//	fileName := "testingss.txt"
//	var isWriteToFile int
//	trackError := 0
//
//	matches, err := client.Glob(filePath + fileName)
//	if err != nil {
//		fmt.Println("		[*] ERROR finding files", err)
//		trackError++
//	}
//	if len(matches) <= 0 {
//		// means not found, then we create it 
//		dstFile, err := client.Create(filePath + fileName)
//		if err != nil {
//			fmt.Println("		[*] ERROR create file", err)
//			trackError++
//		}
//		dstFile.Close()
//	}
//
//	
//	// Check for the file exist or not
//	srcFile, err := client.OpenFile(filePath + fileName, 0666)
//	if err != nil {
//		fmt.Println("Failed to Open a file", err)
//		trackError++
//	}
//	f, err := srcFile.Stat() 
//	if err != nil {
//		fmt.Print("		[*] ERROR stat", err)
//	}
//
//	byteArray := make([]byte, f.Size())
//	if _, err := srcFile.Read(byteArray); err != nil {
//		fmt.Println("	[*] ERROR reading file", err)
//	}
//	
//	if isWriteToFile, err = io.WriteString(srcFile, "\nresxsaxcdc\n"); err != nil {
//		trackError++
//	}
//	srcFile.Close()
//	if isWriteToFile <= 0 && trackError != 0 {
//		// IF has write to file we update uploadFTP in db
//		return 
//	}
	salt := generateSalt(16)
	//fmt.Printf("APPEND SUCCESS %s %s", HashSha512("koo"), salt)
	hashPassword := hashFunc("test", salt)
	fmt.Println(doPasswordsMatch(hashPassword, "test", salt))
}

func HashSha512(input string) (string) {
	sha_512 := sha512.New()
		// sha from a byte array
	sha_512.Write([]byte(input))
	return hex.EncodeToString(sha_512.Sum(nil))
}

func (con sftpDetails) SFTPConnect() (*ssh.Client) {
	port := 22
	addr := fmt.Sprintf("%s:%d", con.Address, port)
	config := ssh.ClientConfig{
		User: con.UserName,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(con.Password),
		},
	}
	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		fmt.Println("Failed to connect to [%s]: %v", addr, err)
	}
	return conn
}

func (conn sftpDetails) NewSFTPClient() (*sftp.Client) {
	client, err := sftp.NewClient(conn.Connection)
	if err != nil {
		fmt.Println("Failed to get NewClient", err)
	}
	return client
}

func RandomNum(maxNum int) (int) {
	if maxNum < 0 {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNum + 1)
}

func RandomNumN(number int64) (*big.Int) {

	num, err := r.Int(r.Reader, big.NewInt(number))
	if err != nil {
		panic(err)
	}
	return num
}

func GenerateId(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("		[*] length must be greater than zero")
	}
	var id []rune
	base10 := []rune("0123456789")
	for i := 0; i < length; i++ {
		if (i % 2) == 0 {
			id = append(id, base10[RandomNum(time.Now().Nanosecond() * (i * time.Now().Second())) % len(base10)])
		} else {
			id = append(id, base10[RandomNumN(10).Bit(i)])
		}
	}
	return string(id), nil
}

func generateRandomBigInt(numBytes int) (string, error) {
	//var num *big.Int
    value := make([]byte, numBytes)
    _, err := rand.Read(value)
    if err != nil {
        return "", err
    }

    for {
        if value[0] != 0 {
            break
        }
        firstByte := value[:1]
        _, err := rand.Read(firstByte)
        if err != nil {
            return "", err
        }
    }

    nums := (&big.Int{}).SetBytes(value).String()
	if l := len(nums); l < 10 {
		fmt.Println("err")
		nums += fmt.Sprint(RandomNum(10))
	}
	return nums, nil
}

func RandomString(length int) (string) {
	var randString []rune
	char := []rune("abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ")

	for i := 0; i < length; i++ {
		randString = append(randString, char[RandomNum(time.Now().Year() + i) % len(char)])
	}

	return string(randString)
} 

const saltSize = 16

// Generate 16 bytes randomly
func generateSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

// Combine password and salt then hash them using the SHA-512
func hashFunc(password string, salt []byte) string {
	// Convert password string to byte slice
	var pwdByte = []byte(password)

	// Create sha-512 hasher
	var sha512 = sha512.New()

	pwdByte = append(pwdByte, salt...)

	sha512.Write(pwdByte)

	// Get the SHA-512 hashed password
	var hashedPassword = sha512.Sum(nil)

	// Convert the hashed to hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPassword)
	return hashedPasswordHex
}

// Check if two passwords match
func doPasswordsMatch(hashedPassword, curPassword string, salt []byte) bool {
	var curPasswordHash = hashFunc(curPassword, salt)

	return hashedPassword == curPasswordHash
}