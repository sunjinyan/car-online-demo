package token

import (
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)


const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1Qp8EA+Lsg1fFZtiJGKg
Qw3oHTPpAKTHlhajZtD3vHnSM+oLiOItTCzvZ24jrrVFN7TJLn5rXtqQxnrnugDK
M+rfVuBhIu3vHittjoMHT/LqgCvQ3fMMfS62JGMk0ypja8LJ/i8uMebjL9RVb1ou
8T2t4PU6fD2TQg4VyzOz4qzv65YLVh1mKcQ4uMP2CRxPEwI93+kIum2uFXTzIoUZ
7r7xP3Qj4tA1PoLSCVCCLC795wh1Klc7tSN+hIOl9JyNqAPyXRFjMVTXaETIvR4l
kQkce8pI40UiK+PJ4oN0H1FYYHuG/gb7m+UDf7PIF5lDrPgUNZURSXG9QFSXCbn9
QwIDAQAB
-----END PUBLIC KEY-----`

func TestJWTTokenVerify_Verify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("parse rsa public key pem error %v",err)
	}

	v := JWTTokenVerify{PublicKey: pubKey}

	token := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjE1MTQwMTUsImlhdCI6MTY2MTUwNjgxNSwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjMwODgyODc3MDU0MWI1ZjI0OTE4NjVkIn0.s-27hBy_ox5DdSMNBkaw-50yOqcAhxa5A9jL02viNe914jFNWDRHqVVhQss7tMYZmO7SOWNkD8f0wqRoYEf2HWCUbvG_D3cHI1jpxhp4W9etXOSoIoH05wt6fHlryFl9mGQXSm6EdTk_H04CVwiWXwtIEzo0IPtnhC0c6mWj1SaSb49HG9uVMojl5qaX2qkOhKcGgbJ0zTuEM_boNmA5pYmFXEfTbNYOGbFUdp4E2CouwksKrQIT3-5ddewY9uMlyZSilIII8sGghDzzWm1f7pjNE2VId7FuKd-dQvvsGLGzTlOgH5UP1xrlWWI8UDScJDGbjCsqtdfgBO-_gzBWDg"


	jwt.TimeFunc = func() time.Time {
		//time.Unix(1316246222,0)
		return time.Now()
	}

	verify, err := v.Verify(token)
	if err != nil {
		t.Errorf("token is not Verified, error %v",err)
	}

	t.Log("verify info ",verify)
}
