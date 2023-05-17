package main

type DatasetPostData struct {
	Name       string
	FileName   string
	FolderPath string
	Container  string
}

type TypeProperties struct {
	Location struct {
		Type       string `json:"type"`
		FileName   string `json:"fileName"`
		FolderPath string `json:"folderPath"`
		Container  string `json:"container"`
	} `json:"location"`
}

type LinkedServiceName struct {
	ReferenceName string `json:"referenceName"`
	Type          string `json:"type"`
}

type Properties struct {
	LinkedServiceName LinkedServiceName `json:"linkedServiceName"`
	Annotations       []interface{}     `json:"annotations"`
	Type              string            `json:"type"`
	TypeProperties    TypeProperties    `json:"typeProperties"`
}

type Dataset struct {
	Name       string     `json:"name"`
	Properties Properties `json:"properties"`
	Type       string     `json:"type"`
}

func NewDataset(name, linkedServiceRef, container, fileName, folderPath string) *Dataset {
	return &Dataset{
		Name: name,
		Type: "Microsoft.DataFactory/factories/datasets",
		Properties: Properties{
			LinkedServiceName: LinkedServiceName{
				ReferenceName: linkedServiceRef,
				Type:          "LinkedServiceReference",
			},
			Annotations: []interface{}{},
			Type:        "Binary",
			TypeProperties: TypeProperties{
				Location: struct {
					Type       string `json:"type"`
					FileName   string `json:"fileName"`
					FolderPath string `json:"folderPath"`
					Container  string `json:"container"`
				}{
					Type:       "AzureBlobStorageLocation",
					FileName:   fileName,
					FolderPath: folderPath,
					Container:  container,
				},
			},
		},
	}
}
