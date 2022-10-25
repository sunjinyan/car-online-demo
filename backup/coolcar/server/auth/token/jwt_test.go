package token

import (
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

const  privateKey = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDVCnwQD4uyDV8V
m2IkYqBDDegdM+kApMeWFqNm0Pe8edIz6guI4i1MLO9nbiOutUU3tMkufmte2pDG
eue6AMoz6t9W4GEi7e8eK22OgwdP8uqAK9Dd8wx9LrYkYyTTKmNrwsn+Ly4x5uMv
1FVvWi7xPa3g9Tp8PZNCDhXLM7PirO/rlgtWHWYpxDi4w/YJHE8TAj3f6Qi6ba4V
dPMihRnuvvE/dCPi0DU+gtIJUIIsLv3nCHUqVzu1I36Eg6X0nI2oA/JdEWMxVNdo
RMi9HiWRCRx7ykjjRSIr48nig3QfUVhge4b+Bvub5QN/s8gXmUOs+BQ1lRFJcb1A
VJcJuf1DAgMBAAECggEANt5Xa4+FbBszhQ4a7cIVApAfWv4WoZrO6KH1/bIOTBWY
BNTfmaafAgKEuj/Mot9+70T5l6bpomlqZiQbcHVpnot8cBSkD+n17vBCbDTk+jtX
GCGwzqkda1RPhvmd72D4cDWf7n5aSHHty1DFnEM+2dUH5mtiX8NAHQAhrs8nIlbt
o68ShPSkviylou+TZ/JoF9FwToZyrKAtuDpkV5dO/gn4XZHnYxocYPk6+nCmu9yK
caM5BYJN3/opO7GHziP9N31YSmIB8A17l1B1LU3GULGhCYitU8AGahu5RoR7zKUt
M7tPGAwERxmWgpBRPtFysCI9jEVwl1/TMRgHvogJWQKBgQDzD8xGJh83HlMFPrMu
DTQz5akZvzniPk5I5ni9q4IKLPwnu+d0HxWrZEVV0PM9B25UVlxfuYeGB6JygyPI
JSb468+CjbDxg7Cfou+VB+qBCCujXs/L65+0lgmVFGvqmGW+QFMgSN2rC9sqxuDf
FHPbpjEKgF1kpxDKRNHqxxLTzQKBgQDgYZgFKRbs49X41U4/N58Ab0kIOKqkjMsM
LcFG+5cM0w00U0UGMy6G1UGYngxKwz3A1+pk34YdWoLNlMTPFqPdVM3tiBafXBVW
ShoetJiF0JbzoMByb2XQ7Pu/1dDDpOfIR42oQMj4bT93j7G2iQNCHapNPN6KkpGQ
s/ckGYYlTwKBgQDHHsURHr/pzqzimhf9/REsBtP+ZLsNzk8hx7VYZL22il3US0p0
xdZhv1ZR1EjAUWsgKlOok/wjX/PD7g0NNOp1d/FamatAX8r9LGUraP3C0Se2wcBn
5VLDhTq6Z6101X5cWqP+ydLXAnJ828py1nWHluGMIdbg00CdA74YoReabQKBgQDP
lnrmi1VCZKcZR37Uzci0xkBy8ros8nUIF2yxFyJwIvIcbs11JvTsmW/OmqoUoP0I
5LleCc4yagKJ4BEU5oVDx6mdEEKdvLug0X3I2pwhcRG25Aon1RgFarqeiFpUJsSp
bkmIWCeWbrxBCRgJIEl7aaadPqtE2MUDcvw60cJh0QKBgGiTc+LDgKtEPrkE8nrh
AmVdQNCeL/kn3hzRK1pgVdkPpoeUMoHknVCWSAGz9gWloVfQ7ewmu1xBAdRy/+1Q
OJTQJpSWBsldtVZno2Dv9XWcAYAyq+zmTU28I7iflPMqF8iV+zZSRqijQDnVBYZ4
ED4Kc807WQM5uXwmYxfVyXqB
-----END PRIVATE KEY-----`

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1Qp8EA+Lsg1fFZtiJGKg
Qw3oHTPpAKTHlhajZtD3vHnSM+oLiOItTCzvZ24jrrVFN7TJLn5rXtqQxnrnugDK
M+rfVuBhIu3vHittjoMHT/LqgCvQ3fMMfS62JGMk0ypja8LJ/i8uMebjL9RVb1ou
8T2t4PU6fD2TQg4VyzOz4qzv65YLVh1mKcQ4uMP2CRxPEwI93+kIum2uFXTzIoUZ
7r7xP3Qj4tA1PoLSCVCCLC795wh1Klc7tSN+hIOl9JyNqAPyXRFjMVTXaETIvR4l
kQkce8pI40UiK+PJ4oN0H1FYYHuG/gb7m+UDf7PIF5lDrPgUNZURSXG9QFSXCbn9
QwIDAQAB
-----END PUBLIC KEY-----`

func TestJWTTokenGenerate_GenerateToken(t *testing.T) {
	priKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("parse rsa private key from PEM error %v",err)
	}
	j := NewJWTTokenGenerate("coolcar/auth",priKey)
	j.nowFunc = func() time.Time {
		return time.Unix(1516239022,0)
	}
	token, err := j.GenerateToken("630734ab70541b5f24913914", 2 * time.Hour)
	if err != nil {
		t.Errorf("create token  error %v",err)
	}
	t.Log("token for string 630734ab70541b5f24913914 : ",token)
}
