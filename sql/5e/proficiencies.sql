DROP TABLE IF EXISTS proficiencies;

CREATE TABLE proficiencies (
    proficiency_index VARCHAR PRIMARY KEY,
    type VARCHAR,
    name VARCHAR,
    classes JSONB,
    races JSONB,
    url VARCHAR,
    reference JSONB
);

INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'medium-armor',
        'Armor',
        'Medium Armor',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/classes/barbarian"},{"index": "cleric", "name": "Cleric", "url": "/api/classes/cleric"},{"index": "druid", "name": "Druid", "url": "/api/classes/druid"},{"index": "ranger", "name": "Ranger", "url": "/api/classes/ranger"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/medium-armor',
        '{"index": "medium-armor", "name": "Medium Armor", "url": "/api/equipment-categories/medium-armor"}' :: jsonb
    );

-- Insert for 'heavy-armor'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'heavy-armor',
        'Armor',
        'Heavy Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/heavy-armor',
        '{"index": "heavy-armor", "name": "Heavy Armor", "url": "/api/equipment-categories/heavy-armor"}' :: jsonb
    );

-- Insert for 'all-armor'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'all-armor',
        'Armor',
        'All Armor',
        '[{"index": "fighter", "name": "Fighter", "url": "/api/classes/fighter"},{"index": "paladin", "name": "Paladin", "url": "/api/classes/paladin"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/all-armor',
        '{"index": "armor", "name": "Armor", "url": "/api/equipment-categories/armor"}' :: jsonb
    );

-- Insert for 'padded-armor'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'padded-armor',
        'Armor',
        'Padded Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/padded-armor',
        '{"index": "padded-armor", "name": "Padded Armor", "url": "/api/equipment/padded-armor"}' :: jsonb
    );

-- Insert for 'leather-armor'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'leather-armor',
        'Armor',
        'Leather Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/leather-armor',
        '{"index": "leather-armor", "name": "Leather Armor", "url": "/api/equipment/leather-armor"}' :: jsonb
    );

-- Insert for 'studded-leather-armor'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'studded-leather-armor',
        'Armor',
        'Studded Leather Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/studded-leather-armor',
        '{"index": "studded-leather-armor", "name": "Studded Leather Armor", "url": "/api/equipment/studded-leather-armor"}' :: jsonb
    );

-- Insert for 'hide-armor'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'hide-armor',
        'Armor',
        'Hide Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/hide-armor',
        '{"index": "hide-armor", "name": "Hide Armor", "url": "/api/equipment/hide-armor"}' :: jsonb
    );

-- Insert for 'chain-shirt'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'chain-shirt',
        'Armor',
        'Chain Shirt',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/chain-shirt',
        '{"index": "chain-shirt", "name": "Chain Shirt", "url": "/api/equipment/chain-shirt"}' :: jsonb
    );

-- Insert for 'scale-mail'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'scale-mail',
        'Armor',
        'Scale Mail',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/scale-mail',
        '{"index": "scale-mail", "name": "Scale Mail", "url": "/api/equipment/scale-mail"}' :: jsonb
    );

-- Insert for 'breastplate'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'breastplate',
        'Armor',
        'Breastplate',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/breastplate',
        '{"index": "breastplate", "name": "Breastplate", "url": "/api/equipment/breastplate"}' :: jsonb
    );

-- Insert for 'half-plate-armor'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'half-plate-armor',
        'Armor',
        'Half Plate Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/half-plate-armor',
        '{"index": "half-plate-armor", "name": "Half Plate Armor", "url": "/api/equipment/half-plate-armor"}' :: jsonb
    );

-- Insert for 'ring-mail'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'ring-mail',
        'Armor',
        'Ring Mail',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/ring-mail',
        '{"index": "ring-mail", "name": "Ring Mail", "url": "/api/equipment/ring-mail"}' :: jsonb
    );

