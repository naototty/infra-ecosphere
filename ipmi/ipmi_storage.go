package ipmi

// port from OpenIPMI
// Storage Network Function
const (
	IPMI_CMD_GET_FRU_INVENTORY_AREA_INFO =		0x10
	IPMI_CMD_READ_FRU_DATA =			0x11
	IPMI_CMD_WRITE_FRU_DATA =			0x12

	IPMI_CMD_GET_SDR_REPOSITORY_INFO =		0x20
	IPMI_CMD_GET_SDR_REPOSITORY_ALLOC_INFO =	0x21
	IPMI_CMD_RESERVE_SDR_REPOSITORY =		0x22
	IPMI_CMD_GET_SDR =				0x23
	IPMI_CMD_ADD_SDR =				0x24
	IPMI_CMD_PARTIAL_ADD_SDR =			0x25
	IPMI_CMD_DELETE_SDR =				0x26
	IPMI_CMD_CLEAR_SDR_REPOSITORY =			0x27
	IPMI_CMD_GET_SDR_REPOSITORY_TIME =		0x28
	IPMI_CMD_SET_SDR_REPOSITORY_TIME =		0x29
	IPMI_CMD_ENTER_SDR_REPOSITORY_UPDATE =		0x2a
	IPMI_CMD_EXIT_SDR_REPOSITORY_UPDATE =		0x2b
	IPMI_CMD_RUN_INITIALIZATION_AGENT =		0x2c

	IPMI_CMD_GET_SEL_INFO =				0x40
	IPMI_CMD_GET_SEL_ALLOCATION_INFO =		0x41
	IPMI_CMD_RESERVE_SEL =				0x42
	IPMI_CMD_GET_SEL_ENTRY =			0x43
	IPMI_CMD_ADD_SEL_ENTRY =			0x44
	IPMI_CMD_PARTIAL_ADD_SEL_ENTRY =		0x45
	IPMI_CMD_DELETE_SEL_ENTRY =			0x46
	IPMI_CMD_CLEAR_SEL =				0x47
	IPMI_CMD_GET_SEL_TIME =				0x48
	IPMI_CMD_SET_SEL_TIME =				0x49
	IPMI_CMD_GET_AUXILIARY_LOG_STATUS =		0x5a
	IPMI_CMD_SET_AUXILIARY_LOG_STATUS =		0x5b
)
