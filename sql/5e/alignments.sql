DROP TABLE IF EXISTS alignments;

CREATE TABLE alignments (
    index VARCHAR PRIMARY KEY,
    alignment_name VARCHAR,
    abbreviation VARCHAR,
    description TEXT,
    url VARCHAR
);

INSERT INTO alignments (index, alignment_name, abbreviation, description, url) 
VALUES
('lawful-good', 'Lawful Good', 'LG', 'Lawful good (LG) creatures can be counted on to do the right thing as expected by society. Gold dragons, paladins, and most dwarves are lawful good.', '/api/ruleset/5e/alignments/lawful-good'),
('neutral-good', 'Neutral Good', 'NG', 'Neutral good (NG) folk do the best they can to help others according to their needs. Many celestials, some cloud giants, and most gnomes are neutral good.', '/api/ruleset/5e/alignments/neutral-good'),
('chaotic-good', 'Chaotic Good', 'CG', 'Chaotic good (CG) creatures act as their conscience directs, with little regard for what others expect. Copper dragons, many elves, and unicorns are chaotic good.', '/api/ruleset/5e/alignments/chaotic-good'),
('lawful-neutral', 'Lawful Neutral', 'LN', 'Lawful neutral (LN) individuals act in accordance with law, tradition, or personal codes. Many monks and some wizards are lawful neutral.', '/api/ruleset/5e/alignments/lawful-neutral'),
('neutral', 'Neutral', 'N', 'Neutral (N) is the alignment of those who prefer to steer clear of moral questions and don''t take sides, doing what seems best at the time. Lizardfolk, most druids, and many humans are neutral.', '/api/ruleset/5e/alignments/neutral'),
('chaotic-neutral', 'Chaotic Neutral', 'CN', 'Chaotic neutral (CN) creatures follow their whims, holding their personal freedom above all else. Many barbarians and rogues, and some bards, are chaotic neutral.', '/api/ruleset/5e/alignments/chaotic-neutral'),
('lawful-evil', 'Lawful Evil', 'LE', 'Lawful evil (LE) creatures methodically take what they want, within the limits of a code of tradition, loyalty, or order. Devils, blue dragons, and hobgoblins are lawful evil.', '/api/ruleset/5e/alignments/lawful-evil'),
('neutral-evil', 'Neutral Evil', 'NE', 'Neutral evil (NE) is the alignment of those who do whatever they can get away with, without compassion or qualms. Many drow, some cloud giants, and goblins are neutral evil.', '/api/ruleset/5e/alignments/neutral-evil'),
('chaotic-evil', 'Chaotic Evil', 'CE', 'Chaotic evil (CE) creatures act with arbitrary violence, spurred by their greed, hatred, or bloodlust. Demons, red dragons, and orcs are chaotic evil.', '/api/ruleset/5e/alignments/chaotic-evil');
