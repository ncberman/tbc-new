package paladin

import (
	"github.com/wowsims/tbc/sim/core"
)

var ItemSetJusticarBattlegear = core.NewItemSet(core.ItemSet{
	ID:   626,
	Name: "Justicar Battlegear",
	Bonuses: map[int32]core.ApplySetBonus{
		2: func(_ core.Agent, setBonusAura *core.Aura) {
			setBonusAura.AttachSpellMod(core.SpellModConfig{
				Kind:       core.SpellMod_DamageDone_Flat,
				FloatValue: 0.15,
				ClassMask:  SpellMaskJudgementOfTheCrusader,
			})
		},
		4: func(agent core.Agent, setBonusAura *core.Aura) {
			setBonusAura.AttachSpellMod(core.SpellModConfig{
				Kind:       core.SpellMod_DamageDone_Pct,
				FloatValue: 0.10,
				ClassMask:  SpellMaskJudgementOfCommand,
			})
		},
	},
})
