package zincsearch

// Document is a struct that represents the document to be indexed.
// It has two fields: Title and Content.
type Document struct {
	Title   string
	Content string
}

// Caller is a struct that represents the caller of the logger.
// It has two fields: Name and Location.
type _Caller struct {
	Name     string
	Location string
}

// GetDefaultProperties returns the default properties for the index.
// It returns a map of properties with their types and settings.
// Example:
//
//	properties := GetDefaultProperties()
//	fmt.Println(properties)
//
// Output:
//
//	map[properties:map[content:map[highlightable:true index:true store:true type:text] status:map[aggregatable:true index:true sortable:true type:keyword] title:map[highlightable:true index:true store:true type:text]]]
//
// The output is a map of properties with their types and settings.
// The properties are title, content, and status.
// The types are text and keyword.
// The settings are index, store, sortable, and aggregatable.
func GetDefaultProperties() map[string]interface{} {
	return map[string]interface{}{
		"properties": map[string]interface{}{
			"title": map[string]interface{}{
				"type":          "text",
				"index":         true,
				"store":         true,
				"highlightable": true,
			},
			"content": map[string]interface{}{
				"type":          "text",
				"index":         true,
				"store":         true,
				"highlightable": true,
			},
			"status": map[string]interface{}{
				"type":         "keyword",
				"index":        true,
				"sortable":     true,
				"aggregatable": true,
			},
		},
	}
}

// kvDocument returns the key-value representation of the document.
// It takes the document, the caller, and the status as input.
// It returns a map of the document with the caller and the status.
func kvDocument(v *Document, caller _Caller, status string) map[string]interface{} {
	return map[string]interface{}{
		"title":   v.Title,
		"content": v.Content,
		"status":  status,
		"caller": map[string]interface{}{
			"name": caller.Name,
			"loc":  caller.Location,
		},
	}
}
