package helpers

import "strconv"

const (
  FormatInt64 = 10
)

func GenerateTenantCode(tenantID int64) string {
  return "NSS" + strconv.FormatInt(tenantID, FormatInt64)
}
