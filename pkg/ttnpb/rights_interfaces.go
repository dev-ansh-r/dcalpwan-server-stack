package ttnpb

func (m *Collaborator) EntityType() string {
	return m.Ids.EntityType()
}

func (m *Collaborator) IDString() string {
	return m.Ids.IDString()
}
