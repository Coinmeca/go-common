package utils

import (
	"coinmeca-go_common/logger"
	"os"
	"os/user"

	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}

func ToHexString(number int64) string {
	return "0x" + strconv.FormatInt(number, 16)
}

func HexStringToInt64(hexString string) (int64, error) {
	if len(hexString) < 2 || strings.HasPrefix(hexString, "0x") == false {
		return -1, errors.New("wrong hex string format. It should start with '0x'")
	}

	value, err := strconv.ParseInt(hexString[2:], 16, 64)

	if err != nil {
		fmt.Printf("Conversion failed: %s\n", err)
		return -1, err
	}

	return value, nil
}

func HexStringToBigInt(hexString string) (*big.Int, error) {
	if len(hexString) < 2 || strings.HasPrefix(hexString, "0x") == false {
		return big.NewInt(-1), errors.New("wrong hex string format. It should start with '0x'")
	}

	value := new(big.Int)
	value, ok := value.SetString(hexString[2:], 16)

	if ok == false {
		return big.NewInt(-1), errors.New("wrong hex string format")
	}

	return value, nil
}

func Keccak256(textSignature string) string {
	hash := sha3.NewLegacyKeccak256()
	var buf []byte
	//hash.Write([]byte{0xcc})
	// "MemberChanged(address,address,address)"
	hash.Write([]byte(textSignature))
	buf = hash.Sum(nil)

	return hex.EncodeToString(buf)
}

func Keccak256ByShake(baseAddress, quoteAddress string) string {
	out := make([]byte, 16)
	cs := sha3.NewCShake128([]byte(baseAddress), []byte(quoteAddress))
	_, err := cs.Write([]byte(`meca`))
	if err != nil {
		return ""
	}
	_, err = cs.Read(out)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(out)
}

// SigRSV signatures R S V returned as arrays
func SigRSV(iSig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := iSig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigStr := common.Bytes2Hex(sig)
	rS := sigStr[0:64]
	sS := sigStr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigStr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}

func PrintPretty(s interface{}) {
	empJSON, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		logger.Error("PrintPretty", "err", err)
	}
	fmt.Printf("config.Config %s\n", string(empJSON))
}

func ExternalData(url, key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequest("GET", url, nil)
	if key != "" {
		req.Header.Add("X-CMC_PRO_API_KEY", key)
	}
	if err != nil {
		return nil, fmt.Errorf("reqeust error: %v", url)
	}
	req = req.WithContext(ctx)
	req.Close = true // req.Header.Add("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, fmt.Errorf("client error: %v", url)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response error")
	}

	return data, nil
}
