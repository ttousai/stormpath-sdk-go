package stormpath

type GroupMembership struct {
	resource
	Account Account `json:"account"`
	Group   Group   `json:"group"`
}

type GroupMemberships struct {
	collectionResource
	Items []GroupMembership `json:"items"`
}

func NewGroupMembership(accountHref string, groupHref string) *GroupMembership {
	account := Account{}
	account.Href = accountHref
	group := Group{}
	group.Href = groupHref
	return &GroupMembership{
		Account: account,
		Group:   group,
	}
}

func (groupmembership *GroupMembership) GetAccount() (*Account, error) {
	account := &Account{}

	err := client.get(groupmembership.Account.Href, emptyPayload(), account)

	return account, err
}

func (groupmembership *GroupMembership) GetGroup() (*Group, error) {
	group := &Group{}

	err := client.get(groupmembership.Group.Href, emptyPayload(), group)

	return group, err
}
