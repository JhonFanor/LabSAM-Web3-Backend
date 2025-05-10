package utils

import (
	"lamsam-web3-backend/internal/consts"
	"strings"
)

func ResolveTableName(model string) string {
	if table, ok := consts.ModelToTable[model]; ok {
		return table
	}
	if strings.HasSuffix(model, "y") {
		return strings.TrimSuffix(model, "y") + "ies"
	}
	return model + "s"
}
