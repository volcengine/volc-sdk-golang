package businessSecurity

const (
	Ak = "AK" // write your access key
	Sk = "SK" // write your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
	SecuritySourceInstance.Client.SetAccessKey(Ak)
	SecuritySourceInstance.Client.SetSecretKey(Sk)
}
