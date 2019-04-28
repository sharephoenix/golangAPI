package database

type ModuleModel struct {
	ModuleId	string 		`json:"module_id" bson:"module_id"`
	ModuleName 	string 		`json:"module_name" bson:"module_name"`
	ModuleDisplayName 	string `json:"module_display_name" bson:"module_display_name"`
}

type EventModel struct {
	EventId 	string 		`json:"event_id" bson:"event_id"`
	ModuleId	string 		`json:"module_id" bson:"module_id"`
	EventName 	string 		`json:"event_name" bson:"event_name"`
	EventDisplayName 	string `json:"event_display_name" bson:"event_display_name"`
	EventParams 	string		`json:"event_params bson:"event_params"`
	EventResponse	string		`json:"event_response" bson:"event_response"`
	IsEnable 		bool		`json:"is_enable" bson:"is_enable"`
}
