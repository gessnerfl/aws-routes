package routes

//IPPrefix type representation of a IP prefix of the AWS IP range API
type IPPrefix struct {
	IPPrefix string `json:"ip_prefix"`
	Region   string `json:"region"`
	Service  string `json:"service"`
}

//IPAddressRanges type representation of the response message of the AWS IP range API
type IPAddressRanges struct {
	SyncToken  string     `json:"syncToken"`
	CreateDate string     `json:"createDate"`
	Prefixes   []IPPrefix `json:"prefixes"`
}