-- Insert for 'chain-mail'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'chain-mail',
        'Armor',
        'Chain Mail',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/chain-mail',
        '{"index": "chain-mail", "name": "Chain Mail", "url": "/api/equipment/chain-mail"}' :: jsonb
    );

-- Insert for 'splint-armor'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'splint-armor',
        'Armor',
        'Splint Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/splint-armor',
        '{"index": "splint-armor", "name": "Splint Armor", "url": "/api/equipment/splint-armor"}' :: jsonb
    );

-- Insert for 'plate-armor'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'plate-armor',
        'Armor',
        'Plate Armor',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/plate-armor',
        '{"index": "plate-armor", "name": "Plate Armor", "url": "/api/equipment/plate-armor"}' :: jsonb
    );

-- Insert for 'shields'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'shields',
        'Armor',
        'Shields',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/classes/barbarian"},{"index": "cleric", "name": "Cleric", "url": "/api/classes/cleric"},{"index": "druid", "name": "Druid", "url": "/api/classes/druid"},{"index": "fighter", "name": "Fighter", "url": "/api/classes/fighter"},{"index": "paladin", "name": "Paladin", "url": "/api/classes/paladin"},{"index": "ranger", "name": "Ranger", "url": "/api/classes/ranger"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/shields',
        '{"index": "shield", "name": "Shield", "url": "/api/equipment/shield"}' :: jsonb
    );

-- Insert for 'simple-weapons'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'simple-weapons',
        'Weapons',
        'Simple Weapons',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/classes/barbarian"},{"index": "bard", "name": "Bard", "url": "/api/classes/bard"},{"index": "cleric", "name": "Cleric", "url": "/api/classes/cleric"},{"index": "fighter", "name": "Fighter", "url": "/api/classes/fighter"},{"index": "monk", "name": "Monk", "url": "/api/classes/monk"},{"index": "paladin", "name": "Paladin", "url": "/api/classes/paladin"},{"index": "ranger", "name": "Ranger", "url": "/api/classes/ranger"},{"index": "rogue", "name": "Rogue", "url": "/api/classes/rogue"},{"index": "warlock", "name": "Warlock", "url": "/api/classes/warlock"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/simple-weapons',
        '{"index": "simple-weapons", "name": "Simple Weapons", "url": "/api/equipment-categories/simple-weapons"}' :: jsonb
    );

-- Insert for 'martial-weapons'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'martial-weapons',
        'Weapons',
        'Martial Weapons',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/classes/barbarian"},{"index": "fighter", "name": "Fighter", "url": "/api/classes/fighter"},{"index": "paladin", "name": "Paladin", "url": "/api/classes/paladin"},{"index": "ranger", "name": "Ranger", "url": "/api/classes/ranger"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/martial-weapons',
        '{"index": "martial-weapons", "name": "Martial Weapons", "url": "/api/equipment-categories/martial-weapons"}' :: jsonb
    );

-- Insert for 'clubs'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'clubs',
        'Weapons',
        'Clubs',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/clubs',
        '{"index": "club", "name": "Club", "url": "/api/equipment/club"}' :: jsonb
    );

-- Insert for 'daggers'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'daggers',
        'Weapons',
        'Daggers',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/classes/sorcerer"},{"index": "wizard", "name": "Wizard", "url": "/api/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/daggers',
        '{"index": "dagger", "name": "Dagger", "url": "/api/equipment/dagger"}' :: jsonb
    );

-- Insert for 'greatclubs'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'greatclubs',
        'Weapons',
        'Greatclubs',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/greatclubs',
        '{"index": "greatclub", "name": "Greatclub", "url": "/api/equipment/greatclub"}' :: jsonb
    );

-- Insert for 'handaxes'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'handaxes',
        'Weapons',
        'Handaxes',
        '[]' :: jsonb,
        '[{"index": "dwarf", "name": "Dwarf", "url": "/api/races/dwarf"}]' :: jsonb,
        '/api/proficiencies/handaxes',
        '{"index": "handaxe", "name": "Handaxe", "url": "/api/equipment/handaxe"}' :: jsonb
    );

