DROP TABLE IF EXISTS skills;
DROP TABLE IF EXISTS ability_scores;


CREATE TABLE ability_scores (
    ability_index VARCHAR PRIMARY KEY,
    ability_name VARCHAR,
    full_name VARCHAR,
    description TEXT[],
    url VARCHAR
);

CREATE TABLE skills (
    skill_index VARCHAR PRIMARY KEY,
    skill_name VARCHAR,
    ability_index VARCHAR REFERENCES ability_scores(ability_index),
    url VARCHAR
);

INSERT INTO ability_scores (ability_index, ability_name, full_name, description, url) 
VALUES
('str', 'STR', 'Strength', ARRAY['Strength measures bodily power, athletic training, and the extent to which you can exert raw physical force.', 'A Strength check can model any attempt to lift, push, pull, or break something, to force your body through a space, or to otherwise apply brute force to a situation. The Athletics skill reflects aptitude in certain kinds of Strength checks.'], '/api/ability-scores/str'),
('dex', 'DEX', 'Dexterity', ARRAY['Dexterity measures agility, reflexes, and balance.', 'A Dexterity check can model any attempt to move nimbly, quickly, or quietly, or to keep from falling on tricky footing. The Acrobatics, Sleight of Hand, and Stealth skills reflect aptitude in certain kinds of Dexterity checks.'], '/api/ability-scores/dex'),
('con', 'CON', 'Constitution', ARRAY['Constitution measures health, stamina, and vital force.', 'Constitution checks are uncommon, and no skills apply to Constitution checks, because the endurance this ability represents is largely passive rather than involving a specific effort on the part of a character or monster.'], '/api/ability-scores/con'),
('int', 'INT', 'Intelligence', ARRAY['Intelligence measures mental acuity, accuracy of recall, and the ability to reason.', 'An Intelligence check comes into play when you need to draw on logic, education, memory, or deductive reasoning. The Arcana, History, Investigation, Nature, and Religion skills reflect aptitude in certain kinds of Intelligence checks.'], '/api/ability-scores/int'),
('wis', 'WIS', 'Wisdom', ARRAY['Wisdom reflects how attuned you are to the world around you and represents perceptiveness and intuition.', 'A Wisdom check might reflect an effort to read body language, understand someone’s feelings, notice things about the environment, or care for an injured person. The Animal Handling, Insight, Medicine, Perception, and Survival skills reflect aptitude in certain kinds of Wisdom checks.'], '/api/ability-scores/wis'),
('cha', 'CHA', 'Charisma', ARRAY['Charisma measures your ability to interact effectively with others. It includes such factors as confidence and eloquence, and it can represent a charming or commanding personality.', 'A Charisma check might arise when you try to influence or entertain others, when you try to make an impression or tell a convincing lie, or when you are navigating a tricky social situation. The Deception, Intimidation, Performance, and Persuasion skills reflect aptitude in certain kinds of Charisma checks.'], '/api/ability-scores/cha');

INSERT INTO skills (skill_index, skill_name, ability_index, url) 
VALUES
('athletics', 'Athletics', 'str', '/api/skills/athletics'),
('acrobatics', 'Acrobatics', 'dex', '/api/skills/acrobatics'),
('sleight-of-hand', 'Sleight of Hand', 'dex', '/api/skills/sleight-of-hand'),
('stealth', 'Stealth', 'dex', '/api/skills/stealth'),
('arcana', 'Arcana', 'int', '/api/skills/arcana'),
('history', 'History', 'int', '/api/skills/history'),
('investigation', 'Investigation', 'int', '/api/skills/investigation'),
('nature', 'Nature', 'int', '/api/skills/nature'),
('religion', 'Religion', 'int', '/api/skills/religion'),
('animal-handling', 'Animal Handling', 'wis', '/api/skills/animal-handling'),
('insight', 'Insight', 'wis', '/api/skills/insight'),
('medicine', 'Medicine', 'wis', '/api/skills/medicine'),
('perception', 'Perception', 'wis', '/api/skills/perception'),
('survival', 'Survival', 'wis', '/api/skills/survival'),
('deception', 'Deception', 'cha', '/api/skills/deception'),
('intimidation', 'Intimidation', 'cha', '/api/skills/intimidation'),
('performance', 'Performance', 'cha', '/api/skills/performance'),
('persuasion', 'Persuasion', 'cha', '/api/skills/persuasion');
