package application

// Delete value from repository
func (app *Librapi) DeleteDataFromStorage(key string) error {
	return app.storage.Delete(key)
}

type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Get value from repository
func (app *Librapi) GetDataFromStorage(key string) (*Entry, error) {

	var body Entry
	body.Key = key

	err := app.storage.Get(key, &body.Value)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

// Save value in Repository
func (app *Librapi) SaveDataInStorage(key, value string) error {
	return app.storage.Put(key, value)
}