-- Insert for 'javelins'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'javelins',
        'Weapons',
        'Javelins',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/javelins',
        '{"index": "javelin", "name": "Javelin", "url": "/api/equipment/javelin"}' :: jsonb
    );

-- Insert for 'light-hammers'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'light-hammers',
        'Weapons',
        'Light Hammers',
        '[]' :: jsonb,
        '[{"index": "dwarf", "name": "Dwarf", "url": "/api/races/dwarf"}]' :: jsonb,
        '/api/proficiencies/light-hammers',
        '{"index": "light-hammer", "name": "Light Hammer", "url": "/api/equipment/light-hammer"}' :: jsonb
    );

-- Insert for 'maces'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'maces',
        'Weapons',
        'Maces',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/maces',
        '{"index": "mace", "name": "Mace", "url": "/api/equipment/mace"}' :: jsonb
    );

-- Insert for 'quarterstaffs'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'quarterstaffs',
        'Weapons',
        'Quarterstaffs',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/classes/sorcerer"},{"index": "wizard", "name": "Wizard", "url": "/api/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/quarterstaffs',
        '{"index": "quarterstaff", "name": "Quarterstaff", "url": "/api/equipment/quarterstaff"}' :: jsonb
    );

-- Insert for 'sickles'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'sickles',
        'Weapons',
        'Sickles',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/sickles',
        '{"index": "sickle", "name": "Sickle", "url": "/api/equipment/sickle"}' :: jsonb
    );

-- Insert for 'spears'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'spears',
        'Weapons',
        'Spears',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/spears',
        '{"index": "spear", "name": "Spear", "url": "/api/equipment/spear"}' :: jsonb
    );

-- Insert for 'crossbows-light'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'crossbows-light',
        'Weapons',
        'Crossbows, Light',
        '[{"index": "sorcerer", "name": "Sorcerer", "url": "/api/classes/sorcerer"},{"index": "wizard", "name": "Wizard", "url": "/api/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/crossbows-light',
        '{"index": "crossbow-light", "name": "Crossbow, Light", "url": "/api/equipment/crossbow-light"}' :: jsonb
    );

-- Insert for 'darts'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'darts',
        'Weapons',
        'Darts',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/classes/sorcerer"},{"index": "wizard", "name": "Wizard", "url": "/api/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/darts',
        '{"index": "dart", "name": "Dart", "url": "/api/equipment/dart"}' :: jsonb
    );

-- Insert for 'shortbows'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'shortbows',
        'Weapons',
        'Shortbows',
        '[]' :: jsonb,
        '[{"index": "high-elf", "name": "High Elf", "url": "/api/subraces/high-elf"}]' :: jsonb,
        '/api/proficiencies/shortbows',
        '{"index": "shortbow", "name": "Shortbow", "url": "/api/equipment/shortbow"}' :: jsonb
    );

-- Insert for 'slings'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'slings',
        'Weapons',
        'Slings',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/classes/sorcerer"},{"index": "wizard", "name": "Wizard", "url": "/api/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/slings',
        '{"index": "sling", "name": "Sling", "url": "/api/equipment/sling"}' :: jsonb
    );

-- Insert for 'bagpipes'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'bagpipes',
        'Musical Instruments',
        'Bagpipes',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/bagpipes',
        '{"index": "bagpipes", "name": "Bagpipes", "url": "/api/equipment/bagpipes"}' :: jsonb
    );

-- Insert for 'drum'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'drum',
        'Musical Instruments',
        'Drum',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/drum',
        '{"index": "drum", "name": "Drum", "url": "/api/equipment/drum"}' :: jsonb
    );

