// Code generated by generator/blockentities/main.go; DO NOT EDIT.

package block

var EntityList = [...]Entity{
	FurnaceEntity{},
	ChestEntity{},
	TrappedChestEntity{},
	EnderChestEntity{},
	JukeboxEntity{},
	DispenserEntity{},
	DropperEntity{},
	SignEntity{},
	HangingSignEntity{},
	MobSpawnerEntity{},
	PistonEntity{},
	BrewingStandEntity{},
	EnchantingTableEntity{},
	EndPortalEntity{},
	BeaconEntity{},
	SkullEntity{},
	DaylightDetectorEntity{},
	HopperEntity{},
	ComparatorEntity{},
	BannerEntity{},
	StructureBlockEntity{},
	EndGatewayEntity{},
	CommandBlockEntity{},
	ShulkerBoxEntity{},
	BedEntity{},
	ConduitEntity{},
	BarrelEntity{},
	SmokerEntity{},
	BlastFurnaceEntity{},
	LecternEntity{},
	BellEntity{},
	JigsawEntity{},
	CampfireEntity{},
	BeehiveEntity{},
	SculkSensorEntity{},
	CalibratedSculkSensorEntity{},
	SculkCatalystEntity{},
	SculkShriekerEntity{},
	ChiseledBookshelfEntity{},
	BrushableBlockEntity{},
	DecoratedPotEntity{},
	CrafterEntity{},
	TrialSpawnerEntity{},
	VaultEntity{},
}

func (FurnaceEntity) ID() string               { return "minecraft:furnace" }
func (ChestEntity) ID() string                 { return "minecraft:chest" }
func (TrappedChestEntity) ID() string          { return "minecraft:trapped_chest" }
func (EnderChestEntity) ID() string            { return "minecraft:ender_chest" }
func (JukeboxEntity) ID() string               { return "minecraft:jukebox" }
func (DispenserEntity) ID() string             { return "minecraft:dispenser" }
func (DropperEntity) ID() string               { return "minecraft:dropper" }
func (SignEntity) ID() string                  { return "minecraft:sign" }
func (HangingSignEntity) ID() string           { return "minecraft:hanging_sign" }
func (MobSpawnerEntity) ID() string            { return "minecraft:mob_spawner" }
func (PistonEntity) ID() string                { return "minecraft:piston" }
func (BrewingStandEntity) ID() string          { return "minecraft:brewing_stand" }
func (EnchantingTableEntity) ID() string       { return "minecraft:enchanting_table" }
func (EndPortalEntity) ID() string             { return "minecraft:end_portal" }
func (BeaconEntity) ID() string                { return "minecraft:beacon" }
func (SkullEntity) ID() string                 { return "minecraft:skull" }
func (DaylightDetectorEntity) ID() string      { return "minecraft:daylight_detector" }
func (HopperEntity) ID() string                { return "minecraft:hopper" }
func (ComparatorEntity) ID() string            { return "minecraft:comparator" }
func (BannerEntity) ID() string                { return "minecraft:banner" }
func (StructureBlockEntity) ID() string        { return "minecraft:structure_block" }
func (EndGatewayEntity) ID() string            { return "minecraft:end_gateway" }
func (CommandBlockEntity) ID() string          { return "minecraft:command_block" }
func (ShulkerBoxEntity) ID() string            { return "minecraft:shulker_box" }
func (BedEntity) ID() string                   { return "minecraft:bed" }
func (ConduitEntity) ID() string               { return "minecraft:conduit" }
func (BarrelEntity) ID() string                { return "minecraft:barrel" }
func (SmokerEntity) ID() string                { return "minecraft:smoker" }
func (BlastFurnaceEntity) ID() string          { return "minecraft:blast_furnace" }
func (LecternEntity) ID() string               { return "minecraft:lectern" }
func (BellEntity) ID() string                  { return "minecraft:bell" }
func (JigsawEntity) ID() string                { return "minecraft:jigsaw" }
func (CampfireEntity) ID() string              { return "minecraft:campfire" }
func (BeehiveEntity) ID() string               { return "minecraft:beehive" }
func (SculkSensorEntity) ID() string           { return "minecraft:sculk_sensor" }
func (CalibratedSculkSensorEntity) ID() string { return "minecraft:calibrated_sculk_sensor" }
func (SculkCatalystEntity) ID() string         { return "minecraft:sculk_catalyst" }
func (SculkShriekerEntity) ID() string         { return "minecraft:sculk_shrieker" }
func (ChiseledBookshelfEntity) ID() string     { return "minecraft:chiseled_bookshelf" }
func (BrushableBlockEntity) ID() string        { return "minecraft:brushable_block" }
func (DecoratedPotEntity) ID() string          { return "minecraft:decorated_pot" }
func (CrafterEntity) ID() string               { return "minecraft:crafter" }
func (TrialSpawnerEntity) ID() string          { return "minecraft:trial_spawner" }
func (VaultEntity) ID() string                 { return "minecraft:vault" }

