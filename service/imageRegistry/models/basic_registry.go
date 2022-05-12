package models

type CreateNamespaceBasicRequest struct {
	Name string `json:"Name"`
}

type CreateNamespaceBasicResponse struct {
}

type DeleteNamespaceBasicRequest struct {
	Name string `json:"Name"`
}

type DeleteNamespaceBasicResponse struct {
}

type GetNamespaceBasicRequest struct {
	Name string `json:"Name"`
}

type GetNamespaceBasicResponse struct {
	Namespace
}

type ValidateNamespaceBasicRequest struct {
	Name string `json:"Name"`
}

type ValidateNamespaceBasicResponse struct {
	Existed bool `json:"Existed"`
}

type ListNamespacesBasicRequest struct {
	Page
	Filter NamespaceFilter `json:"Filter,omitempty"`
}

type ListNamespacesBasicResponse struct {
	Total         int64       `json:"Total"`
	Items         []Namespace `json:"Items"`
	CreatedAmount int64       `json:"CreatedAmount"`
}

type CreateRepositoryBasicRequest struct {
	Namespace   string `json:"Namespace"`
	Name        string `json:"Name"`
	Type        string `json:"Type,omitempty"`
	Description string `json:"Description,omitempty"`
}

type CreateRepositoryBasicResponse struct {
}

type UpdateRepositoryBasicRequest struct {
	Namespace   string `json:"Namespace"`
	Name        string `json:"Name"`
	Type        string `json:"Type,omitempty"`
	Description string `json:"Description,omitempty"`
}

type UpdateRepositoryBasicResponse struct {
}

type DeleteRepositoryBasicRequest struct {
	Namespace string `json:"Namespace"`
	Name      string `json:"Name"`
}

type DeleteRepositoryBasicResponse struct {
}

type GetRepositoryBasicRequest struct {
	Namespace string `json:"Namespace"`
	Name      string `json:"Name"`
}

type GetRepositoryBasicResponse struct {
	Repository
}

type ValidateRepositoryBasicRequest struct {
	Namespace string `json:"Namespace"`
	Name      string `json:"Name"`
}

type ValidateRepositoryBasicResponse struct {
	Existed bool `json:"Existed"`
}

type ListRepositoriesBasicRequest struct {
	Page
	Filter *RepositoryFilter `json:"Filter,omitempty"`
}

type ListRepositoriesBasicResponse struct {
	Total         int64        `json:"Total"`
	Domain        string       `json:"Domain"`
	Items         []Repository `json:"Items"`
	CreatedAmount int64        `json:"CreatedAmount"`
	VpcDomain     string       `json:"VpcDomain"`
}

type DeleteTagBasicRequest struct {
	Namespace string `json:"Namespace"`
	Repo      string `json:"Repo"`
	Name      string `json:"Name"`
}

type DeleteTagBasicResponse struct {
}

type GetTagBasicRequest struct {
	Namespace string `json:"Namespace"`
	Repo      string `json:"Repo"`
	Name      string `json:"Name"`
}

type GetTagBasicResponse struct {
	Namespace string `json:"Namespace"`
	Repo      string `json:"Repo"`
	Tag       Tag    `json:"Tag"`
}

type GetTagAdditionBasicRequest struct {
	Namespace string `json:"Namespace"`
	Repo      string `json:"Repo"`
	Name      string `json:"Name"`
	Digest    string `json:"Digest"`
	Addition  string `json:"Addition"`
}

type GetTagAdditionBasicResponse struct {
	Contents string `json:"Contents"`
}

type ListTagsBasicRequest struct {
	Namespace string `json:"Namespace"`
	Repo      string `json:"Repo"`
	Type      string `json:"Type"`
	Page
	Filter TagFilter `json:"Filter,omitempty"`
}

type ListTagsBasicResponse struct {
	Total     int64  `json:"Total"`
	Namespace string `json:"Namespace"`
	Repo      string `json:"Repo"`
	Items     []Tag  `json:"Items"`
}

type GetAuthorizationTokenBasicRequest struct {
}

type GetAuthorizationTokenBasicResponse struct {
	AuthorizationToken string `json:"AuthorizationToken"`
	AuthorizationUser  string `json:"AuthorizationUser"`
	ExpireTime         string `json:"ExpireTime"`
	Domain             string `json:"Domain"`
	VpcDomain          string `json:"VpcDomain"`
}

type ScanPolicy struct {
	Cron   string `json:"Cron"`
	Status string `json:"Status"`
}

type NamespaceFilter struct {
	Name string `json:"Name,omitempty"`
}

type Namespace struct {
	Name       string     `json:"Name"`
	CreateTime string     `json:"CreateTime"`
	RepoAmount int64      `json:"RepoAmount"`
	ScanPolicy ScanPolicy `json:"ScanPolicy"`
}

type RepositoryFilter struct {
	Namespace string `json:"Namespace,omitempty"`
	Name      string `json:"Name,omitempty"`
	Type      string `json:"Type,omitempty"`
}

type Repository struct {
	Namespace   string `json:"Namespace"`
	Name        string `json:"Name"`
	Type        string `json:"Type"`
	CreateTime  string `json:"CreateTime"`
	UpdateTime  string `json:"UpdateTime"`
	Description string `json:"Description,omitempty"`
}

type TagFilter struct {
	Name string `json:"Name,omitempty"`
}

type Tag struct {
	Type       string      `json:"Type"`
	Name       string      `json:"Name"`
	Digest     string      `json:"Digest"`
	PushTime   string      `json:"PushTime"`
	SizeByte   int64       `json:"SizeByte"`
	ScanResult *ScanResult `json:"ScanResult,omitempty"`
	ImageAttrs *ImageAttrs `json:"ImageAttrs,omitempty"`
	ChartAttrs `json:"ChartAttrs,omitempty"`
}

type ScanResult struct {
	Status  string      `json:"Status"`
	Summary ScanSummary `json:"Summary"`
}

type ScanSummary struct {
	CriticalAmount int64 `json:"CriticalAmount"`
	HighAmount     int64 `json:"HighAmount"`
	MediumAmount   int64 `json:"MediumAmount"`
	LowAmount      int64 `json:"LowAmount"`
	TotalAmount    int64 `json:"TotalAmount"`
}

type ImageAttrs struct {
	Author       string `json:"Author"`
	Architecture string `json:"Architecture"`
	Os           string `json:"Os"`
}

type ChartAttrs struct {
	ApiVersion  string `json:"ApiVersion"`
	AppVersion  string `json:"AppVersion"`
	Description string `json:"Description"`
	Name        string `json:"Name"`
	Type        string `json:"Type"`
	Version     string `json:"Version"`
}