-- Insert for 'dulcimer'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'dulcimer',
        'Musical Instruments',
        'Dulcimer',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/dulcimer',
        '{"index": "dulcimer", "name": "Dulcimer", "url": "/api/equipment/dulcimer"}' :: jsonb
    );

-- Insert for 'flute'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'flute',
        'Musical Instruments',
        'Flute',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/flute',
        '{"index": "flute", "name": "Flute", "url": "/api/equipment/flute"}' :: jsonb
    );

-- Insert for 'lute'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'lute',
        'Musical Instruments',
        'Lute',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/lute',
        '{"index": "lute", "name": "Lute", "url": "/api/equipment/lute"}' :: jsonb
    );

-- Insert for 'lyre'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'lyre',
        'Musical Instruments',
        'Lyre',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/lyre',
        '{"index": "lyre", "name": "Lyre", "url": "/api/equipment/lyre"}' :: jsonb
    );

-- Insert for 'horn'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'horn',
        'Musical Instruments',
        'Horn',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/horn',
        '{"index": "horn", "name": "Horn", "url": "/api/equipment/horn"}' :: jsonb
    );

-- Insert for 'pan-flute'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'pan-flute',
        'Musical Instruments',
        'Pan Flute',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/pan-flute',
        '{"index": "pan-flute", "name": "Pan Flute", "url": "/api/equipment/pan-flute"}' :: jsonb
    );

-- Insert for 'shawm'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'shawm',
        'Musical Instruments',
        'Shawm',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/shawm',
        '{"index": "shawm", "name": "Shawm", "url": "/api/equipment/shawm"}' :: jsonb
    );

-- Insert for 'viol'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'viol',
        'Musical Instruments',
        'Viol',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/viol',
        '{"index": "viol", "name": "Viol", "url": "/api/equipment/viol"}' :: jsonb
    );

-- Insert for 'herbalism-kit'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'herbalism-kit',
        'Other',
        'Herbalism Kit',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/herbalism-kit',
        '{"index": "herbalism-kit", "name": "Herbalism Kit", "url": "/api/equipment/herbalism-kit"}' :: jsonb
    );

-- Insert for 'navigators-tools'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'navigators-tools',
        'Other',
        'Navigator’s Tools',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/navigators-tools',
        '{"index": "navigators-tools", "name": "Navigator’s Tools", "url": "/api/equipment/navigators-tools"}' :: jsonb
    );

-- Insert for 'poisoners-kit'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'poisoners-kit',
        'Other',
        'Poisoner’s Kit',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/poisoners-kit',
        '{"index": "poisoners-kit", "name": "Poisoner’s Kit", "url": "/api/equipment/poisoners-kit"}' :: jsonb
    );

-- Insert for 'thieves-tools'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'thieves-tools',
        'Other',
        'Thieves’ Tools',
        '[{"index": "rogue", "name": "Rogue", "url": "/api/classes/rogue"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/thieves-tools',
        '{"index": "thieves-tools", "name": "Thieves’ Tools", "url": "/api/equipment/thieves-tools"}' :: jsonb
    );

-- Insert for 'land-vehicles'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'land-vehicles',
        'Vehicles',
        'Land Vehicles',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/land-vehicles',
        '{"index": "land-vehicles", "name": "Land Vehicles", "url": "/api/equipment-categories/land-vehicles"}' :: jsonb
    );

-- Insert for 'water-vehicles'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'water-vehicles',
        'Vehicles',
        'Water Vehicles',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/water-vehicles',
        '{"index": "waterborne-vehicles", "name": "Waterborne Vehicles", "url": "/api/equipment-categories/waterborne-vehicles"}' :: jsonb
    );

-- Insert for 'saving-throw-str'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'saving-throw-str',
        'Saving Throws',
        'Saving Throw: STR',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/classes/barbarian"},{"index": "fighter", "name": "Fighter", "url": "/api/classes/fighter"},{"index": "monk", "name": "Monk", "url": "/api/classes/monk"},{"index": "ranger", "name": "Ranger", "url": "/api/classes/ranger"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/saving-throw-str',
        '{"index": "str", "name": "STR", "url": "/api/ability-scores/str"}' :: jsonb
    );

