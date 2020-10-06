package routes

//IPPrefix type representation of a IP prefix of the AWS IP range API
type IPPrefix struct {
	IPPrefix string `json:"ip_prefix"`
	Region   string `json:"region"`
	Service  string `json:"service"`
}

//IsRelevantService returns true if the service is relevant for routing
func (ipPrefix IPPrefix) IsRelevantService() bool {
	return ipPrefix.Service == "AMAZON"
}

//IPAddressRanges type representation of the response message of the AWS IP range API
type IPAddressRanges struct {
	SyncToken  string     `json:"syncToken"`
	CreateDate string     `json:"createDate"`
	Prefixes   []IPPrefix `json:"prefixes"`
}
