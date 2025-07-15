COPY generator.sex_orientations_npc(id, sexual_orientation, sexual_orientation_description)
FROM '/rawdata/csv/Sexual_Orientations.csv'
DELIMITER ','
CSV HEADER;