package stormpath

//collectionResource represent the basic attributes of collection of resources (Application, Group, Account, etc.)
type collectionResource struct {
	Href       string `json:"href,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	ModifiedAt string `json:"modifiedAt,omitempty"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
}

func (r collectionResource) IsCacheable() bool {
	return false
}

//resource resprents the basic attributes of any resource (Application, Group, Account, etc.)
type resource struct {
	Href       string `json:"href,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	ModifiedAt string `json:"modifiedAt,omitempty"`
}

func (r resource) IsCacheable() bool {
	return true
}

//Refresh refreshes the resource by doing a GET to the resource href endpoint
func (r *resource) Refresh() error {
	return client.get(r.Href, emptyPayload(), r)
}

//Save updates the given resource, by doing a POST to the resource Href
func (r *resource) Save() error {
	return client.post(r.Href, r, r)
}

//Delete deletes the given account, it wont modify the calling account
func (r *resource) Delete() error {
	return client.delete(r.Href, emptyPayload())
}

type accountStoreResource struct {
	customDataAwareResource
	Accounts *Accounts `json:"accounts,omitempty"`
}

//GetAccounts returns all the accounts of the application
//
//See: http://docs.stormpath.com/rest/product-guide/#application-accounts
func (r *accountStoreResource) GetAccounts(criteria Criteria) (*Accounts, error) {
	accounts := &Accounts{}

	err := client.get(
		buildAbsoluteURL(r.Accounts.Href, criteria.ToQueryString()),
		emptyPayload(),
		accounts,
	)

	return accounts, err
}
