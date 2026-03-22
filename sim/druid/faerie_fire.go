package druid

import (
	"github.com/wowsims/tbc/sim/core"
)

func (druid *Druid) registerFaerieFireSpell() {
	auras := druid.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return core.FaerieFireAura(target, float64(druid.Talents.ImprovedFaerieFire))
	})

	druid.FaerieFire = druid.RegisterSpell(Humanoid|Moonkin, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 26993},
		SpellSchool: core.SpellSchoolNature,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			FlatCost: 145,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		ThreatMultiplier: 1,
		FlatThreatBonus:  132,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)

			if result.Landed() {
				druid.TryApplyFaerieFireEffect(sim, target, spell)

				if druid.InForm(Bear) && sim.Proc(0.25, "Mangle CD Reset") {
					druid.MangleBear.CD.Reset()
				}
			}
		},

		RelatedAuraArrays: auras.ToMap(),
	})
}

func (druid *Druid) CanApplyFaerieFireDebuff(target *core.Unit, spell *core.Spell) bool {
	return spell.RelatedAuraArrays["Faerie Fire"].Get(target).IsActive() || !spell.RelatedAuraArrays["Faerie Fire"].Get(target).ExclusiveEffects[0].Category.AnyActive()
}

func (druid *Druid) TryApplyFaerieFireEffect(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
	if druid.CanApplyFaerieFireDebuff(target, spell) {
		aura := spell.RelatedAuraArrays["Faerie Fire"].Get(target)
		aura.Activate(sim)
	}
}
