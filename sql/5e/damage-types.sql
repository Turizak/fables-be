DROP TABLE IF EXISTS damage_types;

CREATE TABLE damage_types (
    index VARCHAR PRIMARY KEY,
    name VARCHAR,
    description TEXT,
    url VARCHAR
);

INSERT INTO
    damage_types (index, name, description, url)
VALUES
    (
        'acid',
        'Acid',
        'The corrosive spray of a black dragon''s breath and the dissolving enzymes secreted by a black pudding deal acid damage.',
        '/api/ruleset/5e/damage-types/acid'
    ),
    (
        'bludgeoning',
        'Bludgeoning',
        'Blunt force attacks, falling, constriction, and the like deal bludgeoning damage.',
        '/api/ruleset/5e/damage-types/bludgeoning'
    ),
    (
        'cold',
        'Cold',
        'The infernal chill radiating from an ice devil''s spear and the frigid blast of a white dragon''s breath deal cold damage.',
        '/api/ruleset/5e/damage-types/cold'
    ),
    (
        'fire',
        'Fire',
        'Red dragons breathe fire, and many spells conjure flames to deal fire damage.',
        '/api/ruleset/5e/damage-types/fire'
    ),
    (
        'force',
        'Force',
        'Force is pure magical energy focused into a damaging form. Most effects that deal force damage are spells, including magic missile and spiritual weapon.',
        '/api/ruleset/5e/damage-types/force'
    ),
    (
        'lightning',
        'Lightning',
        'A lightning bolt spell and a blue dragon''s breath deal lightning damage.',
        '/api/ruleset/5e/damage-types/lightning'
    ),
    (
        'necrotic',
        'Necrotic',
        'Necrotic damage, dealt by certain undead and a spell such as chill touch, withers matter and even the soul.',
        '/api/ruleset/5e/damage-types/necrotic'
    ),
    (
        'piercing',
        'Piercing',
        'Puncturing and impaling attacks, including spears and monsters'' bites, deal piercing damage.',
        '/api/ruleset/5e/damage-types/piercing'
    ),
    (
        'poison',
        'Poison',
        'Venomous stings and the toxic gas of a green dragon''s breath deal poison damage.',
        '/api/ruleset/5e/damage-types/poison'
    ),
    (
        'psychic',
        'Psychic',
        'Mental abilities such as a psionic blast deal psychic damage.',
        '/api/ruleset/5e/damage-types/psychic'
    ),
    (
        'radiant',
        'Radiant',
        'Radiant damage, dealt by a cleric''s flame strike spell or an angel''s smiting weapon, sears the flesh like fire and overloads the spirit with power.',
        '/api/ruleset/5e/damage-types/radiant'
    ),
    (
        'slashing',
        'Slashing',
        'Swords, axes, and monsters'' claws deal slashing damage.',
        '/api/ruleset/5e/damage-types/slashing'
    ),
    (
        'thunder',
        'Thunder',
        'A concussive burst of sound, such as the effect of the thunderwave spell, deals thunder damage.',
        '/api/ruleset/5e/damage-types/thunder'
    );