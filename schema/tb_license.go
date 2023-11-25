package schema

type TbLicense struct {
	IdLicense int    `json:"id_license"`
	IdItem    int64  `json:"id_item"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	SPDXID    string `json:"spdx_id"`
	URL       string `json:"url"`
	NodeID    string `json:"node_id"`
}
