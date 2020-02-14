package inventory

func Swords() Inventory {
	return []Item {
		{0,"Broad Sword"},
		{1,"Bastard Sword"},
		{2,"Heat Blade"},
		{3,"Falchion"},
		{4,"Mind Crush"},
		{5,"Fairy Sword"},
		{6,"Claymore"},
		{7,"Soul Eater"},
	}
}

func Axes() Inventory {
	return []Item {
		{8,"Axe"},
		{9,"Tomahawk"},
		{10,"Battle Axe"},
		{11,"Great Axe"},
		{12,"Indora's Axe"},
	}
}

func Daggers() Inventory {
	return []Item{
		{13,"Rapier"},
		{19,"Demon Stiletto"},
		{14,"Shadow Cutter"},
		{15,"Dancing Dagger"},
		{16,"Flamberge"},
		{17,"Gladius"},
		{18,"Dragon Buster"},
	}
}

func Spears() Inventory {
	return []Item {
		{20,"Spear"},
		{21,"Lance"},
		{26,"Twister Glaive"},
		{22,"Glaive"},
		{23,"Spear Of Terror"},
		{24,"Partisan"},
		{25,"Halberd"},
	}
}

func Bows() Inventory {
	return []Item {
		{27,"Short Bow"},
		{28,"Sparkle Arrow"},
		{29,"Long Bow"},
		{30,"Bemusing Arrow"},
		{31,"Virulent Arrow"},
		{33,"Arrow Of Force"},
		{32,"Detonate Arrow"},
	}
}

func Maces() Inventory {
	return []Item {
		{34,"Mace"},
		{35,"Morning Star"},
		{36,"War Hammer"},
		{37,"Heavy Mace"},
		{39,"Pretty Hammer"},
		{38,"Basher"},
	}
}

func Knuckles() Inventory {
	return []Item {
		{40,"Iron Knuckle"},
		{41,"Beast Fang"},
		{42,"Diamond Claw"},
		{45,"Brass Knuckle"},
		{43,"Thunder Fist"},
		{44,"Destroyer Mace"},
	}
}
