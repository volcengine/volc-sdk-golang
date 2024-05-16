package businessSecurity

const (
	Ak = "AK" // fill in your access key
	Sk = "SK" // fill in your secret key
)

func init() {
	DefaultInstance.Client.SetAccessKey(Ak)
	DefaultInstance.Client.SetSecretKey(Sk)
	SecuritySourceInstance.Client.SetAccessKey(Ak)
	SecuritySourceInstance.Client.SetSecretKey(Sk)
}
