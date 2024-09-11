DROP TABLE IF EXISTS proficiency_choices;
DROP TABLE IF EXISTS proficiencies;
DROP TABLE IF EXISTS saving_throws;
DROP TABLE IF EXISTS starting_equipment;
DROP TABLE IF EXISTS starting_equipment_options;
DROP TABLE IF EXISTS classes;


CREATE TABLE classes (
    index VARCHAR PRIMARY KEY,
    name VARCHAR,
    hit_die INTEGER,
    class_levels VARCHAR,
    multi_classing JSONB,
    subclasses JSONB,
    spellcasting JSONB,
    spells VARCHAR,
    url VARCHAR
);

CREATE TABLE proficiency_choices (
    class_index VARCHAR REFERENCES classes(index),
    description VARCHAR,
    choose_number INT,
    proficiency_type VARCHAR,
    proficiency_options JSONB
);

CREATE TABLE proficiencies (
    class_index VARCHAR REFERENCES classes(index),
    proficiency_index VARCHAR,
    proficiency_name VARCHAR,
    proficiency_url VARCHAR
);

CREATE TABLE saving_throws (
    class_index VARCHAR REFERENCES classes(index),
    saving_throw_index VARCHAR,
    saving_throw_name VARCHAR,
    saving_throw_url VARCHAR
);

CREATE TABLE starting_equipment (
    class_index VARCHAR REFERENCES classes(index),
    equipment_index VARCHAR,
    equipment_name VARCHAR,
    equipment_url VARCHAR,
    quantity INTEGER
);

CREATE TABLE starting_equipment_options (
    class_index VARCHAR REFERENCES classes(index),
    description TEXT,
    choose INTEGER,
    equipment_type VARCHAR,
    equipment_options JSONB
);

