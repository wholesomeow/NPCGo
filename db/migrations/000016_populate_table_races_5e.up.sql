COPY generator.races_5e(id, race, subrace, adult_age_min, adult_age_max, covering, convering_alt, incidence, race_size, speed, race_language, strength, dexterity, constitution, intelligence, wisdom, charisma, extra)
FROM '/rawdata/csv/Races_5e.csv'
DELIMITER ','
CSV HEADER;