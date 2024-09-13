DROP TABLE IF EXISTS races;

CREATE TABLE races (
    index VARCHAR PRIMARY KEY,
    name VARCHAR,
    speed INT,
    ability_bonuses JSONB,
    ability_bonus_options JSONB,
    alignment TEXT,
    age TEXT,
    size VARCHAR,
    size_description TEXT,
    starting_proficiencies JSONB,
    starting_proficiency_options JSONB,
    languages JSONB,
    language_desc TEXT,
    language_options JSONB,
    traits JSONB,
    subraces JSONB,
    url VARCHAR
);

INSERT INTO
    races (
        index,
        name,
        speed,
        ability_bonuses,
        ability_bonus_options,
        alignment,
        age,
        size,
        size_description,
        starting_proficiencies,
        starting_proficiency_options,
        languages,
        language_desc,
        language_options,
        traits,
        subraces,
        url
    )
VALUES
    (
        'dwarf',
        'Dwarf',
        25,
        '[{"ability_score": {"index": "con", "name": "CON", "url": "/api/ruleset/5e/ability-scores/con"}, "bonus": 2}]' :: jsonb,
        NULL,
        'Most dwarves are lawful, believing firmly in the benefits of a well-ordered society. They tend toward good as well, with a strong sense of fair play and a belief that everyone deserves to share in the benefits of a just order.',
        'Dwarves mature at the same rate as humans, but they''re considered young until they reach the age of 50. On average, they live about 350 years.',
        'Medium',
        'Dwarves stand between 4 and 5 feet tall and average about 150 pounds. Your size is Medium.',
        '[{"index": "battleaxes", "name": "Battleaxes", "url": "/api/ruleset/5e/proficiencies/battleaxes"}, {"index": "handaxes", "name": "Handaxes", "url": "/api/ruleset/5e/proficiencies/handaxes"}, {"index": "light-hammers", "name": "Light hammers", "url": "/api/ruleset/5e/proficiencies/light-hammers"}, {"index": "warhammers", "name": "Warhammers", "url": "/api/ruleset/5e/proficiencies/warhammers"}]' :: jsonb,
        '{"desc": "You gain proficiency with the artisan’s tools of your choice: smith’s tools, brewer’s supplies, or mason’s tools.", "choose": 1, "type": "proficiencies", "from": {"option_set_type": "options_array", "options": [{"option_type": "reference", "item": {"index": "smiths-tools", "name": "Smith’s Tools", "url": "/api/ruleset/5e/proficiencies/smiths-tools"}}, {"option_type": "reference", "item": {"index": "brewers-supplies", "name": "Brewer’s Supplies", "url": "/api/ruleset/5e/proficiencies/brewers-supplies"}}, {"option_type": "reference", "item": {"index": "masons-tools", "name": "Mason’s Tools", "url": "/api/ruleset/5e/proficiencies/masons-tools"}}]}}' :: jsonb,
        '[{"index": "common", "name": "Common", "url": "/api/ruleset/5e/languages/common"}, {"index": "dwarvish", "name": "Dwarvish", "url": "/api/ruleset/5e/languages/dwarvish"}]' :: jsonb,
        'You can speak, read, and write Common and Dwarvish. Dwarvish is full of hard consonants and guttural sounds, and those characteristics spill over into whatever other language a dwarf might speak.',
        NULL,
        '[{"index": "darkvision", "name": "Darkvision", "url": "/api/ruleset/5e/traits/darkvision"}, {"index": "dwarven-resilience", "name": "Dwarven Resilience", "url": "/api/ruleset/5e/traits/dwarven-resilience"}, {"index": "stonecunning", "name": "Stonecunning", "url": "/api/ruleset/5e/traits/stonecunning"}, {"index": "dwarven-combat-training", "name": "Dwarven Combat Training", "url": "/api/ruleset/5e/traits/dwarven-combat-training"}, {"index": "tool-proficiency", "name": "Tool Proficiency", "url": "/api/ruleset/5e/traits/tool-proficiency"}]' :: jsonb,
        '[{"index": "hill-dwarf", "name": "Hill Dwarf", "url": "/api/ruleset/5e/subraces/hill-dwarf"}]' :: jsonb,
        '/api/ruleset/5e/races/dwarf'
    ),
    (
        'elf',
        'Elf',
        30,
        '[{"ability_score": {"index": "dex", "name": "DEX", "url": "/api/ruleset/5e/ability-scores/dex"}, "bonus": 2}]' :: jsonb,
        NULL,
        'Elves love freedom, variety, and self-expression, so they lean strongly toward the gentler aspects of chaos. They value and protect others'' freedom as well as their own, and they are more often good than not.',
        'Although elves reach physical maturity at about the same age as humans, the elven understanding of adulthood goes beyond physical growth to encompass worldly experience. An elf typically claims adulthood and an adult name around the age of 100 and can live to be 750 years old.',
        'Medium',
        'Elves range from under 5 to over 6 feet tall and have slender builds. Your size is Medium.',
        '[{"index": "skill-perception", "name": "Skill: Perception", "url": "/api/ruleset/5e/proficiencies/skill-perception"}]' :: jsonb,
        NULL,
        '[{"index": "common", "name": "Common", "url": "/api/ruleset/5e/languages/common"}, {"index": "elvish", "name": "Elvish", "url": "/api/ruleset/5e/languages/elvish"}]' :: jsonb,
        'You can speak, read, and write Common and Elvish. Elvish is fluid, with subtle intonations and intricate grammar. Elven literature is rich and varied, and their songs and poems are famous among other races. Many bards learn their language so they can add Elvish ballads to their repertoires.',
        NULL,
        '[{"index": "darkvision", "name": "Darkvision", "url": "/api/ruleset/5e/traits/darkvision"}, {"index": "fey-ancestry", "name": "Fey Ancestry", "url": "/api/ruleset/5e/traits/fey-ancestry"}, {"index": "trance", "name": "Trance", "url": "/api/ruleset/5e/traits/trance"}, {"index": "keen-senses", "name": "Keen Senses", "url": "/api/ruleset/5e/traits/keen-senses"}]' :: jsonb,
        '[{"index": "high-elf", "name": "High Elf", "url": "/api/ruleset/5e/subraces/high-elf"}]' :: jsonb,
        '/api/ruleset/5e/races/elf'
    ),
    (
        'halfling',
        'Halfling',
        25,
        '[{"ability_score": {"index": "dex", "name": "DEX", "url": "/api/ruleset/5e/ability-scores/dex"}, "bonus": 2}]' :: jsonb,
        NULL,
        'Most halflings are lawful good. As a rule, they are good-hearted and kind, hate to see others in pain, and have no tolerance for oppression. They are also very orderly and traditional, leaning heavily on the support of their community and the comfort of their old ways.',
        'A halfling reaches adulthood at the age of 20 and generally lives into the middle of his or her second century.',
        'Small',
        'Halflings average about 3 feet tall and weigh about 40 pounds. Your size is Small.',
        '[]' :: jsonb,
        NULL,
        '[{"index": "common", "name": "Common", "url": "/api/ruleset/5e/languages/common"}, {"index": "halfling", "name": "Halfling", "url": "/api/ruleset/5e/languages/halfling"}]' :: jsonb,
        'You can speak, read, and write Common and Halfling. The Halfling language isn''t secret, but halflings are loath to share it with others. They write very little, so they don''t have a rich body of literature. Their oral tradition, however, is very strong.',
        NULL,
        '[{"index": "brave", "name": "Brave", "url": "/api/ruleset/5e/traits/brave"}, {"index": "halfling-nimbleness", "name": "Halfling Nimbleness", "url": "/api/ruleset/5e/traits/halfling-nimbleness"}, {"index": "lucky", "name": "Lucky", "url": "/api/ruleset/5e/traits/lucky"}]' :: jsonb,
        '[{"index": "lightfoot-halfling", "name": "Lightfoot Halfling", "url": "/api/ruleset/5e/subraces/lightfoot-halfling"}]' :: jsonb,
        '/api/ruleset/5e/races/halfling'
    ),
    (
        'human',
        'Human',
        30,
        '[{"ability_score": {"index": "str", "name": "STR", "url": "/api/ruleset/5e/ability-scores/str"}, "bonus": 1}, 
      {"ability_score": {"index": "dex", "name": "DEX", "url": "/api/ruleset/5e/ability-scores/dex"}, "bonus": 1}, 
      {"ability_score": {"index": "con", "name": "CON", "url": "/api/ruleset/5e/ability-scores/con"}, "bonus": 1}, 
      {"ability_score": {"index": "int", "name": "INT", "url": "/api/ruleset/5e/ability-scores/int"}, "bonus": 1}, 
      {"ability_score": {"index": "wis", "name": "WIS", "url": "/api/ruleset/5e/ability-scores/wis"}, "bonus": 1}, 
      {"ability_score": {"index": "cha", "name": "CHA", "url": "/api/ruleset/5e/ability-scores/cha"}, "bonus": 1}]' :: jsonb,
        NULL,
        'Humans tend toward no particular alignment. The best and the worst are found among them.',
        'Humans reach adulthood in their late teens and live less than a century.',
        'Medium',
        'Humans vary widely in height and build, from barely 5 feet to well over 6 feet tall. Regardless of your position in that range, your size is Medium.',
        '[]' :: jsonb,
        NULL,
        '[{"index": "common", "name": "Common", "url": "/api/ruleset/5e/languages/common"}]' :: jsonb,
        'You can speak, read, and write Common and one extra language of your choice. Humans typically learn the languages of other peoples they deal with, including obscure dialects. They are fond of sprinkling their speech with words borrowed from other tongues: Orc curses, Elvish musical expressions, Dwarvish military phrases, and so on.',
        '{"choose": 1, "type": "languages", "from": {"option_set_type": "options_array", "options": [
        {"option_type": "reference", "item": {"index": "dwarvish", "name": "Dwarvish", "url": "/api/ruleset/5e/languages/dwarvish"}},
        {"option_type": "reference", "item": {"index": "elvish", "name": "Elvish", "url": "/api/ruleset/5e/languages/elvish"}},
        {"option_type": "reference", "item": {"index": "giant", "name": "Giant", "url": "/api/ruleset/5e/languages/giant"}},
        {"option_type": "reference", "item": {"index": "gnomish", "name": "Gnomish", "url": "/api/ruleset/5e/languages/gnomish"}},
        {"option_type": "reference", "item": {"index": "goblin", "name": "Goblin", "url": "/api/ruleset/5e/languages/goblin"}},
        {"option_type": "reference", "item": {"index": "halfling", "name": "Halfling", "url": "/api/ruleset/5e/languages/halfling"}},
        {"option_type": "reference", "item": {"index": "orc", "name": "Orc", "url": "/api/ruleset/5e/languages/orc"}},
        {"option_type": "reference", "item": {"index": "abyssal", "name": "Abyssal", "url": "/api/ruleset/5e/languages/abyssal"}},
        {"option_type": "reference", "item": {"index": "celestial", "name": "Celestial", "url": "/api/ruleset/5e/languages/celestial"}},
        {"option_type": "reference", "item": {"index": "draconic", "name": "Draconic", "url": "/api/ruleset/5e/languages/draconic"}},
        {"option_type": "reference", "item": {"index": "deep-speech", "name": "Deep Speech", "url": "/api/ruleset/5e/languages/deep-speech"}},
        {"option_type": "reference", "item": {"index": "infernal", "name": "Infernal", "url": "/api/ruleset/5e/languages/infernal"}},
        {"option_type": "reference", "item": {"index": "primordial", "name": "Primordial", "url": "/api/ruleset/5e/languages/primordial"}},
        {"option_type": "reference", "item": {"index": "sylvan", "name": "Sylvan", "url": "/api/ruleset/5e/languages/sylvan"}},
        {"option_type": "reference", "item": {"index": "undercommon", "name": "Undercommon", "url": "/api/ruleset/5e/languages/undercommon"}}
    ]}}' :: jsonb,
        '[]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/races/human'
    ),
    (
        'dragonborn',
        'Dragonborn',
        30,
        '[{"ability_score": {"index": "str", "name": "STR", "url": "/api/ruleset/5e/ability-scores/str"}, "bonus": 2}, 
      {"ability_score": {"index": "cha", "name": "CHA", "url": "/api/ruleset/5e/ability-scores/cha"}, "bonus": 1}]' :: jsonb,
        NULL,
        'Dragonborn tend to extremes, making a conscious choice for one side or the other in the cosmic war between good and evil. Most dragonborn are good, but those who side with evil can be terrible villains.',
        'Young dragonborn grow quickly. They walk hours after hatching, attain the size and development of a 10-year-old human child by the age of 3, and reach adulthood by 15. They live to be around 80.',
        'Medium',
        'Dragonborn are taller and heavier than humans, standing well over 6 feet tall and averaging almost 250 pounds. Your size is Medium.',
        '[]' :: jsonb,
        NULL,
        '[{"index": "common", "name": "Common", "url": "/api/ruleset/5e/languages/common"}, {"index": "draconic", "name": "Draconic", "url": "/api/ruleset/5e/languages/draconic"}]' :: jsonb,
        'You can speak, read, and write Common and Draconic. Draconic is thought to be one of the oldest languages and is often used in the study of magic. The language sounds harsh to most other creatures and includes numerous hard consonants and sibilants.',
        NULL,
        '[{"index": "draconic-ancestry", "name": "Draconic Ancestry", "url": "/api/ruleset/5e/traits/draconic-ancestry"}, 
      {"index": "breath-weapon", "name": "Breath Weapon", "url": "/api/ruleset/5e/traits/breath-weapon"}, 
      {"index": "damage-resistance", "name": "Damage Resistance", "url": "/api/ruleset/5e/traits/damage-resistance"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/races/dragonborn'
    ),
    (
        'gnome',
        'Gnome',
        25,
        '[{"ability_score": {"index": "int", "name": "INT", "url": "/api/ruleset/5e/ability-scores/int"}, "bonus": 2}]' :: jsonb,
        NULL,
        'Gnomes are most often good. Those who tend toward law are sages, engineers, researchers, scholars, investigators, or inventors. Those who tend toward chaos are minstrels, tricksters, wanderers, or fanciful jewelers. Gnomes are good-hearted, and even the tricksters among them are more playful than vicious.',
        'Gnomes mature at the same rate humans do, and most are expected to settle down into an adult life by around age 40. They can live 350 to almost 500 years.',
        'Small',
        'Gnomes are between 3 and 4 feet tall and average about 40 pounds. Your size is Small.',
        '[]' :: jsonb,
        NULL,
        '[{"index": "common", "name": "Common", "url": "/api/ruleset/5e/languages/common"}, {"index": "gnomish", "name": "Gnomish", "url": "/api/ruleset/5e/languages/gnomish"}]' :: jsonb,
        'You can speak, read, and write Common and Gnomish. The Gnomish language, which uses the Dwarvish script, is renowned for its technical treatises and its catalogs of knowledge about the natural world.',
        NULL,
        '[{"index": "darkvision", "name": "Darkvision", "url": "/api/ruleset/5e/traits/darkvision"}, {"index": "gnome-cunning", "name": "Gnome Cunning", "url": "/api/ruleset/5e/traits/gnome-cunning"}]' :: jsonb,
        '[{"index": "rock-gnome", "name": "Rock Gnome", "url": "/api/ruleset/5e/subraces/rock-gnome"}]' :: jsonb,
        '/api/ruleset/5e/races/gnome'
    ),
    (
        'half-elf',
        'Half-Elf',
        30,
        '[{"ability_score": {"index": "cha", "name": "CHA", "url": "/api/ruleset/5e/ability-scores/cha"}, "bonus": 2}]' :: jsonb,
        '{"choose": 2, "type": "ability_bonuses", "from": {"option_set_type": "options_array", "options": [
        {"option_type": "ability_bonus", "ability_score": {"index": "str", "name": "STR", "url": "/api/ruleset/5e/ability-scores/str"}, "bonus": 1},
        {"option_type": "ability_bonus", "ability_score": {"index": "dex", "name": "DEX", "url": "/api/ruleset/5e/ability-scores/dex"}, "bonus": 1},
        {"option_type": "ability_bonus", "ability_score": {"index": "con", "name": "CON", "url": "/api/ruleset/5e/ability-scores/con"}, "bonus": 1},
        {"option_type": "ability_bonus", "ability_score": {"index": "int", "name": "INT", "url": "/api/ruleset/5e/ability-scores/int"}, "bonus": 1},
        {"option_type": "ability_bonus", "ability_score": {"index": "wis", "name": "WIS", "url": "/api/ruleset/5e/ability-scores/wis"}, "bonus": 1}
    ]}}' :: jsonb,
        'Half-elves share the chaotic bent of their elven heritage. They value both personal freedom and creative expression, demonstrating neither love of leaders nor desire for followers. They chafe at rules, resent others'' demands, and sometimes prove unreliable, or at least unpredictable.',
        'Half-elves mature at the same rate humans do and reach adulthood around the age of 20. They live much longer than humans, however, often exceeding 180 years.',
        'Medium',
        'Half-elves are about the same size as humans, ranging from 5 to 6 feet tall. Your size is Medium.',
        '[]' :: jsonb,
        '{"choose": 2, "type": "proficiencies", "from": {"option_set_type": "options_array", "options": [
        {"option_type": "reference", "item": {"index": "skill-acrobatics", "name": "Skill: Acrobatics", "url": "/api/ruleset/5e/proficiencies/skill-acrobatics"}},
        {"option_type": "reference", "item": {"index": "skill-animal-handling", "name": "Skill: Animal Handling", "url": "/api/ruleset/5e/proficiencies/skill-animal-handling"}},
        {"option_type": "reference", "item": {"index": "skill-arcana", "name": "Skill: Arcana", "url": "/api/ruleset/5e/proficiencies/skill-arcana"}},
        {"option_type": "reference", "item": {"index": "skill-athletics", "name": "Skill: Athletics", "url": "/api/ruleset/5e/proficiencies/skill-athletics"}},
        {"option_type": "reference", "item": {"index": "skill-deception", "name": "Skill: Deception", "url": "/api/ruleset/5e/proficiencies/skill-deception"}},
        {"option_type": "reference", "item": {"index": "skill-history", "name": "Skill: History", "url": "/api/ruleset/5e/proficiencies/skill-history"}},
        {"option_type": "reference", "item": {"index": "skill-insight", "name": "Skill: Insight", "url": "/api/ruleset/5e/proficiencies/skill-insight"}},
        {"option_type": "reference", "item": {"index": "skill-intimidation", "name": "Skill: Intimidation", "url": "/api/ruleset/5e/proficiencies/skill-intimidation"}},
        {"option_type": "reference", "item": {"index": "skill-investigation", "name": "Skill: Investigation", "url": "/api/ruleset/5e/proficiencies/skill-investigation"}},
        {"option_type": "reference", "item": {"index": "skill-medicine", "name": "Skill: Medicine", "url": "/api/ruleset/5e/proficiencies/skill-medicine"}},
        {"option_type": "reference", "item": {"index": "skill-nature", "name": "Skill: Nature", "url": "/api/ruleset/5e/proficiencies/skill-nature"}},
        {"option_type": "reference", "item": {"index": "skill-perception", "name": "Skill: Perception", "url": "/api/ruleset/5e/proficiencies/skill-perception"}},
        {"option_type": "reference", "item": {"index": "skill-performance", "name": "Skill: Performance", "url": "/api/ruleset/5e/proficiencies/skill-performance"}},
        {"option_type": "reference", "item": {"index": "skill-persuasion", "name": "Skill: Persuasion", "url": "/api/ruleset/5e/proficiencies/skill-persuasion"}},
        {"option_type": "reference", "item": {"index": "skill-religion", "name": "Skill: Religion", "url": "/api/ruleset/5e/proficiencies/skill-religion"}},
        {"option_type": "reference", "item": {"index": "skill-sleight-of-hand", "name": "Skill: Sleight of Hand", "url": "/api/ruleset/5e/proficiencies/skill-sleight-of-hand"}},
        {"option_type": "reference", "item": {"index": "skill-stealth", "name": "Skill: Stealth", "url": "/api/ruleset/5e/proficiencies/skill-stealth"}},
        {"option_type": "reference", "item": {"index": "skill-survival", "name": "Skill: Survival", "url": "/api/ruleset/5e/proficiencies/skill-survival"}}
    ]}}' :: jsonb,
        '[{"index": "common", "name": "Common", "url": "/api/ruleset/5e/languages/common"}, {"index": "elvish", "name": "Elvish", "url": "/api/ruleset/5e/languages/elvish"}]' :: jsonb,
        'You can speak, read, and write Common, Elvish, and one extra language of your choice.',
        '{"choose": 1, "type": "languages", "from": {"option_set_type": "options_array", "options": [
        {"option_type": "reference", "item": {"index": "dwarvish", "name": "Dwarvish", "url": "/api/ruleset/5e/languages/dwarvish"}},
        {"option_type": "reference", "item": {"index": "giant", "name": "Giant", "url": "/api/ruleset/5e/languages/giant"}},
        {"option_type": "reference", "item": {"index": "gnomish", "name": "Gnomish", "url": "/api/ruleset/5e/languages/gnomish"}},
        {"option_type": "reference", "item": {"index": "goblin", "name": "Goblin", "url": "/api/ruleset/5e/languages/goblin"}},
        {"option_type": "reference", "item": {"index": "halfling", "name": "Halfling", "url": "/api/ruleset/5e/languages/halfling"}},
        {"option_type": "reference", "item": {"index": "orc", "name": "Orc", "url": "/api/ruleset/5e/languages/orc"}},
        {"option_type": "reference", "item": {"index": "abyssal", "name": "Abyssal", "url": "/api/ruleset/5e/languages/abyssal"}},
        {"option_type": "reference", "item": {"index": "celestial", "name": "Celestial", "url": "/api/ruleset/5e/languages/celestial"}},
        {"option_type": "reference", "item": {"index": "draconic", "name": "Draconic", "url": "/api/ruleset/5e/languages/draconic"}},
        {"option_type": "reference", "item": {"index": "deep-speech", "name": "Deep Speech", "url": "/api/ruleset/5e/languages/deep-speech"}},
        {"option_type": "reference", "item": {"index": "infernal", "name": "Infernal", "url": "/api/ruleset/5e/languages/infernal"}},
        {"option_type": "reference", "item": {"index": "primordial", "name": "Primordial", "url": "/api/ruleset/5e/languages/primordial"}},
        {"option_type": "reference", "item": {"index": "sylvan", "name": "Sylvan", "url": "/api/ruleset/5e/languages/sylvan"}},
        {"option_type": "reference", "item": {"index": "undercommon", "name": "Undercommon", "url": "/api/ruleset/5e/languages/undercommon"}}
    ]}}' :: jsonb,
        '[{"index": "darkvision", "name": "Darkvision", "url": "/api/ruleset/5e/traits/darkvision"}, {"index": "fey-ancestry", "name": "Fey Ancestry", "url": "/api/ruleset/5e/traits/fey-ancestry"}, {"index": "skill-versatility", "name": "Skill Versatility", "url": "/api/ruleset/5e/traits/skill-versatility"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/races/half-elf'
    ),
    (
        'half-orc',
        'Half-Orc',
        30,
        '[{"ability_score": {"index": "str", "name": "STR", "url": "/api/ruleset/5e/ability-scores/str"}, "bonus": 2}, 
      {"ability_score": {"index": "con", "name": "CON", "url": "/api/ruleset/5e/ability-scores/con"}, "bonus": 1}]' :: jsonb,
        NULL,
        'Half-orcs inherit a tendency toward chaos from their orc parents and are not strongly inclined toward good. Half-orcs raised among orcs and willing to live out their lives among them are usually evil.',
        'Half-orcs mature a little faster than humans, reaching adulthood around age 14. They age noticeably faster and rarely live longer than 75 years.',
        'Medium',
        'Half-orcs are somewhat larger and bulkier than humans, and they range from 5 to well over 6 feet tall. Your size is Medium.',
        '[{"index": "skill-intimidation", "name": "Skill: Intimidation", "url": "/api/ruleset/5e/proficiencies/skill-intimidation"}]' :: jsonb,
        NULL,
        '[{"index": "common", "name": "Common", "url": "/api/ruleset/5e/languages/common"}, {"index": "orc", "name": "Orc", "url": "/api/ruleset/5e/languages/orc"}]' :: jsonb,
        'You can speak, read, and write Common and Orc. Orc is a harsh, grating language with hard consonants. It has no script of its own but is written in the Dwarvish script.',
        NULL,
        '[{"index": "darkvision", "name": "Darkvision", "url": "/api/ruleset/5e/traits/darkvision"}, {"index": "savage-attacks", "name": "Savage Attacks", "url": "/api/ruleset/5e/traits/savage-attacks"}, {"index": "relentless-endurance", "name": "Relentless Endurance", "url": "/api/ruleset/5e/traits/relentless-endurance"}, {"index": "menacing", "name": "Menacing", "url": "/api/ruleset/5e/traits/menacing"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/races/half-orc'
    ),
    (
        'tiefling',
        'Tiefling',
        30,
        '[{"ability_score": {"index": "int", "name": "INT", "url": "/api/ruleset/5e/ability-scores/int"}, "bonus": 1}, 
      {"ability_score": {"index": "cha", "name": "CHA", "url": "/api/ruleset/5e/ability-scores/cha"}, "bonus": 2}]' :: jsonb,
        NULL,
        'Tieflings might not have an innate tendency toward evil, but many of them end up there. Evil or not, an independent nature inclines many tieflings toward a chaotic alignment.',
        'Tieflings mature at the same rate as humans but live a few years longer.',
        'Medium',
        'Tieflings are about the same size and build as humans. Your size is Medium.',
        '[]' :: jsonb,
        NULL,
        '[{"index": "common", "name": "Common", "url": "/api/ruleset/5e/languages/common"}, {"index": "infernal", "name": "Infernal", "url": "/api/ruleset/5e/languages/infernal"}]' :: jsonb,
        'You can speak, read, and write Common and Infernal.',
        NULL,
        '[{"index": "darkvision", "name": "Darkvision", "url": "/api/ruleset/5e/traits/darkvision"}, {"index": "hellish-resistance", "name": "Hellish Resistance", "url": "/api/ruleset/5e/traits/hellish-resistance"}, {"index": "infernal-legacy", "name": "Infernal Legacy", "url": "/api/ruleset/5e/traits/infernal-legacy"}]' :: jsonb,
        '[]' :: jsonb,
        '/api/ruleset/5e/races/tiefling'
    );