INSERT INTO classes (index, name, hit_die, class_levels, multi_classing, subclasses, spellcasting, spells, url)
VALUES (
    'barbarian', 
    'Barbarian', 
    12, 
    '/api/classes/barbarian/levels', 
    '{"prerequisites": [{"ability_score": {"index": "str", "name": "STR", "url": "/api/ability-scores/str"}, "minimum_score": 13}], "proficiencies": [{"index": "shields", "name": "Shields", "url": "/api/proficiencies/shields"}, {"index": "simple-weapons", "name": "Simple Weapons", "url": "/api/proficiencies/simple-weapons"}, {"index": "martial-weapons", "name": "Martial Weapons", "url": "/api/proficiencies/martial-weapons"}]}',
    '[{"index": "berserker", "name": "Berserker", "url": "/api/subclasses/berserker"}]',
    null,
    null,
    '/api/classes/barbarian'
),
(
    'bard',
    'Bard',
    8,
    '/api/classes/bard/levels',
    '{"prerequisites": [{"ability_score": {"index": "cha", "name": "CHA", "url": "/api/ability-scores/cha"}, "minimum_score": 13}], "proficiencies": [{"index": "light-armor", "name": "Light Armor", "url": "/api/proficiencies/light-armor"}], "proficiency_choices": []}'::jsonb,
    '[{"index": "lore", "name": "Lore", "url": "/api/subclasses/lore"}]'::jsonb,
    '{"level": 1, "spellcasting_ability": {"index": "cha", "name": "CHA", "url": "/api/ability-scores/cha"}, "info":[{"name":"Cantrips","desc":["You know two cantrips of your choice from the bard spell list. You learn additional bard cantrips of your choice at higher levels, as shown in the Cantrips Known column of the Bard table."]},{"name":"Spell Slots","desc":["The Bard table shows how many spell slots you have to cast your spells of 1st level and higher. To cast one of these spells, you must expend a slot of the spell’s level or higher. You regain all expended spell slots when you finish a long rest.","For example, if you know the 1st-level spell cure wounds and have a 1st-level and a 2nd-level spell slot available, you can cast cure wounds using either slot."]},{"name":"Spells Known of 1st Level and Higher","desc":["You know four 1st-level spells of your choice from the bard spell list.","The Spells Known column of the Bard table shows when you learn more bard spells of your choice.","Each of these spells must be of a level for which you have spell slots, as shown on the table. For instance, when you reach 3rd level in this class, you can learn one new spell of 1st or 2nd level.","Additionally, when you gain a level in this class, you can choose one of the bard spells you know and replace it with another spell from the bard spell list, which also must be of a level for which you have spell slots."]},{"name":"Spellcasting Ability","desc":["Charisma is your spellcasting ability for your bard spells. Your magic comes from the heart and soul you pour into the performance of your music or oration. You use your Charisma whenever a spell refers to your spellcasting ability. In addition, you use your Charisma modifier when setting the saving throw DC for a bard spell you cast and when making an attack roll with one.","Spell save DC = 8 + your proficiency bonus + your Charisma modifier.","Spell attack modifier = your proficiency bonus + your Charisma modifier."]},{"name":"Ritual Casting","desc":["You can cast any bard spell you know as a ritual if that spell has the ritual tag."]},{"name":"Spellcasting Focus","desc":["You can use a musical instrument (see Equipment) as a spellcasting focus for your bard spells."]}]}'::jsonb,
    '/api/classes/bard/spells',
    '/api/classes/bard'
),
(
    'cleric',
    'Cleric',
    8,
    '/api/classes/cleric/levels',
    '{"prerequisites":[{"ability_score":{"index":"wis","name":"WIS","url":"/api/ability-scores/wis"},"minimum_score":13}],"proficiencies":[{"index":"light-armor","name":"Light Armor","url":"/api/proficiencies/light-armor"},{"index":"medium-armor","name":"Medium Armor","url":"/api/proficiencies/medium-armor"},{"index":"shields","name":"Shields","url":"/api/proficiencies/shields"}]}',
    '[{"index": "life","name": "Life","url": "/api/subclasses/life"}]',
    '{"level":1,"spellcasting_ability":{"index":"wis","name":"WIS","url":"/api/ability-scores/wis"},"info":[{"name":"Cantrips","desc":["At 1st level, you know three cantrips of your choice from the cleric spell list. You learn additional cleric cantrips of your choice at higher levels, as shown in the Cantrips Known column of the Cleric table."]},{"name":"Preparing and Casting Spells","desc":["The Cleric table shows how many spell slots you have to cast your spells of 1st level and higher. To cast one of these spells, you must expend a slot of the spell’s level or higher. You regain all expended spell slots when you finish a long rest.","You prepare the list of cleric spells that are available for you to cast, choosing from the cleric spell list. When you do so, choose a number of cleric spells equal to your Wisdom modifier + your cleric level (minimum of one spell). The spells must be of a level for which you have spell slots."]},{"name":"Spellcasting Ability","desc":["Wisdom is your spellcasting ability for your cleric spells. The power of your spells comes from your devotion to your deity. You use your Wisdom whenever a cleric spell refers to your spellcasting ability. In addition, you use your Wisdom modifier when setting the saving throw DC for a cleric spell you cast and when making an attack roll with one.","Spell save DC = 8 + your proficiency bonus + your Wisdom modifier","Spell attack modifier = your proficiency bonus + your Wisdom modifier"]},{"name":"Ritual Casting","desc":["You can cast a cleric spell as a ritual if that spell has the ritual tag and you have the spell prepared."]},{"name":"Spellcasting Focus","desc":["You can use a holy symbol (see Equipment) as a spellcasting focus for your cleric spells."]}]}',
    '/api/classes/cleric/spells',
    '/api/classes/cleric'
),
(
    'druid',
    'Druid',
    8,
    '/api/classes/druid/levels',
    '{"prerequisites":[{"ability_score":{"index":"wis","name":"WIS","url":"/api/ability-scores/wis"},"minimum_score":13}],"proficiencies":[{"index":"light-armor","name":"Light Armor","url":"/api/proficiencies/light-armor"},{"index":"medium-armor","name":"Medium Armor","url":"/api/proficiencies/medium-armor"},{"index":"shields","name":"Shields","url":"/api/proficiencies/shields"}]}',
    '[{"index":"land","name":"Land","url":"/api/subclasses/land"}]',
    '{"level":1,"spellcasting_ability":{"index":"wis","name":"WIS","url":"/api/ability-scores/wis"},"info":[{"name":"Cantrips","desc":["At 1st level, you know two cantrips of your choice from the druid spell list. You learn additional druid cantrips of your choice at higher levels, as shown in the Cantrips Known column of the Druid table."]},{"name":"Preparing and Casting Spells","desc":["The Druid table shows how many spell slots you have to cast your spells of 1st level and higher. To cast one of these druid spells, you must expend a slot of the spell’s level or higher. You regain all expended spell slots when you finish a long rest.","You prepare the list of druid spells that are available for you to cast, choosing from the druid spell list. When you do so, choose a number of druid spells equal to your Wisdom modifier + your druid level (minimum of one spell). The spells must be of a level for which you have spell slots.","For example, if you are a 3rd-level druid, you have four 1st-level and two 2nd-level spell slots. With a Wisdom of 16, your list of prepared spells can include six spells of 1st or 2nd level, in any combination. If you prepare the 1st-level spell cure wounds, you can cast it using a 1st-level or 2nd-level slot. Casting the spell doesn’t remove it from your list of prepared spells.","You can also change your list of prepared spells when you finish a long rest. Preparing a new list of druid spells requires time spent in prayer and meditation: at least 1 minute per spell level for each spell on your list."]},{"name":"Spellcasting Ability","desc":["Wisdom is your spellcasting ability for your druid spells, since your magic draws upon your devotion and attunement to nature. You use your Wisdom whenever a spell refers to your spellcasting ability. In addition, you use your Wisdom modifier when setting the saving throw DC for a druid spell you cast and when making an attack roll with one.","Spell save DC = 8 + your proficiency bonus + your Wisdom modifier.","Spell attack modifier = your proficiency bonus + your Wisdom modifier."]},{"name":"Ritual Casting","desc":["You can cast a druid spell as a ritual if that spell has the ritual tag and you have the spell prepared."]},{"name":"Spellcasting Focus","desc":["You can use a druidic focus (see chapter 5, ’Equipment’) as a spellcasting focus for your druid spells."]}]}',
    '/api/classes/druid/spells',
    '/api/classes/druid'
),
(
    'fighter',
    'Fighter',
    10,
    '/api/classes/fighter/levels',
    '{"prerequisite_options":{"type":"ability-scores","choose":1,"from":{"option_set_type":"options_array","options":[{"option_type":"score_prerequisite","ability_score":{"index":"str","name":"STR","url":"/api/ability-scores/str"},"minimum_score":13},{"option_type":"score_prerequisite","ability_score":{"index":"dex","name":"DEX","url":"/api/ability-scores/dex"},"minimum_score":13}]}}, "proficiencies":[{"index":"light-armor","name":"Light Armor","url":"/api/proficiencies/light-armor"},{"index":"medium-armor","name":"Medium Armor","url":"/api/proficiencies/medium-armor"},{"index":"shields","name":"Shields","url":"/api/proficiencies/shields"},{"index":"simple-weapons","name":"Simple Weapons","url":"/api/proficiencies/simple-weapons"},{"index":"martial-weapons","name":"Martial Weapons","url":"/api/proficiencies/martial-weapons"}]}',
    '[{"index":"champion","name":"Champion","url":"/api/subclasses/champion"}]',
    null,
    null,
    '/api/classes/fighter'
),
(
    'monk',
    'Monk',
    8,
    '/api/classes/monk/levels',
    '{"prerequisites":[{"ability_score":{"index":"dex","name":"DEX","url":"/api/ability-scores/dex"},"minimum_score":13},{"ability_score":{"index":"wis","name":"WIS","url":"/api/ability-scores/wis"},"minimum_score":13}],"proficiencies":[{"index":"simple-weapons","name":"Simple Weapons","url":"/api/proficiencies/simple-weapons"},{"index":"shortswords","name":"Shortswords","url":"/api/proficiencies/shortswords"}]}',
    '[{"index":"open-hand","name":"Open Hand","url":"/api/subclasses/open-hand"}]',
    null,
    null,
    '/api/classes/monk'
),
(
    'paladin',
    'Paladin',
    10,
    '/api/classes/paladin/levels',
    '{"prerequisites":[{"ability_score":{"index":"str","name":"STR","url":"/api/ability-scores/str"},"minimum_score":13},{"ability_score":{"index":"cha","name":"CHA","url":"/api/ability-scores/cha"},"minimum_score":13}],"proficiencies":[{"index":"light-armor","name":"Light Armor","url":"/api/proficiencies/light-armor"},{"index":"medium-armor","name":"Medium Armor","url":"/api/proficiencies/medium-armor"},{"index":"shields","name":"Shields","url":"/api/proficiencies/shields"},{"index":"simple-weapons","name":"Simple Weapons","url":"/api/proficiencies/simple-weapons"},{"index":"martial-weapons","name":"Martial Weapons","url":"/api/proficiencies/martial-weapons"}]}',
    '[{"index":"devotion","name":"Devotion","url":"/api/subclasses/devotion"}]',
    '{"level":2,"spellcasting_ability":{"index":"cha","name":"CHA","url":"/api/ability-scores/cha"},"info":[{"name":"Preparing and Casting Spells","desc":["The Paladin table shows how many spell slots you have to cast your spells. To cast one of your paladin spells of 1st level or higher, you must expend a slot of the spell’s level or higher. You regain all expended spell slots when you finish a long rest.","You prepare the list of paladin spells that are available for you to cast, choosing from the paladin spell list. When you do so, choose a number of paladin spells equal to your Charisma modifier + half your paladin level, rounded down (minimum of one spell). The spells must be of a level for which you have spell slots.","For example, if you are a 5th-level paladin, you have four 1st-level and two 2nd-level spell slots. With a Charisma of 14, your list of prepared spells can include four spells of 1st or 2nd level, in any combination. If you prepare the 1st-level spell cure wounds, you can cast it using a 1st-level or a 2nd- level slot. Casting the spell doesn’t remove it from your list of prepared spells.","You can change your list of prepared spells when you finish a long rest. Preparing a new list of paladin spells requires time spent in prayer and meditation: at least 1 minute per spell level for each spell on your list."]},{"name":"Spellcasting Ability","desc":["Charisma is your spellcasting ability for your paladin spells, since their power derives from the strength of your convictions. You use your Charisma whenever a spell refers to your spellcasting ability. In addition, you use your Charisma modifier when setting the saving throw DC for a paladin spell you cast and when making an attack roll with one.","Spell save DC = 8 + your proficiency bonus + your Charisma modifier.","Spell attack modifier = your proficiency bonus + your Charisma modifier."]},{"name":"Spellcasting Focus","desc":["You can use a holy symbol as a spellcasting focus for your paladin spells."]}]}',
    '/api/classes/paladin/spells',
    '/api/classes/paladin'
),
(
    'ranger',
    'Ranger',
    10,
    '/api/classes/ranger/levels',
    '{"prerequisites":[{"ability_score":{"index":"dex","name":"DEX","url":"/api/ability-scores/dex"},"minimum_score":13},{"ability_score":{"index":"wis","name":"WIS","url":"/api/ability-scores/wis"},"minimum_score":13}],"proficiencies":[{"index":"light-armor","name":"Light Armor","url":"/api/proficiencies/light-armor"},{"index":"medium-armor","name":"Medium Armor","url":"/api/proficiencies/medium-armor"},{"index":"shields","name":"Shields","url":"/api/proficiencies/shields"},{"index":"simple-weapons","name":"Simple Weapons","url":"/api/proficiencies/simple-weapons"},{"index":"martial-weapons","name":"Martial Weapons","url":"/api/proficiencies/martial-weapons"}],"proficiency_choices":[{"choose":1,"type":"proficiencies","from":{"option_set_type":"options_array","options":[{"option_type":"reference","item":{"index":"skill-animal-handling","name":"Skill: Animal Handling","url":"/api/proficiencies/skill-animal-handling"}},{"option_type":"reference","item":{"index":"skill-athletics","name":"Skill: Athletics","url":"/api/proficiencies/skill-athletics"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-investigation","name":"Skill: Investigation","url":"/api/proficiencies/skill-investigation"}},{"option_type":"reference","item":{"index":"skill-nature","name":"Skill: Nature","url":"/api/proficiencies/skill-nature"}},{"option_type":"reference","item":{"index":"skill-perception","name":"Skill: Perception","url":"/api/proficiencies/skill-perception"}},{"option_type":"reference","item":{"index":"skill-stealth","name":"Skill: Stealth","url":"/api/proficiencies/skill-stealth"}},{"option_type":"reference","item":{"index":"skill-survival","name":"Skill: Survival","url":"/api/proficiencies/skill-survival"}}]}}]}',
    '[{"index":"hunter","name":"Hunter","url":"/api/subclasses/hunter"}]',
    '{"level":2,"spellcasting_ability":{"index":"wis","name":"WIS","url":"/api/ability-scores/wis"},"info":[{"name":"Spell Slots","desc":["The Ranger table shows how many spell slots you have to cast your spells of 1st level and higher. To cast one of these spells, you must expend a slot of the spell’s level or higher. You regain all expended spell slots when you finish a long rest.","For example, if you know the 1st-level spell animal friendship and have a 1st-level and a 2nd-level spell slot available, you can cast animal friendship using either slot."]},{"name":"Spells Known of 1st Level and Higher","desc":["You know two 1st-level spells of your choice from the ranger spell list.","The Spells Known column of the Ranger table shows when you learn more ranger spells of your choice. Each of these spells must be of a level for which you have spell slots. For instance, when you reach 5th level in this class, you can learn one new spell of 1st or 2nd level.","Additionally, when you gain a level in this class, you can choose one of the ranger spells you know and replace it with another spell from the ranger spell list, which also must be of a level for which you have spell slots."]},{"name":"Spellcasting Ability","desc":["Wisdom is your spellcasting ability for your ranger spells, since your magic draws on your attunement to nature. You use your Wisdom whenever a spell refers to your spellcasting ability. In addition, you use your Wisdom modifier when setting the saving throw DC for a ranger spell you cast and when making an attack roll with one.","Spell save DC = 8 + your proficiency bonus + your Wisdom modifier.","Spell attack modifier = your proficiency bonus + your Wisdom modifier."]}]}',
    '/api/classes/ranger/spells',
    '/api/classes/ranger'
),
(
    'rogue',
    'Rogue',
    8,
    '/api/classes/rogue/levels',
    '{"prerequisites":[{"ability_score":{"index":"dex","name":"DEX","url":"/api/ability-scores/dex"},"minimum_score":13}],"proficiencies":[{"index":"light-armor","name":"Light Armor","url":"/api/proficiencies/light-armor"},{"index":"thieves-tools","name":"Thieves’ Tools","url":"/api/proficiencies/thieves-tools"}],"proficiency_choices":[{"choose":1,"type":"proficiencies","from":{"option_set_type":"options_array","options":[{"option_type":"reference","item":{"index":"skill-acrobatics","name":"Skill: Acrobatics","url":"/api/proficiencies/skill-acrobatics"}},{"option_type":"reference","item":{"index":"skill-athletics","name":"Skill: Athletics","url":"/api/proficiencies/skill-athletics"}},{"option_type":"reference","item":{"index":"skill-deception","name":"Skill: Deception","url":"/api/proficiencies/skill-deception"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-intimidation","name":"Skill: Intimidation","url":"/api/proficiencies/skill-intimidation"}},{"option_type":"reference","item":{"index":"skill-investigation","name":"Skill: Investigation","url":"/api/proficiencies/skill-investigation"}},{"option_type":"reference","item":{"index":"skill-perception","name":"Skill: Perception","url":"/api/proficiencies/skill-perception"}},{"option_type":"reference","item":{"index":"skill-performance","name":"Skill: Performance","url":"/api/proficiencies/skill-performance"}},{"option_type":"reference","item":{"index":"skill-persuasion","name":"Skill: Persuasion","url":"/api/proficiencies/skill-persuasion"}},{"option_type":"reference","item":{"index":"skill-sleight-of-hand","name":"Skill: Sleight of Hand","url":"/api/proficiencies/skill-sleight-of-hand"}},{"option_type":"reference","item":{"index":"skill-stealth","name":"Skill: Stealth","url":"/api/proficiencies/skill-stealth"}}]}}]}',
    '[{"index":"thief","name":"Thief","url":"/api/subclasses/thief"}]',
    null,
    null,
    '/api/classes/rogue'
),
(
    'sorcerer',
    'Sorcerer',
    6,
    '/api/classes/sorcerer/levels',
    '{"prerequisites":[{"ability_score":{"index":"cha","name":"CHA","url":"/api/ability-scores/cha"},"minimum_score":13}],"proficiencies":[]}',
    '[{"index":"draconic","name":"Draconic","url":"/api/subclasses/draconic"}]',
    '{"level":1,"spellcasting_ability":{"index":"cha","name":"CHA","url":"/api/ability-scores/cha"},"info":[{"name":"Cantrips","desc":["At 1st level, you know four cantrips of your choice from the sorcerer spell list. You learn additional sorcerer cantrips of your choice at higher levels, as shown in the Cantrips Known column of the Sorcerer table."]},{"name":"Spell Slots","desc":["The Sorcerer table shows how many spell slots you have to cast your spells of 1st level and higher. To cast one of these sorcerer spells, you must expend a slot of the spell’s level or higher. You regain all expended spell slots when you finish a long rest.","For example, if you know the 1st-level spell burning hands and have a 1st-level and a 2nd-level spell slot available, you can cast burning hands using either slot."]},{"name":"Spells Known of 1st Level and Higher","desc":["You know two 1st-level spells of your choice from the sorcerer spell list.","The Spells Known column of the Sorcerer table shows when you learn more sorcerer spells of your choice. Each of these spells must be of a level for which you have spell slots. For instance, when you reach 3rd level in this class, you can learn one new spell of 1st or 2nd level.","Additionally, when you gain a level in this class, you can choose one of the sorcerer spells you know and replace it with another spell from the sorcerer spell list, which also must be of a level for which you have spell slots."]},{"name":"Spellcasting Ability","desc":["Charisma is your spellcasting ability for your sorcerer spells, since the power of your magic relies on your ability to project your will into the world. You use your Charisma whenever a spell refers to your spellcasting ability. In addition, you use your Charisma modifier when setting the saving throw DC for a sorcerer spell you cast and when making an attack roll with one.","Spell save DC = 8 + your proficiency bonus + your Charisma modifier.","Spell attack modifier = your proficiency bonus + your Charisma modifier."]},{"name":"Spellcasting Focus","desc":["You can use an arcane focus as a spellcasting focus for your sorcerer spells."]}]}',
    '/api/classes/sorcerer/spells',
    '/api/classes/sorcerer'
),
(
    'warlock',
    'Warlock',
    8,
    '/api/classes/warlock/levels',
    '[{"index":"fiend","name":"Fiend","url":"/api/subclasses/fiend"}]',
    '{"level":1,"spellcasting_ability":{"index":"cha","name":"CHA","url":"/api/ability-scores/cha"},"info":[{"name":"Cantrips","desc":["You know two cantrips of your choice from the warlock spell list. You learn additional warlock cantrips of your choice at higher levels, as shown in the Cantrips Known column of the Warlock table."]},{"name":"Spell Slots","desc":["The Warlock table shows how many spell slots you have. The table also shows what the level of those slots is; all of your spell slots are the same level. To cast one of your warlock spells of 1st level or higher, you must expend a spell slot. You regain all expended spell slots when you finish a short or long rest.","For example, when you are 5th level, you have two 3rd-level spell slots. To cast the 1st-level spell thunderwave, you must spend one of those slots, and you cast it as a 3rd-level spell."]},{"name":"Spells Known of 1st Level and Higher","desc":["At 1st level, you know two 1st-level spells of your choice from the warlock spell list.","The Spells Known column of the Warlock table shows when you learn more warlock spells of your choice of 1st level and higher. ","A spell you choose must be of a level no higher than what’s shown in the table’s Slot Level column for your level. When you reach 6th level, for example, you learn a new warlock spell, which can be 1st, 2nd, or 3rd level.","Additionally, when you gain a level in this class, you can choose one of the warlock spells you know and replace it with another spell from the warlock spell list, which also must be of a level for which you have spell slots."]},{"name":"Spellcasting Ability","desc":["Charisma is your spellcasting ability for your warlock spells, so you use your Charisma whenever a spell refers to your spellcasting ability. In addition, you use your Charisma modifier when setting the saving throw DC for a warlock spell you cast and when making an attack roll with one.","Spell save DC = 8 + your proficiency bonus + your Charisma modifier.","Spell attack modifier = your proficiency bonus + your Charisma modifier."]},{"name":"Spellcasting Focus","desc":["You can use an arcane focus as a spellcasting focus for your warlock spells."]}]}',
    '{"prerequisites":[{"ability_score":{"index":"cha","name":"CHA","url":"/api/ability-scores/cha"},"minimum_score":13}],"proficiencies":[{"index":"light-armor","name":"Light Armor","url":"/api/proficiencies/light-armor"},{"index":"simple-weapons","name":"Simple Weapons","url":"/api/proficiencies/simple-weapons"}]}',
    '/api/classes/warlock/spells',
    '/api/classes/warlock'
),
(
    'wizard',
    'Wizard',
    6,
    '/api/classes/wizard/levels',
    '[{"index": "evocation", "name": "Evocation", "url": "/api/subclasses/evocation"}]',
    '{"level": 1,"spellcasting_ability": {"index": "int", "name": "INT", "url": "/api/ability-scores/int"},"info": [{"name": "Cantrips", "desc": ["At 1st level, you know three cantrips of your choice from the wizard spell list. You learn additional wizard cantrips of your choice at higher levels, as shown in the Cantrips Known column of the Wizard table."]},{"name": "Spellbook", "desc": ["At 1st level, you have a spellbook containing six 1st-level wizard spells of your choice. Your spellbook is the repository of the wizard spells you know, except your cantrips, which are fixed in your mind."]},{"name": "Preparing and Casting Spells", "desc": ["The Wizard table shows how many spell slots you have to cast your spells of 1st level and higher. To cast one of these spells, you must expend a slot of the spell’s level or higher. You regain all expended spell slots when you finish a long rest.", "You prepare the list of wizard spells that are available for you to cast. To do so, choose a number of wizard spells from your spellbook equal to your Intelligence modifier + your wizard level (minimum of one spell). The spells must be of a level for which you have spell slots.", "You can change your list of prepared spells when you finish a long rest."]},{"name": "Spellcasting Ability", "desc": ["Intelligence is your spellcasting ability for your wizard spells. You use your Intelligence modifier when setting the saving throw DC for a wizard spell you cast and when making an attack roll with one."]},{"name": "Ritual Casting", "desc": ["You can cast a wizard spell as a ritual if that spell has the ritual tag and you have the spell in your spellbook."]},{"name": "Spellcasting Focus", "desc": ["You can use an arcane focus as a spellcasting focus for your wizard spells."]}]}',
    '{"prerequisites": [{"ability_score": {"index": "int", "name": "INT", "url": "/api/ability-scores/int"}, "minimum_score": 13}], "proficiencies": []}',
    '/api/classes/wizard/spells',
    '/api/classes/wizard'
);

