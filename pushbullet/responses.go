package pushbullet

type UploadReqRes struct {
	Data      *UploadReqDest `json:"data"`
	FileName  string         `json:"file_name"`
	FileType  string         `json:"file_type"`
	FileUrl   string         `json:"file_url"`
	UploadUrl string         `json:"upload_url"`
}

type UploadReqDest struct {
	Acl          string `json:"acl"`
	AwsAccessKey string `json:"awsaccesskey"`
	ContentType  string `json:"content-type"`
	Key          string `json:"key"`
	Policy       string `json:"policy"`
	Signature    string `json:"signature"`
}
