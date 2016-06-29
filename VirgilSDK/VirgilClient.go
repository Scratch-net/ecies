package VirgilSDK

type VirgilClient struct {
	BaseUrl string
	Token string
}

func (client *VirgilClient) Request(method string, endpoint string) (err error, result string){
	return nil, "ok"
}