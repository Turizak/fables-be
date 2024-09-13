DROP TABLE IF EXISTS proficiencies;

CREATE TABLE proficiencies (
    index VARCHAR PRIMARY KEY,
    type VARCHAR,
    name VARCHAR,
    classes JSONB,
    races JSONB,
    url VARCHAR,
    reference JSONB
);

INSERT INTO
    proficiencies (
        index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'light-armor',
        'Armor',
        'Light Armor',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/ruleset/5e/classes/barbarian"},{"index": "cleric", "name": "Cleric", "url": "/api/ruleset/5e/classes/cleric"},{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"},{"index": "monk", "name": "Monk", "url": "/api/ruleset/5e/classes/monk"},{"index": "ranger", "name": "Ranger", "url": "/api/ruleset/5e/classes/ranger"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/light-armor',
        '{"index": "light-armor", "name": "Light Armor", "url": "/api/ruleset/5e/equipment-categories/light-armor"}' :: jsonb
    ),
    (
        'medium-armor',
        'Armor',
        'Medium Armor',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/ruleset/5e/classes/barbarian"},{"index": "cleric", "name": "Cleric", "url": "/api/ruleset/5e/classes/cleric"},{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"},{"index": "ranger", "name": "Ranger", "url": "/api/ruleset/5e/classes/ranger"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/medium-armor',
        '{"index": "medium-armor", "name": "Medium Armor", "url": "/api/ruleset/5e/equipment-categories/medium-armor"}' :: jsonb
    ),
    (
        'heavy-armor',
        'Armor',
        'Heavy Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/heavy-armor',
        '{"index": "heavy-armor", "name": "Heavy Armor", "url": "/api/ruleset/5e/equipment-categories/heavy-armor"}' :: jsonb
    ),
    (
        'all-armor',
        'Armor',
        'All Armor',
        '[{"index": "fighter", "name": "Fighter", "url": "/api/ruleset/5e/classes/fighter"},{"index": "paladin", "name": "Paladin", "url": "/api/ruleset/5e/classes/paladin"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/all-armor',
        '{"index": "armor", "name": "Armor", "url": "/api/ruleset/5e/equipment-categories/armor"}' :: jsonb
    ),
    (
        'padded-armor',
        'Armor',
        'Padded Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/padded-armor',
        '{"index": "padded-armor", "name": "Padded Armor", "url": "/api/ruleset/5e/equipment/padded-armor"}' :: jsonb
    ),
    (
        'leather-armor',
        'Armor',
        'Leather Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/leather-armor',
        '{"index": "leather-armor", "name": "Leather Armor", "url": "/api/ruleset/5e/equipment/leather-armor"}' :: jsonb
    ),
    (
        'studded-leather-armor',
        'Armor',
        'Studded Leather Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/studded-leather-armor',
        '{"index": "studded-leather-armor", "name": "Studded Leather Armor", "url": "/api/ruleset/5e/equipment/studded-leather-armor"}' :: jsonb
    ),
    (
        'hide-armor',
        'Armor',
        'Hide Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/hide-armor',
        '{"index": "hide-armor", "name": "Hide Armor", "url": "/api/ruleset/5e/equipment/hide-armor"}' :: jsonb
    ),
    (
        'chain-shirt',
        'Armor',
        'Chain Shirt',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/chain-shirt',
        '{"index": "chain-shirt", "name": "Chain Shirt", "url": "/api/ruleset/5e/equipment/chain-shirt"}' :: jsonb
    ),
    (
        'scale-mail',
        'Armor',
        'Scale Mail',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/scale-mail',
        '{"index": "scale-mail", "name": "Scale Mail", "url": "/api/ruleset/5e/equipment/scale-mail"}' :: jsonb
    ),
    (
        'breastplate',
        'Armor',
        'Breastplate',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/breastplate',
        '{"index": "breastplate", "name": "Breastplate", "url": "/api/ruleset/5e/equipment/breastplate"}' :: jsonb
    ),
    (
        'half-plate-armor',
        'Armor',
        'Half Plate Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/half-plate-armor',
        '{"index": "half-plate-armor", "name": "Half Plate Armor", "url": "/api/ruleset/5e/equipment/half-plate-armor"}' :: jsonb
    ),
    (
        'ring-mail',
        'Armor',
        'Ring Mail',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/ring-mail',
        '{"index": "ring-mail", "name": "Ring Mail", "url": "/api/ruleset/5e/equipment/ring-mail"}' :: jsonb
    ),
    (
        'chain-mail',
        'Armor',
        'Chain Mail',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/chain-mail',
        '{"index": "chain-mail", "name": "Chain Mail", "url": "/api/ruleset/5e/equipment/chain-mail"}' :: jsonb
    ),
    (
        'splint-armor',
        'Armor',
        'Splint Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/splint-armor',
        '{"index": "splint-armor", "name": "Splint Armor", "url": "/api/ruleset/5e/equipment/splint-armor"}' :: jsonb
    ),
    (
        'plate-armor',
        'Armor',
        'Plate Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/plate-armor',
        '{"index": "plate-armor", "name": "Plate Armor", "url": "/api/ruleset/5e/equipment/plate-armor"}' :: jsonb
    ),
    (
        'shields',
        'Armor',
        'Shields',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/ruleset/5e/classes/barbarian"},{"index": "cleric", "name": "Cleric", "url": "/api/ruleset/5e/classes/cleric"},{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"},{"index": "fighter", "name": "Fighter", "url": "/api/ruleset/5e/classes/fighter"},{"index": "paladin", "name": "Paladin", "url": "/api/ruleset/5e/classes/paladin"},{"index": "ranger", "name": "Ranger", "url": "/api/ruleset/5e/classes/ranger"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/shields',
        '{"index": "shield", "name": "Shield", "url": "/api/ruleset/5e/equipment/shield"}' :: jsonb
    ),
    (
        'simple-weapons',
        'Weapons',
        'Simple Weapons',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/ruleset/5e/classes/barbarian"},{"index": "bard", "name": "Bard", "url": "/api/ruleset/5e/classes/bard"},{"index": "cleric", "name": "Cleric", "url": "/api/ruleset/5e/classes/cleric"},{"index": "fighter", "name": "Fighter", "url": "/api/ruleset/5e/classes/fighter"},{"index": "monk", "name": "Monk", "url": "/api/ruleset/5e/classes/monk"},{"index": "paladin", "name": "Paladin", "url": "/api/ruleset/5e/classes/paladin"},{"index": "ranger", "name": "Ranger", "url": "/api/ruleset/5e/classes/ranger"},{"index": "rogue", "name": "Rogue", "url": "/api/ruleset/5e/classes/rogue"},{"index": "warlock", "name": "Warlock", "url": "/api/ruleset/5e/classes/warlock"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/simple-weapons',
        '{"index": "simple-weapons", "name": "Simple Weapons", "url": "/api/ruleset/5e/equipment-categories/simple-weapons"}' :: jsonb
    ),
    (
        'martial-weapons',
        'Weapons',
        'Martial Weapons',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/ruleset/5e/classes/barbarian"},{"index": "fighter", "name": "Fighter", "url": "/api/ruleset/5e/classes/fighter"},{"index": "paladin", "name": "Paladin", "url": "/api/ruleset/5e/classes/paladin"},{"index": "ranger", "name": "Ranger", "url": "/api/ruleset/5e/classes/ranger"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/martial-weapons',
        '{"index": "martial-weapons", "name": "Martial Weapons", "url": "/api/ruleset/5e/equipment-categories/martial-weapons"}' :: jsonb
    ),
    (
        'clubs',
        'Weapons',
        'Clubs',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/clubs',
        '{"index": "club", "name": "Club", "url": "/api/ruleset/5e/equipment/club"}' :: jsonb
    ),
    (
        'daggers',
        'Weapons',
        'Daggers',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/ruleset/5e/classes/sorcerer"},{"index": "wizard", "name": "Wizard", "url": "/api/ruleset/5e/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/daggers',
        '{"index": "dagger", "name": "Dagger", "url": "/api/ruleset/5e/equipment/dagger"}' :: jsonb
    ),
    (
        'greatclubs',
        'Weapons',
        'Greatclubs',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/greatclubs',
        '{"index": "greatclub", "name": "Greatclub", "url": "/api/ruleset/5e/equipment/greatclub"}' :: jsonb
    ),
    (
        'handaxes',
        'Weapons',
        'Handaxes',
        '[]' :: jsonb,
        '[{"index": "dwarf", "name": "Dwarf", "url": "/api/ruleset/5e/races/dwarf"}]' :: jsonb,
        '/api/ruleset/5e/proficiencies/handaxes',
        '{"index": "handaxe", "name": "Handaxe", "url": "/api/ruleset/5e/equipment/handaxe"}' :: jsonb
    ),
    (
        'javelins',
        'Weapons',
        'Javelins',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/javelins',
        '{"index": "javelin", "name": "Javelin", "url": "/api/ruleset/5e/equipment/javelin"}' :: jsonb
    ),
    (
        'light-hammers',
        'Weapons',
        'Light Hammers',
        '[]' :: jsonb,
        '[{"index": "dwarf", "name": "Dwarf", "url": "/api/ruleset/5e/races/dwarf"}]' :: jsonb,
        '/api/ruleset/5e/proficiencies/light-hammers',
        '{"index": "light-hammer", "name": "Light Hammer", "url": "/api/ruleset/5e/equipment/light-hammer"}' :: jsonb
    ),
    (
        'maces',
        'Weapons',
        'Maces',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/maces',
        '{"index": "mace", "name": "Mace", "url": "/api/ruleset/5e/equipment/mace"}' :: jsonb
    ),
    (
        'quarterstaffs',
        'Weapons',
        'Quarterstaffs',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/ruleset/5e/classes/sorcerer"},{"index": "wizard", "name": "Wizard", "url": "/api/ruleset/5e/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/quarterstaffs',
        '{"index": "quarterstaff", "name": "Quarterstaff", "url": "/api/ruleset/5e/equipment/quarterstaff"}' :: jsonb
    ),
    (
        'sickles',
        'Weapons',
        'Sickles',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/sickles',
        '{"index": "sickle", "name": "Sickle", "url": "/api/ruleset/5e/equipment/sickle"}' :: jsonb
    ),
    (
        'spears',
        'Weapons',
        'Spears',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/spears',
        '{"index": "spear", "name": "Spear", "url": "/api/ruleset/5e/equipment/spear"}' :: jsonb
    ),
    (
        'crossbows-light',
        'Weapons',
        'Crossbows, Light',
        '[{"index": "sorcerer", "name": "Sorcerer", "url": "/api/ruleset/5e/classes/sorcerer"},{"index": "wizard", "name": "Wizard", "url": "/api/ruleset/5e/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/crossbows-light',
        '{"index": "crossbow-light", "name": "Crossbow, Light", "url": "/api/ruleset/5e/equipment/crossbow-light"}' :: jsonb
    ),
    (
        'darts',
        'Weapons',
        'Darts',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/ruleset/5e/classes/sorcerer"},{"index": "wizard", "name": "Wizard", "url": "/api/ruleset/5e/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/darts',
        '{"index": "dart", "name": "Dart", "url": "/api/ruleset/5e/equipment/dart"}' :: jsonb
    ),
    (
        'shortbows',
        'Weapons',
        'Shortbows',
        '[]' :: jsonb,
        '[{"index": "high-elf", "name": "High Elf", "url": "/api/ruleset/5e/subraces/high-elf"}]' :: jsonb,
        '/api/ruleset/5e/proficiencies/shortbows',
        '{"index": "shortbow", "name": "Shortbow", "url": "/api/ruleset/5e/equipment/shortbow"}' :: jsonb
    ),
    (
        'slings',
        'Weapons',
        'Slings',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/ruleset/5e/classes/sorcerer"},{"index": "wizard", "name": "Wizard", "url": "/api/ruleset/5e/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/slings',
        '{"index": "sling", "name": "Sling", "url": "/api/ruleset/5e/equipment/sling"}' :: jsonb
    ),
    (
        'bagpipes',
        'Musical Instruments',
        'Bagpipes',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/bagpipes',
        '{"index": "bagpipes", "name": "Bagpipes", "url": "/api/ruleset/5e/equipment/bagpipes"}' :: jsonb
    ),
    (
        'drum',
        'Musical Instruments',
        'Drum',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/drum',
        '{"index": "drum", "name": "Drum", "url": "/api/ruleset/5e/equipment/drum"}' :: jsonb
    ),
    (
        'dulcimer',
        'Musical Instruments',
        'Dulcimer',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/dulcimer',
        '{"index": "dulcimer", "name": "Dulcimer", "url": "/api/ruleset/5e/equipment/dulcimer"}' :: jsonb
    ),
    (
        'flute',
        'Musical Instruments',
        'Flute',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/flute',
        '{"index": "flute", "name": "Flute", "url": "/api/ruleset/5e/equipment/flute"}' :: jsonb
    ),
    (
        'lute',
        'Musical Instruments',
        'Lute',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/lute',
        '{"index": "lute", "name": "Lute", "url": "/api/ruleset/5e/equipment/lute"}' :: jsonb
    ),
    (
        'lyre',
        'Musical Instruments',
        'Lyre',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/lyre',
        '{"index": "lyre", "name": "Lyre", "url": "/api/ruleset/5e/equipment/lyre"}' :: jsonb
    ),
    (
        'horn',
        'Musical Instruments',
        'Horn',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/horn',
        '{"index": "horn", "name": "Horn", "url": "/api/ruleset/5e/equipment/horn"}' :: jsonb
    ),
    (
        'pan-flute',
        'Musical Instruments',
        'Pan Flute',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/pan-flute',
        '{"index": "pan-flute", "name": "Pan Flute", "url": "/api/ruleset/5e/equipment/pan-flute"}' :: jsonb
    ),
    (
        'shawm',
        'Musical Instruments',
        'Shawm',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/shawm',
        '{"index": "shawm", "name": "Shawm", "url": "/api/ruleset/5e/equipment/shawm"}' :: jsonb
    ),
    (
        'viol',
        'Musical Instruments',
        'Viol',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/viol',
        '{"index": "viol", "name": "Viol", "url": "/api/ruleset/5e/equipment/viol"}' :: jsonb
    ),
    (
        'herbalism-kit',
        'Other',
        'Herbalism Kit',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/herbalism-kit',
        '{"index": "herbalism-kit", "name": "Herbalism Kit", "url": "/api/ruleset/5e/equipment/herbalism-kit"}' :: jsonb
    ),
    (
        'navigators-tools',
        'Other',
        'Navigator’s Tools',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/navigators-tools',
        '{"index": "navigators-tools", "name": "Navigator’s Tools", "url": "/api/ruleset/5e/equipment/navigators-tools"}' :: jsonb
    ),
    (
        'poisoners-kit',
        'Other',
        'Poisoner’s Kit',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/poisoners-kit',
        '{"index": "poisoners-kit", "name": "Poisoner’s Kit", "url": "/api/ruleset/5e/equipment/poisoners-kit"}' :: jsonb
    ),
    (
        'thieves-tools',
        'Other',
        'Thieves’ Tools',
        '[{"index": "rogue", "name": "Rogue", "url": "/api/ruleset/5e/classes/rogue"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/thieves-tools',
        '{"index": "thieves-tools", "name": "Thieves’ Tools", "url": "/api/ruleset/5e/equipment/thieves-tools"}' :: jsonb
    ),
    (
        'land-vehicles',
        'Vehicles',
        'Land Vehicles',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/land-vehicles',
        '{"index": "land-vehicles", "name": "Land Vehicles", "url": "/api/ruleset/5e/equipment-categories/land-vehicles"}' :: jsonb
    ),
    (
        'water-vehicles',
        'Vehicles',
        'Water Vehicles',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/water-vehicles',
        '{"index": "waterborne-vehicles", "name": "Waterborne Vehicles", "url": "/api/ruleset/5e/equipment-categories/waterborne-vehicles"}' :: jsonb
    ),
    (
        'saving-throw-str',
        'Saving Throws',
        'Saving Throw: STR',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/ruleset/5e/classes/barbarian"},{"index": "fighter", "name": "Fighter", "url": "/api/ruleset/5e/classes/fighter"},{"index": "monk", "name": "Monk", "url": "/api/ruleset/5e/classes/monk"},{"index": "ranger", "name": "Ranger", "url": "/api/ruleset/5e/classes/ranger"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/saving-throw-str',
        '{"index": "str", "name": "STR", "url": "/api/ruleset/5e/ability-scores/str"}' :: jsonb
    ),
    (
        'saving-throw-dex',
        'Saving Throws',
        'Saving Throw: DEX',
        '[{"index": "bard", "name": "Bard", "url": "/api/ruleset/5e/classes/bard"},{"index": "monk", "name": "Monk", "url": "/api/ruleset/5e/classes/monk"},{"index": "ranger", "name": "Ranger", "url": "/api/ruleset/5e/classes/ranger"},{"index": "rogue", "name": "Rogue", "url": "/api/ruleset/5e/classes/rogue"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/saving-throw-dex',
        '{"index": "dex", "name": "DEX", "url": "/api/ruleset/5e/ability-scores/dex"}' :: jsonb
    ),
    (
        'saving-throw-con',
        'Saving Throws',
        'Saving Throw: CON',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/ruleset/5e/classes/barbarian"},{"index": "fighter", "name": "Fighter", "url": "/api/ruleset/5e/classes/fighter"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/ruleset/5e/classes/sorcerer"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/saving-throw-con',
        '{"index": "con", "name": "CON", "url": "/api/ruleset/5e/ability-scores/con"}' :: jsonb
    ),
    (
        'saving-throw-int',
        'Saving Throws',
        'Saving Throw: INT',
        '[{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"},{"index": "rogue", "name": "Rogue", "url": "/api/ruleset/5e/classes/rogue"},{"index": "wizard", "name": "Wizard", "url": "/api/ruleset/5e/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/saving-throw-int',
        '{"index": "int", "name": "INT", "url": "/api/ruleset/5e/ability-scores/int"}' :: jsonb
    ),
    (
        'saving-throw-wis',
        'Saving Throws',
        'Saving Throw: WIS',
        '[{"index": "cleric", "name": "Cleric", "url": "/api/ruleset/5e/classes/cleric"},{"index": "druid", "name": "Druid", "url": "/api/ruleset/5e/classes/druid"},{"index": "paladin", "name": "Paladin", "url": "/api/ruleset/5e/classes/paladin"},{"index": "warlock", "name": "Warlock", "url": "/api/ruleset/5e/classes/warlock"},{"index": "wizard", "name": "Wizard", "url": "/api/ruleset/5e/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/saving-throw-wis',
        '{"index": "wis", "name": "WIS", "url": "/api/ruleset/5e/ability-scores/wis"}' :: jsonb
    ),
    (
        'saving-throw-cha',
        'Saving Throws',
        'Saving Throw: CHA',
        '[{"index": "bard", "name": "Bard", "url": "/api/ruleset/5e/classes/bard"},{"index": "cleric", "name": "Cleric", "url": "/api/ruleset/5e/classes/cleric"},{"index": "paladin", "name": "Paladin", "url": "/api/ruleset/5e/classes/paladin"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/ruleset/5e/classes/sorcerer"},{"index": "warlock", "name": "Warlock", "url": "/api/ruleset/5e/classes/warlock"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/saving-throw-cha',
        '{"index": "cha", "name": "CHA", "url": "/api/ruleset/5e/ability-scores/cha"}' :: jsonb
    ),
    (
        'skill-acrobatics',
        'Skills',
        'Skill: Acrobatics',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-acrobatics',
        '{"index": "acrobatics", "name": "Acrobatics", "url": "/api/ruleset/5e/skills/acrobatics"}' :: jsonb
    ),
    (
        'skill-animal-handling',
        'Skills',
        'Skill: Animal Handling',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-animal-handling',
        '{"index": "animal-handling", "name": "Animal Handling", "url": "/api/ruleset/5e/skills/animal-handling"}' :: jsonb
    ),
    (
        'skill-arcana',
        'Skills',
        'Skill: Arcana',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-arcana',
        '{"index": "arcana", "name": "Arcana", "url": "/api/ruleset/5e/skills/arcana"}' :: jsonb
    ),
    (
        'skill-athletics',
        'Skills',
        'Skill: Athletics',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-athletics',
        '{"index": "athletics", "name": "Athletics", "url": "/api/ruleset/5e/skills/athletics"}' :: jsonb
    ),
    (
        'skill-deception',
        'Skills',
        'Skill: Deception',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-deception',
        '{"index": "deception", "name": "Deception", "url": "/api/ruleset/5e/skills/deception"}' :: jsonb
    ),
    (
        'skill-history',
        'Skills',
        'Skill: History',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-history',
        '{"index": "history", "name": "History", "url": "/api/ruleset/5e/skills/history"}' :: jsonb
    ),
    (
        'skill-insight',
        'Skills',
        'Skill: Insight',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-insight',
        '{"index": "insight", "name": "Insight", "url": "/api/ruleset/5e/skills/insight"}' :: jsonb
    ),
    (
        'skill-intimidation',
        'Skills',
        'Skill: Intimidation',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-intimidation',
        '{"index": "intimidation", "name": "Intimidation", "url": "/api/ruleset/5e/skills/intimidation"}' :: jsonb
    ),
    (
        'skill-investigation',
        'Skills',
        'Skill: Investigation',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-investigation',
        '{"index": "investigation", "name": "Investigation", "url": "/api/ruleset/5e/skills/investigation"}' :: jsonb
    ),
    (
        'skill-medicine',
        'Skills',
        'Skill: Medicine',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-medicine',
        '{"index": "medicine", "name": "Medicine", "url": "/api/ruleset/5e/skills/medicine"}' :: jsonb
    ),
    (
        'skill-nature',
        'Skills',
        'Skill: Nature',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-nature',
        '{"index": "nature", "name": "Nature", "url": "/api/ruleset/5e/skills/nature"}' :: jsonb
    ),
    (
        'skill-perception',
        'Skills',
        'Skill: Perception',
        '[]' :: jsonb,
        '[{"index": "elf", "name": "Elf", "url": "/api/ruleset/5e/races/elf"}]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-perception',
        '{"index": "perception", "name": "Perception", "url": "/api/ruleset/5e/skills/perception"}' :: jsonb
    ),
    (
        'skill-performance',
        'Skills',
        'Skill: Performance',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-performance',
        '{"index": "performance", "name": "Performance", "url": "/api/ruleset/5e/skills/performance"}' :: jsonb
    ),
    (
        'skill-persuasion',
        'Skills',
        'Skill: Persuasion',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-persuasion',
        '{"index": "persuasion", "name": "Persuasion", "url": "/api/ruleset/5e/skills/persuasion"}' :: jsonb
    ),
    (
        'skill-religion',
        'Skills',
        'Skill: Religion',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-religion',
        '{"index": "religion", "name": "Religion", "url": "/api/ruleset/5e/skills/religion"}' :: jsonb
    ),
    (
        'skill-sleight-of-hand',
        'Skills',
        'Skill: Sleight of Hand',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-sleight-of-hand',
        '{"index": "sleight-of-hand", "name": "Sleight of Hand", "url": "/api/ruleset/5e/skills/sleight-of-hand"}' :: jsonb
    ),
    (
        'skill-stealth',
        'Skills',
        'Skill: Stealth',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-stealth',
        '{"index": "stealth", "name": "Stealth", "url": "/api/ruleset/5e/skills/stealth"}' :: jsonb
    ),
    (
        'skill-survival',
        'Skills',
        'Skill: Survival',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/proficiencies/skill-survival',
        '{"index": "survival", "name": "Survival", "url": "/api/ruleset/5e/skills/survival"}' :: jsonb
    );