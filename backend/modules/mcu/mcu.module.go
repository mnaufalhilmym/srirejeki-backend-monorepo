package mcu

func Module(m *Microcontroller) {
	getAllMicrocontrollers(m)
	getFarmlandMicrocontrollers(m)
	getMicrocontroller(m)
	postMicrocontroller(m)
	patchMicrocontroller(m)
	deleteMicrocontroller(m)
	authMicrocontroller(m)
	postSendDataToMcu(m)
}
