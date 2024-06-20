package request

type InsertMessageModel struct {
	Id                            uint64 `json:"id"`
	Type                          string `json:"type"` // Nombre del campo en MySQL
	Content                       string `json:"content"`
	MobileNumber                  string `json:"mobile_number"`
	MobileCountryISOCode          string `json:"mobile_country_iso_code"`
	ShortNumber                   string `json:"short_number"`
	Telco                         string `json:"telco"`
	Created                       string `json:"created"`
	RoutingType                   string `json:"routing_type"`
	MatchedPattern                string `json:"matched_pattern"`
	ServiceID                     string `json:"service_id"`
	TelcoID                       string `json:"telco_id"`
	SessionAction                 string `json:"session_action"`
	SessionParametersMap          string `json:"session_parameters_map"`
	SessionTimeoutSeconds         uint64 `json:"session_timeout_seconds"`
	Priority                      uint64 `json:"priority"`
	ClientID                      string `json:"client_id"`
	URL                           string `json:"url"`
	AccessTimeoutSeconds          uint64 `json:"access_timeout_seconds"`
	RequestID                     uint64 `json:"request_id"`
	DefaultActionID               uint64 `json:"default_action_id"`
	ApplicationID                 uint64 `json:"application_id"`
	SessionID                     uint64 `json:"session_id"`
	Processed                     string `json:"processed"`
	MillisSinceRequest            uint64 `json:"millis_since_request"`
	SessionApplicationName        string `json:"session_application_name"`
	Sendafter                     string `json:"sendafter"`
	Sendbefore                    string `json:"sendbefore"`
	Sent                          string `json:"sent"`
	Status                        string `json:"status"`
	AccessTimeoutHandlerQueuename string `json:"access_timeout_handler_queuename"`
	UseUnsupportedMobilesRegistry uint64 `json:"use_unsupported_mobiles_registry"`
	OriginName                    string `json:"origin_name"`
}