func (f FurnaceEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:furnace"
}

func (c ChestEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:chest"
}

func (t TrappedChestEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:trapped_chest"
}

func (e EnderChestEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:ender_chest"
}

func (j JukeboxEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:jukebox"
}

func (d DispenserEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:dispenser"
}

func (d DropperEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:dropper"
}

func (s SignEntity) IsValidBlock(block Block) bool {
	switch block.ID() {
	case "minecraft:oak_sign",
		"minecraft:spruce_sign",
		"minecraft:birch_sign",
		"minecraft:acacia_sign",
		"minecraft:cherry_sign",
		"minecraft:jungle_sign",
		"minecraft:dark_oak_sign",
		"minecraft:oak_wall_sign",
		"minecraft:spruce_wall_sign",
		"minecraft:birch_wall_sign",
		"minecraft:acacia_wall_sign",
		"minecraft:cherry_wall_sign",
		"minecraft:jungle_wall_sign",
		"minecraft:dark_oak_wall_sign",
		"minecraft:crimson_sign",
		"minecraft:crimson_wall_sign",
		"minecraft:warped_sign",
		"minecraft:warped_wall_sign",
		"minecraft:mangrove_sign",
		"minecraft:mangrove_wall_sign",
		"minecraft:bamboo_sign",
		"minecraft:bamboo_wall_sign":
		return true
	default:
		return false
	}
}

func (h HangingSignEntity) IsValidBlock(block Block) bool {
	switch block.ID() {
	case "minecraft:oak_hanging_sign",
		"minecraft:spruce_hanging_sign",
		"minecraft:birch_hanging_sign",
		"minecraft:acacia_hanging_sign",
		"minecraft:cherry_hanging_sign",
		"minecraft:jungle_hanging_sign",
		"minecraft:dark_oak_hanging_sign",
		"minecraft:crimson_hanging_sign",
		"minecraft:warped_hanging_sign",
		"minecraft:mangrove_hanging_sign",
		"minecraft:bamboo_hanging_sign",
		"minecraft:oak_wall_hanging_sign",
		"minecraft:spruce_wall_hanging_sign",
		"minecraft:birch_wall_hanging_sign",
		"minecraft:acacia_wall_hanging_sign",
		"minecraft:cherry_wall_hanging_sign",
		"minecraft:jungle_wall_hanging_sign",
		"minecraft:dark_oak_wall_hanging_sign",
		"minecraft:crimson_wall_hanging_sign",
		"minecraft:warped_wall_hanging_sign",
		"minecraft:mangrove_wall_hanging_sign",
		"minecraft:bamboo_wall_hanging_sign":
		return true
	default:
		return false
	}
}

func (m MobSpawnerEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:spawner"
}

func (p PistonEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:moving_piston"
}

func (b BrewingStandEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:brewing_stand"
}

func (e EnchantingTableEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:enchanting_table"
}

func (e EndPortalEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:end_portal"
}

func (b BeaconEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:beacon"
}

func (s SkullEntity) IsValidBlock(block Block) bool {
	switch block.ID() {
	case "minecraft:skeleton_skull",
		"minecraft:skeleton_wall_skull",
		"minecraft:creeper_head",
		"minecraft:creeper_wall_head",
		"minecraft:dragon_head",
		"minecraft:dragon_wall_head",
		"minecraft:zombie_head",
		"minecraft:zombie_wall_head",
		"minecraft:wither_skeleton_skull",
		"minecraft:wither_skeleton_wall_skull",
		"minecraft:player_head",
		"minecraft:player_wall_head",
		"minecraft:piglin_head",
		"minecraft:piglin_wall_head":
		return true
	default:
		return false
	}
}

func (d DaylightDetectorEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:daylight_detector"
}

func (h HopperEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:hopper"
}

func (c ComparatorEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:comparator"
}

