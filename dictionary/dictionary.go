package dictionary

type Dictionary struct {
	translations map[string]string
}

func NewDictionary() Dictionary {
	return Dictionary{
		translations: make(map[string]string),
	}
}

func (d *Dictionary) Get(key string) string {
	return d.translations[key]
}

func (d *Dictionary) Remove(key string) {
	delete(d.translations, key)
}

func (d *Dictionary) List() []string {
	result := make([]string, 0, len(d.translations))
	for key, value := range d.translations {
		result = append(result, key+": "+value+",")
	}
	return result
}

func (d *Dictionary) Add(key, value string) {
	d.translations[key] = value
}
