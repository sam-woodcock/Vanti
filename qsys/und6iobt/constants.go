package und6iobt

//https://www.qsys.com/resource-files/productresources/dn/attero_tech/api_documentation/q_dn_attero_tech_by_qsc_un_series_3rd_party_api.pdf
// Constants used for storing commands to send to device

const (
	ActivatePairing           = "BTB<CR>"
	ActivatePairingResponseOk = "ACK BTB<CR>"
	ActivatePairingResponseNo = "NACK BTB<CR>"
	BTStatus                  = "BTS<CR>"
	Connected                 = "ACK BTS 2<CR>" //Bluetooth® Interface has an active connection

	ReadyForPairing = "ready For Pairing"
	PairingFailed   = "pairing Failed"
	ErrorMessage    = "something went wrong"
)