func (b BannerEntity) IsValidBlock(block Block) bool {
	switch block.ID() {
	case "minecraft:white_banner",
		"minecraft:orange_banner",
		"minecraft:magenta_banner",
		"minecraft:light_blue_banner",
		"minecraft:yellow_banner",
		"minecraft:lime_banner",
		"minecraft:pink_banner",
		"minecraft:gray_banner",
		"minecraft:light_gray_banner",
		"minecraft:cyan_banner",
		"minecraft:purple_banner",
		"minecraft:blue_banner",
		"minecraft:brown_banner",
		"minecraft:green_banner",
		"minecraft:red_banner",
		"minecraft:black_banner",
		"minecraft:white_wall_banner",
		"minecraft:orange_wall_banner",
		"minecraft:magenta_wall_banner",
		"minecraft:light_blue_wall_banner",
		"minecraft:yellow_wall_banner",
		"minecraft:lime_wall_banner",
		"minecraft:pink_wall_banner",
		"minecraft:gray_wall_banner",
		"minecraft:light_gray_wall_banner",
		"minecraft:cyan_wall_banner",
		"minecraft:purple_wall_banner",
		"minecraft:blue_wall_banner",
		"minecraft:brown_wall_banner",
		"minecraft:green_wall_banner",
		"minecraft:red_wall_banner",
		"minecraft:black_wall_banner":
		return true
	default:
		return false
	}
}

func (s StructureBlockEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:structure_block"
}

func (e EndGatewayEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:end_gateway"
}

func (c CommandBlockEntity) IsValidBlock(block Block) bool {
	switch block.ID() {
	case "minecraft:command_block",
		"minecraft:chain_command_block",
		"minecraft:repeating_command_block":
		return true
	default:
		return false
	}
}

func (s ShulkerBoxEntity) IsValidBlock(block Block) bool {
	switch block.ID() {
	case "minecraft:shulker_box",
		"minecraft:black_shulker_box",
		"minecraft:blue_shulker_box",
		"minecraft:brown_shulker_box",
		"minecraft:cyan_shulker_box",
		"minecraft:gray_shulker_box",
		"minecraft:green_shulker_box",
		"minecraft:light_blue_shulker_box",
		"minecraft:light_gray_shulker_box",
		"minecraft:lime_shulker_box",
		"minecraft:magenta_shulker_box",
		"minecraft:orange_shulker_box",
		"minecraft:pink_shulker_box",
		"minecraft:purple_shulker_box",
		"minecraft:red_shulker_box",
		"minecraft:white_shulker_box",
		"minecraft:yellow_shulker_box":
		return true
	default:
		return false
	}
}

func (b BedEntity) IsValidBlock(block Block) bool {
	switch block.ID() {
	case "minecraft:red_bed",
		"minecraft:black_bed",
		"minecraft:blue_bed",
		"minecraft:brown_bed",
		"minecraft:cyan_bed",
		"minecraft:gray_bed",
		"minecraft:green_bed",
		"minecraft:light_blue_bed",
		"minecraft:light_gray_bed",
		"minecraft:lime_bed",
		"minecraft:magenta_bed",
		"minecraft:orange_bed",
		"minecraft:pink_bed",
		"minecraft:purple_bed",
		"minecraft:white_bed",
		"minecraft:yellow_bed":
		return true
	default:
		return false
	}
}

func (c ConduitEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:conduit"
}

func (b BarrelEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:barrel"
}

func (s SmokerEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:smoker"
}

func (b BlastFurnaceEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:blast_furnace"
}

func (l LecternEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:lectern"
}

func (b BellEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:bell"
}

func (j JigsawEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:jigsaw"
}

func (c CampfireEntity) IsValidBlock(block Block) bool {
	switch block.ID() {
	case "minecraft:campfire",
		"minecraft:soul_campfire":
		return true
	default:
		return false
	}
}

func (b BeehiveEntity) IsValidBlock(block Block) bool {
	switch block.ID() {
	case "minecraft:bee_nest",
		"minecraft:beehive":
		return true
	default:
		return false
	}
}

func (s SculkSensorEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:sculk_sensor"
}

func (c CalibratedSculkSensorEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:calibrated_sculk_sensor"
}

func (s SculkCatalystEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:sculk_catalyst"
}

func (s SculkShriekerEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:sculk_shrieker"
}

func (c ChiseledBookshelfEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:chiseled_bookshelf"
}

func (b BrushableBlockEntity) IsValidBlock(block Block) bool {
	switch block.ID() {
	case "minecraft:suspicious_sand",
		"minecraft:suspicious_gravel":
		return true
	default:
		return false
	}
}

func (d DecoratedPotEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:decorated_pot"
}

func (c CrafterEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:crafter"
}

func (t TrialSpawnerEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:trial_spawner"
}

func (v VaultEntity) IsValidBlock(block Block) bool {
	return block.ID() == "minecraft:vault"
}