-- Insert for 'saving-throw-dex'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'saving-throw-dex',
        'Saving Throws',
        'Saving Throw: DEX',
        '[{"index": "bard", "name": "Bard", "url": "/api/classes/bard"},{"index": "monk", "name": "Monk", "url": "/api/classes/monk"},{"index": "ranger", "name": "Ranger", "url": "/api/classes/ranger"},{"index": "rogue", "name": "Rogue", "url": "/api/classes/rogue"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/saving-throw-dex',
        '{"index": "dex", "name": "DEX", "url": "/api/ability-scores/dex"}' :: jsonb
    );

-- Insert for 'saving-throw-con'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'saving-throw-con',
        'Saving Throws',
        'Saving Throw: CON',
        '[{"index": "barbarian", "name": "Barbarian", "url": "/api/classes/barbarian"},{"index": "fighter", "name": "Fighter", "url": "/api/classes/fighter"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/classes/sorcerer"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/saving-throw-con',
        '{"index": "con", "name": "CON", "url": "/api/ability-scores/con"}' :: jsonb
    );

-- Insert for 'saving-throw-int'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'saving-throw-int',
        'Saving Throws',
        'Saving Throw: INT',
        '[{"index": "druid", "name": "Druid", "url": "/api/classes/druid"},{"index": "rogue", "name": "Rogue", "url": "/api/classes/rogue"},{"index": "wizard", "name": "Wizard", "url": "/api/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/saving-throw-int',
        '{"index": "int", "name": "INT", "url": "/api/ability-scores/int"}' :: jsonb
    );

-- Insert for 'saving-throw-wis'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'saving-throw-wis',
        'Saving Throws',
        'Saving Throw: WIS',
        '[{"index": "cleric", "name": "Cleric", "url": "/api/classes/cleric"},{"index": "druid", "name": "Druid", "url": "/api/classes/druid"},{"index": "paladin", "name": "Paladin", "url": "/api/classes/paladin"},{"index": "warlock", "name": "Warlock", "url": "/api/classes/warlock"},{"index": "wizard", "name": "Wizard", "url": "/api/classes/wizard"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/saving-throw-wis',
        '{"index": "wis", "name": "WIS", "url": "/api/ability-scores/wis"}' :: jsonb
    );

-- Insert for 'saving-throw-cha'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'saving-throw-cha',
        'Saving Throws',
        'Saving Throw: CHA',
        '[{"index": "bard", "name": "Bard", "url": "/api/classes/bard"},{"index": "cleric", "name": "Cleric", "url": "/api/classes/cleric"},{"index": "paladin", "name": "Paladin", "url": "/api/classes/paladin"},{"index": "sorcerer", "name": "Sorcerer", "url": "/api/classes/sorcerer"},{"index": "warlock", "name": "Warlock", "url": "/api/classes/warlock"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/saving-throw-cha',
        '{"index": "cha", "name": "CHA", "url": "/api/ability-scores/cha"}' :: jsonb
    );

-- Insert for 'skill-acrobatics'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-acrobatics',
        'Skills',
        'Skill: Acrobatics',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-acrobatics',
        '{"index": "acrobatics", "name": "Acrobatics", "url": "/api/skills/acrobatics"}' :: jsonb
    );

-- Insert for 'skill-animal-handling'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-animal-handling',
        'Skills',
        'Skill: Animal Handling',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-animal-handling',
        '{"index": "animal-handling", "name": "Animal Handling", "url": "/api/skills/animal-handling"}' :: jsonb
    );

-- Insert for 'skill-arcana'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-arcana',
        'Skills',
        'Skill: Arcana',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-arcana',
        '{"index": "arcana", "name": "Arcana", "url": "/api/skills/arcana"}' :: jsonb
    );

