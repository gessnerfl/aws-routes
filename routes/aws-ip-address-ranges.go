package routes

//AWSIPPrefix type representation of a IP prefix of the AWS IP range API
type AWSIPPrefix struct {
	IPPrefix string `json:"ip_prefix"`
	Region   string `json:"region"`
	Service  string `json:"service"`
}

//AWSIPAddressRanges type representation of the response message of the AWS IP range API
type AWSIPAddressRanges struct {
	SyncToken  string        `json:"syncToken"`
	CreateDate string        `json:"createDate"`
	Prefixes   []AWSIPPrefix `json:"prefixes"`
}
