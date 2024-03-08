package ethics

//go:generate stringer -type=EthicKind

type EthicKind int32

const (
	EthicGestaltConsciousness EthicKind = 1 // a gestalt consciousness, a hive mind, e.g. the zerg
	EthicAuthoritarian        EthicKind = 2 // an authoritarian faction, e.g. the empire
	EthicEgalitarian          EthicKind = 3 // an egalitarian faction, e.g. the federation
	EthicXenophobic           EthicKind = 4 // a xenophobic faction, e.g. the covenant
	EthicXenophilic           EthicKind = 5 // a xenophilic faction, e.g. the alliance
	EthicMilitarist           EthicKind = 6 // a militarist faction, e.g. the horde
	EthicPacifist             EthicKind = 7 // a pacifist faction, e.g. the republic
	EthicSpiritualist         EthicKind = 8 // a spiritualist faction, e.g. the church
	EthicMaterialist          EthicKind = 9 // a materialist faction, e.g. the corporation
	EthicKindMAX              EthicKind = 10
)

// Inspired by the ethics of Stellaris.
// Each faction may have three ethic points. Each ethic point may be spent on one ethic.
// Going fanatic on an ethic doubles the effect of the ethic, and also costs one point.
// Having an ethic besides gestalt consciousness prevents the faction from taking on an opposite ethic.
//
// Mutually exclusive ethics are:
// - Authoritarian and Egalitarian
// - Xenophobic and Xenophilic
// - Militarist and Pacifist
// - Spiritualist and Materialist
//

type Ethic struct {
	EthicKind EthicKind // the kind of ethic
	IsFanatic bool      // double the effect of the ethic
}

func (e Ethic) String() string {
	repr := "("
	if e.IsFanatic {
		repr = "(Fanatic "
	}
	repr += e.EthicKind.String()
	repr += ")"
	return repr
}
