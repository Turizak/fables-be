DROP TABLE IF EXISTS languages; 

CREATE TABLE languages (
    index VARCHAR PRIMARY KEY,
    name VARCHAR,
    description TEXT,
    type VARCHAR,
    typical_speakers JSONB,
    script VARCHAR,
    url VARCHAR
);

INSERT INTO languages (index, name, description, type, typical_speakers, script, url)
VALUES 
('common', 'Common', NULL, 'Standard', '["Humans"]', 'Common', '/api/ruleset/5e/languages/common'),
('dwarvish', 'Dwarvish', 'Dwarvish is full of hard consonants and guttural sounds.', 'Standard', '["Dwarves"]', 'Dwarvish', '/api/ruleset/5e/languages/dwarvish'),
('elvish', 'Elvish', 'Elvish is fluid, with subtle intonations and intricate grammar. Elven literature is rich and varied, and their songs and poems are famous among other races. Many bards learn their language so they can add Elvish ballads to their repertoires.', 'Standard', '["Elves"]', 'Elvish', '/api/ruleset/5e/languages/elvish'),
('giant', 'Giant', NULL, 'Standard', '["Ogres", "Giants"]', 'Dwarvish', '/api/ruleset/5e/languages/giant'),
('gnomish', 'Gnomish', 'The Gnomish language, which uses the Dwarvish script, is renowned for its technical treatises and its catalogs of knowledge about the natural world.', 'Standard', '["Gnomes"]', 'Dwarvish', '/api/ruleset/5e/languages/gnomish'),
('goblin', 'Goblin', NULL, 'Standard', '["Goblinoids"]', 'Dwarvish', '/api/ruleset/5e/languages/goblin'),
('halfling', 'Halfling', 'The Halfling language isn’t secret, but halflings are loath to share it with others. They write very little, so they don’t have a rich body of literature. Their oral tradition, however, is very strong.', 'Standard', '["Halflings"]', 'Common', '/api/ruleset/5e/languages/halfling'),
('orc', 'Orc', 'Orc is a harsh, grating language with hard consonants. It has no script of its own but is written in the Dwarvish script.', 'Standard', '["Orcs"]', 'Dwarvish', '/api/ruleset/5e/languages/orc'),
('abyssal', 'Abyssal', NULL, 'Exotic', '["Demons"]', 'Infernal', '/api/ruleset/5e/languages/abyssal'),
('celestial', 'Celestial', NULL, 'Exotic', '["Celestials"]', 'Celestial', '/api/ruleset/5e/languages/celestial'),
('draconic', 'Draconic', 'Draconic is thought to be one of the oldest languages and is often used in the study of magic. The language sounds harsh to most other creatures and includes numerous hard consonants and sibilants.', 'Exotic', '["Dragons", "Dragonborn"]', 'Draconic', '/api/ruleset/5e/languages/draconic'),
('deep-speech', 'Deep Speech', NULL, 'Exotic', '["Aboleths", "Cloakers"]', NULL, '/api/ruleset/5e/languages/deep-speech'),
('infernal', 'Infernal', NULL, 'Exotic', '["Devils"]', 'Infernal', '/api/ruleset/5e/languages/infernal'),
('primordial', 'Primordial', NULL, 'Exotic', '["Elementals"]', 'Dwarvish', '/api/ruleset/5e/languages/primordial'),
('sylvan', 'Sylvan', NULL, 'Exotic', '["Fey creatures"]', 'Elvish', '/api/ruleset/5e/languages/sylvan'),
('undercommon', 'Undercommon', NULL, 'Exotic', '["Underdark traders"]', 'Elvish', '/api/ruleset/5e/languages/undercommon');
