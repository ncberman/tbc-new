package warlock

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterWarlock()
}

func TestAffliction(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassWarlock,
			Race:       proto.Race_RaceOrc,
			OtherRaces: []proto.Race{proto.Race_RaceHuman, proto.Race_RaceGnome, proto.Race_RaceUndead, proto.Race_RaceBloodElf},
			SpecOptions: core.SpecOptionsCombo{Label: "Affliction", SpecOptions: &proto.Player_Warlock{
				Warlock: &proto.Warlock{
					Options: &proto.Warlock_Options{
						ClassOptions: &proto.WarlockOptions{
							Summon:          proto.WarlockOptions_Imp,
							SacrificeSummon: false,
							Armor:           proto.WarlockOptions_FelArmor,
							CurseOptions:    proto.WarlockOptions_Elements,
						},
					},
				},
			}},
			GearSet:  core.GetGearSet("../../ui/warlock/dps/gear_sets", "preraid"),
			Talents:  "05022221112351055003--50500051220001",
			Rotation: core.GetAplRotation("../../ui/warlock/dps/apls", "affliction"),
			ItemFilter: core.ItemFilter{
				WeaponTypes: []proto.WeaponType{
					proto.WeaponType_WeaponTypeDagger,
					proto.WeaponType_WeaponTypeStaff,
					proto.WeaponType_WeaponTypeSword,
				},
				ArmorType: proto.ArmorType_ArmorTypeCloth,
				RangedWeaponTypes: []proto.RangedWeaponType{
					proto.RangedWeaponType_RangedWeaponTypeWand,
				},
			},
		},
	}))
}

func TestDemonology(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassWarlock,
			Race:       proto.Race_RaceOrc,
			OtherRaces: []proto.Race{proto.Race_RaceHuman, proto.Race_RaceGnome, proto.Race_RaceUndead, proto.Race_RaceBloodElf},
			SpecOptions: core.SpecOptionsCombo{Label: "Demo/Ruin", SpecOptions: &proto.Player_Warlock{
				Warlock: &proto.Warlock{
					Options: &proto.Warlock_Options{
						ClassOptions: &proto.WarlockOptions{
							Summon:          proto.WarlockOptions_Succubus,
							SacrificeSummon: false,
							Armor:           proto.WarlockOptions_FelArmor,
							CurseOptions:    proto.WarlockOptions_Agony,
						},
					},
				},
			}},
			GearSet:  core.GetGearSet("../../ui/warlock/dps/gear_sets", "preraid"),
			Talents:  "01-205003213305010150134-50500251020001",
			Rotation: core.GetAplRotation("../../ui/warlock/dps/apls", "demonology"),
			ItemFilter: core.ItemFilter{
				WeaponTypes: []proto.WeaponType{
					proto.WeaponType_WeaponTypeDagger,
					proto.WeaponType_WeaponTypeStaff,
					proto.WeaponType_WeaponTypeSword,
				},
				ArmorType: proto.ArmorType_ArmorTypeCloth,
				RangedWeaponTypes: []proto.RangedWeaponType{
					proto.RangedWeaponType_RangedWeaponTypeWand,
				},
			},
		},
	}))
}

func TestDestruction(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassWarlock,
			Race:       proto.Race_RaceOrc,
			OtherRaces: []proto.Race{proto.Race_RaceHuman, proto.Race_RaceGnome, proto.Race_RaceUndead, proto.Race_RaceBloodElf},
			SpecOptions: core.SpecOptionsCombo{Label: "Destruction", SpecOptions: &proto.Player_Warlock{
				Warlock: &proto.Warlock{
					Options: &proto.Warlock_Options{
						ClassOptions: &proto.WarlockOptions{
							Summon:          proto.WarlockOptions_Succubus,
							SacrificeSummon: true,
							Armor:           proto.WarlockOptions_FelArmor,
							CurseOptions:    proto.WarlockOptions_Agony,
						},
					},
				},
			}},
			GearSet:  core.GetGearSet("../../ui/warlock/dps/gear_sets", "preraid"),
			Talents:  "-20500301332101-50500051220051053105",
			Rotation: core.GetAplRotation("../../ui/warlock/dps/apls", "destruction"),
			ItemFilter: core.ItemFilter{
				WeaponTypes: []proto.WeaponType{
					proto.WeaponType_WeaponTypeDagger,
					proto.WeaponType_WeaponTypeStaff,
					proto.WeaponType_WeaponTypeSword,
				},
				ArmorType: proto.ArmorType_ArmorTypeCloth,
				RangedWeaponTypes: []proto.RangedWeaponType{
					proto.RangedWeaponType_RangedWeaponTypeWand,
				},
			},
		},
	}))
}
