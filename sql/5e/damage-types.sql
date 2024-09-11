DROP TABLE IF EXISTS damage_type_descriptions;
DROP TABLE IF EXISTS damage_types;

CREATE TABLE damage_types (
    index VARCHAR PRIMARY KEY,
    name VARCHAR,
    url VARCHAR
);

CREATE TABLE damage_type_descriptions (
    damage_type_index VARCHAR REFERENCES damage_types(index),
    description TEXT
);

INSERT INTO damage_types (index, name, url) 
VALUES
('acid', 'Acid', '/api/damage-types/acid'),
('bludgeoning', 'Bludgeoning', '/api/damage-types/bludgeoning'),
('cold', 'Cold', '/api/damage-types/cold'),
('fire', 'Fire', '/api/damage-types/fire'),
('force', 'Force', '/api/damage-types/force'),
('lightning', 'Lightning', '/api/damage-types/lightning'),
('necrotic', 'Necrotic', '/api/damage-types/necrotic'),
('piercing', 'Piercing', '/api/damage-types/piercing'),
('poison', 'Poison', '/api/damage-types/poison'),
('psychic', 'Psychic', '/api/damage-types/psychic'),
('radiant', 'Radiant', '/api/damage-types/radiant'),
('slashing', 'Slashing', '/api/damage-types/slashing'),
('thunder', 'Thunder', '/api/damage-types/thunder');

INSERT INTO damage_type_descriptions (damage_type_index, description) 
VALUES
('acid', 'The corrosive spray of a black dragon''s breath and the dissolving enzymes secreted by a black pudding deal acid damage.'),
('bludgeoning', 'Blunt force attacks, falling, constriction, and the like deal bludgeoning damage.'),
('cold', 'The infernal chill radiating from an ice devil''s spear and the frigid blast of a white dragon''s breath deal cold damage.'),
('fire', 'Red dragons breathe fire, and many spells conjure flames to deal fire damage.'),
('force', 'Force is pure magical energy focused into a damaging form. Most effects that deal force damage are spells, including magic missile and spiritual weapon.'),
('lightning', 'A lightning bolt spell and a blue dragon''s breath deal lightning damage.'),
('necrotic', 'Necrotic damage, dealt by certain undead and a spell such as chill touch, withers matter and even the soul.'),
('piercing', 'Puncturing and impaling attacks, including spears and monsters'' bites, deal piercing damage.'),
('poison', 'Venomous stings and the toxic gas of a green dragon''s breath deal poison damage.'),
('psychic', 'Mental abilities such as a psionic blast deal psychic damage.'),
('radiant', 'Radiant damage, dealt by a cleric''s flame strike spell or an angel''s smiting weapon, sears the flesh like fire and overloads the spirit with power.'),
('slashing', 'Swords, axes, and monsters'' claws deal slashing damage.'),
('thunder', 'A concussive burst of sound, such as the effect of the thunderwave spell, deals thunder damage.');