INSERT INTO proficiency_choices (class_index, description, choose_number, proficiency_type, proficiency_options)
VALUES 
('barbarian', 'Choose two from Animal Handling, Athletics, Intimidation, Nature, Perception, and Survival', 2, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-animal-handling","name":"Skill: Animal Handling","url":"/api/proficiencies/skill-animal-handling"}},{"option_type":"reference","item":{"index":"skill-athletics","name":"Skill: Athletics","url":"/api/proficiencies/skill-athletics"}},{"option_type":"reference","item":{"index":"skill-intimidation","name":"Skill: Intimidation","url":"/api/proficiencies/skill-intimidation"}},{"option_type":"reference","item":{"index":"skill-nature","name":"Skill: Nature","url":"/api/proficiencies/skill-nature"}},{"option_type":"reference","item":{"index":"skill-perception","name":"Skill: Perception","url":"/api/proficiencies/skill-perception"}},{"option_type":"reference","item":{"index":"skill-survival","name":"Skill: Survival","url":"/api/proficiencies/skill-survival"}}]'),
('bard', 'Choose any three', 3, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-acrobatics","name":"Skill: Acrobatics","url":"/api/proficiencies/skill-acrobatics"}},{"option_type":"reference","item":{"index":"skill-animal-handling","name":"Skill: Animal Handling","url":"/api/proficiencies/skill-animal-handling"}},{"option_type":"reference","item":{"index":"skill-arcana","name":"Skill: Arcana","url":"/api/proficiencies/skill-arcana"}},{"option_type":"reference","item":{"index":"skill-athletics","name":"Skill: Athletics","url":"/api/proficiencies/skill-athletics"}},{"option_type":"reference","item":{"index":"skill-deception","name":"Skill: Deception","url":"/api/proficiencies/skill-deception"}},{"option_type":"reference","item":{"index":"skill-history","name":"Skill: History","url":"/api/proficiencies/skill-history"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-intimidation","name":"Skill: Intimidation","url":"/api/proficiencies/skill-intimidation"}},{"option_type":"reference","item":{"index":"skill-investigation","name":"Skill: Investigation","url":"/api/proficiencies/skill-investigation"}},{"option_type":"reference","item":{"index":"skill-medicine","name":"Skill: Medicine","url":"/api/proficiencies/skill-medicine"}},{"option_type":"reference","item":{"index":"skill-nature","name":"Skill: Nature","url":"/api/proficiencies/skill-nature"}},{"option_type":"reference","item":{"index":"skill-perception","name":"Skill: Perception","url":"/api/proficiencies/skill-perception"}},{"option_type":"reference","item":{"index":"skill-performance","name":"Skill: Performance","url":"/api/proficiencies/skill-performance"}},{"option_type":"reference","item":{"index":"skill-persuasion","name":"Skill: Persuasion","url":"/api/proficiencies/skill-persuasion"}},{"option_type":"reference","item":{"index":"skill-religion","name":"Skill: Religion","url":"/api/proficiencies/skill-religion"}},{"option_type":"reference","item":{"index":"skill-sleight-of-hand","name":"Skill: Sleight of Hand","url":"/api/proficiencies/skill-sleight-of-hand"}},{"option_type":"reference","item":{"index":"skill-stealth","name":"Skill: Stealth","url":"/api/proficiencies/skill-stealth"}},{"option_type":"reference","item":{"index":"skill-survival","name":"Skill: Survival","url":"/api/proficiencies/skill-survival"}}]'),
('bard', 'Three musical instruments of your choice', 3, 'proficiencies', '[{"option_type":"reference","item":{"index":"bagpipes","name":"Bagpipes","url":"/api/proficiencies/bagpipes"}},{"option_type":"reference","item":{"index":"drum","name":"Drum","url":"/api/proficiencies/drum"}},{"option_type":"reference","item":{"index":"dulcimer","name":"Dulcimer","url":"/api/proficiencies/dulcimer"}},{"option_type":"reference","item":{"index":"flute","name":"Flute","url":"/api/proficiencies/flute"}},{"option_type":"reference","item":{"index":"lute","name":"Lute","url":"/api/proficiencies/lute"}},{"option_type":"reference","item":{"index":"lyre","name":"Lyre","url":"/api/proficiencies/lyre"}},{"option_type":"reference","item":{"index":"horn","name":"Horn","url":"/api/proficiencies/horn"}},{"option_type":"reference","item":{"index":"pan-flute","name":"Pan flute","url":"/api/proficiencies/pan-flute"}},{"option_type":"reference","item":{"index":"shawm","name":"Shawm","url":"/api/proficiencies/shawm"}},{"option_type":"reference","item":{"index":"viol","name":"Viol","url":"/api/proficiencies/viol"}}]'),
('cleric', 'Choose two from History, Insight, Medicine, Persuasion, and Religion', 2, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-history","name":"Skill: History","url":"/api/proficiencies/skill-history"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-medicine","name":"Skill: Medicine","url":"/api/proficiencies/skill-medicine"}},{"option_type":"reference","item":{"index":"skill-persuasion","name":"Skill: Persuasion","url":"/api/proficiencies/skill-persuasion"}},{"option_type":"reference","item":{"index":"skill-religion","name":"Skill: Religion","url":"/api/proficiencies/skill-religion"}}]'),
('druid', 'Choose two from Arcana, Animal Handling, Insight, Medicine, Nature, Perception, Religion, and Survival', 2, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-arcana","name":"Skill: Arcana","url":"/api/proficiencies/skill-arcana"}},{"option_type":"reference","item":{"index":"skill-animal-handling","name":"Skill: Animal Handling","url":"/api/proficiencies/skill-animal-handling"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-medicine","name":"Skill: Medicine","url":"/api/proficiencies/skill-medicine"}},{"option_type":"reference","item":{"index":"skill-nature","name":"Skill: Nature","url":"/api/proficiencies/skill-nature"}},{"option_type":"reference","item":{"index":"skill-perception","name":"Skill: Perception","url":"/api/proficiencies/skill-perception"}},{"option_type":"reference","item":{"index":"skill-religion","name":"Skill: Religion","url":"/api/proficiencies/skill-religion"}},{"option_type":"reference","item":{"index":"skill-survival","name":"Skill: Survival","url":"/api/proficiencies/skill-survival"}}]'),
('fighter', 'Choose two skills from Acrobatics, Animal Handling, Athletics, History, Insight, Intimidation, Perception, and Survival', 2, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-acrobatics","name":"Skill: Acrobatics","url":"/api/proficiencies/skill-acrobatics"}},{"option_type":"reference","item":{"index":"skill-animal-handling","name":"Skill: Animal Handling","url":"/api/proficiencies/skill-animal-handling"}},{"option_type":"reference","item":{"index":"skill-athletics","name":"Skill: Athletics","url":"/api/proficiencies/skill-athletics"}},{"option_type":"reference","item":{"index":"skill-history","name":"Skill: History","url":"/api/proficiencies/skill-history"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-intimidation","name":"Skill: Intimidation","url":"/api/proficiencies/skill-intimidation"}},{"option_type":"reference","item":{"index":"skill-perception","name":"Skill: Perception","url":"/api/proficiencies/skill-perception"}},{"option_type":"reference","item":{"index":"skill-survival","name":"Skill: Survival","url":"/api/proficiencies/skill-survival"}}]'),
('monk', 'Choose two from Acrobatics, Athletics, History, Insight, Religion, and Stealth', 2, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-acrobatics","name":"Skill: Acrobatics","url":"/api/proficiencies/skill-acrobatics"}},{"option_type":"reference","item":{"index":"skill-athletics","name":"Skill: Athletics","url":"/api/proficiencies/skill-athletics"}},{"option_type":"reference","item":{"index":"skill-history","name":"Skill: History","url":"/api/proficiencies/skill-history"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-religion","name":"Skill: Religion","url":"/api/proficiencies/skill-religion"}},{"option_type":"reference","item":{"index":"skill-stealth","name":"Skill: Stealth","url":"/api/proficiencies/skill-stealth"}}]'),
('monk', 'Choose one type of artisan’s tools or one musical instrument', 1, 'proficiencies', '[{"option_type":"choice","choice":{"desc":"artisan’s tools","type":"proficiencies","choose":1,"from":{"option_set_type":"options_array","options":[{"option_type":"reference","item":{"index":"alchemists-supplies","name":"Alchemist’s Supplies","url":"/api/proficiencies/alchemists-supplies"}},{"option_type":"reference","item":{"index":"brewers-supplies","name":"Brewer’s Supplies","url":"/api/proficiencies/brewers-supplies"}},{"option_type":"reference","item":{"index":"calligraphers-supplies","name":"Calligrapher’s Supplies","url":"/api/proficiencies/calligraphers-supplies"}},{"option_type":"reference","item":{"index":"carpenters-tools","name":"Carpenter’s Tools","url":"/api/proficiencies/carpenters-tools"}},{"option_type":"reference","item":{"index":"cartographers-tools","name":"Cartographer’s Tools","url":"/api/proficiencies/cartographers-tools"}},{"option_type":"reference","item":{"index":"cobblers-tools","name":"Cobbler’s Tools","url":"/api/proficiencies/cobblers-tools"}},{"option_type":"reference","item":{"index":"cooks-utensils","name":"Cook’s utensils","url":"/api/proficiencies/cooks-utensils"}},{"option_type":"reference","item":{"index":"glassblowers-tools","name":"Glassblower’s Tools","url":"/api/proficiencies/glassblowers-tools"}},{"option_type":"reference","item":{"index":"jewelers-tools","name":"Jeweler’s Tools","url":"/api/proficiencies/jewelers-tools"}},{"option_type":"reference","item":{"index":"leatherworkers-tools","name":"Leatherworker’s Tools","url":"/api/proficiencies/leatherworkers-tools"}},{"option_type":"reference","item":{"index":"masons-tools","name":"Mason’s Tools","url":"/api/proficiencies/masons-tools"}},{"option_type":"reference","item":{"index":"painters-supplies","name":"Painter’s Supplies","url":"/api/proficiencies/painters-supplies"}},{"option_type":"reference","item":{"index":"potters-tools","name":"Potter’s Tools","url":"/api/proficiencies/potters-tools"}},{"option_type":"reference","item":{"index":"smiths-tools","name":"Smith’s Tools","url":"/api/proficiencies/smiths-tools"}},{"option_type":"reference","item":{"index":"tinkers-tools","name":"Tinker’s Tools","url":"/api/proficiencies/tinkers-tools"}},{"option_type":"reference","item":{"index":"weavers-tools","name":"Weaver’s Tools","url":"/api/proficiencies/weavers-tools"}},{"option_type":"reference","item":{"index":"woodcarvers-tools","name":"Woodcarver’s Tools","url":"/api/proficiencies/woodcarvers-tools"}},{"option_type":"reference","item":{"index":"disguise-kit","name":"Disguise Kit","url":"/api/proficiencies/disguise-kit"}},{"option_type":"reference","item":{"index":"forgery-kit","name":"Forgery Kit","url":"/api/proficiencies/forgery-kit"}}]}}},{"option_type":"choice","choice":{"desc":"musical instrument","type":"proficiencies","choose":1,"from":{"option_set_type":"options_array","options":[{"option_type":"reference","item":{"index":"bagpipes","name":"Bagpipes","url":"/api/proficiencies/bagpipes"}},{"option_type":"reference","item":{"index":"drum","name":"Drum","url":"/api/proficiencies/drum"}},{"option_type":"reference","item":{"index":"dulcimer","name":"Dulcimer","url":"/api/proficiencies/dulcimer"}},{"option_type":"reference","item":{"index":"flute","name":"Flute","url":"/api/proficiencies/flute"}},{"option_type":"reference","item":{"index":"lute","name":"Lute","url":"/api/proficiencies/lute"}},{"option_type":"reference","item":{"index":"lyre","name":"Lyre","url":"/api/proficiencies/lyre"}},{"option_type":"reference","item":{"index":"horn","name":"Horn","url":"/api/proficiencies/horn"}},{"option_type":"reference","item":{"index":"pan-flute","name":"Pan flute","url":"/api/proficiencies/pan-flute"}},{"option_type":"reference","item":{"index":"shawm","name":"Shawm","url":"/api/proficiencies/shawm"}},{"option_type":"reference","item":{"index":"viol","name":"Viol","url":"/api/proficiencies/viol"}}]}}}]'),
('paladin', 'Choose two from Athletics, Insight, Intimidation, Medicine, Persuasion, and Religion', 2, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-athletics","name":"Skill: Athletics","url":"/api/proficiencies/skill-athletics"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-intimidation","name":"Skill: Intimidation","url":"/api/proficiencies/skill-intimidation"}},{"option_type":"reference","item":{"index":"skill-medicine","name":"Skill: Medicine","url":"/api/proficiencies/skill-medicine"}},{"option_type":"reference","item":{"index":"skill-persuasion","name":"Skill: Persuasion","url":"/api/proficiencies/skill-persuasion"}},{"option_type":"reference","item":{"index":"skill-religion","name":"Skill: Religion","url":"/api/proficiencies/skill-religion"}}]'),
('ranger', 'Choose three from Animal Handling, Athletics, Insight, Investigation, Nature, Perception, Stealth, and Survival', 3, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-animal-handling","name":"Skill: Animal Handling","url":"/api/proficiencies/skill-animal-handling"}},{"option_type":"reference","item":{"index":"skill-athletics","name":"Skill: Athletics","url":"/api/proficiencies/skill-athletics"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-investigation","name":"Skill: Investigation","url":"/api/proficiencies/skill-investigation"}},{"option_type":"reference","item":{"index":"skill-nature","name":"Skill: Nature","url":"/api/proficiencies/skill-nature"}},{"option_type":"reference","item":{"index":"skill-perception","name":"Skill: Perception","url":"/api/proficiencies/skill-perception"}},{"option_type":"reference","item":{"index":"skill-stealth","name":"Skill: Stealth","url":"/api/proficiencies/skill-stealth"}},{"option_type":"reference","item":{"index":"skill-survival","name":"Skill: Survival","url":"/api/proficiencies/skill-survival"}}]'),
('rogue', 'Choose four from Acrobatics, Athletics, Deception, Insight, Intimidation, Investigation, Perception, Performance, Persuasion, Sleight of Hand, and Stealth', 4, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-acrobatics","name":"Skill: Acrobatics","url":"/api/proficiencies/skill-acrobatics"}},{"option_type":"reference","item":{"index":"skill-athletics","name":"Skill: Athletics","url":"/api/proficiencies/skill-athletics"}},{"option_type":"reference","item":{"index":"skill-deception","name":"Skill: Deception","url":"/api/proficiencies/skill-deception"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-intimidation","name":"Skill: Intimidation","url":"/api/proficiencies/skill-intimidation"}},{"option_type":"reference","item":{"index":"skill-investigation","name":"Skill: Investigation","url":"/api/proficiencies/skill-investigation"}},{"option_type":"reference","item":{"index":"skill-perception","name":"Skill: Perception","url":"/api/proficiencies/skill-perception"}},{"option_type":"reference","item":{"index":"skill-performance","name":"Skill: Performance","url":"/api/proficiencies/skill-performance"}},{"option_type":"reference","item":{"index":"skill-persuasion","name":"Skill: Persuasion","url":"/api/proficiencies/skill-persuasion"}},{"option_type":"reference","item":{"index":"skill-sleight-of-hand","name":"Skill: Sleight of Hand","url":"/api/proficiencies/skill-sleight-of-hand"}},{"option_type":"reference","item":{"index":"skill-stealth","name":"Skill: Stealth","url":"/api/proficiencies/skill-stealth"}}]'),
('sorcerer', 'Choose two from Arcana, Deception, Insight, Intimidation, Persuasion, and Religion', 2, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-arcana","name":"Skill: Arcana","url":"/api/proficiencies/skill-arcana"}},{"option_type":"reference","item":{"index":"skill-deception","name":"Skill: Deception","url":"/api/proficiencies/skill-deception"}},{"option_type":"reference","item":{"index":"skill-insight","name":"Skill: Insight","url":"/api/proficiencies/skill-insight"}},{"option_type":"reference","item":{"index":"skill-intimidation","name":"Skill: Intimidation","url":"/api/proficiencies/skill-intimidation"}},{"option_type":"reference","item":{"index":"skill-persuasion","name":"Skill: Persuasion","url":"/api/proficiencies/skill-persuasion"}},{"option_type":"reference","item":{"index":"skill-religion","name":"Skill: Religion","url":"/api/proficiencies/skill-religion"}}]'),
('warlock', 'Choose two skills from Arcana, Deception, History, Intimidation, Investigation, Nature, and Religion', 2, 'proficiencies', '[{"option_type":"reference","item":{"index":"skill-arcana","name":"Skill: Arcana","url":"/api/proficiencies/skill-arcana"}},{"option_type":"reference","item":{"index":"skill-deception","name":"Skill: Deception","url":"/api/proficiencies/skill-deception"}},{"option_type":"reference","item":{"index":"skill-history","name":"Skill: History","url":"/api/proficiencies/skill-history"}},{"option_type":"reference","item":{"index":"skill-intimidation","name":"Skill: Intimidation","url":"/api/proficiencies/skill-intimidation"}},{"option_type":"reference","item":{"index":"skill-investigation","name":"Skill: Investigation","url":"/api/proficiencies/skill-investigation"}},{"option_type":"reference","item":{"index":"skill-nature","name":"Skill: Nature","url":"/api/proficiencies/skill-nature"}},{"option_type":"reference","item":{"index":"skill-religion","name":"Skill: Religion","url":"/api/proficiencies/skill-religion"}}]'),
('wizard', 'Choose two from Arcana, History, Insight, Investigation, Medicine, and Religion', 2, 'proficiencies','[{"option_type": "reference", "item": {"index": "skill-arcana", "name": "Skill: Arcana", "url": "/api/proficiencies/skill-arcana"}},{"option_type": "reference", "item": {"index": "skill-history", "name": "Skill: History", "url": "/api/proficiencies/skill-history"}},{"option_type": "reference", "item": {"index": "skill-insight", "name": "Skill: Insight", "url": "/api/proficiencies/skill-insight"}},{"option_type": "reference", "item": {"index": "skill-investigation", "name": "Skill: Investigation", "url": "/api/proficiencies/skill-investigation"}},{"option_type": "reference", "item": {"index": "skill-medicine", "name": "Skill: Medicine", "url": "/api/proficiencies/skill-medicine"}},{"option_type": "reference", "item": {"index": "skill-religion", "name": "Skill: Religion", "url": "/api/proficiencies/skill-religion"}}]');

INSERT INTO proficiencies (class_index, proficiency_index, proficiency_name, proficiency_url)
VALUES
('barbarian', 'light-armor', 'Light Armor', '/api/proficiencies/light-armor'),
('barbarian', 'medium-armor', 'Medium Armor', '/api/proficiencies/medium-armor'),
('barbarian', 'shields', 'Shields', '/api/proficiencies/shields'),
('barbarian', 'simple-weapons', 'Simple Weapons', '/api/proficiencies/simple-weapons'),
('barbarian', 'martial-weapons', 'Martial Weapons', '/api/proficiencies/martial-weapons'),
('barbarian', 'saving-throw-str', 'Saving Throw: STR', '/api/proficiencies/saving-throw-str'),
('barbarian', 'saving-throw-con', 'Saving Throw: CON', '/api/proficiencies/saving-throw-con'),
('bard', 'light-armor', 'Light Armor', '/api/proficiencies/light-armor'),
('bard', 'simple-weapons', 'Simple Weapons', '/api/proficiencies/simple-weapons'),
('bard', 'longswords', 'Longswords', '/api/proficiencies/longswords'),
('bard', 'rapiers', 'Rapiers', '/api/proficiencies/rapiers'),
('bard', 'shortswords', 'Shortswords', '/api/proficiencies/shortswords'),
('bard', 'hand-crossbows', 'Hand crossbows', '/api/proficiencies/hand-crossbows'),
('bard', 'saving-throw-dex', 'Saving Throw: DEX', '/api/proficiencies/saving-throw-dex'),
('bard', 'saving-throw-cha', 'Saving Throw: CHA', '/api/proficiencies/saving-throw-cha'),
('cleric', 'light-armor', 'Light Armor', '/api/proficiencies/light-armor'),
('cleric', 'medium-armor', 'Medium Armor', '/api/proficiencies/medium-armor'),
('cleric', 'shields', 'Shields', '/api/proficiencies/shields'),
('cleric', 'simple-weapons', 'Simple Weapons', '/api/proficiencies/simple-weapons'),
('cleric', 'saving-throw-wis', 'Saving Throw: WIS', '/api/proficiencies/saving-throw-wis'),
('cleric', 'saving-throw-cha', 'Saving Throw: CHA', '/api/proficiencies/saving-throw-cha'),
('druid', 'light-armor', 'Light Armor', '/api/proficiencies/light-armor'),
('druid', 'medium-armor', 'Medium Armor', '/api/proficiencies/medium-armor'),
('druid', 'shields', 'Shields', '/api/proficiencies/shields'),
('druid', 'clubs', 'Clubs', '/api/proficiencies/clubs'),
('druid', 'daggers', 'Daggers', '/api/proficiencies/daggers'),
('druid', 'javelins', 'Javelins', '/api/proficiencies/javelins'),
('druid', 'maces', 'Maces', '/api/proficiencies/maces'),
('druid', 'quarterstaffs', 'Quarterstaffs', '/api/proficiencies/quarterstaffs'),
('druid', 'sickles', 'Sickles', '/api/proficiencies/sickles'),
('druid', 'spears', 'Spears', '/api/proficiencies/spears'),
('druid', 'darts', 'Darts', '/api/proficiencies/darts'),
('druid', 'slings', 'Slings', '/api/proficiencies/slings'),
('druid', 'scimitars', 'Scimitars', '/api/proficiencies/scimitars'),
('druid', 'herbalism-kit', 'Herbalism Kit', '/api/proficiencies/herbalism-kit'),
('druid', 'saving-throw-int', 'Saving Throw: INT', '/api/proficiencies/saving-throw-int'),
('druid', 'saving-throw-wis', 'Saving Throw: WIS', '/api/proficiencies/saving-throw-wis'),
('fighter', 'all-armor', 'All armor', '/api/proficiencies/all-armor'),
('fighter', 'shields', 'Shields', '/api/proficiencies/shields'),
('fighter', 'simple-weapons', 'Simple Weapons', '/api/proficiencies/simple-weapons'),
('fighter', 'martial-weapons', 'Martial Weapons', '/api/proficiencies/martial-weapons'),
('fighter', 'saving-throw-str', 'Saving Throw: STR', '/api/proficiencies/saving-throw-str'),
('fighter', 'saving-throw-con', 'Saving Throw: CON', '/api/proficiencies/saving-throw-con'),
('monk', 'simple-weapons', 'Simple Weapons', '/api/proficiencies/simple-weapons'),
('monk', 'shortswords', 'Shortswords', '/api/proficiencies/shortswords'),
('monk', 'saving-throw-dex', 'Saving Throw: DEX', '/api/proficiencies/saving-throw-dex'),
('monk', 'saving-throw-str', 'Saving Throw: STR', '/api/proficiencies/saving-throw-str'),
('paladin', 'all-armor', 'All armor', '/api/proficiencies/all-armor'),
('paladin', 'shields', 'Shields', '/api/proficiencies/shields'),
('paladin', 'simple-weapons', 'Simple Weapons', '/api/proficiencies/simple-weapons'),
('paladin', 'martial-weapons', 'Martial Weapons', '/api/proficiencies/martial-weapons'),
('paladin', 'saving-throw-wis', 'Saving Throw: WIS', '/api/proficiencies/saving-throw-wis'),
('paladin', 'saving-throw-cha', 'Saving Throw: CHA', '/api/proficiencies/saving-throw-cha'),
('ranger', 'light-armor', 'Light Armor', '/api/proficiencies/light-armor'),
('ranger', 'medium-armor', 'Medium Armor', '/api/proficiencies/medium-armor'),
('ranger', 'shields', 'Shields', '/api/proficiencies/shields'),
('ranger', 'simple-weapons', 'Simple Weapons', '/api/proficiencies/simple-weapons'),
('ranger', 'martial-weapons', 'Martial Weapons', '/api/proficiencies/martial-weapons'),
('ranger', 'saving-throw-dex', 'Saving Throw: DEX', '/api/proficiencies/saving-throw-dex'),
('ranger', 'saving-throw-str', 'Saving Throw: STR', '/api/proficiencies/saving-throw-str'),
('rogue', 'light-armor', 'Light Armor', '/api/proficiencies/light-armor'),
('rogue', 'simple-weapons', 'Simple Weapons', '/api/proficiencies/simple-weapons'),
('rogue', 'longswords', 'Longswords', '/api/proficiencies/longswords'),
('rogue', 'rapiers', 'Rapiers', '/api/proficiencies/rapiers'),
('rogue', 'shortswords', 'Shortswords', '/api/proficiencies/shortswords'),
('rogue', 'hand-crossbows', 'Hand crossbows', '/api/proficiencies/hand-crossbows'),
('rogue', 'thieves-tools', 'Thieves’ Tools', '/api/proficiencies/thieves-tools'),
('rogue', 'saving-throw-dex', 'Saving Throw: DEX', '/api/proficiencies/saving-throw-dex'),
('rogue', 'saving-throw-int', 'Saving Throw: INT', '/api/proficiencies/saving-throw-int'),
('sorcerer', 'daggers', 'Daggers', '/api/proficiencies/daggers'),
('sorcerer', 'darts', 'Darts', '/api/proficiencies/darts'),
('sorcerer', 'slings', 'Slings', '/api/proficiencies/slings'),
('sorcerer', 'quarterstaffs', 'Quarterstaffs', '/api/proficiencies/quarterstaffs'),
('sorcerer', 'crossbows-light', 'Crossbows, light', '/api/proficiencies/crossbows-light'),
('sorcerer', 'saving-throw-con', 'Saving Throw: CON', '/api/proficiencies/saving-throw-con'),
('sorcerer', 'saving-throw-cha', 'Saving Throw: CHA', '/api/proficiencies/saving-throw-cha'),
('warlock', 'light-armor', 'Light Armor', '/api/proficiencies/light-armor'),
('warlock', 'simple-weapons', 'Simple Weapons', '/api/proficiencies/simple-weapons'),
('warlock', 'saving-throw-wis', 'Saving Throw: WIS', '/api/proficiencies/saving-throw-wis'),
('warlock', 'saving-throw-cha', 'Saving Throw: CHA', '/api/proficiencies/saving-throw-cha'),
('wizard', 'daggers', 'Daggers', '/api/proficiencies/daggers'),
('wizard', 'darts', 'Darts', '/api/proficiencies/darts'),
('wizard', 'slings', 'Slings', '/api/proficiencies/slings'),
('wizard', 'quarterstaffs', 'Quarterstaffs', '/api/proficiencies/quarterstaffs'),
('wizard', 'crossbows-light', 'Crossbows, light', '/api/proficiencies/crossbows-light'),
('wizard', 'saving-throw-int', 'Saving Throw: INT', '/api/proficiencies/saving-throw-int'),
('wizard', 'saving-throw-wis', 'Saving Throw: WIS', '/api/proficiencies/saving-throw-wis');

INSERT INTO saving_throws (class_index, saving_throw_index, saving_throw_name, saving_throw_url)
VALUES 
('barbarian', 'str', 'STR', '/api/ability-scores/str'),
('barbarian', 'con', 'CON', '/api/ability-scores/con'),
('bard', 'dex', 'DEX', '/api/ability-scores/dex'),
('bard', 'cha', 'CHA', '/api/ability-scores/cha'),
('cleric', 'wis', 'WIS', '/api/ability-scores/wis'),
('cleric', 'cha', 'CHA', '/api/ability-scores/cha'),
('druid', 'int', 'INT', '/api/ability-scores/int'),
('druid', 'wis', 'WIS', '/api/ability-scores/wis'),
('fighter', 'str', 'STR', '/api/ability-scores/str'),
('fighter', 'con', 'CON', '/api/ability-scores/con'),
('monk', 'str', 'STR', '/api/ability-scores/str'),
('monk', 'dex', 'DEX', '/api/ability-scores/dex'),
('paladin', 'wis', 'WIS', '/api/ability-scores/wis'),
('paladin', 'cha', 'CHA', '/api/ability-scores/cha'),
('ranger', 'str', 'STR', '/api/ability-scores/str'),
('ranger', 'dex', 'DEX', '/api/ability-scores/dex'),
('rogue', 'dex', 'DEX', '/api/ability-scores/dex'),
('rogue', 'int', 'INT', '/api/ability-scores/int'),
('sorcerer', 'con', 'CON', '/api/ability-scores/con'),
('sorcerer', 'cha', 'CHA', '/api/ability-scores/cha'),
('warlock', 'wis', 'WIS', '/api/ability-scores/wis'),
('warlock', 'cha', 'CHA', '/api/ability-scores/cha'),
('wizard', 'int', 'INT', '/api/ability-scores/int'),
('wizard', 'wis', 'WIS', '/api/ability-scores/wis');

INSERT INTO starting_equipment (class_index, equipment_index, equipment_name, equipment_url, quantity)
VALUES
('barbarian', 'explorers-pack', 'Explorer''s Pack', '/api/equipment/explorers-pack', 1),
('barbarian', 'javelin', 'Javelin', '/api/equipment/javelin', 4),
('bard', 'leather-armor', 'Leather Armor', '/api/equipment/leather-armor', 1),
('bard', 'dagger', 'Dagger', '/api/equipment/dagger', 1),
('cleric', 'shield', 'Shield', '/api/equipment/shield', 1),
('druid', 'leather-armor', 'Leather Armor', '/api/equipment/leather-armor', 1),
('druid', 'explorers-pack', 'Explorer’s Pack', '/api/equipment/explorers-pack', 1),
('monk', 'dart', 'Dart', '/api/equipment/dart', 10),
('paladin', 'chain-mail', 'Chain Mail', '/api/equipment/chain-mail', 1),
('ranger', 'longbow', 'Longbow', '/api/equipment/longbow', 1),
('ranger', 'arrow', 'Arrow', '/api/equipment/arrow', 20),
('rogue', 'leather-armor', 'Leather Armor', '/api/equipment/leather-armor', 1),
('rogue', 'dagger', 'Dagger', '/api/equipment/dagger', 2),
('rogue', 'thieves-tools', 'Thieves’ Tools', '/api/equipment/thieves-tools', 1),
('sorcerer', 'dagger', 'Dagger', '/api/equipment/dagger', 2),
('warlock', 'dagger', 'Dagger', '/api/equipment/dagger', 2),
('warlock', 'leather-armor', 'Leather Armor', '/api/equipment/leather-armor', 1),
('wizard', 'spellbook', 'Spellbook', '/api/equipment/spellbook', 1);

INSERT INTO starting_equipment_options (class_index, description, choose, equipment_type, equipment_options)
VALUES
('barbarian', '(a) a greataxe or (b) any martial melee weapon', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"greataxe","name":"Greataxe","url":"/api/equipment/greataxe"}},{"option_type":"choice","choice":{"desc":"any martial melee weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"martial-melee-weapons","name":"Martial Melee Weapons","url":"/api/equipment-categories/martial-melee-weapons"}}}}]'),
('barbarian', '(a) two handaxes or (b) any simple weapon', 1, 'equipment', '[{"option_type":"counted_reference","count":2,"of":{"index":"handaxe","name":"Handaxe","url":"/api/equipment/handaxe"}},{"option_type":"choice","choice":{"desc":"any simple weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"simple-weapons","name":"Simple Weapons","url":"/api/equipment-categories/simple-weapons"}}}}]'),
('bard', '(a) a rapier, (b) a longsword, or (c) any simple weapon', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"rapier","name":"Rapier","url":"/api/equipment/rapier"}},{"option_type":"counted_reference","count":1,"of":{"index":"longsword","name":"Longsword","url":"/api/equipment/longsword"}},{"option_type":"choice","choice":{"desc":"any simple weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"simple-weapons","name":"Simple Weapons","url":"/api/equipment-categories/simple-weapons"}}}}]'),
('bard', '(a) a diplomat’s pack or (b) an entertainer’s pack', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"diplomats-pack","name":"Diplomat’s Pack","url":"/api/equipment/diplomats-pack"}},{"option_type":"counted_reference","count":1,"of":{"index":"entertainers-pack","name":"Entertainer’s Pack","url":"/api/equipment/entertainers-pack"}}]'),
('bard', '(a) a lute or (b) any other musical instrument', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"lute","name":"Lute","url":"/api/equipment/lute"}},{"option_type":"choice","choice":{"desc":"any other musical instrument","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"musical-instruments","name":"Musical Instruments","url":"/api/equipment-categories/musical-instruments"}}}}]'),
('cleric', '(a) a mace or (b) a warhammer (if proficient)', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"mace","name":"Mace","url":"/api/equipment/mace"}},{"option_type":"counted_reference","count":1,"of":{"index":"warhammer","name":"Warhammer","url":"/api/equipment/warhammer"},"prerequisites":[{"type":"proficiency","proficiency":{"index":"warhammers","name":"Warhammers","url":"/api/proficiencies/warhammers"}}]}]'),
('cleric', '(a) scale mail, (b) leather armor, or (c) chain mail (if proficient)', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"scale-mail","name":"Scale Mail","url":"/api/equipment/scale-mail"}},{"option_type":"counted_reference","count":1,"of":{"index":"leather-armor","name":"Leather Armor","url":"/api/equipment/leather-armor"}},{"option_type":"counted_reference","count":1,"of":{"index":"chain-mail","name":"Chain Mail","url":"/api/equipment/chain-mail"},"prerequisites":[{"type":"proficiency","proficiency":{"index":"chain-mail","name":"Chain Mail","url":"/api/proficiencies/chain-mail"}}]}]'),
('cleric', '(a) a light crossbow and 20 bolts or (b) any simple weapon', 1, 'equipment', '[{"option_type":"multiple","items":[{"option_type":"counted_reference","count":1,"of":{"index":"crossbow-light","name":"Crossbow, light","url":"/api/equipment/crossbow-light"}},{"option_type":"counted_reference","count":20,"of":{"index":"crossbow-bolt","name":"Crossbow bolt","url":"/api/equipment/crossbow-bolt"}}]},{"option_type":"choice","choice":{"desc":"any simple weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"simple-weapons","name":"Simple Weapons","url":"/api/equipment-categories/simple-weapons"}}}}]'),
('cleric', '(a) a priest’s pack or (b) an explorer’s pack', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"priests-pack","name":"Priest’s Pack","url":"/api/equipment/priests-pack"}},{"option_type":"counted_reference","count":1,"of":{"index":"explorers-pack","name":"Explorer’s Pack","url":"/api/equipment/explorers-pack"}}]'),
('cleric', 'holy symbol', 1, 'equipment', null),
('druid', '(a) a wooden shield or (b) any simple weapon', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"shield","name":"Shield","url":"/api/equipment/shield"}},{"option_type":"choice","choice":{"desc":"any simple weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"simple-weapons","name":"Simple Weapons","url":"/api/equipment-categories/simple-weapons"}}}}]'),
('druid', '(a) a scimitar or (b) any simple melee weapon', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"scimitar","name":"Scimitar","url":"/api/equipment/scimitar"}},{"option_type":"choice","choice":{"desc":"any simple melee weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"simple-melee-weapons","name":"Simple Melee Weapons","url":"/api/equipment-categories/simple-melee-weapons"}}}}]'),
('druid', 'druidic focus', 1, 'equipment', null),
('fighter', '(a) chain mail or (b) leather armor, longbow, and 20 arrows', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"chain-mail","name":"Chain Mail","url":"/api/equipment/chain-mail"}},{"option_type":"multiple","items":[{"option_type":"counted_reference","count":1,"of":{"index":"leather-armor","name":"Leather Armor","url":"/api/equipment/leather-armor"}},{"option_type":"counted_reference","count":1,"of":{"index":"longbow","name":"Longbow","url":"/api/equipment/longbow"}},{"option_type":"counted_reference","count":20,"of":{"index":"arrow","name":"Arrow","url":"/api/equipment/arrow"}}]}]'),
('fighter', '(a) a martial weapon and a shield or (b) two martial weapons', 1, 'equipment', '[{"option_type":"multiple","items":[{"option_type":"choice","choice":{"desc":"a martial weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"martial-weapons","name":"Martial Weapons","url":"/api/equipment-categories/martial-weapons"}}}},{"option_type":"counted_reference","count":1,"of":{"index":"shield","name":"Shield","url":"/api/equipment/shield"}}]},{"option_type":"choice","choice":{"desc":"two martial weapons","choose":2,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"martial-weapons","name":"Martial Weapons","url":"/api/equipment-categories/martial-weapons"}}}}]'),
('fighter', '(a) a light crossbow and 20 bolts or (b) two handaxes', 1, 'equipment', '[{"option_type":"multiple","items":[{"option_type":"counted_reference","count":1,"of":{"index":"crossbow-light","name":"Crossbow, light","url":"/api/equipment/crossbow-light"}},{"option_type":"counted_reference","count":20,"of":{"index":"crossbow-bolt","name":"Crossbow bolt","url":"/api/equipment/crossbow-bolt"}}]},{"option_type":"counted_reference","count":2,"of":{"index":"handaxe","name":"Handaxe","url":"/api/equipment/handaxe"}}]'),
('fighter', '(a) a dungeoneer’s pack or (b) an explorer’s pack', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"dungeoneers-pack","name":"Dungeoneer’s Pack","url":"/api/equipment/dungeoneers-pack"}},{"option_type":"counted_reference","count":1,"of":{"index":"explorers-pack","name":"Explorer’s Pack","url":"/api/equipment/explorers-pack"}}]'),
('monk', '(a) a shortsword or (b) any simple weapon', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"shortsword","name":"Shortsword","url":"/api/equipment/shortsword"}},{"option_type":"choice","choice":{"desc":"any simple weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"simple-weapons","name":"Simple Weapons","url":"/api/equipment-categories/simple-weapons"}}}}]'),
('monk', '(a) a dungeoneer’s pack or (b) an explorer’s pack', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"dungeoneers-pack","name":"Dungeoneer’s Pack","url":"/api/equipment/dungeoneers-pack"}},{"option_type":"counted_reference","count":1,"of":{"index":"explorers-pack","name":"Explorer’s Pack","url":"/api/equipment/explorers-pack"}}]'),
('paladin', '(a) a martial weapon and a shield or (b) two martial weapons', 1, 'equipment', '[{"option_type":"multiple","items":[{"option_type":"choice","choice":{"desc":"a martial weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"name":"Martial Weapons","index":"martial-weapons","url":"/api/equipment-categories/martial-weapons"}}}},{"option_type":"counted_reference","count":1,"of":{"index":"shield","name":"Shield","url":"/api/equipment/shield"}}]},{"option_type":"choice","choice":{"desc":"two martial weapons","choose":2,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"martial-weapons","name":"Martial Weapons","url":"/api/equipment-categories/martial-weapons"}}}}]'),
('paladin', '(a) five javelins or (b) any simple melee weapon', 1, 'equipment', '[{"option_type":"counted_reference","count":5,"of":{"index":"javelin","name":"Javelin","url":"/api/equipment/javelin"}},{"option_type":"choice","choice":{"desc":"any simple weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"simple-weapons","name":"Simple Weapons","url":"/api/equipment-categories/simple-weapons"}}}}]'),
('paladin', '(a) a priest’s pack or (b) an explorer’s pack', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"priests-pack","name":"Priest’s Pack","url":"/api/equipment/priests-pack"}},{"option_type":"counted_reference","count":1,"of":{"index":"explorers-pack","name":"Explorer’s Pack","url":"/api/equipment/explorers-pack"}}]'),
('paladin', 'holy symbol', 1, 'equipment', null),
('ranger', '(a) scale mail or (b) leather armor', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"scale-mail","name":"Scale Mail","url":"/api/equipment/scale-mail"}},{"option_type":"counted_reference","count":1,"of":{"index":"leather-armor","name":"Leather Armor","url":"/api/equipment/leather-armor"}}]'),
('ranger', '(a) two shortswords or (b) two simple melee weapons', 1, 'equipment', '[{"option_type":"counted_reference","count":2,"of":{"index":"shortsword","name":"Shortsword","url":"/api/equipment/shortsword"}},{"option_type":"choice","choice":{"desc":"two simple melee weapons","choose":2,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"simple-melee-weapons","name":"Simple Melee Weapons","url":"/api/equipment-categories/simple-melee-weapons"}}}}]'),
('ranger', '(a) a dungeoneer’s pack or (b) an explorer’s pack', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"dungeoneers-pack","name":"Dungeoneer’s Pack","url":"/api/equipment/dungeoneers-pack"}},{"option_type":"counted_reference","count":1,"of":{"index":"explorers-pack","name":"Explorer’s Pack","url":"/api/equipment/explorers-pack"}}]'),
('rogue', '(a) a rapier or (b) a shortsword', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"rapier","name":"Rapier","url":"/api/equipment/rapier"}},{"option_type":"counted_reference","count":1,"of":{"index":"shortsword","name":"Shortsword","url":"/api/equipment/shortsword"}}]'),
('rogue', '(a) a shortbow and quiver of 20 arrows or (b) a shortsword', 1, 'equipment', '[{"option_type":"multiple","items":[{"option_type":"counted_reference","count":1,"of":{"index":"shortbow","name":"Shortbow","url":"/api/equipment/shortbow"}},{"option_type":"counted_reference","count":20,"of":{"index":"arrow","name":"Arrow","url":"/api/equipment/arrow"}}]},{"option_type":"counted_reference","count":1,"of":{"index":"shortsword","name":"Shortsword","url":"/api/equipment/shortsword"}}]'),
('rogue', '(a) a burglar’s pack, (b) a dungeoneer’s pack, or (c) an explorer’s pack', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"burglars-pack","name":"Burglar’s Pack","url":"/api/equipment/burglars-pack"}},{"option_type":"counted_reference","count":1,"of":{"index":"dungeoneers-pack","name":"Dungeoneer’s Pack","url":"/api/equipment/dungeoneers-pack"}},{"option_type":"counted_reference","count":1,"of":{"index":"explorers-pack","name":"Explorer’s Pack","url":"/api/equipment/explorers-pack"}}]'),
('sorcerer', '(a) a light crossbow and 20 bolts or (b) any simple weapon', 1, 'equipment', '[{"option_type":"multiple","items":[{"option_type":"counted_reference","count":1,"of":{"index":"crossbow-light","name":"Crossbow, light","url":"/api/equipment/crossbow-light"}},{"option_type":"counted_reference","count":20,"of":{"index":"crossbow-bolt","name":"Crossbow bolt","url":"/api/equipment/crossbow-bolt"}}]},{"option_type":"choice","choice":{"desc":"any simple weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"simple-weapons","name":"Simple Weapons","url":"/api/equipment-categories/simple-weapons"}}}}]'),
('sorcerer', '(a) a component pouch or (b) an arcane focus', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"component-pouch","name":"Component pouch","url":"/api/equipment/component-pouch"}},{"option_type":"choice","choice":{"desc":"arcane focus","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"arcane-foci","name":"Arcane Foci","url":"/api/equipment-categories/arcane-foci"}}}}]'),
('sorcerer', '(a) a dungeoneer’s pack or (b) an explorer’s pack', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"dungeoneers-pack","name":"Dungeoneer’s Pack","url":"/api/equipment/dungeoneers-pack"}},{"option_type":"counted_reference","count":1,"of":{"index":"explorers-pack","name":"Explorer’s Pack","url":"/api/equipment/explorers-pack"}}]'),
('warlock', '(a) a light crossbow and 20 bolts or (b) any simple weapon', 1, 'equipment', '[{"option_type":"multiple","items":[{"option_type":"counted_reference","count":1,"of":{"index":"crossbow-light","name":"Crossbow, light","url":"/api/equipment/crossbow-light"}},{"option_type":"counted_reference","count":20,"of":{"index":"crossbow-bolt","name":"Crossbow bolt","url":"/api/equipment/crossbow-bolt"}}]},{"option_type":"choice","choice":{"desc":"any simple weapon","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"simple-weapons","name":"Simple Weapons","url":"/api/equipment-categories/simple-weapons"}}}}]'),
('warlock', '(a) a component pouch or (b) an arcane focus', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"component-pouch","name":"Component pouch","url":"/api/equipment/component-pouch"}},{"option_type":"choice","choice":{"desc":"arcane focus","choose":1,"type":"equipment","from":{"option_set_type":"equipment_category","equipment_category":{"index":"arcane-foci","name":"Arcane Foci","url":"/api/equipment-categories/arcane-foci"}}}}]'),
('warlock', '(a) a scholar’s pack or (b) a dungeoneer’s pack', 1, 'equipment', '[{"option_type":"counted_reference","count":1,"of":{"index":"scholars-pack","name":"Scholar’s Pack","url":"/api/equipment/scholars-pack"}},{"option_type":"counted_reference","count":1,"of":{"index":"dungeoneers-pack","name":"Dungeoneer’s Pack","url":"/api/equipment/dungeoneers-pack"}}]'),
('warlock', 'any simple weapon', 1, 'equipment', null),
('wizard', '(a) a quarterstaff or (b) a dagger', 1, 'equipment','[{"option_type": "counted_reference", "count": 1, "of": {"index": "quarterstaff", "name": "Quarterstaff", "url": "/api/equipment/quarterstaff"}},{"option_type": "counted_reference", "count": 1, "of": {"index": "dagger", "name": "Dagger", "url": "/api/equipment/dagger"}}]'),
('wizard', '(a) a component pouch or (b) an arcane focus', 1, 'equipment','[{"option_type": "counted_reference", "count": 1, "of": {"index": "component-pouch", "name": "Component pouch", "url": "/api/equipment/component-pouch"}},{"option_type": "choice", "choice": {"desc": "arcane focus", "choose": 1, "type": "equipment", "from": {"option_set_type": "equipment_category", "equipment_category": {"index": "arcane-foci", "name": "Arcane Foci", "url": "/api/equipment-categories/arcane-foci"}}}}]'),
('wizard', '(a) a scholar’s pack or (b) an explorer’s pack', 1, 'equipment','[{"option_type": "counted_reference", "count": 1, "of": {"index": "scholars-pack", "name": "Scholar’s Pack", "url": "/api/equipment/scholars-pack"}},{"option_type": "counted_reference", "count": 1, "of": {"index": "explorers-pack", "name": "Explorer’s Pack", "url": "/api/equipment/explorers-pack"}}]');