-- Insert for 'skill-athletics'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-athletics',
        'Skills',
        'Skill: Athletics',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-athletics',
        '{"index": "athletics", "name": "Athletics", "url": "/api/skills/athletics"}' :: jsonb
    );

-- Insert for 'skill-deception'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-deception',
        'Skills',
        'Skill: Deception',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-deception',
        '{"index": "deception", "name": "Deception", "url": "/api/skills/deception"}' :: jsonb
    );

-- Insert for 'skill-history'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-history',
        'Skills',
        'Skill: History',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-history',
        '{"index": "history", "name": "History", "url": "/api/skills/history"}' :: jsonb
    );

-- Insert for 'skill-insight'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-insight',
        'Skills',
        'Skill: Insight',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-insight',
        '{"index": "insight", "name": "Insight", "url": "/api/skills/insight"}' :: jsonb
    );

-- Insert for 'skill-intimidation'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-intimidation',
        'Skills',
        'Skill: Intimidation',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-intimidation',
        '{"index": "intimidation", "name": "Intimidation", "url": "/api/skills/intimidation"}' :: jsonb
    );

-- Insert for 'skill-investigation'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-investigation',
        'Skills',
        'Skill: Investigation',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-investigation',
        '{"index": "investigation", "name": "Investigation", "url": "/api/skills/investigation"}' :: jsonb
    );

-- Insert for 'skill-medicine'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-medicine',
        'Skills',
        'Skill: Medicine',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-medicine',
        '{"index": "medicine", "name": "Medicine", "url": "/api/skills/medicine"}' :: jsonb
    );

-- Insert for 'skill-nature'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-nature',
        'Skills',
        'Skill: Nature',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-nature',
        '{"index": "nature", "name": "Nature", "url": "/api/skills/nature"}' :: jsonb
    );

-- Insert for 'skill-perception'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-perception',
        'Skills',
        'Skill: Perception',
        '[]' :: jsonb,
        '[{"index": "elf", "name": "Elf", "url": "/api/races/elf"}]' :: jsonb,
        '/api/proficiencies/skill-perception',
        '{"index": "perception", "name": "Perception", "url": "/api/skills/perception"}' :: jsonb
    );

-- Insert for 'skill-performance'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-performance',
        'Skills',
        'Skill: Performance',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-performance',
        '{"index": "performance", "name": "Performance", "url": "/api/skills/performance"}' :: jsonb
    );

-- Insert for 'skill-persuasion'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-persuasion',
        'Skills',
        'Skill: Persuasion',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-persuasion',
        '{"index": "persuasion", "name": "Persuasion", "url": "/api/skills/persuasion"}' :: jsonb
    );

-- Insert for 'skill-religion'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-religion',
        'Skills',
        'Skill: Religion',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-religion',
        '{"index": "religion", "name": "Religion", "url": "/api/skills/religion"}' :: jsonb
    );

-- Insert for 'skill-sleight-of-hand'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-sleight-of-hand',
        'Skills',
        'Skill: Sleight of Hand',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-sleight-of-hand',
        '{"index": "sleight-of-hand", "name": "Sleight of Hand", "url": "/api/skills/sleight-of-hand"}' :: jsonb
    );

-- Insert for 'skill-stealth'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-stealth',
        'Skills',
        'Skill: Stealth',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-stealth',
        '{"index": "stealth", "name": "Stealth", "url": "/api/skills/stealth"}' :: jsonb
    );

-- Insert for 'skill-survival'
INSERT INTO
    proficiencies (
        proficiency_index,
        type,
        name,
        classes,
        races,
        url,
        reference
    )
VALUES
    (
        'skill-survival',
        'Skills',
        'Skill: Survival',
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/proficiencies/skill-survival',
        '{"index": "survival", "name": "Survival", "url": "/api/skills/survival"}' :: jsonb
    );