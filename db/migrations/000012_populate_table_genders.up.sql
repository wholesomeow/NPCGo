COPY generator.genders(id, gender, gender_description, pronouns, secondary_pronouns, tirtiary_pronouns)
FROM '/rawdata/csv/Genders.csv'
DELIMITER ','
CSV HEADER;