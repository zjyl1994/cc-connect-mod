package codex

import (
	"testing"

	"github.com/chenhg5/cc-connect/core"
)

func TestAutoReviewMode_ConfigSwitchAndWorkspaceOptions(t *testing.T) {
	for _, raw := range []string{"auto-review", " AUTO-REVIEW ", "auto_review", "autoreview", "approve-for-me"} {
		t.Run(raw, func(t *testing.T) {
			agent, err := New(map[string]any{"cmd": "go", "mode": raw})
			if err != nil {
				t.Fatal(err)
			}
			a := agent.(*Agent)
			if got := a.GetMode(); got != "auto-review" {
				t.Fatalf("configured mode = %q", got)
			}
			a.SetMode("suggest")
			a.SetMode(raw)
			if got := a.WorkspaceAgentOptions()["mode"]; got != "auto-review" {
				t.Fatalf("workspace mode = %v", got)
			}
			for _, mode := range a.PermissionModes() {
				if mode.Key == "auto-review" {
					for _, lang := range []core.Language{
						core.LangEnglish, core.LangChinese, core.LangTraditionalChinese,
						core.LangJapanese, core.LangSpanish,
					} {
						i18n := core.NewI18n(lang)
						if mode.LocalizedName(i18n) == "" || mode.LocalizedDescription(i18n) == "" {
							t.Fatalf("missing mode translation for %s", lang)
						}
					}
					return
				}
			}
			t.Fatal("auto-review is absent from the mode menu")
		})
	}
}
