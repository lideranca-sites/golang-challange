package features

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getUserID(c *fiber.Ctx) (int, bool) {
	tryKeys := []string{"user_id", "userId", "userid", "uid", "id"}

	for _, k := range tryKeys {
		if v := c.Locals(k); v != nil {
			switch t := v.(type) {
			case int:
				return t, true
			case int64:
				return int(t), true
			case uint:
				return int(t), true
			case uint64:
				return int(t), true
			case float64:
				return int(t), true
			case string:
				if n, err := strconv.Atoi(t); err == nil {
					return n, true
				}
			}
		}
	}

	for _, k := range []string{"user", "claims"} {
		if v := c.Locals(k); v != nil {
			if m, ok := v.(map[string]any); ok {
				if raw, ok := m["user_id"]; ok {
					switch t := raw.(type) {
					case int:
						return t, true
					case float64:
						return int(t), true
					case string:
						if n, err := strconv.Atoi(t); err == nil {
							return n, true
						}
					}
				}
			}
		}
	}

	return 0, false
}
