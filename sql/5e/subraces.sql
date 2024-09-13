DROP TABLE IF EXISTS subraces;

CREATE TABLE subraces (
    index VARCHAR PRIMARY KEY,
    name VARCHAR,
    race JSONB,
    description TEXT,
    ability_bonuses JSONB,
    starting_proficiencies JSONB,
    languages JSONB,
    language_options JSONB,
    racial_traits JSONB,
    url VARCHAR
);

INSERT INTO subraces (index, name, race, description, ability_bonuses, starting_proficiencies, languages, language_options, racial_traits, url)
VALUES
(
    'hill-dwarf',
    'Hill Dwarf',
    '{"index": "dwarf", "name": "Dwarf", "url": "/api/ruleset/5e/races/dwarf"}',
    'As a hill dwarf, you have keen senses, deep intuition, and remarkable resilience.',
    '[{"ability_score": {"index": "wis", "name": "WIS", "url": "/api/ruleset/5e/ability-scores/wis"}, "bonus": 1}]',
    '[]',
    '[]',
    NULL,
    '[{"index": "dwarven-toughness", "name": "Dwarven Toughness", "url": "/api/ruleset/5e/traits/dwarven-toughness"}]',
    '/api/ruleset/5e/subraces/hill-dwarf'
),
(
    'high-elf',
    'High Elf',
    '{"index": "elf", "name": "Elf", "url": "/api/ruleset/5e/races/elf"}',
    'As a high elf, you have a keen mind and a mastery of at least the basics of magic...',
    '[{"ability_score": {"index": "int", "name": "INT", "url": "/api/ruleset/5e/ability-scores/int"}, "bonus": 1}]',
    '[{"index": "longswords", "name": "Longswords", "url": "/api/ruleset/5e/proficiencies/longswords"}, {"index": "shortswords", "name": "Shortswords", "url": "/api/ruleset/5e/proficiencies/shortswords"}, {"index": "shortbows", "name": "Shortbows", "url": "/api/ruleset/5e/proficiencies/shortbows"}, {"index": "longbows", "name": "Longbows", "url": "/api/ruleset/5e/proficiencies/longbows"}]',
    '[]',
    '{"choose": 1, "from": {"option_set_type": "options_array", "options": [{"option_type": "reference", "item": {"index": "dwarvish", "name": "Dwarvish", "url": "/api/ruleset/5e/languages/dwarvish"}}, {"option_type": "reference", "item": {"index": "giant", "name": "Giant", "url": "/api/ruleset/5e/languages/giant"}}, {"option_type": "reference", "item": {"index": "gnomish", "name": "Gnomish", "url": "/api/ruleset/5e/languages/gnomish"}}, {"option_type": "reference", "item": {"index": "goblin", "name": "Goblin", "url": "/api/ruleset/5e/languages/goblin"}}, {"option_type": "reference", "item": {"index": "halfling", "name": "Halfling", "url": "/api/ruleset/5e/languages/halfling"}}, {"option_type": "reference", "item": {"index": "orc", "name": "Orc", "url": "/api/ruleset/5e/languages/orc"}}, {"option_type": "reference", "item": {"index": "abyssal", "name": "Abyssal", "url": "/api/ruleset/5e/languages/abyssal"}}, {"option_type": "reference", "item": {"index": "celestial", "name": "Celestial", "url": "/api/ruleset/5e/languages/celestial"}}, {"option_type": "reference", "item": {"index": "draconic", "name": "Draconic", "url": "/api/ruleset/5e/languages/draconic"}}, {"option_type": "reference", "item": {"index": "deep-speech", "name": "Deep Speech", "url": "/api/ruleset/5e/languages/deep-speech"}}, {"option_type": "reference", "item": {"index": "infernal", "name": "Infernal", "url": "/api/ruleset/5e/languages/infernal"}}, {"option_type": "reference", "item": {"index": "primordial", "name": "Primordial", "url": "/api/ruleset/5e/languages/primordial"}}, {"option_type": "reference", "item": {"index": "sylvan", "name": "Sylvan", "url": "/api/ruleset/5e/languages/sylvan"}}, {"option_type": "reference", "item": {"index": "undercommon", "name": "Undercommon", "url": "/api/ruleset/5e/languages/undercommon"}}]}, "type": "language"}',
    '[{"index": "elf-weapon-training", "name": "Elf Weapon Training", "url": "/api/ruleset/5e/traits/elf-weapon-training"}, {"index": "high-elf-cantrip", "name": "High Elf Cantrip", "url": "/api/ruleset/5e/traits/high-elf-cantrip"}, {"index": "extra-language", "name": "Extra Language", "url": "/api/ruleset/5e/traits/extra-language"}]',
    '/api/ruleset/5e/subraces/high-elf'
),
(
    'lightfoot-halfling',
    'Lightfoot Halfling',
    '{"index": "halfling", "name": "Halfling", "url": "/api/ruleset/5e/races/halfling"}',
    'As a lightfoot halfling, you can easily hide from notice, even using other people as cover...',
    '[{"ability_score": {"index": "cha", "name": "CHA", "url": "/api/ruleset/5e/ability-scores/cha"}, "bonus": 1}]',
    '[]',
    '[]',
    NULL,
    '[{"index": "naturally-stealthy", "name": "Naturally Stealthy", "url": "/api/ruleset/5e/traits/naturally-stealthy"}]',
    '/api/ruleset/5e/subraces/lightfoot-halfling'
),
(
    'rock-gnome',
    'Rock Gnome',
    '{"index": "gnome", "name": "Gnome", "url": "/api/ruleset/5e/races/gnome"}',
    'As a rock gnome, you have a natural inventiveness and hardiness beyond that of other gnomes.',
    '[{"ability_score": {"index": "con", "name": "CON", "url": "/api/ruleset/5e/ability-scores/con"}, "bonus": 1}]',
    '[{"index": "tinkers-tools", "name": "Tinkers Tools", "url": "/api/ruleset/5e/proficiencies/tinkers-tools"}]',
    '[]',
    NULL,
    '[{"index": "artificers-lore", "name": "Artificer’s Lore", "url": "/api/ruleset/5e/traits/artificers-lore"}, {"index": "tinker", "name": "Tinker", "url": "/api/ruleset/5e/traits/tinker"}]',
    '/api/ruleset/5e/subraces/rock-gnome'
);
