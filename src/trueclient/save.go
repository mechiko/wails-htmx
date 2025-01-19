package trueclient

// сохраняем в конфиге токены и время получения
func (t *trueClient) Save() {
	// _ = t.Config().Set("trueclient.tokengis", t.tokenGis, true)
	// _ = t.Config().Set("trueclient.tokensuz", t.tokenSuz, true)
	// aTime := t.authTime.Format(t.layout)
	// _ = t.Config().Set("trueclient.authtime", aTime, true)
	_ = t.IApp.Config().Set("trueclient.deviceid", t.deviсeId, true)
	_ = t.IApp.Config().Set("trueclient.omsid", t.omsId, true)
	_ = t.IApp.Config().Set("trueclient.hashkey", t.hash, true)
}
