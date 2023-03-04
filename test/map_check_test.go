package test

import (
	"testing"

	"github.com/mgechev/revive/lint"
	"github.com/mgechev/revive/rule"
)

func TestMapCheck(t *testing.T) {
	testRule(t, "map-check", &rule.MapCheckRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(1)},
	})
}
