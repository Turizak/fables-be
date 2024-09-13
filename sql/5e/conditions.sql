DROP TABLE IF EXISTS condition_descriptions;
DROP TABLE IF EXISTS conditions;

CREATE TABLE conditions (
    index VARCHAR PRIMARY KEY,
    name VARCHAR,
    description TEXT,
    url VARCHAR
);

CREATE TABLE condition_descriptions (
    index VARCHAR REFERENCES conditions(index),
    description TEXT
);

INSERT INTO conditions (index, name, url) 
VALUES
('blinded', 'Blinded', '/api/conditions/blinded'),
('charmed', 'Charmed', '/api/conditions/charmed'),
('deafened', 'Deafened', '/api/conditions/deafened'),
('frightened', 'Frightened', '/api/conditions/frightened'),
('grappled', 'Grappled', '/api/conditions/grappled'),
('incapacitated', 'Incapacitated', '/api/conditions/incapacitated'),
('invisible', 'Invisible', '/api/conditions/invisible'),
('paralyzed', 'Paralyzed', '/api/conditions/paralyzed'),
('petrified', 'Petrified', '/api/conditions/petrified'),
('poisoned', 'Poisoned', '/api/conditions/poisoned'),
('prone', 'Prone', '/api/conditions/prone'),
('restrained', 'Restrained', '/api/conditions/restrained'),
('stunned', 'Stunned', '/api/conditions/stunned'),
('unconscious', 'Unconscious', '/api/conditions/unconscious'),
('exhaustion', 'Exhaustion', '/api/conditions/exhaustion');

INSERT INTO condition_descriptions (index, description) 
VALUES
('blinded', '- A blinded creature can''t see and automatically fails any ability check that requires sight.'),
('blinded', '- Attack rolls against the creature have advantage, and the creature''s attack rolls have disadvantage.'),
('charmed', '- A charmed creature can''t attack the charmer or target the charmer with harmful abilities or magical effects.'),
('charmed', '- The charmer has advantage on any ability check to interact socially with the creature.'),
('deafened', '- A deafened creature can''t hear and automatically fails any ability check that requires hearing.'),
('frightened', '- A frightened creature has disadvantage on ability checks and attack rolls while the source of its fear is within line of sight.'),
('frightened', '- The creature can''t willingly move closer to the source of its fear.'),
('grappled', '- A grappled creature''s speed becomes 0, and it can''t benefit from any bonus to its speed.'),
('grappled', '- The condition ends if the grappler is incapacitated (see the condition).'),
('grappled', '- The condition also ends if an effect removes the grappled creature from the reach of the grappler or grappling effect, such as when a creature is hurled away by the thunderwave spell.'),
('incapacitated', '- An incapacitated creature can''t take actions or reactions.'),
('invisible', '- An invisible creature is impossible to see without the aid of magic or a special sense. For the purpose of hiding, the creature is heavily obscured. The creature''s location can be detected by any noise it makes or any tracks it leaves.'),
('invisible', '- Attack rolls against the creature have disadvantage, and the creature''s attack rolls have advantage.'),
('paralyzed', '- A paralyzed creature is incapacitated (see the condition) and can''t move or speak.'),
('paralyzed', '- The creature automatically fails Strength and Dexterity saving throws.'),
('paralyzed', '- Attack rolls against the creature have advantage.'),
('paralyzed', '- Any attack that hits the creature is a critical hit if the attacker is within 5 feet of the creature.'),
('petrified', '- A petrified creature is transformed, along with any nonmagical object it is wearing or carrying, into a solid inanimate substance (usually stone). Its weight increases by a factor of ten, and it ceases aging.'),
('petrified', '- The creature is incapacitated (see the condition), can''t move or speak, and is unaware of its surroundings.'),
('petrified', '- Attack rolls against the creature have advantage.'),
('petrified', '- The creature automatically fails Strength and Dexterity saving throws.'),
('petrified', '- The creature has resistance to all damage.'),
('petrified', '- The creature is immune to poison and disease, although a poison or disease already in its system is suspended, not neutralized.'),
('poisoned', '- A poisoned creature has disadvantage on attack rolls and ability checks.'),
('prone', '- A prone creature''s only movement option is to crawl, unless it stands up and thereby ends the condition.'),
('prone', '- The creature has disadvantage on attack rolls.'),
('prone', '- An attack roll against the creature has advantage if the attacker is within 5 feet of the creature. Otherwise, the attack roll has disadvantage.'),
('restrained', '- A restrained creature''s speed becomes 0, and it can''t benefit from any bonus to its speed.'),
('restrained', '- Attack rolls against the creature have advantage, and the creature''s attack rolls have disadvantage.'),
('restrained', '- The creature has disadvantage on Dexterity saving throws.'),
('stunned', '- A stunned creature is incapacitated (see the condition), can''t move, and can speak only falteringly.'),
('stunned', '- The creature automatically fails Strength and Dexterity saving throws.'),
('stunned', '- Attack rolls against the creature have advantage.'),
('unconscious', '- An unconscious creature is incapacitated (see the condition), can''t move or speak, and is unaware of its surroundings.'),
('unconscious', '- The creature drops whatever it''s holding and falls prone.'),
('unconscious', '- The creature automatically fails Strength and Dexterity saving throws.'),
('unconscious', '- Attack rolls against the creature have advantage.'),
('unconscious', '- Any attack that hits the creature is a critical hit if the attacker is within 5 feet of the creature.'),
('exhaustion', 'Some special abilities and environmental hazards, such as starvation and the long-term effects of freezing or scorching temperatures, can lead to a special condition called exhaustion. Exhaustion is measured in six levels. An effect can give a creature one or more levels of exhaustion, as specified in the effect''s description.'),
('exhaustion', '1 - Disadvantage on ability checks'),
('exhaustion', '2 - Speed halved'),
('exhaustion', '3 - Disadvantage on attack rolls and saving throws'),
('exhaustion', '4 - Hit point maximum halved'),
('exhaustion', '5 - Speed reduced to 0'),
('exhaustion', '6 - Death'),
('exhaustion', 'If an already exhausted creature suffers another effect that causes exhaustion, its current level of exhaustion increases by the amount specified in the effect''s description.'),
('exhaustion', 'A creature suffers the effect of its current level of exhaustion as well as all lower levels. For example, a creature suffering level 2 exhaustion has its speed halved and has disadvantage on ability checks.'),
('exhaustion', 'An effect that removes exhaustion reduces its level as specified in the effect''s description, with all exhaustion effects ending if a creature''s exhaustion level is reduced below 1.'),
('exhaustion', 'Finishing a long rest reduces a creature''s exhaustion level by 1, provided that the creature has also ingested some food and drink.');
